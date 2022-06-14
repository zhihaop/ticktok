package clip_service

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zhihaop/ticktok/core"
	"github.com/zhihaop/ticktok/core/service"
	"github.com/zhihaop/ticktok/entity"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"
)

const httpPrefix = "http"

type clipServiceImpl struct {
	clipRepository entity.ClipRepository
	userService    entity.UserService
	address        string
}

func (p *clipServiceImpl) GetByID(userID *int64, clipID int64) (*entity.ClipInfo, error) {
	clip, err := p.clipRepository.GetByID(clipID)
	if err != nil {
		return nil, service.ErrInternalServerError
	} else if clip == nil {
		return nil, nil
	}

	infos, err := p.GetVideoInfos(userID, []entity.Clip{*clip})
	if err != nil {
		return nil, err
	} else if len(infos) == 0 {
		return nil, nil
	}
	return &infos[0], err
}

func NewClipService(clipRepository entity.ClipRepository, userService entity.UserService) entity.ClipService {
	return &clipServiceImpl{
		clipRepository: clipRepository,
		userService:    userService,
		address:        viper.GetString("server.address"),
	}
}

func (p *clipServiceImpl) HasClip(clipID int64) (bool, error) {
	exist, err := p.clipRepository.HasClip(clipID)
	if err != nil {
		return false, service.ErrInternalServerError
	}
	return exist, nil
}

func (p *clipServiceImpl) Store(uuid string, dataLength int64, reader io.Reader) error {
	file, err := os.OpenFile(filepath.Join("resources", uuid), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return service.ErrInternalServerError
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	buffer := make([]byte, 1024)
	for {
		read, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			return service.ErrInternalServerError
		}
		if read == 0 {
			break
		}
		_, err = file.Write(buffer[0:read])
		if err != nil {
			return service.ErrInternalServerError
		}
		dataLength = dataLength - int64(read)
	}

	if dataLength != 0 {
		return service.ErrVideoFileInValid
	}
	return nil
}

func (p *clipServiceImpl) GetUUID() (string, error) {
	for retry := 0; retry < 16; retry++ {
		uuid := core.GetUUID()
		exist, err := p.clipRepository.HasUUID(uuid)
		if err != nil {
			return "", service.ErrInternalServerError
		} else if !exist {
			return uuid, nil
		}
	}
	return "", service.ErrInternalServerError
}

func (p *clipServiceImpl) Publish(userID int64, title string, dataLength int64, reader io.Reader) error {
	uuid, err := p.GetUUID()
	if err != nil {
		return err
	}

	if err := p.Store(uuid, dataLength, reader); err != nil {
		return err
	}

	if err := p.clipRepository.Save(&entity.Clip{
		UserID:    userID,
		Title:     title,
		ClipUUID:  uuid,
		CoverUUID: uuid,
		CreateAt:  time.Now(),
	}); err != nil {
		return service.ErrInternalServerError
	}

	return nil
}

func (p *clipServiceImpl) GetVideoInfos(userID *int64, clips []entity.Clip) ([]entity.ClipInfo, error) {
	infos := make([]entity.ClipInfo, len(clips))
	for i := range infos {
		infos[i].ID = clips[i].ID

		infos[i].PlayURL = fmt.Sprintf("%s://%s/douyin/static/%s", httpPrefix, p.address, clips[i].ClipUUID)
		infos[i].CoverURL = fmt.Sprintf("%s://%s/douyin/static/%s", httpPrefix, p.address, clips[i].ClipUUID)

		// TODO implement comment repository
		infos[i].CommentCount = 0

		// TODO implement favourite repository
		infos[i].FavoriteCount = 0
		infos[i].IsFavorite = false

		author, err := p.userService.GetUserInfo(userID, clips[i].UserID)
		if err != nil {
			return nil, service.ErrInternalServerError
		}
		infos[i].Author = *author
	}
	return infos, nil
}

func (p *clipServiceImpl) List(userID int64) ([]entity.ClipInfo, error) {
	clips, err := p.clipRepository.FetchByID(userID, math.MaxInt, time.Now())
	if err != nil {
		return nil, service.ErrInternalServerError
	}

	return p.GetVideoInfos(&userID, clips)
}

func (p *clipServiceImpl) Fetch(userID *int64, limit int, offset time.Time) ([]entity.ClipInfo, error) {
	clips, err := p.clipRepository.Fetch(limit, offset)
	if err != nil {
		return nil, service.ErrInternalServerError
	}

	return p.GetVideoInfos(userID, clips)
}
