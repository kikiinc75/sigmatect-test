package transaction

import "sigmatech-test/model"

type Service interface {
	GetTransactions(userID int) ([]model.Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactions(userID int) ([]model.Transaction, error) {
	transactions, err := s.repository.FindByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
