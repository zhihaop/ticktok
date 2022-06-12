package user_service

import (
	"encoding/base64"
	"encoding/json"
	"github.com/zhihaop/ticktok/core"
	"github.com/zhihaop/ticktok/core/service"
	"github.com/zhihaop/ticktok/entity"
	"gopkg.in/errgo.v2/errors"
	"log"
)

// UserServiceImpl is an implementation of UserService
type UserServiceImpl struct {
	entity.UserService
	userRepository entity.UserRepository
}

// Token is a representation of user's token.
type Token struct {
	ID   int64  `json:"id"`
	UUID string `json:"UUID"`
}

// NewUserService creates and initializes an instance of UserServiceImpl
func NewUserService(userRepository entity.UserRepository) entity.UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

// GenerateToken generates the token for specific user.
func GenerateToken(id int64) Token {
	return Token{ID: id, UUID: core.GetUUID()}
}

// EncodeToken encode the token to string.
func EncodeToken(token Token) string {
	s, _ := json.Marshal(&token)
	return base64.URLEncoding.EncodeToString(s)
}

// DecodeToken decode the string to user token
func DecodeToken(s string) (Token, error) {
	token := Token{}
	decoded, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return token, err
	}
	err = json.Unmarshal(decoded, &token)
	if err != nil {
		return token, err
	}
	return token, nil
}

func (u *UserServiceImpl) Register(username string, password string) (*entity.UserLoginToken, error) {
	// the length of username should in range [5, 32]
	if len(username) < 5 || len(username) > 32 {
		return nil, service.ErrUsernameLengthInvalid
	}
	// the length of password should in range [6, 32]
	if len(password) < 6 || len(password) > 32 {
		return nil, service.ErrPasswordLengthInvalid
	}

	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		log.Printf("Register(...): %s\n", errors.Details(err))
		return nil, service.ErrInternalServerError
	} else if user != nil {
		return nil, service.ErrUsernameExists
	}

	salt := core.GetUUID()
	encoded := core.Encoded(password, salt)
	if err := u.userRepository.CreateUser(username, encoded, salt); err != nil {
		log.Printf("Register(...): %s\n", errors.Details(err))
		return nil, service.ErrInternalServerError
	}
	return u.Login(username, password)
}

func (u *UserServiceImpl) Login(username string, password string) (*entity.UserLoginToken, error) {
	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		log.Printf("Login(...): %s\n", errors.Details(err))
		return nil, service.ErrInternalServerError
	} else if user == nil {
		return nil, service.ErrUserNotExist
	} else if core.Encoded(password, user.Salt) != user.Password {
		return nil, service.ErrUsernameOrPasswordInvalid
	} else if user.Token != nil {
		return &entity.UserLoginToken{ID: user.ID, Token: *user.Token}, nil
	}

	token := GenerateToken(user.ID)
	tokenString := EncodeToken(token)
	if err := u.userRepository.UpdateTokenByID(user.ID, tokenString); err != nil {
		log.Printf("Login(...): %s\n", errors.Details(err))
		return nil, service.ErrInternalServerError
	}
	return &entity.UserLoginToken{ID: user.ID, Token: tokenString}, nil
}

func (u *UserServiceImpl) GetUsername(id int64) (string, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		log.Printf("GetUsername(...): %s\n", errors.Details(err))
		return "", service.ErrInternalServerError
	} else if user == nil {
		return "", service.ErrUserNotExist
	}
	return user.Name, nil
}

func (u *UserServiceImpl) GetUserID(token string) (int64, error) {
	user, err := u.userRepository.FindByToken(token)
	if err != nil {
		log.Printf("GetUserID(...): %s\n", errors.Details(err))
		return 0, service.ErrInternalServerError
	} else if user == nil {
		return 0, service.ErrUserNotExist
	}
	return user.ID, nil
}
