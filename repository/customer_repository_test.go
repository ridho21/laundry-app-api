package repository

import (
	"database/sql"
	"errors"
	"go-enigma-laundry/model"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
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

type CustomerRepoTestSuite struct {
	suite.Suite
	repo CustomerRepository
	mockDb *sql.DB
	mockSql sqlmock.Sqlmock
}

func (suite *CustomerRepoTestSuite) SetupTest(){
	mockDb,mockSql, _ := sqlmock.New()
	customerRepo := NewCustomerRepository(mockDb);
	suite.mockDb = mockDb
	suite.mockSql = mockSql
	suite.repo = customerRepo
}

func(suite *CustomerRepoTestSuite) TestCustomerCreate_Success(){
	suite.mockSql.ExpectExec("INSERT INTO mst_customers").WithArgs(
		"C001",
		"Robby",
		"087268937892",
		"Cianjur",
		"robby@gmail.com",
	).WillReturnResult(sqlmock.NewResult(1,1))
	err := suite.repo.InsertCustomer(dummyCustomers[0])
	suite.Assert().NoError(err)
}

func(suite *CustomerRepoTestSuite) TestCustomerCreate_Fail(){
	expecErr := errors.New("Failed")
	suite.mockSql.ExpectExec("INSERT INTO mst_customers").WithArgs(
		"C001",
		"Robby",
		"087268937892",
		"Cianjur",
		"robby@gmail.com",
	).WillReturnError(expecErr)
	err := suite.repo.InsertCustomer(dummyCustomers[0])
	suite.Assert().Equal(expecErr,err)
}

func(suite *CustomerRepoTestSuite) TestCustomerGetList_Success(){
	// Set Column
	rows := suite.mockSql.NewRows([]string{"id","name","address","phone_number","email"})
	// Isi Column
	for _, customer := range dummyCustomers {
		rows.AddRow(
			customer.Id,
			customer.Name,
			customer.Address,
			customer.PhoneNumber,
			customer.Email,
		)
	}
	suite.mockSql.ExpectQuery("SELECT \\* FROM mst_customers").WillReturnRows(rows)
	actualCustomers,err := suite.repo.GetListCustomer()	
	suite.Assert().NoError(err)
	suite.Assert().Equal(3,len(actualCustomers))
	suite.Assert().Equal(dummyCustomers[0],actualCustomers[0])
}

func(suite *CustomerRepoTestSuite) TestCustomerUpdate_Success(){
	suite.mockSql.ExpectExec("UPDATE mst_customers SET name=\\$1, address=\\$2, phone_number=\\$3, email=\\$4 WHERE id=\\$5").WithArgs(
		"Robby",
		"Cianjur",
		"087268937892",
		"robby@gmail.com",
		"C001",
	).WillReturnResult(sqlmock.NewResult(1,1))
	err := suite.repo.UpdateCustomer(dummyCustomers[0])
	suite.Assert().NoError(err)
}

func TestCustomerRepoTestSuite(t *testing.T){
	suite.Run(t,new(CustomerRepoTestSuite))
}