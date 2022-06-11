package entity

// Video is a row in `Videos` table
type Video struct {
	ID       int64 `gorm:"autoIncrement"`
	UserID   int64
	Title    string
	VideoUrl string
	CoverUrl string
}
