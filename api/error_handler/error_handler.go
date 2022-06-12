package errorhandler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/toshkentov01/template/api/models"
)

const (
	//ErrorCodeInvalidURL ...
	ErrorCodeInvalidURL = "INVALID_URL"
	//ErrorCodeInvalidJSON ...
	ErrorCodeInvalidJSON = "INVALID_JSON"
	//ErrorCodeInvalidParams ...
	ErrorCodeInvalidParams = "INVALID_PARAMS"
	//ErrorCodeInternal ...
	ErrorCodeInternal = "INTERNAL"
	//ErrorCodeUnauthorized ...
	ErrorCodeUnauthorized = "UNAUTHORIZED"
	//ErrorCodeAlreadyExists ...
	ErrorCodeAlreadyExists = "ALREADY_EXISTS"
	//ErrorCodeNotFound ...
	ErrorCodeNotFound = "NOT_FOUND"
	//ErrorCodeInvalidCode ...
	ErrorCodeInvalidCode = "INVALID_CODE"
	//ErrorBadRequest ...
	ErrorBadRequest = "BAD_REQUEST"
	//ErrorCodeForbidden ...
	ErrorCodeForbidden = "FORBIDDEN"
	//ErrorCodeNotApproved ...
	ErrorCodeNotApproved = "NOT_APPROVED"
	//ErrorCodeWrongClub ...
	ErrorCodeWrongClub = "WRONG_CLUB"
	//ErrorCodePasswordsNotEqual ...
	ErrorCodePasswordsNotEqual = "PASSWORDS_NOT_EQUAL"
	// ErrorExpectationFailed ...
	ErrorExpectationFailed = "EXPECTATION_FAILED"
	// ErrorUpgradeRequired ...
	ErrorUpgradeRequired = "UPGRADE_REQUIRED"
	// ErrorInvalidCredentials ...
	ErrorInvalidCredentials = "INVALID_CREDENTIALS"
)

func HandleInternalWithMessage(c *fiber.Ctx, err error, message string) error {
	log.Println(message+" ", err)
	return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
		Error: true,
		Data: models.Error{
			Status:  "Internal Server Error",
			Message: message,
		},
	})
}

func HandleBadRequestErrWithMessage(c *fiber.Ctx, err error, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(models.Response{
		Error: true,
		Data: models.Error{
			Status:  "Bad Request",
			Message: message,
		},
	})
}

func HandlerStatusOk(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(models.Response{
		Error: false,
		Data: models.SuccessMessage{
			Success: true,
		},
	})
}

// ParsePageQueryParam ...
func ParsePageQueryParam(c *fiber.Ctx) (int, error) {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return 0, err
	}
	if page < 0 {
		return 0, errors.New("page must be an positive integer")
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

//ParseLimitQueryParam ...
func ParseLimitQueryParam(c *fiber.Ctx) (int, error) {
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return 0, err
	}
	if limit < 0 {
		return 0, errors.New("limit must be an positive integer")
	}
	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}

// UUIDValidator - Validator for UUID
func UUIDValidator(ID string) bool {
	_, err := uuid.Parse(ID)

	return err == nil
}

// UsernameValidator - Validator for UUID
func UsernameValidator(usrname string) bool {
	return len(usrname) >= 5 && len(usrname) <= 30 && !strings.ContainsAny(usrname, " ")
}
