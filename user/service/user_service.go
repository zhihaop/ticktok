package service

import (
	"encoding/json"
	"github.com/zhihaop/ticktok/core"
	"github.com/zhihaop/ticktok/entity"
	"log"
)

// UserServiceImpl is an implementation of UserService
type UserServiceImpl struct {
	userRepository   entity.UserRepository
	followRepository entity.FollowRepository
}

// Token is a representation of user's token.
type Token struct {
	ID   int64  `json:"id"`
	UUID string `json:"UUID"`
}

// NewUserService creates and initializes an instance of UserServiceImpl
func NewUserService(userRepository entity.UserRepository, followRepository entity.FollowRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository:   userRepository,
		followRepository: followRepository,
	}
}

// getTokenString generates token for a specific user.
func getTokenString(id int64) string {
	token, _ := json.Marshal(&Token{ID: id, UUID: core.GetUUID()})
	return string(token)
}

func (u *UserServiceImpl) Register(username string, password string) (*entity.UserLoginToken, error) {
	// the length of username should in range [5, 32]
	if len(username) < 5 || len(username) > 32 {
		return nil, core.ErrUsernameLengthInvalid
	}
	// the length of password should in range [5, 32]
	if len(password) < 5 || len(password) > 32 {
		return nil, core.ErrPasswordLengthInvalid
	}

	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		log.Println(err)
		return nil, core.ErrInternalServerError
	}

	if user != nil {
		return nil, core.ErrUsernameExists
	}

	salt := core.GetUUID()
	encoded := core.Encoded(password, salt)
	if err := u.userRepository.CreateUser(username, encoded, salt); err != nil {
		log.Println(err)
		return nil, core.ErrInternalServerError
	}
	return u.Login(username, password)
}

func (u *UserServiceImpl) Login(username string, password string) (*entity.UserLoginToken, error) {
	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		log.Println(err)
		return nil, core.ErrInternalServerError
	}

	if user == nil {
		return nil, core.ErrUserNotExist
	}

	if core.Encoded(password, user.Salt) != user.Password {
		return nil, core.ErrUsernameOrPasswordInvalid
	}

	if user.Token != nil {
		return &entity.UserLoginToken{ID: user.ID, Token: *user.Token}, nil
	}

	token := getTokenString(user.ID)
	if err := u.userRepository.UpdateTokenByID(user.ID, token); err != nil {
		log.Println(err)
		return nil, core.ErrInternalServerError
	}
	return &entity.UserLoginToken{ID: user.ID, Token: token}, nil
}

func (u *UserServiceImpl) GetUsername(id int64) (string, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		log.Println(err)
		return "", core.ErrInternalServerError
	}

	if user == nil {
		return "", core.ErrUserNotExist
	}
	return user.Name, nil
}

func (u *UserServiceImpl) GetFollowerCount(id int64) (int64, error) {
	followerCount, err := u.followRepository.CountFollowerByID(id)
	if err != nil {
		log.Println(err)
		return 0, core.ErrInternalServerError
	}
	return followerCount, nil
}

func (u *UserServiceImpl) GetFollowCount(id int64) (int64, error) {
	followingCount, err := u.followRepository.CountFollowByID(id)
	if err != nil {
		log.Println(err)
		return 0, core.ErrInternalServerError
	}
	return followingCount, nil
}

func (u *UserServiceImpl) GetUserID(token string) (int64, error) {
	user, err := u.userRepository.FindByToken(token)
	if err != nil {
		log.Println(err)
		return 0, core.ErrInternalServerError
	}

	if user == nil {
		return 0, core.ErrUserNotExist
	}
	return user.ID, nil
}

func (u *UserServiceImpl) IsFollow(followerID int64, followID int64) (bool, error) {
	if followerID == followID {
		return false, nil
	}

	isFollow, err := u.followRepository.HasFollow(followerID, followID)
	if err != nil {
		log.Println(err)
		return false, core.ErrInternalServerError
	}
	return isFollow, nil
}
