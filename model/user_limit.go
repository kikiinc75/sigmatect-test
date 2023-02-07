package model

import "time"

type UserLimit struct {
	ID        int
	UserID    int
	Tenor     int
	Amount    float64
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
