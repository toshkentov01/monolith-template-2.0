package controllers

import (
	"github.com/gofiber/fiber/v2"
	service "github.com/toshkentov01/template/service/user"
)

// API ...
type API struct {
	userService service.UserService
}

// APIInterface ...
type APIInterface interface {
	SignUp(ctx *fiber.Ctx) error
	GetProfile(ctx *fiber.Ctx) error
	GetMyProfile(ctx *fiber.Ctx) error
}

// NewAPI ...
func NewAPI(userService service.UserService) APIInterface {
	return &API{
		userService: userService,
	}
}
