package repository

import (
	"github.com/zhihaop/ticktok/entity"
	"gorm.io/gorm"
)

// FollowRepositoryImpl is an implementation of FollowRepository
type FollowRepositoryImpl struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) *FollowRepositoryImpl {
	return &FollowRepositoryImpl{db: db}
}

func (f FollowRepositoryImpl) CountFollowerByID(followID int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (f FollowRepositoryImpl) CountFollowByID(followerID int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (f FollowRepositoryImpl) FetchFollowerByID(followID int64, offset int, limit int) ([]entity.Follow, error) {
	//TODO implement me
	panic("implement me")
}

func (f FollowRepositoryImpl) FetchFollowByID(followerID int64, offset int, limit int) ([]entity.Follow, error) {
	//TODO implement me
	panic("implement me")
}

func (f FollowRepositoryImpl) InsertFollow(followerID int64, followingID int64) error {
	//TODO implement me
	panic("implement me")
}

func (f FollowRepositoryImpl) DeleteFollow(followerID int64, followingID int64) error {
	//TODO implement me
	panic("implement me")
}

func (f FollowRepositoryImpl) HasFollow(followerID int64, followID int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}
