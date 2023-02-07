package payload

import "sigmatech-test/model"

type CreateTransactionInput struct {
	ContractNumber string  `json:"contract_number" binding:"required"`
	OTR            float64 `json:"otr" binding:"required,numeric"`
	AdminFee       float64 `json:"admin_fee" binding:"required,numeric"`
	Installment    int     `json:"installment" binding:"required,numeric"`
	InterestAmount float64 `json:"interest_amount" binding:"required,numeric"`
	AssetName      string  `json:"asset_name" binding:"required"`
	User           model.User
}
