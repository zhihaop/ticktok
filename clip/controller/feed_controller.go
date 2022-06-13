package clip_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhihaop/ticktok/core/controller"
	"github.com/zhihaop/ticktok/entity"
	"net/http"
	"strconv"
	"time"
)

type feedControllerImpl struct {
	userService    entity.UserService
	publishService entity.ClipService
}

func NewFeedController(userService entity.UserService, publishService entity.ClipService) controller.Controller {
	return &feedControllerImpl{
		userService:    userService,
		publishService: publishService,
	}
}

type feedResponse struct {
	controller.Response
	VideoList []entity.ClipInfo `json:"video_list,omitempty"`
}

func (f *feedControllerImpl) InitRouter(g *gin.RouterGroup) {
	g.GET("/", f.Feed)
}

func (f *feedControllerImpl) Feed(c *gin.Context) {
	sLatestTime := c.Query("latest_time")
	token := c.Query("token")

	latestTime := time.Now()
	if len(sLatestTime) != 0 {
		timestamp, err := strconv.ParseInt(sLatestTime, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, controller.ResponseError(err))
			return
		}
		latestTime = time.UnixMilli(timestamp * 1000)
	}

	var userID *int64 = nil
	if len(token) != 0 {
		result, err := f.userService.GetUserID(token)
		if err != nil {
			c.JSON(http.StatusOK, controller.ResponseError(err))
			return
		}
		userID = &result
	}

	fetch, err := f.publishService.Fetch(userID, 30, latestTime)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, feedResponse{
		Response:  controller.ResponseOK(),
		VideoList: fetch,
	})
}
