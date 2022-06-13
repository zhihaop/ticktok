package mocks

import (
	"github.com/zhihaop/ticktok/entity"
	"github.com/zhihaop/ticktok/user/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// NewMockUserRepository creates a mock repository for UserRepository
func NewMockUserRepository() entity.UserRepository {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return user_repository.NewUserRepository(db)
}
