package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/amha-mersha/sanqa-suq/internal/database"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/jackc/pgx/v5"
)

type CategoryRepository struct {
	DB *database.DB
}

func NewCategoryRepository(db *database.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (repositories *CategoryRepository) GetAllCategories(ctx context.Context) ([]models.Categories, error) {
	query := `SELECT category_id, category_name, parent_category_id FROM categories;`
	rows, err := repositories.DB.Pool.Query(ctx, query)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.NotFound("Category not found", nil)
		}
		return nil, errs.InternalError("Failed to fetch category", err)
	}
	defer rows.Close()

	categories, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Categories, error) {
		var c models.Categories
		err := row.Scan(&c.CategoryID, &c.Name, &c.ParentCategoryID)
		return c, err
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.NotFound("Category not found", nil)
		}
		return nil, errs.InternalError("Failed to fetch category", err)
	}
	return categories, nil
}

func (r *CategoryRepository) InsertCategory(ctx context.Context, newCategory *models.Categories) (*models.Categories, error) {
	query := `INSERT INTO categories (category_name, parent_category_id) 
	          VALUES ($1, $2) RETURNING category_id;`
	err := r.DB.Pool.QueryRow(ctx, query, newCategory.Name, newCategory.ParentCategoryID).
		Scan(&newCategory.CategoryID)
	if err != nil {
		return nil, errs.InternalError("Failed to insert category", err)
	}
	return newCategory, nil
}

func (r *CategoryRepository) GetCategoryById(ctx context.Context, categoryId int) (*models.Categories, error) {
	query := `SELECT category_id, category_name, parent_category_id 
	          FROM categories WHERE category_id = $1;`
	row := r.DB.Pool.QueryRow(ctx, query, categoryId)
	var category models.Categories
	err := row.Scan(&category.CategoryID, &category.Name, &category.ParentCategoryID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.NotFound("Category not found", nil)
		}
		return nil, errs.InternalError("Failed to fetch category", err)
	}
	return &category, nil
}

func (r *CategoryRepository) FetchCategoryChildren(ctx context.Context, categoryId string) ([]models.Categories, error) {
	query := `SELECT category_id, category_name, parent_category_id FROM categories WHERE parent_category_id = $1;`
	rows, err := r.DB.Pool.Query(ctx, query, categoryId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.NotFound("Category not found", nil)
		}
		return nil, errs.InternalError("Failed to fetch category", err)
	}
	defer rows.Close()
	childCategories, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Categories, error) {
		var c models.Categories
		err := row.Scan(&c.CategoryID, &c.Name, &c.ParentCategoryID)
		return c, err
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.NotFound("Category not found", nil)
		}
		return nil, errs.InternalError("Failed to fetch category", err)
	}
	return childCategories, nil
}

func (r *CategoryRepository) FetchCategoryTree(ctx context.Context, categoryId int, limit int) ([]*models.CategoryNode, error) {
	if categoryId <= 0 {
		return nil, errs.BadRequest("invalid category ID", fmt.Errorf("categoryId must be positive, got %d", categoryId))
	}
	if limit < 0 {
		return nil, errs.BadRequest("invalid limit", fmt.Errorf("limit must be non-negative, got %d", limit))
	}

	query := `
		WITH RECURSIVE category_tree AS (
			-- Anchor: Select the starting category
			SELECT category_id, category_name, parent_category_id, 0 AS level
			FROM categories
			WHERE category_id = $1
			UNION ALL
			-- Recursive: Join children with the previous step
			SELECT c.category_id, c.category_name, c.parent_category_id, ct.level + 1
			FROM categories c
			JOIN category_tree ct ON c.parent_category_id = ct.category_id
			WHERE ct.level < $2
		)
		SELECT category_id, category_name, parent_category_id FROM category_tree
		ORDER BY level, category_id;
	`

	rows, err := r.DB.Pool.Query(ctx, query, categoryId, limit)
	if err != nil {
		return nil, errs.InternalError(fmt.Sprintf("failed to query category tree for category ID %d", categoryId), err)
	}
	defer rows.Close()

	// Use a map to track all nodes by their ID for easy lookup
	nodeMap := make(map[int]*models.CategoryNode)
	var rootNode *models.CategoryNode

	for rows.Next() {
		var id int
		var name string
		var parentId *int

		if err := rows.Scan(&id, &name, &parentId); err != nil {
			return nil, errs.InternalError(fmt.Sprintf("failed to scan category row for category ID %d", categoryId), err)
		}

		node := &models.CategoryNode{
			CategoryID:       id,
			Name:             name,
			ParentCategoryID: parentId,
			Children:         make([]*models.CategoryNode, 0),
		}
		nodeMap[id] = node

		// Identify the root category
		if id == categoryId {
			rootNode = node
		}
	}

	if rootNode == nil {
		return nil, errs.NotFound(fmt.Sprintf("category with ID %d not found", categoryId), nil)
	}

	// Build the tree structure
	for _, node := range nodeMap {
		if node.ParentCategoryID != nil {
			if parentNode, ok := nodeMap[*node.ParentCategoryID]; ok {
				parentNode.Children = append(parentNode.Children, node)
			}
		}
	}

	// Return the root node in an array
	return []*models.CategoryNode{rootNode}, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, categoryId int, fields map[string]any) error {
	setClause := []string{}
	args := []any{}
	i := 1
	for key, value := range fields {
		setClause = append(setClause, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}
	args = append(args, fields["category_id"])
	query := fmt.Sprintf("UPDATE categories SET %s WHERE category_id = $%d", strings.Join(setClause, ", "), i)
	_, err := r.DB.Pool.Exec(ctx, query, args...)
	if err != nil {
		return errs.InternalError("Failed to update category", err)
	}
	return nil
}
func (r *CategoryRepository) DeleteCategory(ctx context.Context, categoryId int) error {
	query := `DELETE FROM categories WHERE category_id = $1;`
	result, err := r.DB.Pool.Exec(ctx, query, categoryId)
	if err != nil {
		return errs.InternalError("Failed to delete category", err)
	}
	if result.RowsAffected() == 0 {
		return errs.NotFound("Category not found", nil)
	}
	return nil
}

// FetchCategoryAncestors retrieves the lineage of ancestors for a given category,
// starting from the child and ending with the highest ancestor.
func (r *CategoryRepository) FetchCategoryAncestors(ctx context.Context, categoryId int) ([]models.Categories, error) {
	// The SQL query uses a Common Table Expression (CTE) with recursion to find ancestors.
	//
	// Explanation of the SQL:
	// 1. WITH RECURSIVE category_ancestors AS (...): Defines a recursive CTE named category_ancestors.
	// 2. Base Case (SELECT ... FROM categories WHERE category_id = $1):
	//    - This is the starting point of the recursion. It selects the initial category
	//      identified by categoryId ($1).
	// 3. UNION ALL: Combines the result of the base case with the recursive step.
	// 4. Recursive Step (SELECT c.category_id, ... FROM categories c JOIN category_ancestors ca ON ca.parent_category_id = c.category_id):
	//    - 'categories c': Refers to the main categories table.
	//    - 'category_ancestors ca': Refers to the results accumulated so far in the CTE (which are the categories found in the previous step, starting with the initial child).
	//    - 'ON ca.parent_category_id = c.category_id': This is the crucial part for finding ancestors.
	//      It says: "Find a category 'c' in the main 'categories' table whose 'category_id'
	//      matches the 'parent_category_id' of an entry 'ca' that was found in the previous
	//      step of the recursion."
	//      Essentially, if 'ca' is a child, 'c' will be its parent.
	//    - The recursion continues until `ca.parent_category_id` is NULL, indicating the top-most ancestor has been reached.
	// 5. Final SELECT (SELECT category_id, category_name, parent_category_id FROM category_ancestors;):
	//    - Selects all the accumulated rows from the `category_ancestors` CTE.
	//    - The order will naturally be from the initial child up to the highest ancestor,
	//      as that's how the recursion builds the set.
	query := `
		WITH RECURSIVE category_ancestors AS (
			-- Base case: Select the initial category
			SELECT
				category_id,
				category_name,
				parent_category_id
			FROM categories
			WHERE category_id = $1

			UNION ALL

			-- Recursive step: Find the parent of the categories found in the previous step
			SELECT
				c.category_id,
				c.category_name,
				c.parent_category_id
			FROM categories c
			JOIN category_ancestors ca ON c.category_id = ca.parent_category_id
			WHERE ca.parent_category_id IS NOT NULL -- Stop recursion when we reach a top-level category (parent_category_id is NULL)
		)
		SELECT category_id, category_name, parent_category_id
		FROM category_ancestors;
	`

	rows, err := r.DB.Pool.Query(ctx, query, categoryId)
	if err != nil {
		return nil, errs.InternalError("Failed to fetch ancestors", err)
	}
	defer rows.Close()

	ancestors, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Categories, error) {
		var c models.Categories
		err := row.Scan(&c.CategoryID, &c.Name, &c.ParentCategoryID)
		return c, err
	})
	if err != nil {
		return nil, errs.InternalError("Failed to collect ancestors", err)
	}

	return ancestors, nil
}

// FetchCategoryDescendants retrieves all direct and indirect children (descendants)
// of a given category, including the initial category itself.
func (r *CategoryRepository) FetchCategoryDescendants(ctx context.Context, categoryId int) ([]models.Categories, error) {
	// The SQL query uses a Common Table Expression (CTE) with recursion to find descendants.
	//
	// Explanation of the SQL:
	// 1. WITH RECURSIVE category_descendants AS (...): Defines a recursive CTE named category_descendants.
	// 2. Base Case (SELECT ... FROM categories WHERE category_id = $1):
	//    - This is the starting point of the recursion. It selects the initial category
	//      identified by categoryId ($1). This category is considered a descendant of itself.
	// 3. UNION ALL: Combines the result of the base case with the recursive step.
	// 4. Recursive Step (SELECT c.category_id, ... FROM categories c JOIN category_descendants cd ON c.parent_category_id = cd.category_id):
	//    - 'categories c': Refers to the main categories table.
	//    - 'category_descendants cd': Refers to the results accumulated so far in the CTE
	//      (which are the categories found in the previous step, starting with the initial category).
	//    - 'ON c.parent_category_id = cd.category_id': This is the crucial part for finding descendants.
	//      It says: "Find a category 'c' in the main 'categories' table whose 'parent_category_id'
	//      matches the 'category_id' of an entry 'cd' that was found in the previous
	//      step of the recursion."
	//      Essentially, if 'cd' is a parent, 'c' will be its child.
	//    - The recursion continues as long as children are found.
	// 5. Final SELECT (SELECT category_id, category_name, parent_category_id FROM category_descendants;):
	//    - Selects all the accumulated rows from the `category_descendants` CTE.
	//    - The order might not be strictly hierarchical by default, but it will include all descendants.
	query := `
		WITH RECURSIVE category_descendants AS (
			-- Base case: Start with the initial category
			SELECT
				category_id,
				category_name,
				parent_category_id
			FROM categories
			WHERE category_id = $1

			UNION ALL

			-- Recursive step: Find children of the categories found in the previous step
			SELECT
				c.category_id,
				c.category_name,
				c.parent_category_id
			FROM categories c
			JOIN category_descendants cd ON c.parent_category_id = cd.category_id
		)
		SELECT category_id, category_name, parent_category_id
		FROM category_descendants;
	`

	rows, err := r.DB.Pool.Query(ctx, query, categoryId)
	if err != nil {
		return nil, errs.InternalError("Failed to fetch descendants", err)
	}
	defer rows.Close()

	descendants, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Categories, error) {
		var c models.Categories
		err := row.Scan(&c.CategoryID, &c.Name, &c.ParentCategoryID)
		return c, err
	})
	if err != nil {
		return nil, errs.InternalError("Failed to collect descendants", err)
	}

	return descendants, nil
}
