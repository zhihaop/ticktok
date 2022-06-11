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
	entity.FollowRepository
}

func NewMockFollowRepository() entity.FollowRepository {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	mock := &MockFollowRepository{repository.NewFollowRepository(db)}
	return mock
}
