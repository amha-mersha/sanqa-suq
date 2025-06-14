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
	CategoryID    *int     `json:"categroy_id"`
	BrandID       *int     `json:"brand_id"`
	Name          *string  `json:"name"`
	Description   *string  `json:"description"`
	Price         *float64 `json:"price"`
	StockQuantity *int     `json:"stock_quantity"`
}

type CreateProductDTO struct {
	CategoryID    int     `json:"categroy_id"`
	BrandID       int     `json:"brand_id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
}
