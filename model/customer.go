package model

// type struct untuk di mapping sebagai object json ke client
type Customer struct {
	Id      	string		`json:"id"`
	Name    	string		`json:"name" binding:"required"`
	Address 	string		`json:"address" binding:"required"`
	PhoneNumber string		`json:"phoneNumber" binding:"required"`
	Email 		string		`json:"email" binding:"required"`
}

// Model/entity/domain -> repository -> usecase/service -> controller