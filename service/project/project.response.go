package project

import "github.com/jansennn/golang_heroku/entity"

type ProjectResponse struct {
	ID  		int64  		`json:"id"`
	Name 		string 		`json:"name"`
	Image 		string		`json:"image"`
	Description string 		`json:"description"`
	LinkCode 	string 		`json:"link_code"`
}

func NewProjectResponse(project entity.Project) ProjectResponse {
	return ProjectResponse{
		ID: project.ID,
		Name: project.Name,
		Image: project.Image,
		Description: project.Description,
		LinkCode: project.LinkCode,
	}
}

func NewProjectArrayResponse(projects []entity.Project) []ProjectResponse {
	projectRes := []ProjectResponse{}
	for _, v := range projects {
		p := ProjectResponse{
			ID: v.ID,
			Name: v.Name,
			Image: v.Image,
			Description: v.Description,
			LinkCode: v.LinkCode,
		}
		projectRes = append(projectRes, p)
	}

	return projectRes
}