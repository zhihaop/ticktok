package follow_service

import (
	"github.com/zhihaop/ticktok/core/service"
	"github.com/zhihaop/ticktok/entity"
	"math"
)

// FollowServiceImpl is an implementation of FollowService
type FollowServiceImpl struct {
	entity.UserService
	FollowRepository entity.FollowRepository
	UserRepository   entity.UserRepository
}

func NewFollowService(followRepository entity.FollowRepository, userRepository entity.UserRepository) entity.FollowService {
	return &FollowServiceImpl{
		FollowRepository: followRepository,
		UserRepository:   userRepository,
	}
}

func (u *FollowServiceImpl) HasFollow(followerID int64, followID int64) (bool, error) {
	if followerID == followID {
		return false, nil
	}

	hasFollow, err := u.FollowRepository.HasFollow(followerID, followID)
	if err != nil {
		return false, service.ErrInternalServerError
	}
	return hasFollow, nil
}

func (u *FollowServiceImpl) checkUser(id int64) (bool, error) {
	user, err := u.UserRepository.FindByID(id)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}

func (u *FollowServiceImpl) ListFollow(userID int64) ([]entity.UserInfo, error) {
	follow, err := u.FollowRepository.FetchFollow(userID, 0, math.MaxInt64)
	if err != nil {
		return nil, err
	}

	users := make([]entity.UserInfo, len(follow))
	for i := range follow {
		user, err := u.UserRepository.FindByID(follow[i].FollowID)
		if err != nil {
			return nil, err
		}

		followCount, err := u.GetFollowCount(follow[i].FollowID)
		if err != nil {
			return nil, err
		}

		followerCount, err := u.GetFollowerCount(follow[i].FollowID)
		if err != nil {
			return nil, err
		}

		users[i] = entity.UserInfo{
			ID:            follow[i].FollowID,
			Name:          user.Name,
			FollowerCount: followerCount,
			FollowCount:   followCount,
			IsFollow:      true,
		}
	}
	return users, nil
}

func (u *FollowServiceImpl) ListFollower(userID int64) ([]entity.UserInfo, error) {
	follow, err := u.FollowRepository.FetchFollower(userID, 0, math.MaxInt64)
	if err != nil {
		return nil, err
	}

	users := make([]entity.UserInfo, len(follow))
	for i := range follow {
		user, err := u.UserRepository.FindByID(follow[i].FollowerID)
		if err != nil {
			return nil, err
		}

		followCount, err := u.GetFollowCount(follow[i].FollowerID)
		if err != nil {
			return nil, err
		}

		followerCount, err := u.GetFollowerCount(follow[i].FollowerID)
		if err != nil {
			return nil, err
		}

		hasFollow, err := u.HasFollow(userID, follow[i].FollowerID)
		if err != nil {
			return nil, err
		}

		users[i] = entity.UserInfo{
			ID:            follow[i].FollowerID,
			Name:          user.Name,
			FollowerCount: followerCount,
			FollowCount:   followCount,
			IsFollow:      hasFollow,
		}
	}
	return users, nil
}

func (u *FollowServiceImpl) Follow(followerID int64, followID int64) error {
	if followID == followerID {
		return service.ErrSelfFollowing
	}

	if valid, err := u.checkUser(followerID); err != nil {
		return service.ErrInternalServerError
	} else if !valid {
		return service.ErrUserNotExist
	}

	if valid, err := u.checkUser(followID); err != nil {
		return service.ErrInternalServerError
	} else if !valid {
		return service.ErrUserNotExist
	}

	hasFollow, err := u.FollowRepository.HasFollow(followerID, followID)
	if err != nil {
		return service.ErrInternalServerError
	} else if hasFollow {
		return service.ErrRelationExist
	}

	err = u.FollowRepository.InsertFollow(followerID, followID)
	if err != nil {
		return service.ErrInternalServerError
	}
	return nil
}

func (u *FollowServiceImpl) UnFollow(followerID int64, followID int64) error {
	if followID == followerID {
		return service.ErrSelfUnFollowing
	}

	if valid, err := u.checkUser(followerID); err != nil {
		return service.ErrInternalServerError
	} else if !valid {
		return service.ErrUserNotExist
	}

	if valid, err := u.checkUser(followID); err != nil {
		return service.ErrInternalServerError
	} else if !valid {
		return service.ErrUserNotExist
	}

	hasFollow, err := u.FollowRepository.HasFollow(followerID, followID)
	if err != nil {
		return service.ErrInternalServerError
	} else if !hasFollow {
		return service.ErrRelationNotExist
	}

	err = u.FollowRepository.DeleteFollow(followerID, followID)
	if err != nil {
		return service.ErrInternalServerError
	}
	return nil
}

func (u *FollowServiceImpl) GetFollowerCount(id int64) (int64, error) {
	followerCount, err := u.FollowRepository.CountFollowerByID(id)
	if err != nil {
		return 0, service.ErrInternalServerError
	}
	return followerCount, nil
}

func (u *FollowServiceImpl) GetFollowCount(id int64) (int64, error) {
	followingCount, err := u.FollowRepository.CountFollowByID(id)
	if err != nil {
		return 0, service.ErrInternalServerError
	}
	return followingCount, nil
}
