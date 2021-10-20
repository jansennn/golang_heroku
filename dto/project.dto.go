package dto

type CreateProjectRequest struct {
	Name string `json:"name" form:"name" binding:"required,min=1"`
	Image  string `json:"image" form:"image" binding:"required,min=1"`
	Description string `json:"description" form:"description" binding:"required,min=1"`
	LinkCode string `json:"link_code" form:"link_code" binding:"required,min=1"`
}

type UpdateProjectRequest struct {
	ID    int64  `json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required,min=1"`
	Image  string `json:"image" form:"image" binding:"required,min=1"`
	Description string `json:"description" form:"description" binding:"required,min=1"`
	LinkCode string `json:"link_code" form:"link_code" binding:"required,min=1"`
}