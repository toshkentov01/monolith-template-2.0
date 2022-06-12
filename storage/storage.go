package storage

import "github.com/toshkentov01/template/storage/user"

// Interface ...
type Interface interface {
	User() user.UserRepository
}

// storage ...
type storage struct {
	userRepo user.UserRepository
}

// NewStorage ...
func NewStorage() Interface {
	return &storage{
		userRepo: user.NewUserRepo(),
	}
}

func (s storage) User() user.UserRepository {
	return s.userRepo
}
