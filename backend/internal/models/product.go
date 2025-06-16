package models

import "time"

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

type Brands struct {
	BrandID     int    `json:"brand_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Products struct {
	ProductID     int       `json:"product_id"`
	CategoryID    int       `json:"categroy_id"`
	BrandID       int       `json:"brand_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         *float64  `json:"price"`
	StockQuantity *int      `json:"stock_quantity"`
	CreatedAt     time.Time `json:"created_at"`
}
