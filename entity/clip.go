package entity

import (
	"io"
	"time"
)

// Clip is a row in `clips` table
type Clip struct {
	ID        int64 `gorm:"autoIncrement"`
	UserID    int64
	Title     string
	ClipUUID  string
	CoverUUID string
	CreateAt  time.Time
}

// ClipInfo represents the clip information
type ClipInfo struct {
	Author        UserInfo `json:"author"`
	CommentCount  int64    `json:"comment_count"`
	CoverURL      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	ID            int64    `json:"id"`
	IsFavorite    bool     `json:"is_favorite"`
	PlayURL       string   `json:"play_url"`
	Title         string   `json:"title"`
}

// ClipService represents the service for clips publishing, clips fetching
type ClipService interface {
	Publish(userID int64, title string, dataLength int64, reader io.Reader) error
	List(userID int64) ([]ClipInfo, error)
	Fetch(userID *int64, limit int, offset time.Time) ([]ClipInfo, error)
}

// ClipRepository represents the repository for clips
type ClipRepository interface {
	FetchByID(userID int64, limit int, offset time.Time) ([]Clip, error)
	Fetch(limit int, offset time.Time) ([]Clip, error)
	Save(video *Clip) error
	HasUUID(uuid string) (bool, error)
}
