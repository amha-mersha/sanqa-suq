package models

import "time"

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

type ProductSpecifications struct {
	ProductID int    `json:"product_id"`
	SpecName  string `json:"spec_name"`
	SpecValue string `json:"spec_value"`
}
type Review struct {
	ReviewId   int       `json:"review_id"`
	UserId     int       `json:"user_id"`
	ProductId  int       `json:"product_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
	ReviewDate time.Time `json:"review_date"`
}
