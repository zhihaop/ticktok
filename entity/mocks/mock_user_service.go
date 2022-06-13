package mocks

import (
	"github.com/zhihaop/ticktok/entity"
	user_service "github.com/zhihaop/ticktok/user/service"
)

// NewMockUserService creates a mock service for UserService
func NewMockUserService() entity.UserService {
	return user_service.NewUserService(
		NewMockUserRepository(),
		NewMockFollowRepository(),
	)
}
