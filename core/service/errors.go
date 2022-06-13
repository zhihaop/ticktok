package service

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

// FollowService error type
var (
	ErrActionInValid    = errors.New("action is invalid")
	ErrRelationExist    = errors.New("action follow failed: relation exists")
	ErrRelationNotExist = errors.New("action unfollow failed: relation not exist")
	ErrSelfFollowing    = errors.New("action follow failed: self following")
	ErrSelfUnFollowing  = errors.New("action unfollow failed: self following")
)

// ClipService error type
var (
	ErrVideoFileInValid = errors.New("the uploaded video file is invalid")
)
