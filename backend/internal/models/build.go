package models

import (
	"time"
)

type CustomBuild struct {
	BuildID    string    `json:"build_id"`
	UserID     string    `json:"user_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	TotalPrice float64   `json:"total_price"`
}

type BuildItem struct {
	BuildID      string  `json:"build_id"`
	ProductID    int     `json:"product_id"`
	Quantity     int     `json:"quantity"`
	ProductName  string  `json:"product_name"`
	Price        float64 `json:"price"`
	Description  string  `json:"description"`
	BrandName    string  `json:"brand_name"`
	CategoryName string  `json:"category_name"`
}

type BuildWithItems struct {
	CustomBuild
	Items []BuildItem `json:"items"`
}
