package models

import (
	entities "github.com/toshkentov01/template/entities/user"
)

func GetProfileTransfromator(user *entities.UserProfileModel) *UserProfile {
	return &UserProfile{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt,
	}
}
