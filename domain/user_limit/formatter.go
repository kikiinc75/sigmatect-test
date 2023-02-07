package user_limit

import (
	"sigmatech-test/model"
)

type UserLimitFormatter struct {
	ID     int     `json:"id"`
	Tenor  int     `json:"tenor"`
	Amount float64 `json:"amount"`
}

func FormatUserLimit(userLimit *model.UserLimit) UserLimitFormatter {
	formatter := UserLimitFormatter{
		ID:     userLimit.ID,
		Tenor:  userLimit.Tenor,
		Amount: userLimit.Amount,
	}

	return formatter
}

func FormatUserLimits(userLimits []model.UserLimit) []UserLimitFormatter {
	userLimitFormatter := []UserLimitFormatter{}
	for _, userLimit := range userLimits {
		formatter := FormatUserLimit(&userLimit)
		userLimitFormatter = append(userLimitFormatter, formatter)
	}

	return userLimitFormatter
}
