package dtos

var ModelToDatabaseMap = map[string]string{
	"CategoryID":    "categroy_id",
	"BrandID":       "brand_id",
	"Name":          "name",
	"Description":   "description",
	"Price":         "price",
	"StockQuantity": "stock_quantity",
}

type ProductUpdateDTO struct {
	CategoryID    *int     `json:"categroy_id" binding:"omitempty"`
	BrandID       *int     `json:"brand_id" binding:"omitempty"`
	Name          *string  `json:"name" binding:"omitempty"`
	Description   *string  `json:"description" binding:"omitempty"`
	Price         *float64 `json:"price" binding:"omitempty"`
	StockQuantity *int     `json:"stock_quantity" binding:"omitempty"`
}

type CreateProductDTO struct {
	CategoryID    int     `json:"category_id" binding:"required"`
	BrandID       int     `json:"brand_id" binding:"required"`
	Name          string  `json:"name" binding:"required"`
	Description   string  `json:"description" binding:"required"`
	Price         float64 `json:"price" binding:"required"`
	StockQuantity int     `json:"stock_quantity" binding:"required"`
}

type CreateReviewDTO struct {
	UserID    string `json:"user_id" binding:"required,uuid"`
	ProductID int    `json:"product_id" binding:"required"`
	Rating    int    `json:"rating" binding:"required,min=1,max=5"`
	Comment   string `json:"comment" binding:"required"`
}
type UpdateReviewDTO struct {
	Rating  *int    `json:"rating" binding:"omitempty,min=1,max=5"`
	Comment *string `json:"comment" binding:"omitempty"`
}
