package model

type Service struct {
	Id          string  `json:"id"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

// CRUD
// Desc
// Cuci Kering Lipat
// Curi Kering Setrika
// Satuan
