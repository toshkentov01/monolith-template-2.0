package controllers

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	errorhandler "github.com/toshkentov01/template/api/error_handler"
	"github.com/toshkentov01/template/api/models"
	"github.com/toshkentov01/template/config"
	entities "github.com/toshkentov01/template/entities/user"
	"github.com/toshkentov01/template/pkg/errs"
	"github.com/toshkentov01/template/pkg/utils"
)

// SignUp method for sign up.
// @Description SignUp API used for signing up.
// @Tags register
// @Accept json
// @Produce json
// @Param veridy_model body models.SignupRequestModel true "SignupRequestModel"
// @Success 200 {object} models.SignupResponseModel
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /v1/register/signup/ [post]
func (api *API) SignUp(c *fiber.Ctx) error {
	// Get request body.
	var (
		signupRequest models.SignupRequestModel
	)

	if err := c.BodyParser(&signupRequest); err != nil {
		return errorhandler.HandleBadRequestErrWithMessage(c, err, "Invalid request body")
	}

	signupRequest.Username = strings.ToLower(signupRequest.Username)

	if valid := errorhandler.UsernameValidator(signupRequest.Username); !valid {
		return errorhandler.HandleBadRequestErrWithMessage(c, nil, "Invalid username")
	}

	// Declaring context with timeOut
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Config().CtxTimeout))
	defer cancel()

	userID, _ := uuid.NewRandom()

	err := api.userService.CreateUser(ctx, &entities.CreateUserModel{
		ID:       userID.String(),
		Username: signupRequest.Username,
		FullName: signupRequest.FullName,
		Email:    signupRequest.Email,
		Password: signupRequest.Password,
	})

	if err != nil {
		if err == errs.ErrEmailExists {
			return errorhandler.HandleBadRequestErrWithMessage(c, err, "Email already exists")

		} else if err == errs.ErrUsernameExists {
			return errorhandler.HandleBadRequestErrWithMessage(c, err, "Username already exists")
		}

		return errorhandler.HandleInternalWithMessage(c, err, err.Error())
	}

	//Creating access and refresh tokens
	tokens, err := utils.GenerateNewTokens(userID.String(), map[string]string{"role": "user"})

	if err != nil {
		log.Println("Error while generating tokens! ", err)
		return errorhandler.HandleBadRequestErrWithMessage(c, err, err.Error())
	}

	return c.Status(http.StatusOK).JSON(models.SignupResponseModel{
		ID:           userID.String(),
		AccessToken:  tokens.Access,
		RefreshToken: tokens.Refresh,
	})
}

// GetProfile ...
// @Security ApiKeyAuth
// @Summary Get Profile
// @Description GetProfile
// @Tags user
// @Accept json
// @Produce json
// @Param user_id path string true "user_id"
// @Success 200 {object} models.UserProfile
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /v1/user/profile/{user_id}/ [get]
func (api *API) GetProfile(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	if valid := errorhandler.UUIDValidator(userID); !valid {
		return errorhandler.HandleBadRequestErrWithMessage(c, nil, "Invalid user id")
	}

	// Declaring context with timeOut
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Config().CtxTimeout))
	defer cancel()

	user, err := api.userService.GetUser(ctx, &entities.GetUserModel{
		ID: userID,
	})

	if err != nil {
		if err == errs.ErrUserNotFound {
			return errorhandler.HandleBadRequestErrWithMessage(c, err, "User not found")
		}

		return errorhandler.HandleInternalWithMessage(c, err, err.Error())
	}

	return c.Status(http.StatusOK).JSON(models.GetProfileTransfromator(user))
}

// GetMyProfile ...
// @Security ApiKeyAuth
// @Summary Gets My Profile
// @Description GetProfile API gets my profile.
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.UserProfile
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /v1/user/get-my-profile/ [get]
func (api *API) GetMyProfile(c *fiber.Ctx) error {

	user, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return errorhandler.HandleBadRequestErrWithMessage(c, err, err.Error())
	}

	// Declaring context with timeOut
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Config().CtxTimeout))
	defer cancel()

	profile, err := api.userService.GetUser(ctx, &entities.GetUserModel{
		ID: user.UserID.String(),
	})

	if err != nil {
		if err == errs.ErrUserNotFound {
			return errorhandler.HandleBadRequestErrWithMessage(c, err, "User not found")
		}

		return errorhandler.HandleInternalWithMessage(c, err, err.Error())
	}

	return c.Status(http.StatusOK).JSON(models.GetProfileTransfromator(profile))
}
