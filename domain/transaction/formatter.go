package transaction

import (
	"sigmatech-test/model"
	"time"
)

type TransactionFormatter struct {
	ID             int        `json:"id"`
	ContractNumber string     `json:"contract_number"`
	OTR            float64    `json:"otr"`
	AdminFee       float64    `json:"admin_fee"`
	Installment    int        `json:"installment"`
	InterestAmount float64    `json:"interest_amount"`
	AssetName      string     `json:"asset_name"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

func FormatTransaction(transaction *model.Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		ID:             transaction.ID,
		ContractNumber: transaction.ContractNumber,
		OTR:            transaction.OTR,
		AdminFee:       transaction.AdminFee,
		Installment:    transaction.Installment,
		InterestAmount: transaction.InterestAmount,
		AssetName:      transaction.AssetName,
		CreatedAt:      transaction.CreatedAt,
		UpdatedAt:      transaction.UpdatedAt,
	}

	return formatter
}

func FormatTransactions(transactions []model.Transaction) []TransactionFormatter {
	transactionFormatter := []TransactionFormatter{}
	for _, transaction := range transactions {
		formatter := FormatTransaction(&transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}
