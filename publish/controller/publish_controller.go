package publish_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhihaop/ticktok/core/controller"
	"github.com/zhihaop/ticktok/core/service"
	"github.com/zhihaop/ticktok/entity"
	"net/http"
	"strconv"
)

type PublishController struct {
	controller.Controller
	PublishService entity.PublishService
	UserService    entity.UserService
}

type VideoInfoResponse struct {
	controller.Response
	VideoList []entity.VideoInfo `json:"video_list"`
}

func (p *PublishController) InitRouter(g *gin.RouterGroup) {
	g.POST("/action/", p.Action)
	g.GET("/list/", p.List)
}

func NewPublishController(publishService entity.PublishService, userService entity.UserService) controller.Controller {
	return &PublishController{
		PublishService: publishService,
		UserService:    userService,
	}
}

func (p *PublishController) Action(c *gin.Context) {
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

func (p *PublishController) List(c *gin.Context) {
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

	c.JSON(http.StatusOK, VideoInfoResponse{
		Response:  controller.ResponseOK(),
		VideoList: result,
	})
}
