package model

import "time"

type User struct {
	ID           int
	FullName     string
	LegalName    string
	Email        string
	Password     string
	NIK          int
	BirthPlace   string
	BirthDate    *time.Time
	Salary       float64
	PhotosCard   string
	PhotosSelfie string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
