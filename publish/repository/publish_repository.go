package publish_repository

import (
	"github.com/zhihaop/ticktok/core/repository"
	"github.com/zhihaop/ticktok/entity"
	"gorm.io/gorm"
	"log"
)

// PublishRepositoryImpl is an implementation of PublishRepository
type PublishRepositoryImpl struct {
	db *gorm.DB
}

func (p *PublishRepositoryImpl) HasUUID(uuid string) (bool, error) {
	count := int64(0)
	query := p.db.Model(&entity.Video{}).Where("video_uuid = ? OR cover_uuid = ?", uuid, uuid).Count(&count)
	if query.Error != nil {
		return false, query.Error
	}
	return count != 0, nil
}

func (p *PublishRepositoryImpl) FetchByID(userID int64) ([]entity.Video, error) {
	videos := make([]entity.Video, 0)
	result := p.db.Model(&entity.Video{}).Where("user_id = ?", userID).Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	return videos, nil
}

func (p *PublishRepositoryImpl) Save(video entity.Video) error {
	return p.db.Save(&video).Error
}

func NewPublishRepository(db *gorm.DB) entity.PublishRepository {
	if err := repository.CheckOrCreateTable(db, &entity.Video{}); err != nil {
		log.Fatalln(err)
	}
	return &PublishRepositoryImpl{db: db}
}
