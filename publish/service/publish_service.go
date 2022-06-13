package publish_service

import (
	"github.com/zhihaop/ticktok/core"
	"github.com/zhihaop/ticktok/core/service"
	"github.com/zhihaop/ticktok/entity"
	"io"
	"log"
	"os"
	"path/filepath"
)

type PublishServiceImpl struct {
	PublishRepository entity.PublishRepository
	UserService       entity.UserService
}

func NewPublishService(publishRepository entity.PublishRepository, userService entity.UserService) entity.PublishService {
	return &PublishServiceImpl{PublishRepository: publishRepository, UserService: userService}
}

func (p *PublishServiceImpl) Store(uuid string, dataLength int64, reader io.Reader) error {
	file, err := os.OpenFile(filepath.Join("resources", uuid), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
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
			return err
		}
		if read == 0 {
			break
		}
		_, err = file.Write(buffer[0:read])
		if err != nil {
			return err
		}
		dataLength = dataLength - int64(read)
	}

	if dataLength != 0 {
		return service.ErrVideoFileInValid
	}
	return nil
}

func (p *PublishServiceImpl) GetUUID() (string, error) {
	for retry := 0; retry < 16; retry++ {
		uuid := core.GetUUID()
		exist, err := p.PublishRepository.HasUUID(uuid)
		if err != nil {
			return "", service.ErrInternalServerError
		} else if !exist {
			return uuid, nil
		}
	}
	return "", service.ErrInternalServerError
}

func (p *PublishServiceImpl) Publish(userID int64, title string, dataLength int64, reader io.Reader) error {
	uuid, err := p.GetUUID()
	if err != nil {
		return err
	}

	if err := p.Store(uuid, dataLength, reader); err != nil {
		return err
	}

	if err := p.PublishRepository.Save(entity.Video{
		UserID:    userID,
		Title:     title,
		VideoUUID: uuid,
		CoverUUID: uuid,
	}); err != nil {
		return service.ErrInternalServerError
	}

	return nil
}

func (p *PublishServiceImpl) List(userID int64) ([]entity.VideoInfo, error) {
	videos, err := p.PublishRepository.FetchByID(userID)
	if err != nil {
		return nil, service.ErrInternalServerError
	}

	infos := make([]entity.VideoInfo, len(videos))
	for i := range infos {
		infos[i].ID = videos[i].ID

		// TODO implement uuid to url
		infos[i].PlayURL = videos[i].VideoUUID
		infos[i].CoverURL = videos[i].CoverUUID

		// TODO implement comment repository
		infos[i].CommentCount = 0

		// TODO implement favourite repository
		infos[i].FavoriteCount = 0
		infos[i].IsFavorite = false

		author, err := p.UserService.GetUserInfo(userID, videos[i].UserID)
		if err != nil {
			return nil, service.ErrInternalServerError
		}
		infos[i].Author = *author
	}
	return infos, nil
}
