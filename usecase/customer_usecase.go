package usecase

import (
	"go-enigma-laundry/model"
	"go-enigma-laundry/model/dto"
	"go-enigma-laundry/repository"
)

type CustomerUsecase interface { // layer untuk komunikasi | jembatan antar layer
	RegisterCustomer(newCustomer model.Customer) error
    FindAllCustomer() ([]model.Customer,error)
    UpdateCustomer(newCustomer model.Customer) error
    FindCustomerById(id string) (model.Customer,error)
    DeleteCustomerById(id string) error
    FindAllPaging(page int,size int) ([]model.Customer,dto.Paging,error)
}

type customerUsecase struct {
    repo repository.CustomerRepository
}

func (c *customerUsecase) FindAllPaging(page int,size int) ([]model.Customer,dto.Paging,error) {
    return c.repo.GetListPaging(page,size)
}

func (c *customerUsecase) DeleteCustomerById(id string) error {
    return c.repo.DeleteCustomerById(id)
}

func (c *customerUsecase) FindCustomerById(id string) (model.Customer,error) {
    return c.repo.GetCustomerById(id)
}

func (c *customerUsecase) UpdateCustomer(newCustomer model.Customer) error {    
    return c.repo.UpdateCustomer(newCustomer)
}

func (c *customerUsecase) RegisterCustomer(newCustomer model.Customer) error {
    return c.repo.InsertCustomer(newCustomer)
}

func (c *customerUsecase)  FindAllCustomer() ([]model.Customer,error){
    return c.repo.GetListCustomer()
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
    return &customerUsecase{
        repo: repo,
    }
}