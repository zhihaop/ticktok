package repository

import (
	"github.com/zhihaop/ticktok/core"
	utils "github.com/zhihaop/ticktok/core/repository"
	"github.com/zhihaop/ticktok/entity"
	"gorm.io/gorm"
	"log"
)

// FollowRepositoryImpl is an implementation of FollowRepository
type FollowRepositoryImpl struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) entity.FollowRepository {
	if err := utils.CheckOrCreateTable(db, &entity.Follow{}); err != nil {
		log.Fatalln(err)
	}
	return &FollowRepositoryImpl{db: db}
}

func (f *FollowRepositoryImpl) CountFollowerByID(followID int64) (int64, error) {
	followerCount := int64(0)
	db := f.db.Where("followID = ?", followID).Count(&followerCount)
	if db.Error != nil {
		return -1, db.Error
	}

	return followerCount, nil
}

func (f *FollowRepositoryImpl) CountFollowByID(followerID int64) (int64, error) {
	followCount := int64(0)
	db := f.db.Where("followerID = ?", followerID).Count(&followCount)
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
	db := f.db.Delete(&entity.Follow{
		FollowerID: followerID,
		FollowID:   followID,
	}, 1)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (f *FollowRepositoryImpl) HasFollow(followerID int64, followID int64) (bool, error) {
	count := int64(0)
	db := f.db.Where("followerID = ? AND followID = ?", followerID, followID).Count(&count)
	if db.Error != nil {
		return false, db.Error
	}

	if count < 0 || count > 1 {
		return false, core.ErrInternalServerError
	}

	return count == 1, nil
}
