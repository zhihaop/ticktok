package mocks

import (
	"github.com/zhihaop/ticktok/entity"
	"github.com/zhihaop/ticktok/user/service"
)

// MockUserService provides a mock service for UserService
type MockUserService struct {
	service entity.UserService
}

func NewMockUserService() *MockUserService {
	userRepository := NewMockUserRepository()
	followRepository := NewMockFollowRepository()

	return &MockUserService{service: service.NewUserService(userRepository, followRepository)}
}

func (m *MockUserService) Register(username string, password string) (*entity.UserLoginToken, error) {
	return m.service.Register(username, password)
}

func (m *MockUserService) Login(username string, password string) (*entity.UserLoginToken, error) {
	return m.service.Login(username, password)
}

func (m *MockUserService) GetUsername(id int64) (string, error) {
	return m.service.GetUsername(id)
}

func (m *MockUserService) GetFollowerCount(id int64) (int64, error) {
	return m.service.GetFollowerCount(id)
}

func (m *MockUserService) GetFollowCount(id int64) (int64, error) {
	return m.service.GetFollowCount(id)
}

func (m *MockUserService) GetUserID(token string) (int64, error) {
	return m.service.GetUserID(token)
}

func (m *MockUserService) IsFollow(followerID int64, followID int64) (bool, error) {
	return m.service.IsFollow(followerID, followID)
}
