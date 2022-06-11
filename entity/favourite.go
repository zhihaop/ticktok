package entity

import "time"

// Favourite is a row in `Favourites` table
type Favourite struct {
	ID       int64 `gorm:"autoIncrement"`
	UserID   int64
	VideoID  int64
	CreateAt time.Time
}
