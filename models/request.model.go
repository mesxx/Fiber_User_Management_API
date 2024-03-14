package models

type (
	RequestCreateUser struct {
		Name     string `json:"name" validate:"required,min=5,max=20"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=5"`
	}

	RequestLoginUser struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=5"`
	}

	RequestForgotPasswordUser struct {
		Email string `json:"email" validate:"required,email"`
	}

	RequestResetPasswordUser struct {
		Password string `json:"password" validate:"required,min=5"`
	}
)
