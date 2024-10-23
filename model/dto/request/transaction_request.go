package request

type TransactionRequest struct {
	CustomerId string                      `json:"customerId" binding:"required"`
	TrxDetails []TransactionDetailsRequest `json:"transactionDetails" binding:"required"`
}

type TransactionDetailsRequest struct {
	ServiceId string `json:"serviceId"`
	Qty       int    `json:"qty"`
}
