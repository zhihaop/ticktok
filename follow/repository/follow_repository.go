package follow_repository

import (
	utils "github.com/zhihaop/ticktok/core/repository"
	"github.com/zhihaop/ticktok/core/service"
	"github.com/zhihaop/ticktok/entity"
	"gorm.io/gorm"
	"log"
)

// FollowRepositoryImpl is an implementation of FollowRepository
type FollowRepositoryImpl struct {
	entity.FollowRepository
	db *gorm.DB
}

// FetchFollow TODO add `offset` and `limit` support
func (f *FollowRepositoryImpl) FetchFollow(followerID int64, offset int64, limit int64) ([]entity.Follow, error) {
	follows := make([]entity.Follow, 0)
	db := f.db.Model(&entity.Follow{}).Where("follower_id = ?", followerID).Find(&follows)
	if db.Error != nil {
		return nil, db.Error
	}
	return follows, nil
}

// FetchFollower TODO add `offset` and `limit` support
func (f *FollowRepositoryImpl) FetchFollower(followID int64, offset int64, limit int64) ([]entity.Follow, error) {
	follows := make([]entity.Follow, 0)
	db := f.db.Model(&entity.Follow{}).Where("follow_id = ?", followID).Find(&follows)
	if db.Error != nil {
		return nil, db.Error
	}
	return follows, nil
}

func NewFollowRepository(db *gorm.DB) entity.FollowRepository {
	if err := utils.CheckOrCreateTable(db, &entity.Follow{}); err != nil {
		log.Fatalln(err)
	}
	return &FollowRepositoryImpl{db: db}
}

func (f *FollowRepositoryImpl) CountFollowerByID(followID int64) (int64, error) {
	followerCount := int64(0)
	db := f.db.Model(&entity.Follow{}).Where("follow_id = ?", followID).Count(&followerCount)
	if db.Error != nil {
		return -1, db.Error
	}
	return followerCount, nil
}

func (f *FollowRepositoryImpl) CountFollowByID(followerID int64) (int64, error) {
	followCount := int64(0)
	db := f.db.Model(&entity.Follow{}).Where("follower_id = ?", followerID).Count(&followCount)
	if db.Error != nil {
		return -1, db.Error
	}
	return followCount, nil
}

func (f *FollowRepositoryImpl) InsertFollow(followerID int64, followID int64) error {
	db := f.db.Save(&entity.Follow{
		FollowerID: followerID,
		FollowID:   followID,
	})
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (f *FollowRepositoryImpl) DeleteFollow(followerID int64, followID int64) error {
	db := f.db.Where("follower_id = ? AND follow_id = ?", followerID, followID).Delete(&entity.Follow{})
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (f *FollowRepositoryImpl) HasFollow(followerID int64, followID int64) (bool, error) {
	count := int64(0)
	db := f.db.Model(&entity.Follow{}).Where("follower_id = ? AND follow_id = ?", followerID, followID).Count(&count)
	if db.Error != nil {
		return false, db.Error
	} else if count < 0 || count > 1 {
		return false, service.ErrInternalServerError
	}
	return count == 1, nil
}
