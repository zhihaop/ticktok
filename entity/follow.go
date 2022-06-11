package entity

import "time"

// Follow is a row in `Follows` table
type Follow struct {
	ID         int64 `gorm:"autoIncrement"`
	FollowerID int64
	FollowID   int64
	CreateAt   time.Time
}

// FollowRepository represents the user's follow repository
type FollowRepository interface {
	CountFollowerByID(followID int64) (int64, error)
	CountFollowByID(followerID int64) (int64, error)
	FetchFollowerByID(followID int64, offset int, limit int) ([]Follow, error)
	FetchFollowByID(followerID int64, offset int, limit int) ([]Follow, error)
	InsertFollow(followerID int64, followingID int64) error
	DeleteFollow(followerID int64, followingID int64) error
	HasFollow(followerID int64, followID int64) (bool, error)
}
