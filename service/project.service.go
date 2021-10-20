package service

import (
	"fmt"
	"github.com/jansennn/golang_heroku/dto"
	"github.com/jansennn/golang_heroku/entity"
	"github.com/jansennn/golang_heroku/repo"
	"github.com/mashingan/smapping"
	"log"

	_project "github.com/jansennn/golang_heroku/service/project"
	)

type ProjectService interface {
	All() (*[]_project.ProjectResponse, error)
	FindOneProjectById(projectID string) (*_project.ProjectResponse, error)
	CreateProject(projectRequest dto.CreateProjectRequest) (*_project.ProjectResponse, error)
	UpdateProject(updateProjectRequest dto.UpdateProjectRequest) (*_project.ProjectResponse, error)
}

type projectService struct {
	projectRepo repo.ProjectRepository
}

func NewProjectService(projectRepo repo.ProjectRepository) ProjectService{
	return &projectService{
		projectRepo: projectRepo,
	}
}

func (c *projectService) All() (*[]_project.ProjectResponse, error) {
	projects, err := c.projectRepo.All()
	if err != nil {
		return nil, err
	}

	projs := _project.NewProjectArrayResponse(projects)
	return &projs, nil
}

func (c *projectService) FindOneProjectById(projectID string) (*_project.ProjectResponse, error) {
	project, err := c.projectRepo.FindOneProjectById(projectID)

	if err != nil{
		return nil, err
	}

	res := _project.NewProjectResponse(project)
	return &res, nil
}



func (c *projectService) CreateProject(projectRequest dto.CreateProjectRequest) (*_project.ProjectResponse, error) {
	project := entity.Project{}
	err := smapping.FillStruct(&project, smapping.MapFields(&projectRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	p, err := c.projectRepo.InsertProject(project)
	if err != nil {
		return nil, err
	}

	res := _project.NewProjectResponse(p)
	return &res, nil
}

func (c *projectService) UpdateProject(updateProjectRequest dto.UpdateProjectRequest) (*_project.ProjectResponse, error){
	project, err := c.projectRepo.FindOneProjectById(fmt.Sprintf("%d", updateProjectRequest.ID))

	if err != nil{
		return nil, err
	}

	project = entity.Project{}

	err = smapping.FillStruct(&project, smapping.MapFields(&updateProjectRequest))

	if err != nil{
		return nil, err
	}

	project, err = c.projectRepo.UpdateProject(project)

	if err != nil {
		return nil, err
	}

	res := _project.NewProjectResponse(project)
	return &res, err
}
