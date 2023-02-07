package user

import "sigmatech-test/model"

type UserFormatter struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	NIK      string `json:"nik"`
	Token    string `json:"token"`
}

func FormatUser(user *model.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		NIK:      user.NIK,
		Token:    token,
	}

	return formatter
}
