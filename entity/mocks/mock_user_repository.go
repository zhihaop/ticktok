package mocks

import (
	"github.com/zhihaop/ticktok/entity"
	"github.com/zhihaop/ticktok/user/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// MockUserRepository provides a mock repository for UserRepository
type MockUserRepository struct {
	entity.UserRepository
}

func NewMockUserRepository() entity.UserRepository {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	mock := &MockUserRepository{user_repository.NewUserRepository(db)}
	return mock
}
