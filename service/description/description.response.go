package _description

import (
	"github.com/jansennn/golang_heroku/entity"
)

type DescriptionResponse struct {
	 ID           int64         `json:"id"`
	 Description  string        `json:"description"`
}

func NewDescriptionResponse(description entity.Description) DescriptionResponse {
	return DescriptionResponse{
		ID: description.ID,
		Description: description.Description,
	}
}

func NewDescriptionArrayResponse(description []entity.Description) []DescriptionResponse {
	descriptionRes := []DescriptionResponse{}
	for _, v := range description {
		p := DescriptionResponse{
			ID:          v.ID,
			Description: v.Description,
		}
		descriptionRes = append(descriptionRes, p)
	}
	return descriptionRes
}