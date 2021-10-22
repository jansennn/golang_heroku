package repo

import (
	"github.com/jansennn/golang_heroku/entity"
	"gorm.io/gorm"
)

type CareerRepository interface {
	All() ([]entity.Career, error)
	InsertCareer(career entity.Career) (entity.Career, error)
	UpdateCareer(career entity.Career) (entity.Career, error)
	FindOneCareerById(careerID string) (entity.Career, error)
}

type careerRepo struct {
	connection *gorm.DB
}

func NewCareerRepo(connection *gorm.DB) CareerRepository{
	return &careerRepo{
		connection: connection,
	}
}

func (c *careerRepo) All() ([]entity.Career, error){
	career := []entity.Career{}
	res := c.connection.Find(&career)
	if res.Error != nil {
		return career, res.Error
	}

	return career, nil
}

func (c *careerRepo) InsertCareer(career entity.Career) (entity.Career, error){
	res := c.connection.Save(&career)
	if res.Error != nil {
		return career, res.Error
	}
	return career, nil
}

func (c *careerRepo) UpdateCareer(career entity.Career) (entity.Career, error){
	res := c.connection.Save(&career)
	if res.Error != nil {
		return career, res.Error
	}
	return career, nil
}

func (c *careerRepo) FindOneCareerById(careerId string) (entity.Career, error){
	var career entity.Career
	res := c.connection.Find(&career, careerId)
	if res.Error != nil {
		return career, res.Error
	}
	return career, nil
}