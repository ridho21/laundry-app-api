package usecase

import (
	"errors"
	"go-enigma-laundry/model"
	"go-enigma-laundry/model/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyCustomers = []model.Customer {
	{
		Id: "C001",
		Name: "Robby",
		PhoneNumber: "087268937892",
		Address: "Cianjur",
		Email: "robby@gmail.com",
	},
	{
		Id: "C002",
		Name: "Sena",
		PhoneNumber: "087268937123",
		Address: "Jakarta",
		Email: "sena@gmail.com",
	},
	{
		Id: "C003",
		Name: "Herman",
		PhoneNumber: "087268936782",
		Address: "Depok",
		Email: "herman@gmail.com",
	},
}

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) GetListCustomer() ([]model.Customer, error) {
	return nil,nil
}

func (r *RepoMock) InsertCustomer(newCustomer model.Customer) error{
	args := r.Called(newCustomer)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (r *RepoMock) UpdateCustomer(newCustomer model.Customer) error{
	return nil
}

func (r *RepoMock) GetCustomerById(id string) (model.Customer, error){
	args := r.Called(id)
	if args.Get(1) != nil {
		return model.Customer{},args.Error(1)
	}
	return args.Get(0).(model.Customer),nil
}

func (r *RepoMock) DeleteCustomerById(id string) error{
	return nil
}

func (r *RepoMock) GetListPaging(page int, size int) ([]model.Customer, dto.Paging, error){
	return nil,dto.Paging{},nil
}

type CustomerUsecaseTestSuite struct {
	suite.Suite
	repoMock *RepoMock
	usecase CustomerUsecase
}

// Positive Test Case
func (suite *CustomerUsecaseTestSuite) TestCustomerRegister_Success(){
	suite.repoMock.On("InsertCustomer",dummyCustomers[0]).Return(nil)
	err := suite.usecase.RegisterCustomer(dummyCustomers[0])
	assert.Nil(suite.T(),err)
}

// Negative Test Case
func (suite *CustomerUsecaseTestSuite) TestCustomerRegister_Fail(){
	suite.repoMock.On("InsertCustomer",dummyCustomers[0]).Return(errors.New("Error"))
	err := suite.usecase.RegisterCustomer(dummyCustomers[0])
	assert.Error(suite.T(),err)
}

func (suite *CustomerUsecaseTestSuite) TestCustomerFindById_Success(){
	customerId := dummyCustomers[0].Id
	suite.repoMock.On("GetCustomerById",customerId).Return(dummyCustomers[0],nil)
	customer,err := suite.usecase.FindCustomerById(customerId)
	assert.Nil(suite.T(),err)
	assert.Equal(suite.T(),customer,dummyCustomers[0])
}

func (suite *CustomerUsecaseTestSuite) TestCustomerFindById_Fail(){
	customerId := dummyCustomers[0].Id
	suite.repoMock.On("GetCustomerById",customerId).Return(model.Customer{},errors.New("Error"))
	customer,err := suite.usecase.FindCustomerById(customerId)
	assert.Error(suite.T(),err)
	assert.Equal(suite.T(),model.Customer{},customer)
}

func (suite *CustomerUsecaseTestSuite) SetupTest(){
	suite.repoMock = new(RepoMock)
	suite.usecase = NewCustomerUsecase(suite.repoMock)
}

func TestCustomerUsecaseTestSuite(t *testing.T){
	suite.Run(t,new(CustomerUsecaseTestSuite))
}

