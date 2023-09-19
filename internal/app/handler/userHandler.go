package handler

import (
	"encoding/json"
	"gitlab.com/M.darvish/funtory/internal/app/service/accountService"
	"gitlab.com/M.darvish/funtory/internal/app/validator/user"
	"gitlab.com/M.darvish/funtory/internal/model"
	"gitlab.com/M.darvish/funtory/internal/model/repository"
	"gitlab.com/M.darvish/funtory/internal/util/response"
	"gitlab.com/M.darvish/funtory/internal/util/security"
	"gitlab.com/M.darvish/funtory/pkg/jwt"
	"net/http"
)

type UserHandler struct {
	ur repository.IUser
}

func NewUserHandler(userRepo repository.IUser) *UserHandler {
	return &UserHandler{
		ur: userRepo,
	}
}

func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	message := "request was processed successfully"

	var inputs map[string]string
	err := json.NewDecoder(r.Body).Decode(&inputs)
	if err != nil {
		message = "Failed to parse request body"
		_ = response.NewResponse(err, message, 422).Failed(w)
		return
	}

	inputValidationErr := user.NewRegisterUserValidator(inputs).RegisterUserValidation()
	if len(inputValidationErr) > 0 {
		message = "Failed to validated inputs"
		_ = response.NewResponse(inputValidationErr, message, 422).Failed(w)
		return
	}

	password, err := security.EncryptPassword(inputs["password"])
	if err != nil {
		message = "can not secure the password"
		_ = response.NewResponse(err, message, 422).Failed(w)
		return
	}

	err = u.ur.Register(&model.User{
		Username: inputs["username"],
		Password: password,
		Phone:    inputs["phone"],
	})

	if err != nil {
		message = "failed to insert user"
		_ = response.NewResponse(err, message, 400).Failed(w)
		return
	}

	token, err := jwt.GetAuthToken(inputs["username"])

	if err != nil {
		message = "failed to create authentication token"
		_ = response.NewResponse(err, message, 200).Success(w)
		return
	}

	_ = response.NewResponse(
		map[string]string{
			"token": token,
		}, message, 200).Success(w)
	return
}

func (u *UserHandler) ConnectAccount(w http.ResponseWriter, r *http.Request) {
	message := "request was processed successfully"

	ctx := r.Context()
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		message = "Failed to parse request body"
		_ = response.NewResponse("", message, 422).Failed(w)
		return
	}

	userPhone, ok := ctx.Value("phone").(string)
	if !ok {
		message = "Failed to parse request body"
		_ = response.NewResponse("", message, 422).Failed(w)
		return
	}

	err := accountService.NewWhatsAppService(u.ur).Connect(userId, userPhone)
	if err != nil {
		message = "the connection was not established"
		_ = response.NewResponse(err, message, 404).Failed(w)
		return
	}

	_ = response.NewResponse("", message, 200).Success(w)
	return
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	message := "request was processed successfully"
	var inputs map[string]string
	err := json.NewDecoder(r.Body).Decode(&inputs)
	if err != nil {
		message = "Failed to parse request body"
		_ = response.NewResponse(err, message, 422).Failed(w)
		return
	}

	inputValidationErr := user.NewLoginUserValidator(inputs).LoginUserValidation()

	if inputValidationErr != nil {
		message = "inputs are invalid"
		_ = response.NewResponse(inputValidationErr, message, 422).Failed(w)
		return
	}

	userInfo, err := u.ur.GetByUsername(inputs["username"])
	if err != nil {
		message = "invalid inputs"
		_ = response.NewResponse(err, message, 422).Failed(w)
		return
	}

	if !security.CheckPasswordHash(inputs["password"], userInfo.Password) {
		message = "invalid username or password"
		_ = response.NewResponse(err, message, 422).Failed(w)
		return
	}

	inputs["password"] = userInfo.Password

	err = u.ur.Login(inputs)
	if err != nil {
		message = "invalid username or password"
		_ = response.NewResponse(err, message, 422).Failed(w)
		return
	}

	authToken, err := jwt.GetAuthToken(inputs["username"])
	if err != nil {
		message = "please try again later"
		_ = response.NewResponse(err, message, 422).Failed(w)
		return
	}

	_ = response.NewResponse(
		map[string]string{
			"token": authToken,
		}, message, 200).Success(w)
	return
}
