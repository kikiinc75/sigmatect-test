package user

import (
	"sigmatech-test/model"
	"time"
)

type UserFormatter struct {
	ID           int        `json:"id"`
	FullName     string     `json:"full_name"`
	Email        string     `json:"email"`
	NIK          int        `json:"nik"`
	Token        string     `json:"token"`
	LegalName    string     `json:"legal_name"`
	BirthPlace   string     `json:"birth_place"`
	BirthDate    *time.Time `json:"birth_date"`
	Salary       float64    `json:"salary"`
	PhotosCard   string     `json:"photos_card"`
	PhotosSelfie string     `json:"photos_selfie"`
}

func FormatUser(user *model.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:           user.ID,
		FullName:     user.FullName,
		Email:        user.Email,
		NIK:          user.NIK,
		LegalName:    user.LegalName,
		BirthPlace:   user.BirthPlace,
		BirthDate:    user.BirthDate,
		Salary:       user.Salary,
		PhotosCard:   user.PhotosCard,
		PhotosSelfie: user.PhotosSelfie,
		Token:        token,
	}

	return formatter
}
