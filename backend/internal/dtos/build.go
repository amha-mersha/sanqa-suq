package dtos

type BuildItemDTO struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required,min=1"`
}

type CreateBuildRequestDTO struct {
	Name  string         `json:"name" binding:"required"`
	Items []BuildItemDTO `json:"items" binding:"required,min=1"`
}

type UpdateBuildRequestDTO struct {
	Name  string         `json:"name" binding:"omitempty"`
	Items []BuildItemDTO `json:"items" binding:"omitempty,min=1"`
}

type BuildItemResponseDTO struct {
	ProductID    int     `json:"product_id"`
	Quantity     int     `json:"quantity"`
	ProductName  string  `json:"product_name"`
	Price        float64 `json:"price"`
	Description  string  `json:"description"`
	BrandName    string  `json:"brand_name"`
	CategoryName string  `json:"category_name"`
}

type BuildResponseDTO struct {
	BuildID    string                 `json:"build_id"`
	UserID     string                 `json:"user_id"`
	Name       string                 `json:"name"`
	CreatedAt  string                 `json:"created_at"`
	TotalPrice float64                `json:"total_price"`
	Items      []BuildItemResponseDTO `json:"items"`
}

type CompatibleProductsRequestDTO struct {
	CategoryID    int   `json:"category_id" binding:"required"`
	SelectedItems []int `json:"selected_items" binding:"omitempty"`
}

type CompatibleProductDTO struct {
	ProductID    int               `json:"product_id"`
	ProductName  string            `json:"product_name"`
	Price        float64           `json:"price"`
	Description  string            `json:"description"`
	BrandName    string            `json:"brand_name"`
	CategoryName string            `json:"category_name"`
	Specs        map[string]string `json:"specs"`
}
