package transaction

import (
	"errors"
	"sigmatech-test/domain/user_limit"
	"sigmatech-test/model"
	"sigmatech-test/payload"
)

type Service interface {
	GetTransactionsByUserID(userID int) ([]model.Transaction, error)
	CreateTransaction(input payload.CreateTransactionInput) (*model.Transaction, error)
}

type service struct {
	repository          Repository
	userLimitRepository user_limit.Repository
}

func NewService(repository Repository, userLimitRepository user_limit.Repository) *service {
	return &service{repository, userLimitRepository}
}

func (s *service) GetTransactionsByUserID(userID int) ([]model.Transaction, error) {
	transactions, err := s.repository.FindByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input payload.CreateTransactionInput) (*model.Transaction, error) {
	userLimitData, err := s.userLimitRepository.FindByUserIDAndTenor(input.User.ID, input.Installment)
	if err != nil {
		return nil, err
	}

	totalAmount := input.OTR + input.AdminFee + (input.InterestAmount * float64(input.Installment))
	println(userLimitData.Amount, totalAmount)
	if userLimitData.Amount < totalAmount {
		return nil, errors.New("Limit exceeded")
	}

	transaction := model.Transaction{}
	transaction.UserID = input.User.ID
	transaction.ContractNumber = input.ContractNumber
	transaction.OTR = input.OTR
	transaction.AdminFee = input.AdminFee
	transaction.Installment = input.Installment
	transaction.InterestAmount = input.InterestAmount
	transaction.AssetName = input.AssetName

	newTransaction, err := s.repository.Save(&transaction)
	if err != nil {
		return newTransaction, err
	}

	userLimitData.Amount = userLimitData.Amount - totalAmount
	_, err = s.userLimitRepository.Update(userLimitData)
	if err != nil {
		return nil, err
	}

	return newTransaction, nil
}
