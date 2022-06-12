package errs

import "errors"

var (
	// ErrInternal ...
	ErrInternal = errors.New("internal server error")

	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("user not found")

	// ErrUserAlreadyExists ...
	ErrUserAlreadyExists = errors.New("user already exists")

	// ErrUsernameExists ...
	ErrUsernameExists = errors.New("username already exists")

	// ErrEmailExists ...
	ErrEmailExists = errors.New("email already exists")
)
