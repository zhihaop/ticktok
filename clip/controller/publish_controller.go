package clip_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhihaop/ticktok/core/controller"
	"github.com/zhihaop/ticktok/core/service"
	"github.com/zhihaop/ticktok/entity"
	"net/http"
	"strconv"
)

type publishControllerImpl struct {
	controller.Controller
	PublishService entity.ClipService
	UserService    entity.UserService
}

type clipInfoResponse struct {
	controller.Response
	VideoList []entity.ClipInfo `json:"video_list"`
}

func (p *publishControllerImpl) InitRouter(g *gin.RouterGroup) {
	g.POST("/action/", p.Action)
	g.GET("/list/", p.List)
}

func NewPublishController(publishService entity.ClipService, userService entity.UserService) controller.Controller {
	return &publishControllerImpl{
		PublishService: publishService,
		UserService:    userService,
	}
}

func (p *publishControllerImpl) Action(c *gin.Context) {
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	} else if data == nil {
		c.JSON(http.StatusOK, controller.ResponseError(service.ErrVideoFileInValid))
		return
	}
	token := c.PostForm("token")
	title := c.PostForm("title")

	userID, err := p.UserService.GetUserID(token)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	reader, err := data.Open()
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	err = p.PublishService.Publish(userID, title, data.Size, reader)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, controller.ResponseOK())
}

func (p *publishControllerImpl) List(c *gin.Context) {
	token := c.Query("token")
	sUserID := c.Query("user_id")

	userID, err := strconv.ParseInt(sUserID, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	if _, err := p.UserService.GetUserID(token); err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	result, err := p.PublishService.List(userID)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, clipInfoResponse{
		Response:  controller.ResponseOK(),
		VideoList: result,
	})
}
