package repository

import (
	"gitlab.com/M.darvish/funtory/internal/model"
	_interface "gitlab.com/M.darvish/funtory/internal/model/repository/interface"
	"gorm.io/gorm"
)

type IUser _interface.IUser

type UserRepositoryImp struct {
	db *gorm.DB
}

func NewUserImp(db *gorm.DB) IUser {
	return &UserRepositoryImp{db}
}

func (u *UserRepositoryImp) Register(inputs *model.User) error {
	result := u.db.Create(inputs)
	return result.Error
}

func (u *UserRepositoryImp) GetByUserId(userId int) (*model.User, error) {
	var user model.User
	result := u.db.Where("id = ? AND deleted_at IS NULL", userId).First(&user)
	return &user, result.Error
}

func (u *UserRepositoryImp) GetByUsername(username string) (*model.User, error) {
	var user model.User
	result := u.db.Where("username = ? AND deleted_at IS NULL", username).First(&user)
	return &user, result.Error
}

func (u *UserRepositoryImp) UpdateWhatsAppSession(id int, session string) error {
	updates := map[string]interface{}{
		"session": session, // Assuming "session" is the name of the column.
	}
	result := u.db.Table("users").Where("id = ? AND deleted_at IS NULL", id).Updates(updates)
	return result.Error
}

func (u *UserRepositoryImp) Login(inputs map[string]string) error {
	var user model.User
	result := u.db.Where("username = ? AND password = ? AND deleted_at IS NULL", inputs["username"], inputs["password"]).First(&user)
	return result.Error
}
