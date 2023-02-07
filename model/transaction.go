package model

import (
	"time"
)

type Transaction struct {
	ID             int
	UserID         int
	ContractNumber string
	OTR            float64
	AdminFee       float64
	Installment    int
	InterestAmount float64
	AssetName      string
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
