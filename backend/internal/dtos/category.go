package dtos

type CreateCategoryDTO struct {
	Name             string `json:"name"`
	ParentCategoryID *int   `json:"parent_category_id,omitempty"`
}

type UpdateCategoryDTO struct {
	Name             *string `json:"name,omitempty"`
	ParentCategoryID *int    `json:"parent_category_id,omitempty"`
}
