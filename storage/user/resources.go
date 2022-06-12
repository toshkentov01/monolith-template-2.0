package user

import (
	entities "github.com/toshkentov01/template/entities/user"
)

// UserRepository ...
type UserRepository interface {
	Reader
	Writer
}

// Reader ...
type Reader interface {
	GetUser(id string) (*entities.UserProfileModel, error)
}

// Writer ...
type Writer interface {
	CreateUser(user *entities.CreateUserModel) error
}
