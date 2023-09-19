package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type RegisterUser struct {
	Username          string `json:"username" validate:"required"`
	Password          string `json:"password" validate:"required,eqfield=ConfirmedPassword"`
	ConfirmedPassword string `json:"confirmed_password" validate:"required"`
	Phone             string `json:"phone" validate:"required"`
}

// NewRegisterUserValidator NewRegisterUser register validation constructor
func NewRegisterUserValidator(inputs map[string]string) *RegisterUser {
	return &RegisterUser{
		Username:          inputs["username"],
		Password:          inputs["password"],
		ConfirmedPassword: inputs["confirmed_password"],
		Phone:             inputs["phone"],
	}
}

// RegisterUserValidation  validate inputs for registering a User
func (r *RegisterUser) RegisterUserValidation() []string {

	var errorRes []string

	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, validationErr := range validationErrs {
				param := fmt.Sprintf(" ,value  is %s", validationErr.Param())
				if validationErr.Param() == "" {
					param = ""
				}
				errorRes = append(errorRes, fmt.Sprintf("the field %s does not passed %s validation %s", validationErr.Field(), validationErr.Tag(), param))
			}
		}
		return errorRes
	}
	return nil
}
