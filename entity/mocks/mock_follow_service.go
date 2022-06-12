package mocks

import (
	"github.com/zhihaop/ticktok/entity"
	"github.com/zhihaop/ticktok/follow/service"
)

// MockFollowService provides a mock service for FollowService
type MockFollowService struct {
	entity.FollowService
}

func NewMockFollowService() entity.FollowService {
	return &follow_service.FollowServiceImpl{
		FollowRepository: NewMockFollowRepository(),
		UserRepository:   NewMockUserRepository(),
	}
}
