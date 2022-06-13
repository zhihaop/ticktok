package mocks

import (
	"github.com/zhihaop/ticktok/entity"
	"github.com/zhihaop/ticktok/user/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// NewMockFollowRepository creates a mock repository for FollowRepository
func NewMockFollowRepository() entity.FollowRepository {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return user_repository.NewFollowRepository(db)
}
