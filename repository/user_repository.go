package repository

import (
	"database/sql"
	"go-enigma-laundry/model"
	"go-enigma-laundry/utils"
)


type UserRepository interface {
	GetById(id string) (model.User,error)
	Create(payload model.User) (model.User,error)
	GetByUsername(username string) (model.User,error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository)GetById(id string) (model.User,error) {
	var user model.User
	err := u.db.QueryRow(utils.SELECT_USER_ID,id).Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return model.User{},err
	}
	return user,nil
}
func (u *userRepository)Create(payload model.User) (model.User,error){
	var user model.User
	err := u.db.QueryRow(utils.INSERT_USER,
		payload.Id,
		payload.FullName,
		payload.Email,
		payload.Username,
		payload.Password,
		payload.Role,
		).Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return model.User{},err
	}
	return user,nil
}
func (u *userRepository)GetByUsername(username string) (model.User,error){
	var user model.User
	err := u.db.QueryRow(utils.SELECT_USER_USERNAME,username).Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return model.User{},err
	}
	return user,nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db :db,
	}
}