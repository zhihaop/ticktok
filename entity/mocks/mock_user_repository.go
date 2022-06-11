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
	repo *repository.UserRepositoryImpl
}

func NewMockUserRepository() *MockUserRepository {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	mock := &MockUserRepository{repo: repository.NewUserRepository(db)}
	return mock
}

func (m *MockUserRepository) CreateUser(username string, password string, salt string) error {
	return m.repo.CreateUser(username, password, salt)
}

func (m *MockUserRepository) UpdateTokenByID(id int64, token string) error {
	return m.repo.UpdateTokenByID(id, token)
}

func (m *MockUserRepository) FindByUsername(username string) (*entity.User, error) {
	return m.repo.FindByUsername(username)
}

func (m *MockUserRepository) FindByID(id int64) (*entity.User, error) {
	return m.repo.FindByID(id)
}

func (m *MockUserRepository) FindByToken(token string) (*entity.User, error) {
	return m.repo.FindByToken(token)
}
