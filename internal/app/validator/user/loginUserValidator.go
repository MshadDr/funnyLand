package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type LoginUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// NewLoginUserValidator NewLoginUser login validator constructor
func NewLoginUserValidator(inputs map[string]string) *LoginUser {

	return &LoginUser{
		Username: inputs["username"],
		Password: inputs["password"],
	}
}

// LoginUserValidation validate input for login User
func (l *LoginUser) LoginUserValidation() []string {

	var errorRes []string

	validate := validator.New()
	err := validate.Struct(l)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, validationErr := range validationErrs {
				param := fmt.Sprintf(" is %s", validationErr.Param())
				if validationErr.Param() == "" {
					param = ""
				}
				errorRes = append(errorRes, fmt.Sprintf("the field %s %s%s", validationErr.Field(), validationErr.Tag(), param))
			}
		}
		return errorRes
	}
	return nil
}
