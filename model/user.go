package model

import "time"

type User struct {
	ID        int
	FullName  string
	Email     string
	Password  string
	NIK       string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
