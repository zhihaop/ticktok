package mocks

import (
	"github.com/zhihaop/ticktok/entity"
	"github.com/zhihaop/ticktok/user/service"
)

// MockFollowService provides a mock service for FollowService
type MockFollowService struct {
	entity.FollowService
}

func NewMockFollowService() entity.FollowService {
	return &user_service.FollowServiceImpl{
		FollowRepository: NewMockFollowRepository(),
		UserRepository:   NewMockUserRepository(),
	}
}
