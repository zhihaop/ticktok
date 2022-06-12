package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhihaop/ticktok/core"
	"github.com/zhihaop/ticktok/entity"
	"net/http"
	"strconv"
)

// UserController handles all the request mapping to '/douyin/user'
type UserController struct {
	UserService     entity.UserService
	FollowerService entity.FollowService
}

// UserLoginResponse is the response type for '/douyin/user/login' api
type UserLoginResponse struct {
	core.Response
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

// UserInfoResponse is the response type for '/douyin/user' api
type UserInfoResponse struct {
	core.Response
	entity.UserInfo
}

// NewUserController creates an instance of UserController
func NewUserController(userService entity.UserService, followService entity.FollowService) *UserController {
	return &UserController{
		UserService:     userService,
		FollowerService: followService,
	}
}

// InitRouter register handlers to gin.RouterGroup
func (u *UserController) InitRouter(g *gin.RouterGroup) {
	g.POST("/register/", u.Register)
	g.POST("/login/", u.Login)
	g.GET("/", u.Info)
}

func (u *UserController) Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, err := u.UserService.Register(username, password)

	if err != nil {
		c.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: core.ResponseOK(),
		UserID:   token.ID,
		Token:    token.Token,
	})
}

func (u *UserController) Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, err := u.UserService.Login(username, password)

	if err != nil {
		c.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: core.ResponseOK(),
		UserID:   token.ID,
		Token:    token.Token,
	})
}

func (u *UserController) Info(c *gin.Context) {
	token := c.Query("token")
	stringID := c.Query("user_id")

	id, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	userID, err := u.UserService.GetUserID(token)
	if err != nil {
		c.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	username, err := u.UserService.GetUsername(id)
	if err != nil {
		c.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	followerCount, err := u.FollowerService.GetFollowerCount(id)
	if err != nil {
		c.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	followCount, err := u.FollowerService.GetFollowCount(id)
	if err != nil {
		c.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	hasFollow, err := u.FollowerService.HasFollow(userID, id)
	if err != nil {
		c.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, UserInfoResponse{
		Response: core.ResponseOK(),
		UserInfo: entity.UserInfo{
			ID:            userID,
			Name:          username,
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      hasFollow,
		},
	})
}
