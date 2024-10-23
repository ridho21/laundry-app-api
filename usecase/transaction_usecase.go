package usecase

import (
	"go-enigma-laundry/model"
	"go-enigma-laundry/model/dto/request"
	"go-enigma-laundry/model/dto/response"
	"go-enigma-laundry/repository"
	"time"

	"github.com/google/uuid"
)

type TransactionUsecase interface {
	CreateTransaction(newTransaction request.TransactionRequest) (response.TransactionResponse,error)
	FindAllTrx()([]model.Transaction,error)
}

type transactionUsecase struct {
	repo repository.TransactionRepository
	customerUsecase CustomerUsecase
	serviceUsecase ServiceUsecase
}

func (t *transactionUsecase) FindAllTrx()([]model.Transaction,error) {
	return t.repo.GetList()
}

func (t *transactionUsecase) CreateTransaction(newTransaction request.TransactionRequest) (response.TransactionResponse,error) {
	var transaction model.Transaction
	var transactionDetails []model.TransactionDetails

	// Response
	var transactionResponse response.TransactionResponse
	
	// Get Customer
	findCustomer, err := t.customerUsecase.FindCustomerById(newTransaction.CustomerId)
	if err != nil {
		return response.TransactionResponse{},err
	}
	transaction.Id = uuid.NewString()
	transaction.Customer = findCustomer
	transaction.PickupDate = time.Now().AddDate(0,0,2) // Set 2 hari setelah transaction dibuat
	transaction.Status = model.StatusUnPaid

	// set response
	transactionResponse.Id = transaction.Id
	transactionResponse.Customer = transaction.Customer
	transactionResponse.PickupDate = transaction.PickupDate
	transactionResponse.Status = transaction.Status

	var totalPrice float64

	// set service object setiap trxDetail
	for _, trxDetail := range newTransaction.TrxDetails {
		var transactionDetail model.TransactionDetails
		findService, err := t.serviceUsecase.FindServiceById(trxDetail.ServiceId)
		if err != nil {
			return response.TransactionResponse{},err
		}
		transactionDetail.Id = uuid.NewString()
		transactionDetail.TransactionId = transaction.Id
		transactionDetail.Service = findService
		transactionDetail.Qty = trxDetail.Qty
		totalPrice += transactionDetail.Service.Price * float64(transactionDetail.Qty)
		transactionDetails = append(transactionDetails,transactionDetail)
	}
	transaction.TrxDetails = transactionDetails
	// set response
	transactionResponse.TrxDetails = transaction.TrxDetails
	trxCreated, errTrx := t.repo.Create(transaction)
	if errTrx != nil {
		return response.TransactionResponse{},errTrx
	}
	transaction.Date = trxCreated.Date
	transaction.IsPickedUp = trxCreated.IsPickedUp

	// set response
	transactionResponse.Date = transaction.Date
	transactionResponse.IsPickedUp = transaction.IsPickedUp
	transactionResponse.TotalPrice = totalPrice
	return transactionResponse,nil
}

func NewTransactionUsecase(
	repo repository.TransactionRepository,
	customerUc CustomerUsecase,
	serviceUc ServiceUsecase,
	) TransactionUsecase {	
	return &transactionUsecase {
		repo: repo,
		customerUsecase: customerUc,
		serviceUsecase: serviceUc,
	}
}