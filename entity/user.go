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

// UserService represents the user's service
type UserService interface {
	Register(username string, password string) (*UserLoginToken, error)
	Login(username string, password string) (*UserLoginToken, error)
	GetUsername(id int64) (string, error)
	GetFollowerCount(id int64) (int64, error)
	GetFollowCount(id int64) (int64, error)
	GetUserID(token string) (int64, error)
	IsFollow(followerID int64, followID int64) (bool, error)
}

// UserRepository represents the user's repository
type UserRepository interface {
	CreateUser(username string, password string, salt string) error
	UpdateTokenByID(id int64, token string) error
	FindByUsername(username string) (*User, error)
	FindByID(id int64) (*User, error)
	FindByToken(token string) (*User, error)
}
