package entity

import "time"

// Favourite is a row in `Favourites` table
type Favourite struct {
	UserID   int64 `gorm:"primaryKey"`
	ClipID   int64 `gorm:"primaryKey"`
	CreateAt time.Time
}

type FavouriteService interface {
	Favourite(userID int64, clipID int64) error
	UndoFavourite(userID int64, clipID int64) error
	ListFavourite(userID int64, queryID int64) ([]ClipInfo, error)
}

type FavouriteRepository interface {
	HasFavourite(userID int64, clipID int64) (bool, error)
	Favourite(userID int64, clipID int64) error
	UndoFavourite(userID int64, clipID int64) error
	FetchFavouriteByUserID(userID int64) ([]Favourite, error)
}
