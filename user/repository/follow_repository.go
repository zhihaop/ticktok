package user_repository

import (
	utils "github.com/zhihaop/ticktok/core/repository"
	"github.com/zhihaop/ticktok/core/service"
	"github.com/zhihaop/ticktok/entity"
	"gorm.io/gorm"
	"log"
)

// followRepositoryImpl is an implementation of FollowRepository
type followRepositoryImpl struct {
	entity.FollowRepository
	db *gorm.DB
}

func (f *followRepositoryImpl) FetchFollow(followerID int64, offset int, limit int) ([]entity.Follow, error) {
	follows := make([]entity.Follow, 0)
	query := f.db.Model(&entity.Follow{})

	query = query.Where("follower_id = ?", followerID).Offset(offset).Limit(limit)
	if err := query.Find(&follows).Error; err != nil {
		return nil, err
	}
	return follows, nil
}

func (f *followRepositoryImpl) FetchFollower(followID int64, offset int, limit int) ([]entity.Follow, error) {
	follows := make([]entity.Follow, 0)
	query := f.db.Model(&entity.Follow{})

	query = query.Where("follow_id = ?", followID).Offset(offset).Limit(limit)
	if err := query.Find(&follows).Error; err != nil {
		return nil, err
	}
	return follows, nil
}

func NewFollowRepository(db *gorm.DB) entity.FollowRepository {
	if err := utils.CheckOrCreateTable(db, &entity.Follow{}); err != nil {
		log.Fatalln(err)
	}
	return &followRepositoryImpl{db: db}
}

func (f *followRepositoryImpl) CountFollowerByID(followID int64) (int64, error) {
	followerCount := int64(0)
	db := f.db.Model(&entity.Follow{}).Where("follow_id = ?", followID).Count(&followerCount)
	if db.Error != nil {
		return -1, db.Error
	}
	return followerCount, nil
}

func (f *followRepositoryImpl) CountFollowByID(followerID int64) (int64, error) {
	followCount := int64(0)
	db := f.db.Model(&entity.Follow{}).Where("follower_id = ?", followerID).Count(&followCount)
	if db.Error != nil {
		return -1, db.Error
	}
	return followCount, nil
}

func (f *followRepositoryImpl) InsertFollow(followerID int64, followID int64) error {
	db := f.db.Save(&entity.Follow{
		FollowerID: followerID,
		FollowID:   followID,
	})
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (f *followRepositoryImpl) DeleteFollow(followerID int64, followID int64) error {
	db := f.db.Where("follower_id = ? AND follow_id = ?", followerID, followID).Delete(&entity.Follow{})
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (f *followRepositoryImpl) HasFollow(followerID int64, followID int64) (bool, error) {
	count := int64(0)
	db := f.db.Model(&entity.Follow{}).Where("follower_id = ? AND follow_id = ?", followerID, followID).Count(&count)
	if db.Error != nil {
		return false, db.Error
	} else if count < 0 || count > 1 {
		return false, service.ErrInternalServerError
	}
	return count == 1, nil
}
