package mocks

import (
	"github.com/zhihaop/ticktok/entity"
	"github.com/zhihaop/ticktok/user/service"
)

// NewMockFollowService creates a mock service for FollowService
func NewMockFollowService() entity.FollowService {
	return user_service.NewFollowService(
		NewMockFollowRepository(),
		NewMockUserRepository(),
	)
}
