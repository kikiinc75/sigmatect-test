package user

import (
	"errors"
	"sigmatech-test/model"
	"sigmatech-test/payload"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input payload.RegisterUserInput) (*model.User, error)
	Login(input payload.LoginInput) (*model.User, error)
	IsEmailAvailable(input payload.CheckEmailInput) (bool, error)
	GetUserByID(ID int) (*model.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input payload.RegisterUserInput) (*model.User, error) {
	user := model.User{}
	user.FullName = input.FullName
	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return &user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(&user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input payload.LoginInput) (*model.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input payload.CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) GetUserByID(ID int) (*model.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}
