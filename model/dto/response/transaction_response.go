package response

import (
	"go-enigma-laundry/model"
	"time"
)

type TransactionResponse struct {
	Id         string                     `json:"id"`
	Customer   model.Customer             `json:"customer"`
	Date       time.Time                  `json:"date"`
	PickupDate time.Time                  `json:"pickupDate"`
	Status     model.TransactionStatus    `json:"status"`
	IsPickedUp bool                       `json:"isPickedUp"`
	TrxDetails []model.TransactionDetails `json:"transactionDetails"`
	TotalPrice float64                    `json:"totalPrice"`
}
