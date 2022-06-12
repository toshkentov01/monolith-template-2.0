package user

import (
	"context"

	entities "github.com/toshkentov01/template/entities/user"
	"github.com/toshkentov01/template/pkg/logger"
	"github.com/toshkentov01/template/storage"
)

// UserService ...
type UserService struct {
	storage storage.Interface
	logger  logger.Logger
}

// NewUserService ...
func NewUserService(logger logger.Logger) *UserService {
	return &UserService{
		storage: storage.NewStorage(),
		logger:  logger,
	}
}

// CreateUser ...
func (ur *UserService) CreateUser(ctx context.Context, request *entities.CreateUserModel) error {
	err := ur.storage.User().CreateUser(request)
	if err != nil {
		ur.logger.Error("Error while creating user, error: " + err.Error())
		return err
	}

	return nil
}

// GetUser ...
func (ur *UserService) GetUser(ctx context.Context, request *entities.GetUserModel) (*entities.UserProfileModel, error) {
	user, err := ur.storage.User().GetUser(request.ID)
	if err != nil {
		ur.logger.Error("Error while getting user, error: " + err.Error())
		return nil, err
	}

	return user, nil
}
