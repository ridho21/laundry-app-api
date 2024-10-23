package repository

import (
	"database/sql"
	"go-enigma-laundry/model"
	"go-enigma-laundry/utils"
)

type TransactionRepository interface {
	GetList() ([]model.Transaction, error)
	Create(newTrx model.Transaction) (model.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func (t *transactionRepository) GetList() ([]model.Transaction, error) {
	var transactions []model.Transaction

	rows, err := t.db.Query(`SELECT * FROM transactions`)
	if err != nil {
		return nil, nil
	}
	for rows.Next() {
		var transaction model.Transaction
		var transactionDetails []model.TransactionDetails
		var customerDetails model.Customer
		// var serviceDetails model.Service
		err = rows.Scan(
			&transaction.Id,
			&transaction.Customer.Id,
			&transaction.Date,
			&transaction.PickupDate,
			&transaction.Status,
			&transaction.IsPickedUp,
		)
		if err != nil {
			return nil, err
		}

		rowsCustomerDetails, err := t.db.Query(`SELECT * FROM mst_customers WHERE id = $1`, transaction.Customer.Id)

		if err != nil {
			return nil, err
		}

		for rowsCustomerDetails.Next() {
			var customerDetail model.Customer

			err = rowsCustomerDetails.Scan(
				&customerDetail.Id,
				&customerDetail.Name,
				&customerDetail.Address,
				&customerDetail.PhoneNumber,
				&customerDetail.Email,
			)

			if err != nil {
				return nil, err
			}
			customerDetails = customerDetail
		}

		// Ambil data trx detail berdasar transactino id
		rowsTrxDetails, err := t.db.Query(`SELECT * FROM transaction_details WHERE transaction_id = $1`, transaction.Id)

		if err != nil {
			return nil, err
		}
		for rowsTrxDetails.Next() {
			var transactionDetail model.TransactionDetails
			var serviceDetail model.Service
			err = rowsTrxDetails.Scan(
				&transactionDetail.Id,
				&transactionDetail.TransactionId,
				&transactionDetail.Service.Id,
				&transactionDetail.Qty,
			)
			if err != nil {
				return nil, err
			}

			rowService, err := t.db.Query(`SELECT * FROM mst_services WHERE id = $1`, transactionDetail.Service.Id)

			if err != nil {
				return nil, err
			}

			for rowService.Next() {
				err = rowService.Scan(
					&serviceDetail.Id,
					&serviceDetail.Description,
					&serviceDetail.Price,
				)

				if err != nil {
					return nil, err
				}
			}

			transactionDetail.Service = serviceDetail
			transactionDetails = append(transactionDetails, transactionDetail)
		}
		transaction.TrxDetails = transactionDetails
		transaction.Customer = customerDetails
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (t *transactionRepository) Create(newTrx model.Transaction) (model.Transaction, error) {
	// Begin
	// DML
	// Commit/Rollback
	// akan melakukan insert data kedalam 2 table

	//Siapkan model untuk di return nantinya
	var transaction model.Transaction

	tx, err := t.db.Begin()
	if err != nil {
		return model.Transaction{}, err
	}
	// Insert Data Transaction
	// Prepared Statement => untuk menghindari Sql Injection
	stmt, err := tx.Prepare(utils.INSERT_TRANSACTION)
	if err != nil {
		tx.Rollback()
		return model.Transaction{}, err
	}
	err = stmt.QueryRow(
		newTrx.Id,
		newTrx.Customer.Id,
		newTrx.PickupDate,
		newTrx.Status,
	).Scan(
		&transaction.Id,
		&transaction.Customer.Id,
		&transaction.Date,
		&transaction.PickupDate,
		&transaction.Status,
	)
	if err != nil {
		tx.Rollback()
		return model.Transaction{}, err
	}

	// Insert Data Transaction Details
	var transactionDetails []model.TransactionDetails
	for _, TrxDetail := range newTrx.TrxDetails {
		var transactionDetail model.TransactionDetails
		stmtDetail, err := tx.Prepare(utils.INSERT_TRX_DETAILS)
		if err != nil {
			tx.Rollback()
			return model.Transaction{}, err
		}
		err = stmtDetail.QueryRow(
			TrxDetail.Id,
			newTrx.Id,
			TrxDetail.Service.Id,
			TrxDetail.Qty,
		).Scan(
			&transactionDetail.Id,
			&transactionDetail.TransactionId,
			&transactionDetail.Service.Id,
			&transactionDetail.Qty,
		)
		if err != nil {
			tx.Rollback()
			return model.Transaction{}, err
		}
		transactionDetails = append(transactionDetails, transactionDetail)

	}
	transaction.TrxDetails = transactionDetails
	err = tx.Commit()
	if err != nil {
		return model.Transaction{}, err
	}
	return newTrx, nil
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
