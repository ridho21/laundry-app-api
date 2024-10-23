package repository

import (
	"database/sql"
	"go-enigma-laundry/model"
	"go-enigma-laundry/model/dto"
	"go-enigma-laundry/utils"
)

type CustomerRepository interface {
	GetListCustomer() ([]model.Customer, error)
	InsertCustomer(newCustomer model.Customer) error
	UpdateCustomer(newCustomer model.Customer) error
	GetCustomerById(id string) (model.Customer, error)
	DeleteCustomerById(id string) error
	GetListPaging(page int, size int) ([]model.Customer, dto.Paging, error)
}

type customerRepository struct {
	db *sql.DB
}

// page (halaman)
// limit/size (banyak data perhalaman)

// 10
// 2
// 2
// 1      2      3     4     5
// 2      2      2     2     2

func (c *customerRepository) GetListPaging(page int, size int) ([]model.Customer, dto.Paging, error) {
	skip := (page - 1) * size
	rows, err := c.db.Query(utils.SELECT_CUSTOMER_PAGING, size, skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err = rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.Address,
			&customer.PhoneNumber,
			&customer.Email,
		)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		customers = append(customers, customer)
	}

	var totalRows int
	err = c.db.QueryRow(utils.SELECT_COUNT_CUSTOMER).Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	resultPagingDto := utils.Paginate(page, size, totalRows)
	return customers, resultPagingDto, nil
}

func (c *customerRepository) UpdateCustomer(newCustomer model.Customer) error {
	_, err := c.db.Exec(utils.UPDATE_CUSTOMER,
		newCustomer.Name,
		newCustomer.Address,
		newCustomer.PhoneNumber,
		newCustomer.Email,
		newCustomer.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) GetCustomerById(id string) (model.Customer, error) {
	var customer model.Customer
	err := c.db.QueryRow(utils.SELECT_CUSTOMER_ID, id).Scan(
		&customer.Id,
		&customer.Name,
		&customer.Address,
		&customer.PhoneNumber,
		&customer.Email,
	)

	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

func (c *customerRepository) DeleteCustomerById(id string) error {
	_, err := c.db.Exec(utils.DELETE_CUSTOMER, id)
	if err != nil {
		return err
	}
	return nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (c *customerRepository) GetListCustomer() ([]model.Customer, error) {
	rows, err := c.db.Query(utils.SELECT_CUSTOMER)
	if err != nil {
		return nil, err
	}
	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err = rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.Address,
			&customer.PhoneNumber,
			&customer.Email,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c *customerRepository) InsertCustomer(newCustomer model.Customer) error {
	_, err := c.db.Exec(utils.INSERT_CUSTOMER,
		newCustomer.Id, // Placeholder parameter
		newCustomer.Name,
		newCustomer.PhoneNumber,
		newCustomer.Address,
		newCustomer.Email,
	)
	if err != nil {
		return err
	}
	return nil
}

// ada kebutuhan pengambilan data/query -> Query(list) /QueryRow(single Value)
// tidak kebutuhan untuk ngambil data/record dari db -> Exec() : contoh : insert, Update, Delete,dan lain lain.
