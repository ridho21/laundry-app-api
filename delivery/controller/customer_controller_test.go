package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-enigma-laundry/model"
	"go-enigma-laundry/model/dto"
	"go-enigma-laundry/model/dto/response"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyCustomers = []model.Customer{
	{
		Name:        "Robby",
		PhoneNumber: "087268937892",
		Address:     "Cianjur",
		Email:       "robby@gmail.com",
	},
	{
		Id:          "C002",
		Name:        "Sena",
		PhoneNumber: "087268937123",
		Address:     "Jakarta",
		Email:       "sena@gmail.com",
	},
	{
		Id:          "C003",
		Name:        "Herman",
		PhoneNumber: "087268936782",
		Address:     "Depok",
		Email:       "herman@gmail.com",
	},
}

type authMock struct {
	mock.Mock
}

func(a *authMock) RequireToken(roles ...string) gin.HandlerFunc{
	return func(c *gin.Context) {}
}

type UsecaseMock struct {
	mock.Mock
}

func (u *UsecaseMock) RegisterCustomer(newCustomer model.Customer) error{
	args := u.Called(newCustomer)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (u *UsecaseMock) FindAllCustomer() ([]model.Customer,error) {
	return nil,nil
}
func (u *UsecaseMock) UpdateCustomer(newCustomer model.Customer) error {
	return nil
}
func (u *UsecaseMock) FindCustomerById(id string) (model.Customer,error) {
	return model.Customer{},nil
}
func (u *UsecaseMock) DeleteCustomerById(id string) error {
	return nil
}
func (u *UsecaseMock) FindAllPaging(page int,size int) ([]model.Customer,dto.Paging,error){
	return nil,dto.Paging{},nil
}

type CustomerControllerTestSuite struct {
	suite.Suite
	routerMock *gin.Engine
	usecaseMock *UsecaseMock
	authMock *authMock
	controller *CustomerController
}

func (suite *CustomerControllerTestSuite) TestRegisterHandler_Success(){
	// Model Customer yang kana dijadikan sebagai request body nantinya
	customer := dummyCustomers[0]
	// lifecycle mock ketika dipanggil, driver, 
	suite.usecaseMock.On("RegisterCustomer",customer).Return(nil)
	// inisiasi recorder kumpulan response context httpnya
	recorder := httptest.NewRecorder()
	// membuat object of json dari model customer
	reqBody,_:= json.Marshal(customer)
	// simulasi hit endpoint http
	request,_ := http.NewRequest(http.MethodPost,"/customer",bytes.NewBuffer(reqBody))
	// membuat gin testnya
	ctx, _ := gin.CreateTestContext(recorder)
	// set ctx 
	ctx.Request = request

	// eksekusi api nya
	suite.controller.registerHandler(ctx)

	var response response.SingleResponse
	json.Unmarshal([]byte(recorder.Body.Bytes()),&response)

	// proses assertion
	// ekspetasi http status codenya 200
	assert.Equal(suite.T(),http.StatusOK,recorder.Code)		
	assert.NotEmpty(suite.T(),response.Data)	
}

func (suite *CustomerControllerTestSuite) TestRegisterHandler_FailUsecase(){
	// Model Customer yang kana dijadikan sebagai request body nantinya
	customer := dummyCustomers[0]
	// lifecycle mock ketika dipanggil, driver, 
	suite.usecaseMock.On("RegisterCustomer",customer).Return(errors.New("Error"))
	// inisiasi recorder kumpulan response context httpnya
	recorder := httptest.NewRecorder()
	// membuat object of json dari model customer
	reqBody,_:= json.Marshal(customer)
	// simulasi hit endpoint http
	request,_ := http.NewRequest(http.MethodPost,"/customer",bytes.NewBuffer(reqBody))
	// membuat gin testnya
	ctx, _ := gin.CreateTestContext(recorder)
	// set ctx 
	ctx.Request = request
	// eksekusi api nya
	suite.controller.registerHandler(ctx)
	// proses assertion
	// ekspetasi http status codenya 200
	assert.Equal(suite.T(),http.StatusBadRequest,recorder.Code)
}

func (suite *CustomerControllerTestSuite) TestRegisterHandler_FailStructBinding(){
	// inisiasi recorder kumpulan response context httpnya
	recorder := httptest.NewRecorder()
	reqBody,_:= json.Marshal(dummyCustomers[0].Name)
	// simulasi hit endpoint http
	request,_ := http.NewRequest(http.MethodPost,"/customer",bytes.NewBuffer(reqBody))
	// membuat gin testnya
	ctx, _ := gin.CreateTestContext(recorder)
	// set ctx 
	ctx.Request = request
	// instance object controller
	// eksekusi api nya
	suite.controller.registerHandler(ctx)
	// proses assertion
	// ekspetasi http status codenya 400
	assert.Equal(suite.T(),http.StatusBadRequest,recorder.Code)
}

func (suite *CustomerControllerTestSuite) SetupTest(){
	suite.usecaseMock = new(UsecaseMock)
	suite.routerMock = gin.Default()
	suite.controller = NewCustomerController(suite.usecaseMock,suite.routerMock,suite.authMock)
}

func TestCustomerUsecaseTestSuite(t *testing.T){
	suite.Run(t,new(CustomerControllerTestSuite))
}

