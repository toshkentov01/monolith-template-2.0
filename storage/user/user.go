package user

import (
	"database/sql"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	entities "github.com/toshkentov01/template/entities/user"
	"github.com/toshkentov01/template/pkg/errs"
	"github.com/toshkentov01/template/pkg/platform/postgres"
)

// userRepo ...
type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo() UserRepository {
	return &userRepo{
		db: postgres.DB(),
	}
}

// CreateUser ...
func (ur *userRepo) CreateUser(user *entities.CreateUserModel) error {
	if _, err := ur.db.NamedExec(CreateUserSQL, newDbUserProfile(*user)); err != nil {

		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "users_username_key"`) {
			return errs.ErrUsernameExists

		} else if strings.Contains(err.Error(), `duplicate key value violates unique constraint "users_email_key"`) {
			return errs.ErrEmailExists
		}

		log.Println("error: ", err.Error())
		return errs.ErrInternal
	}

	return nil
}

// GetUser ...
func (ur *userRepo) GetUser(id string) (*entities.UserProfileModel, error) {
	var user dbUserProfile

	if err := ur.db.Get(&user, GetUserSQL, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ErrUserNotFound
		}

		return nil, errs.ErrInternal
	}

	return user.toModel(), nil
}
