package _interface

import "gitlab.com/M.darvish/funtory/internal/model"

type IUser interface {
	Register(inputs *model.User) error
	GetByUserId(id int) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	UpdateWhatsAppSession(id int, session string) error
	Login(inputs map[string]string) error
}
