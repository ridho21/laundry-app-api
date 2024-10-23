package model

type User struct {
	Id			string	`json:"id"`
	FullName 	string	`json:"fullName" binding:"required"`
	Email 		string	`json:"email" binding:"required"`
	Username 	string	`json:"username" binding:"required"`
	Password 	string	`json:"password" binding:"required"`
	Role 		string	`json:"role" binding:"required"`
}