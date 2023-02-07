package payload

type RegisterUserInput struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdateUserInput struct {
	ID           int
	FullName     string  `json:"full_name" binding:"required"`
	LegalName    string  `json:"legal_name" binding:"required"`
	NIK          int     `json:"nik" binding:"required,numeric"`
	BirthPlace   string  `json:"birth_place" binding:"required"`
	BirthDate    string  `json:"birth_date" binding:"required" time_format:"2006-01-02"`
	Salary       float64 `json:"salary" binding:"required,numeric"`
	PhotosCard   string  `json:"photos_card" binding:"required"`
	PhotosSelfie string  `json:"photos_selfie" binding:"required"`
}
