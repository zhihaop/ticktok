package favourite_service

import (
	"github.com/zhihaop/ticktok/entity"
	"gopkg.in/errgo.v2/errors"
)

type favouriteServiceImpl struct {
	favouriteRepository entity.FavouriteRepository
	clipService         entity.ClipService
}

func (f *favouriteServiceImpl) Favourite(userID int64, clipID int64) error {
	clip, err := f.clipService.HasClip(clipID)
	if err != nil {
		return err
	} else if !clip {
		return errors.New("clip not exist")
	}

	favourite, err := f.favouriteRepository.HasFavourite(userID, clipID)
	if err != nil {
		return err
	} else if favourite {
		return errors.New("has favourite")
	}

	return f.favouriteRepository.Favourite(userID, clipID)
}

func (f *favouriteServiceImpl) UndoFavourite(userID int64, clipID int64) error {
	clip, err := f.clipService.HasClip(clipID)
	if err != nil {
		return err
	} else if !clip {
		return errors.New("clip not exist")
	}

	favourite, err := f.favouriteRepository.HasFavourite(userID, clipID)
	if err != nil {
		return err
	} else if !favourite {
		return errors.New("favourite not exist")
	}

	return f.favouriteRepository.UndoFavourite(userID, clipID)
}

func (f *favouriteServiceImpl) ListFavourite(userID int64, queryID int64) ([]entity.ClipInfo, error) {
	favourite, err := f.favouriteRepository.FetchFavouriteByUserID(queryID)
	if err != nil {
		return nil, err
	}

	clips := make([]entity.ClipInfo, len(favourite))
	for i := range clips {
		clip, err := f.clipService.GetByID(&userID, favourite[i].ClipID)
		if err != nil {
			return clips, err
		}
		clips[i] = *clip
	}
	return clips, nil
}

func NewFavouriteService(favouriteRepository entity.FavouriteRepository, clipService entity.ClipService) entity.FavouriteService {
	return &favouriteServiceImpl{
		favouriteRepository: favouriteRepository,
		clipService:         clipService,
	}
}
