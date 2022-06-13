package entity

import (
	"time"
)

// User is a row in `Users` table
type User struct {
	ID       int64 `gorm:"autoIncrement"`
	Name     string
	Salt     string
	Password string
	Token    *string
	CreateAt time.Time
	UpdateAt time.Time
}

// UserLoginToken represents the user's token
type UserLoginToken struct {
	ID    int64
	Token string
}

// UserInfo includes all the info about user's follow relationship
type UserInfo struct {
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	ID            int64  `json:"id"`
	IsFollow      bool   `json:"is_follow"`
	Name          string `json:"name"`
}

// UserService represents the user's service
type UserService interface {
	Register(username string, password string) (*UserLoginToken, error)
	Login(username string, password string) (*UserLoginToken, error)
	GetUsername(id int64) (string, error)
	GetUserID(token string) (int64, error)
	GetUserInfo(userID int64, queryID int64) (*UserInfo, error)
}

// UserRepository represents the user's repository
type UserRepository interface {
	CreateUser(username string, password string, salt string) error
	UpdateTokenByID(id int64, token string) error
	FindByUsername(username string) (*User, error)
	FindByID(id int64) (*User, error)
	FindByToken(token string) (*User, error)
}
