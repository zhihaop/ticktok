package entity

import "io"

// Video is a row in `videos` table
type Video struct {
	ID        int64 `gorm:"autoIncrement"`
	UserID    int64
	Title     string
	VideoUUID string
	CoverUUID string
}

// VideoInfo is the video information required by PublishController
type VideoInfo struct {
	Author        UserInfo `json:"author"`
	CommentCount  int64    `json:"comment_count"`
	CoverURL      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	ID            int64    `json:"id"`
	IsFavorite    bool     `json:"is_favorite"`
	PlayURL       string   `json:"play_url"`
	Title         string   `json:"title"`
}

// PublishService represents the service for video publishing
type PublishService interface {
	Publish(userID int64, title string, dataLength int64, reader io.Reader) error
	List(userID int64) ([]VideoInfo, error)
}

// PublishRepository represents the repository for video publishing
type PublishRepository interface {
	FetchByID(userID int64) ([]Video, error)
	Save(video Video) error
	HasUUID(uuid string) (bool, error)
}
