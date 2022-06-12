package user

import (
	"database/sql"
	"time"

	entities "github.com/toshkentov01/template/entities/user"
	"github.com/toshkentov01/template/pkg/utils"
)

// dbUserProfile ...
type dbUserProfile struct {
	ID        string         `db:"id"`
	Username  string         `db:"username"`
	FullName  sql.NullString `db:"full_name"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}

// newDbUserProfile ...
func newDbUserProfile(user entities.CreateUserModel) *dbUserProfile {
	return &dbUserProfile{
		ID:       user.ID,
		Username: user.Username,
		FullName: utils.StringToNullString(user.FullName),
		Email:    user.Email,
		Password: user.Password,
	}
}

// toModel ...
func (u dbUserProfile) toModel() *entities.UserProfileModel {
	return &entities.UserProfileModel{
		ID:        u.ID,
		Username:  u.Username,
		FullName:  u.FullName.String,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.Time.String(),
	}
}
