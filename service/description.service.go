package service

import (
	"fmt"
	"github.com/jansennn/golang_heroku/dto"
	"github.com/jansennn/golang_heroku/entity"
	"github.com/jansennn/golang_heroku/repo"
	_description "github.com/jansennn/golang_heroku/service/description"
	"github.com/mashingan/smapping"
	"log"
)

type DescriptionService interface {
	FindOneDescriptionById(descID string) (*_description.DescriptionResponse, error)
	CreateDescription(descriptionRequest dto.CreateDescriptionRequest) (*_description.DescriptionResponse, error)
	UpdateDescription(updateDescriptionRequest dto.UpdateDescriptionRequest) (*_description.DescriptionResponse, error)
}

type descriptionService struct {
	descriptionRepo repo.DescriptionRepository
}

func NewDescriptionService(descriptionRepo repo.DescriptionRepository) DescriptionService {
	return &descriptionService{
		descriptionRepo: descriptionRepo,
	}
}

func (c *descriptionService) FindOneDescriptionById(descID string) (*_description.DescriptionResponse, error){
	description, err := c.descriptionRepo.FindOneDescriptionById(descID)

	if err != nil {
		return nil, err
	}

	res := _description.NewDescriptionResponse(description)
	return &res, nil
}

func (c *descriptionService) CreateDescription(descriptionRequest dto.CreateDescriptionRequest) (*_description.DescriptionResponse, error) {
	description := entity.Description{}
	err := smapping.FillStruct(&description, smapping.MapFields(&descriptionRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	d, err := c.descriptionRepo.InsertDescription(description)
	if err != nil {
		return nil, err
	}

	res := _description.NewDescriptionResponse(d)
	return &res, nil
}

func (c *descriptionService) UpdateDescription(updateDescriptionRequest dto.UpdateDescriptionRequest) (*_description.DescriptionResponse, error) {
	description, err := c.descriptionRepo.FindOneDescriptionById(fmt.Sprintf("%d", updateDescriptionRequest.ID))

	if err != nil {
		return nil, err
	}

	description = entity.Description{}
	err = smapping.FillStruct(&description, smapping.MapFields(&updateDescriptionRequest))

	if err != nil {
		return nil, err
	}

	description, err = c.descriptionRepo.UpdateDescription(description)

	if err != nil {
		return nil, err
	}

	res := _description.NewDescriptionResponse(description)
	return &res, err

}
