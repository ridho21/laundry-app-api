package model

import "time"

type TransactionStatus string

const (
	StatusPaid   TransactionStatus = "Paid"
	StatusUnPaid TransactionStatus = "UnPaid"
)

type Transaction struct {
	Id         string               `json:"id"`
	Customer   Customer             `json:"customer"`
	Date       time.Time            `json:"date"`
	PickupDate time.Time            `json:"pickupDate"`
	Status     TransactionStatus    `json:"status"`
	IsPickedUp bool                 `json:"isPickedUp"`
	TrxDetails []TransactionDetails `json:"transactionDetails"`
}

type TransactionDetails struct {
	Id            string  `json:"id"`
	TransactionId string  `json:"transactionId"`
	Service       Service `json:"service"`
	Qty           int     `json:"qty"`
}
