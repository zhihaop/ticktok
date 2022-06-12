package entity

// Follow is a row in `follows` table
type Follow struct {
	FollowerID int64 `gorm:"primaryKey"`
	FollowID   int64 `gorm:"primaryKey"`
}

// FollowService represents the user's follow service
type FollowService interface {
	Follow(followerID int64, followID int64) error
	UnFollow(followerID int64, followID int64) error
	ListFollow(userID int64) ([]UserInfo, error)
	ListFollower(userID int64) ([]UserInfo, error)
	GetFollowerCount(userID int64) (int64, error)
	GetFollowCount(userID int64) (int64, error)
	HasFollow(followerID int64, followID int64) (bool, error)
}

// FollowRepository represents the user's follow repository
type FollowRepository interface {
	CountFollowerByID(followID int64) (int64, error)
	CountFollowByID(followerID int64) (int64, error)
	InsertFollow(followerID int64, followingID int64) error
	DeleteFollow(followerID int64, followingID int64) error
	HasFollow(followerID int64, followID int64) (bool, error)
	FetchFollow(followerID int64, offset int64, limit int64) ([]Follow, error)
	FetchFollower(followID int64, offset int64, limit int64) ([]Follow, error)
}
