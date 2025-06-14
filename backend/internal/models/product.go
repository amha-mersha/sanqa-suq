package models

import "time"

type Categories struct {
	CategoryID       int    `json:"category_id"`
	Name             string `json:"name"`
	ParentCategoryID *int   `json:"parent_category_id"`
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
