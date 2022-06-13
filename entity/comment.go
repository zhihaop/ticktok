package entity

import (
	"time"
)

// Comment is a row in `Comments` table
type Comment struct {
	ID       int64 `gorm:"autoIncrement"`
	UserID   int64
	ClipID   int64
	Content  string
	CreateAt time.Time
	UpdateAt time.Time
}
