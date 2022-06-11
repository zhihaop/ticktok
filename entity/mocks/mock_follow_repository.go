package mocks

import (
	"github.com/zhihaop/ticktok/entity"
)

// MockFollowRepository provides a mock repository for FollowRepository
type MockFollowRepository struct {
}

func NewMockFollowRepository() *MockFollowRepository {
	//TODO implement me
	return &MockFollowRepository{}
}

func (m *MockFollowRepository) CountFollowerByID(followID int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFollowRepository) CountFollowByID(followerID int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFollowRepository) FetchFollowerByID(followID int64, offset int, limit int) ([]entity.Follow, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFollowRepository) FetchFollowByID(followerID int64, offset int, limit int) ([]entity.Follow, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockFollowRepository) InsertFollow(followerID int64, followingID int64) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockFollowRepository) DeleteFollow(followerID int64, followingID int64) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockFollowRepository) HasFollow(followerID int64, followID int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}
