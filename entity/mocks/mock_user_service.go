package mocks

import (
	"github.com/zhihaop/ticktok/entity"
	user_service "github.com/zhihaop/ticktok/user/service"
)

// MockUserService provides a mock service for us
type MockUserService struct {
	entity.UserService
}

func NewMockUserService() entity.UserService {
	userRepository := NewMockUserRepository()

	return &MockUserService{user_service.NewUserService(userRepository)}
}
