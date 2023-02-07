package user_limit

import (
	"sigmatech-test/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindByUserIDAndTenor(ID, tenor int) (*model.UserLimit, error)
	FindByUserID(ID int) ([]model.UserLimit, error)
	Update(userLimit *model.UserLimit) (*model.UserLimit, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByUserID(userID int) ([]model.UserLimit, error) {
	var userLimits []model.UserLimit

	err := r.db.Where("user_id = ?", userID).Find(&userLimits).Error
	if err != nil {
		return userLimits, err
	}

	return userLimits, nil
}

func (r *repository) FindByUserIDAndTenor(ID, tenor int) (*model.UserLimit, error) {
	var userLimit model.UserLimit

	err := r.db.Where("user_id = ? and tenor = ?", ID, tenor).Find(&userLimit).Error
	if err != nil {
		return &userLimit, err
	}

	return &userLimit, nil
}

func (r *repository) Update(userLimit *model.UserLimit) (*model.UserLimit, error) {
	err := r.db.Save(&userLimit).Error

	if err != nil {
		return userLimit, err
	}

	return userLimit, nil
}
