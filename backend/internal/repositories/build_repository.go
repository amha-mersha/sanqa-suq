package repositories

import (
	"context"

	"github.com/amha-mersha/sanqa-suq/internal/database"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/jackc/pgx/v5"
)

type BuildRepository struct {
	DB *database.DB
}

func NewBuildRepository(db *database.DB) *BuildRepository {
	return &BuildRepository{DB: db}
}

func (r *BuildRepository) CreateBuild(ctx context.Context, build *models.CustomBuild, items []models.BuildItem) (*models.BuildWithItems, error) {
	tx, err := r.DB.Pool.Begin(ctx)
	if err != nil {
		return nil, errs.InternalError("failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	// Insert the build
	var buildID string
	err = tx.QueryRow(ctx,
		`INSERT INTO custom_builds (user_id, name) 
		 VALUES ($1, $2) 
		 RETURNING build_id`,
		build.UserID, build.Name,
	).Scan(&buildID)
	if err != nil {
		return nil, errs.InternalError("failed to create build", err)
	}

	// Insert build items
	for _, item := range items {
		_, err = tx.Exec(ctx,
			`INSERT INTO build_items (build_id, product_id, quantity)
			 VALUES ($1, $2, $3)`,
			buildID, item.ProductID, item.Quantity,
		)
		if err != nil {
			return nil, errs.InternalError("failed to create build item", err)
		}
	}

	// Validate build compatibility using the database function
	var isCompatible bool
	var message string
	err = tx.QueryRow(ctx,
		`SELECT is_compatible, message FROM validate_build($1)`,
		buildID,
	).Scan(&isCompatible, &message)
	if err != nil {
		return nil, errs.InternalError("failed to validate build", err)
	}

	if !isCompatible {
		return nil, errs.UnprocessableEntity(message, nil)
	}

	// Get the created build with items
	var result models.BuildWithItems
	err = tx.QueryRow(ctx,
		`SELECT b.build_id, b.user_id, b.name, b.created_at, b.total_price
		 FROM custom_builds b
		 WHERE b.build_id = $1`,
		buildID,
	).Scan(&result.BuildID, &result.UserID, &result.Name, &result.CreatedAt, &result.TotalPrice)
	if err != nil {
		return nil, errs.InternalError("failed to fetch created build", err)
	}

	// Get build items with product details
	rows, err := tx.Query(ctx,
		`SELECT bi.product_id, bi.quantity, p.name, p.price, p.description,
				b.name as brand_name, c.name as category_name
		 FROM build_items bi
		 JOIN products p ON bi.product_id = p.product_id
		 JOIN brands b ON p.brand_id = b.brand_id
		 JOIN categories c ON p.category_id = c.category_id
		 WHERE bi.build_id = $1`,
		buildID,
	)
	if err != nil {
		return nil, errs.InternalError("failed to fetch build items", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.BuildItem
		if err := rows.Scan(&item.ProductID, &item.Quantity, &item.ProductName, &item.Price,
			&item.Description, &item.BrandName, &item.CategoryName); err != nil {
			return nil, errs.InternalError("failed to scan build item", err)
		}
		buildID := buildID // Create a new variable to avoid closure issues
		item.BuildID = buildID
		result.Items = append(result.Items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, errs.InternalError("error iterating build items", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, errs.InternalError("failed to commit transaction", err)
	}

	return &result, nil
}

func (r *BuildRepository) GetUserBuilds(ctx context.Context, userID string) ([]models.BuildWithItems, error) {
	rows, err := r.DB.Pool.Query(ctx,
		`SELECT b.build_id, b.user_id, b.name, b.created_at, b.total_price
		 FROM custom_builds b
		 WHERE b.user_id = $1
		 ORDER BY b.created_at DESC`,
		userID,
	)
	if err != nil {
		return nil, errs.InternalError("failed to fetch user builds", err)
	}
	defer rows.Close()

	var builds []models.BuildWithItems
	for rows.Next() {
		var build models.BuildWithItems
		if err := rows.Scan(&build.BuildID, &build.UserID, &build.Name, &build.CreatedAt, &build.TotalPrice); err != nil {
			return nil, errs.InternalError("failed to scan build", err)
		}

		// Get build items with product details
		itemRows, err := r.DB.Pool.Query(ctx,
			`SELECT bi.product_id, bi.quantity, p.name, p.price, p.description,
					b.name as brand_name, c.name as category_name
			 FROM build_items bi
			 JOIN products p ON bi.product_id = p.product_id
			 JOIN brands b ON p.brand_id = b.brand_id
			 JOIN categories c ON p.category_id = c.category_id
			 WHERE bi.build_id = $1`,
			build.BuildID,
		)
		if err != nil {
			return nil, errs.InternalError("failed to fetch build items", err)
		}
		defer itemRows.Close()

		for itemRows.Next() {
			var item models.BuildItem
			if err := itemRows.Scan(&item.ProductID, &item.Quantity, &item.ProductName, &item.Price,
				&item.Description, &item.BrandName, &item.CategoryName); err != nil {
				return nil, errs.InternalError("failed to scan build item", err)
			}
			item.BuildID = build.BuildID
			build.Items = append(build.Items, item)
		}

		if err = itemRows.Err(); err != nil {
			return nil, errs.InternalError("error iterating build items", err)
		}

		builds = append(builds, build)
	}

	if err = rows.Err(); err != nil {
		return nil, errs.InternalError("error iterating builds", err)
	}

	return builds, nil
}

func (r *BuildRepository) GetBuildByID(ctx context.Context, buildID string) (*models.BuildWithItems, error) {
	var build models.BuildWithItems
	err := r.DB.Pool.QueryRow(ctx,
		`SELECT b.build_id, b.user_id, b.name, b.created_at, b.total_price
		 FROM custom_builds b
		 WHERE b.build_id = $1`,
		buildID,
	).Scan(&build.BuildID, &build.UserID, &build.Name, &build.CreatedAt, &build.TotalPrice)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.NotFound("build not found", nil)
		}
		return nil, errs.InternalError("failed to fetch build", err)
	}

	// Get build items with product details
	rows, err := r.DB.Pool.Query(ctx,
		`SELECT bi.product_id, bi.quantity, p.name, p.price, p.description,
				b.name as brand_name, c.name as category_name
		 FROM build_items bi
		 JOIN products p ON bi.product_id = p.product_id
		 JOIN brands b ON p.brand_id = b.brand_id
		 JOIN categories c ON p.category_id = c.category_id
		 WHERE bi.build_id = $1`,
		buildID,
	)
	if err != nil {
		return nil, errs.InternalError("failed to fetch build items", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.BuildItem
		if err := rows.Scan(&item.ProductID, &item.Quantity, &item.ProductName, &item.Price,
			&item.Description, &item.BrandName, &item.CategoryName); err != nil {
			return nil, errs.InternalError("failed to scan build item", err)
		}
		item.BuildID = buildID
		build.Items = append(build.Items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, errs.InternalError("error iterating build items", err)
	}

	return &build, nil
}

func (r *BuildRepository) UpdateBuild(ctx context.Context, buildID string, userID string, name *string, items []models.BuildItem) (*models.BuildWithItems, error) {
	// First verify the build belongs to the user
	var existingBuild models.CustomBuild
	err := r.DB.Pool.QueryRow(ctx,
		`SELECT build_id, user_id, name 
		 FROM custom_builds 
		 WHERE build_id = $1`,
		buildID,
	).Scan(&existingBuild.BuildID, &existingBuild.UserID, &existingBuild.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.NotFound("build not found", nil)
		}
		return nil, errs.InternalError("failed to fetch build", err)
	}

	if existingBuild.UserID != userID {
		return nil, errs.Forbidden("you can only update your own builds", nil)
	}

	tx, err := r.DB.Pool.Begin(ctx)
	if err != nil {
		return nil, errs.InternalError("failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	// Update build name if provided
	if name != nil {
		_, err = tx.Exec(ctx,
			`UPDATE custom_builds 
			 SET name = $1 
			 WHERE build_id = $2`,
			*name, buildID,
		)
		if err != nil {
			return nil, errs.InternalError("failed to update build name", err)
		}
	}

	// Update items if provided
	if len(items) > 0 {
		// Delete existing items
		_, err = tx.Exec(ctx,
			`DELETE FROM build_items 
			 WHERE build_id = $1`,
			buildID,
		)
		if err != nil {
			return nil, errs.InternalError("failed to delete existing build items", err)
		}

		// Insert new items
		for _, item := range items {
			_, err = tx.Exec(ctx,
				`INSERT INTO build_items (build_id, product_id, quantity)
				 VALUES ($1, $2, $3)`,
				buildID, item.ProductID, item.Quantity,
			)
			if err != nil {
				return nil, errs.InternalError("failed to insert build item", err)
			}
		}

		// Validate build compatibility
		var isCompatible bool
		var message string
		err = tx.QueryRow(ctx,
			`SELECT is_compatible, message FROM validate_build($1)`,
			buildID,
		).Scan(&isCompatible, &message)
		if err != nil {
			return nil, errs.InternalError("failed to validate build", err)
		}

		if !isCompatible {
			return nil, errs.UnprocessableEntity(message, nil)
		}
	}

	// Get the updated build with items
	var result models.BuildWithItems
	err = tx.QueryRow(ctx,
		`SELECT b.build_id, b.user_id, b.name, b.created_at, b.total_price
		 FROM custom_builds b
		 WHERE b.build_id = $1`,
		buildID,
	).Scan(&result.BuildID, &result.UserID, &result.Name, &result.CreatedAt, &result.TotalPrice)
	if err != nil {
		return nil, errs.InternalError("failed to fetch updated build", err)
	}

	// Get build items with product details
	rows, err := tx.Query(ctx,
		`SELECT bi.product_id, bi.quantity, p.name, p.price, p.description,
				b.name as brand_name, c.name as category_name
		 FROM build_items bi
		 JOIN products p ON bi.product_id = p.product_id
		 JOIN brands b ON p.brand_id = b.brand_id
		 JOIN categories c ON p.category_id = c.category_id
		 WHERE bi.build_id = $1`,
		buildID,
	)
	if err != nil {
		return nil, errs.InternalError("failed to fetch build items", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.BuildItem
		if err := rows.Scan(&item.ProductID, &item.Quantity, &item.ProductName, &item.Price,
			&item.Description, &item.BrandName, &item.CategoryName); err != nil {
			return nil, errs.InternalError("failed to scan build item", err)
		}
		item.BuildID = buildID
		result.Items = append(result.Items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, errs.InternalError("error iterating build items", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, errs.InternalError("failed to commit transaction", err)
	}

	return &result, nil
}
