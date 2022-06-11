package core

import "errors"

// UserService error type
var (
	ErrInternalServerError       = errors.New("internal server error")
	ErrUserNotExist              = errors.New("user not exist")
	ErrUsernameOrPasswordInvalid = errors.New("username or password is invalid")
	ErrUsernameExists            = errors.New("username already exists")
	ErrUsernameLengthInvalid     = errors.New("the length of username is invalid")
	ErrPasswordLengthInvalid     = errors.New("the length of password is invalid")
)
