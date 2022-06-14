package clip_repository

import (
	"github.com/zhihaop/ticktok/core/repository"
	"github.com/zhihaop/ticktok/entity"
	"gorm.io/gorm"
	"log"
	"time"
)

// publishRepositoryImpl is an implementation of publishRepository
type publishRepositoryImpl struct {
	db *gorm.DB
}

func (p *publishRepositoryImpl) GetByID(clipID int64) (*entity.Clip, error) {
	clip := make([]entity.Clip, 0)
	if err := p.db.Model(&entity.Clip{}).Where("id = ?", clipID).Find(&clip).Error; err != nil {
		return nil, err
	}
	return &clip[0], nil
}

func (p *publishRepositoryImpl) HasClip(clipID int64) (bool, error) {
	count := int64(0)
	if err := p.db.Model(&entity.Clip{}).Where("id = ?", clipID).Count(&count).Error; err != nil {
		return false, err
	}
	return count != 0, nil
}

func (p *publishRepositoryImpl) FetchByID(userID int64, limit int, offset time.Time) ([]entity.Clip, error) {
	clips := make([]entity.Clip, 0)
	table := p.db.Model(&entity.Clip{})
	result := table.Where("user_id = ? AND create_at <= ?", userID, offset).Find(&clips).Limit(limit)
	if result.Error != nil {
		return nil, result.Error
	}
	return clips, nil
}

func (p *publishRepositoryImpl) Fetch(limit int, offset time.Time) ([]entity.Clip, error) {
	clips := make([]entity.Clip, 0)
	table := p.db.Model(&entity.Clip{})
	result := table.Where("create_at <= ?", offset).Find(&clips).Limit(limit)
	if result.Error != nil {
		return nil, result.Error
	}
	return clips, nil
}

func (p *publishRepositoryImpl) HasUUID(uuid string) (bool, error) {
	count := int64(0)
	query := p.db.Model(&entity.Clip{}).Where("Clip_uuid = ? OR cover_uuid = ?", uuid, uuid).Count(&count)
	if query.Error != nil {
		return false, query.Error
	}
	return count != 0, nil
}

func (p *publishRepositoryImpl) Save(Clip *entity.Clip) error {
	return p.db.Save(&Clip).Error
}

func NewClipRepository(db *gorm.DB) entity.ClipRepository {
	if err := repository.CheckOrCreateTable(db, &entity.Clip{}); err != nil {
		log.Fatalln(err)
	}
	return &publishRepositoryImpl{db: db}
}
