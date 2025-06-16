package models

// Categories is the database model struct, representing a single category entry.
type Categories struct {
	CategoryID       int    `json:"category_id"`
	Name             string `json:"name"`
	ParentCategoryID *int   `json:"parent_category_id"` // Pointer for nullable parent ID from DB
}

// CategoryNode is used for the API response, representing a node in the category tree.
// It includes a Children field for nested subcategories.
type CategoryNode struct {
	CategoryID       int             `json:"category_id"`
	Name             string          `json:"name"`
	ParentCategoryID *int            `json:"parent_category_id,omitempty"` // Omit if nil in JSON
	Children         []*CategoryNode `json:"children,omitempty"`           // Omit if empty in JSON
}
