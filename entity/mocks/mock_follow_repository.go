package mocks

import (
	"github.com/zhihaop/ticktok/entity"
	"github.com/zhihaop/ticktok/follow/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// MockFollowRepository provides a mock repository for FollowRepository
type MockFollowRepository struct {
	repo entity.FollowRepository
}

func NewMockFollowRepository() entity.FollowRepository {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	mock := &MockFollowRepository{repo: repository.NewFollowRepository(db)}
	return mock
}

func (m *MockFollowRepository) CountFollowerByID(followID int64) (int64, error) {
	return m.repo.CountFollowByID(followID)
}

func (m *MockFollowRepository) CountFollowByID(followerID int64) (int64, error) {
	return m.repo.CountFollowByID(followerID)
}

func (m *MockFollowRepository) InsertFollow(followerID int64, followID int64) error {
	return m.repo.InsertFollow(followerID, followID)
}

func (m *MockFollowRepository) DeleteFollow(followerID int64, followID int64) error {
	return m.repo.DeleteFollow(followerID, followID)
}

func (m *MockFollowRepository) HasFollow(followerID int64, followID int64) (bool, error) {
	return m.repo.HasFollow(followerID, followID)
}
