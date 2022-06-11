package mocks

import (
	"github.com/zhihaop/ticktok/entity"
	"github.com/zhihaop/ticktok/user/service"
)

// MockUserService provides a mock service for UserService
type MockUserService struct {
	entity.UserService
}

func NewMockUserService() entity.UserService {
	userRepository := NewMockUserRepository()
	followRepository := NewMockFollowRepository()

	return &MockUserService{service.NewUserService(userRepository, followRepository)}
}
