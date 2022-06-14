package favourite_repository

import (
	"github.com/zhihaop/ticktok/core/repository"
	"github.com/zhihaop/ticktok/entity"
	"gorm.io/gorm"
	"log"
	"time"
)

type favouriteRepositoryImpl struct {
	db *gorm.DB
}

func (f favouriteRepositoryImpl) HasFavourite(userID int64, clipID int64) (bool, error) {
	count := int64(0)
	db := f.db.Model(&entity.Favourite{}).Where("user_id = ? AND clip_id = ?", userID, clipID).Count(&count)
	if db.Error != nil {
		return false, db.Error
	}
	return count != 0, nil
}

func (f favouriteRepositoryImpl) Favourite(userID int64, clipID int64) error {
	model := f.db.Save(&entity.Favourite{
		UserID:   userID,
		ClipID:   clipID,
		CreateAt: time.Now(),
	})
	return model.Error
}

func (f favouriteRepositoryImpl) UndoFavourite(userID int64, clipID int64) error {
	db := f.db
	db = db.Where("user_id = ? AND clip_id = ?", userID, clipID).Delete(&entity.Favourite{})
	return db.Error
}

func (f favouriteRepositoryImpl) FetchFavouriteByUserID(userID int64) ([]entity.Favourite, error) {
	favourites := make([]entity.Favourite, 0)
	find := f.db.Where("user_id = ?", userID).Find(&favourites)
	if find.Error != nil {
		return nil, find.Error
	}
	return favourites, nil
}

func NewFavouriteRepository(db *gorm.DB) entity.FavouriteRepository {
	if err := repository.CheckOrCreateTable(db, &entity.Favourite{}); err != nil {
		log.Fatalln(err)
	}
	return &favouriteRepositoryImpl{db}
}
