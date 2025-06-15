package dtos

type CreateCategoryDTO struct {
	Name             string `json:"name" binding:"required"`
	ParentCategoryID *int   `json:"parent_category_id" binding:"omitempty"`
}

type UpdateCategoryDTO struct {
	Name             *string `json:"name,omitempty"`
	ParentCategoryID *int    `json:"parent_category_id,omitempty"`
}
