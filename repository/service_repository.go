package repository

import (
	"database/sql"
	"go-enigma-laundry/model"
	"go-enigma-laundry/utils"
)

type ServiceRepository interface{
	GetListService() ([]model.Service,error)
	InsertService(newService model.Service) error	
	UpdateService(newService model.Service) error
	GetServiceById(id string) (model.Service,error)
	DeleteServiceById(id string) error
}

type serviceRepository struct {
	db *sql.DB
}


func NewServiceRepository(db *sql.DB) ServiceRepository {
	return &serviceRepository{
		db: db ,
	}
}

func (c *serviceRepository) UpdateService(newService model.Service) error {
	_,err := c.db.Exec(utils.UPDATE_SERVICE,
			newService.Description,
			newService.Price,
			newService.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *serviceRepository) GetServiceById(id string) (model.Service,error) {
	var Service model.Service
	err := c.db.QueryRow(utils.SELECT_SERVICE_ID,id).Scan(
		&Service.Id,
		&Service.Description,
		&Service.Price,
	)

	if err != nil {
		return model.Service{},err
	}
	return Service,nil
}

func (c *serviceRepository) DeleteServiceById(id string) error {
	_,err := c.db.Exec(utils.DELETE_SERVICE,id)
	if err != nil {
		return err
	}
	return nil
}


func (c *serviceRepository) GetListService() ([]model.Service,error) {
	rows, err := c.db.Query(utils.SELECT_SERVICE)
	if err != nil {
		return nil,err
	}
	var Services []model.Service
	for rows.Next() {
		var Service model.Service
		err = rows.Scan(
			&Service.Id,
			&Service.Description,
			&Service.Price,
		)
		if err != nil {
			return nil,err
		}
		Services = append(Services, Service)
	}
	return Services,nil
}

func (c *serviceRepository) InsertService(newService model.Service) error {
	_, err := c.db.Exec(utils.INSERT_SERVICE,
			newService.Id, // Placeholder parameter
			newService.Description,
			newService.Price,
	)
	if err != nil {
		return err
	}
	return nil
}

// ada kebutuhan pengambilan data/query -> Query(list) /QueryRow(single Value)
// tidak kebutuhan untuk ngambil data/record dari db -> Exec() : contoh : insert, Update, Delete,dan lain lain.