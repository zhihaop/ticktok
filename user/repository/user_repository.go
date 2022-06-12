package user_repository

import (
	utils "github.com/zhihaop/ticktok/core/repository"
	"github.com/zhihaop/ticktok/entity"
	"gorm.io/gorm"
	"log"
)

// UserRepositoryImpl is an implementation of UserRepository
type UserRepositoryImpl struct {
	entity.UserRepository
	db *gorm.DB
}

// NewUserRepository creates a user's repository
func NewUserRepository(db *gorm.DB) entity.UserRepository {
	if err := utils.CheckOrCreateTable(db, &entity.User{}); err != nil {
		log.Fatalln(err)
	}
	return &UserRepositoryImpl{db: db}
}

func (u *UserRepositoryImpl) CreateUser(username string, password string, salt string) error {
	user := &entity.User{
		Name:     username,
		Salt:     salt,
		Password: password,
	}
	db := u.db.Save(&user)
	return db.Error
}

func (u *UserRepositoryImpl) UpdateTokenByID(id int64, token string) error {
	db := u.db.Model(&entity.User{}).Where("id = ?", id).Update("token", token)
	return db.Error
}

func (u *UserRepositoryImpl) FindByUsername(username string) (*entity.User, error) {
	user := make([]entity.User, 0)
	db := u.db.Where("name = ?", username).Find(&user).Limit(1)
	if db.Error != nil {
		return nil, db.Error
	}

	if len(user) == 0 {
		return nil, nil
	}
	return &user[0], nil
}

func (u *UserRepositoryImpl) FindByID(id int64) (*entity.User, error) {
	user := make([]entity.User, 0)
	db := u.db.Where("id = ?", id).Find(&user).Limit(1)
	if db.Error != nil {
		return nil, db.Error
	}

	if len(user) == 0 {
		return nil, nil
	}
	return &user[0], nil
}

func (u *UserRepositoryImpl) FindByToken(token string) (*entity.User, error) {
	user := make([]entity.User, 0)
	db := u.db.Where("token = ?", token).Find(&user).Limit(1)
	if db.Error != nil {
		return nil, db.Error
	}

	if len(user) == 0 {
		return nil, nil
	}
	return &user[0], nil
}
