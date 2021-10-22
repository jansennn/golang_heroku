package service

import (
	"fmt"
	"github.com/jansennn/golang_heroku/dto"
	"github.com/jansennn/golang_heroku/entity"
	"github.com/jansennn/golang_heroku/repo"
	_career "github.com/jansennn/golang_heroku/service/career"
	"github.com/mashingan/smapping"
	"log"
)

type CareerService interface {
	All() (*[]_career.CareerResponse, error)
	FindOneCareerById(careerID string) (*_career.CareerResponse, error)
	CreateCareer(careerRequest dto.CreateCareerRequest) (*_career.CareerResponse, error)
	UpdateCareer(updateCareerRequest dto.UpdateCareerRequest) (*_career.CareerResponse, error)
}

type careerService struct {
	careerRepo repo.CareerRepository
}

func NewCareerService(careerRepo repo.CareerRepository) CareerService{
	return &careerService{
		careerRepo: careerRepo,
	}
}

func (c *careerService) All() (*[]_career.CareerResponse, error){
	careers, err := c.careerRepo.All()
	if err != nil{
		return nil, err
	}

	cars := _career.NewCareerArrayResponse(careers)
	return &cars, nil
}

func (c *careerService) FindOneCareerById(careerID string) (*_career.CareerResponse, error){
	career, err := c.careerRepo.FindOneCareerById(careerID)

	if err != nil {
		return nil, err
	}

	res := _career.NewCareerResponse(career)
	return &res, nil
}

func (c *careerService) CreateCareer(careerRequest dto.CreateCareerRequest) (*_career.CareerResponse, error){
	career := entity.Career{}
	err := smapping.FillStruct(&career, smapping.MapFields(&careerRequest))

	if err != nil{
		log.Fatalf("Failed Map %v", err)
		return nil, err
	}

	p, err := c.careerRepo.InsertCareer(career)
	if err != nil{
		return nil, err
	}

	res:= _career.NewCareerResponse(p)
	return &res, nil

}

func (c *careerService) UpdateCareer(updateCareerRequest dto.UpdateCareerRequest) (*_career.CareerResponse, error){
	career, err := c.careerRepo.FindOneCareerById(fmt.Sprintf("%d", updateCareerRequest.ID))

	if err != nil {
		return nil, err
	}

	career = entity.Career{}

	err = smapping.FillStruct(&career, smapping.MapFields(&updateCareerRequest))

	if err != nil{
		return nil, err
	}

	career, err = c.careerRepo.UpdateCareer(career)

	if err != nil{
		return nil, err
	}
	res := _career.NewCareerResponse(career)
	return &res, nil

}