package usecase

import (
	"go-enigma-laundry/model"
	"go-enigma-laundry/repository"

	"github.com/google/uuid"
)

type ServiceUsecase interface { // layer untuk komunikasi | jembatan antar layer
	RegisterService(newService model.Service) (model.Service,error)
    FindAllService() ([]model.Service,error)
    UpdateService(newService model.Service) error
    FindServiceById(id string) (model.Service,error)
    DeleteServiceById(id string) error
}

type serviceUsecase struct {
    repo repository.ServiceRepository
}

func (c *serviceUsecase) DeleteServiceById(id string) error {
    return c.repo.DeleteServiceById(id)
}

func (c *serviceUsecase) FindServiceById(id string) (model.Service,error) {
    return c.repo.GetServiceById(id)
}

func (c *serviceUsecase) UpdateService(newService model.Service) error {    
    return c.repo.UpdateService(newService)
}

func (c *serviceUsecase) RegisterService(newService model.Service) (model.Service,error) {
    newService.Id = uuid.NewString();
    err := c.repo.InsertService(newService)
    if err != nil {
        return model.Service{},err
    }
    return newService,nil
}

func (c *serviceUsecase)  FindAllService() ([]model.Service,error){
    return c.repo.GetListService()
}

func NewServiceUsecase(repo repository.ServiceRepository) ServiceUsecase {
    return &serviceUsecase{
        repo: repo,
    }
}