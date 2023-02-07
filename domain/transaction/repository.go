package transaction

import (
	"sigmatech-test/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindByUserID(userID int) ([]model.Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByUserID(userID int) ([]model.Transaction, error) {
	var transactions []model.Transaction

	err := r.db.Where("user_id = ?", userID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
