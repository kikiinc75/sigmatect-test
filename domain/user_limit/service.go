package user_limit

import "sigmatech-test/model"

type Service interface {
	GetUserLimitsByUserID(userID int) ([]model.UserLimit, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserLimitsByUserID(userID int) ([]model.UserLimit, error) {
	userLimits, err := s.repository.FindByUserID(userID)
	if err != nil {
		return userLimits, err
	}

	return userLimits, nil
}
