package dtos

type CreateBuildItemDTO struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required,min=1"`
}

type CreateBuildRequestDTO struct {
	Name  string `json:"name" binding:"required"`
	Items []struct {
		ProductID int `json:"product_id" binding:"required"`
		Quantity  int `json:"quantity" binding:"required,min=1"`
	} `json:"items" binding:"required,min=1"`
}

type UpdateBuildRequestDTO struct {
	Name  string `json:"name" binding:"omitempty"`
	Items []struct {
		ProductID int `json:"product_id" binding:"required"`
		Quantity  int `json:"quantity" binding:"required,min=1"`
	} `json:"items" binding:"omitempty,min=1"`
}

type BuildResponseDTO struct {
	BuildID    string  `json:"build_id"`
	UserID     string  `json:"user_id"`
	Name       string  `json:"name"`
	CreatedAt  string  `json:"created_at"`
	TotalPrice float64 `json:"total_price"`
	Items      []struct {
		ProductID    int     `json:"product_id"`
		Quantity     int     `json:"quantity"`
		ProductName  string  `json:"product_name"`
		Price        float64 `json:"price"`
		Description  string  `json:"description"`
		BrandName    string  `json:"brand_name"`
		CategoryName string  `json:"category_name"`
	} `json:"items"`
}

type BuildItemResponseDTO struct {
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}
