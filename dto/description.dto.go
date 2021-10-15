package dto

type CreateDescriptionRequest struct {
	Description  string  `json:"description" form:"description" binding:"required,min=1"`
}

type UpdateDescriptionRequest struct {
	ID    int64  `json:"id" form:"id"`
	Description  string  `json:"description" form:"description" binding:"required,min=1"`
}