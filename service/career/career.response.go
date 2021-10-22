package career

import "github.com/jansennn/golang_heroku/entity"

type CareerResponse struct {
	ID  		int64  		`json:"id"`
	Year		string		`json:"year"`
	Company 	string		`json:"company"`
	Position 	string		`json:"position"`
	Color 		string		`json:"color"`
	Description string		`json:"description"`
}

func NewCareerResponse(career entity.Career) CareerResponse{
	return CareerResponse{
		ID:          career.ID,
		Year:        career.Year,
		Company:     career.Company,
		Position:    career.Position,
		Color:       career.Color,
		Description: career.Description,
	}
}

func NewCareerArrayResponse(careers []entity.Career) []CareerResponse{
	careerRes := []CareerResponse{}
	for _, v := range careers {
		p := CareerResponse{
			ID: v.ID,
			Year: v.Year,
			Company: v.Company,
			Position: v.Position,
			Color: v.Color,
			Description: v.Description,
		}
		careerRes = append(careerRes, p)
	}

	return careerRes
}