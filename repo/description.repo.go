package repo

import (
	"github.com/jansennn/golang_heroku/entity"
	"gorm.io/gorm"
)

type DescriptionRepository interface {
	FindOneDescriptionById(ID string) (entity.Description, error)
	InsertDescription(description entity.Description) (entity.Description, error)
	UpdateDescription(description entity.Description) (entity.Description, error)
}

type descriptionRepo struct {
	connection *gorm.DB
}

func NewDescriptionRepo(connection *gorm.DB) DescriptionRepository {
	return &descriptionRepo{
		connection: connection,
	}
}

func (c *descriptionRepo) FindOneDescriptionById(descID string) (entity.Description, error) {
	var description entity.Description
	res := c.connection.Find(&description, descID)
	if res.Error != nil {
		return description, res.Error
	}
	return description, nil
}

func (c *descriptionRepo) InsertDescription(description entity.Description) (entity.Description, error) {
	c.connection.Save(&description)
	c.connection.Find(&description)
	return description, nil
}

func (c *descriptionRepo) UpdateDescription(description entity.Description) (entity.Description, error) {
	c.connection.Save(&description)
	c.connection.Find(&description)
	return description, nil
}