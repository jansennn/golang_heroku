package repo

import (
	"github.com/jansennn/golang_heroku/entity"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	All() ([]entity.Project, error)
	InsertProject(project entity.Project) (entity.Project, error)
	UpdateProject(project entity.Project) (entity.Project, error)
	FindOneProjectById(ID string) (entity.Project, error)
}

type projectRepo struct{
	connection *gorm.DB
}

func NewProjectRepo(connection *gorm.DB) ProjectRepository{
	return &projectRepo{
		connection: connection,
	}
}

func (c *projectRepo) All() ([]entity.Project, error) {
	project := []entity.Project{}
	c.connection.Find(&project)
	return project, nil
}

func (c *projectRepo) InsertProject(project entity.Project) (entity.Project, error){
	c.connection.Save(&project)
	return project, nil
}

func (c *projectRepo) UpdateProject(project entity.Project) (entity.Project, error){
	c.connection.Save(&project)
	return project, nil
}

func (c *projectRepo) FindOneProjectById(projectID string) (entity.Project, error){
	var project entity.Project
	res := c.connection.Find(&project, projectID)
	if res.Error != nil{
		return project, res.Error
	}
	return project, nil
}
