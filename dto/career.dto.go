package dto

type CreateCareerRequest struct {
	Year 			string `json:"year" form:"year" binding:"required,min=1"`
	Company 		string `json:"company" form:"company" binding:"required,min=1"`
	Position 		string `json:"position" form:"position" binding:"required,min=1"`
	Color 			string `json:"color" form:"color" binding:"required,min=1"`
	Description 	string `json:"description" form:"description" binding:"required,min=1"`
}

type UpdateCareerRequest struct {
	ID 				int64 `json:"id" form:"id"`
	Year 			string `json:"year" form:"year" binding:"required,min=1"`
	Company 		string `json:"company" form:"company" binding:"required,min=1"`
	Position 		string `json:"position" form:"position" binding:"required,min=1"`
	Color 			string `json:"color" form:"color" binding:"required,min=1"`
	Description 	string `json:"description" form:"description" binding:"required,min=1"`
}