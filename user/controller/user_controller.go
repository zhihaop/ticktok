package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhihaop/ticktok/core/controller"
	"github.com/zhihaop/ticktok/entity"
	"net/http"
	"strconv"
)

// userController handles all the request mapping to '/douyin/user'
type userController struct {
	controller.Controller
	UserService     entity.UserService
	FollowerService entity.FollowService
}

// loginResponse is the response type for '/douyin/user/login' api
type loginResponse struct {
	controller.Response
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

// infoResponse is the response type for '/douyin/user' api
type infoResponse struct {
	controller.Response
	entity.UserInfo
}

// NewUserController creates an instance of userController
func NewUserController(userService entity.UserService, followService entity.FollowService) controller.Controller {
	return &userController{
		UserService:     userService,
		FollowerService: followService,
	}
}

// InitRouter register handlers to gin.RouterGroup
func (u *userController) InitRouter(g *gin.RouterGroup) {
	g.POST("/register/", u.Register)
	g.POST("/login/", u.Login)
	g.GET("/", u.Info)
}

func (u *userController) Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, err := u.UserService.Register(username, password)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, loginResponse{
		Response: controller.ResponseOK(),
		UserID:   token.ID,
		Token:    token.Token,
	})
}

func (u *userController) Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, err := u.UserService.Login(username, password)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, loginResponse{
		Response: controller.ResponseOK(),
		UserID:   token.ID,
		Token:    token.Token,
	})
}

func (u *userController) Info(c *gin.Context) {
	token := c.Query("token")
	sQueryID := c.Query("user_id")

	queryID, err := strconv.ParseInt(sQueryID, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	userID, err := u.UserService.GetUserID(token)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	userInfo, err := u.UserService.GetUserInfo(&userID, queryID)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, infoResponse{
		Response: controller.ResponseOK(),
		UserInfo: *userInfo,
	})
}
