package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhihaop/ticktok/core"
	"github.com/zhihaop/ticktok/entity"
	"net/http"
	"strconv"
)

// UserController handles all the request mapping to '/douyin/user'
type UserController struct {
	UserService entity.UserService
}

// UserLoginResponse is the response type for '/douyin/user/login' api
type UserLoginResponse struct {
	core.Response
	UserID int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// UserInfoResponse is the response type for '/douyin/user' api
type UserInfoResponse struct {
	core.Response
	ID            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// NewUserController creates an instance of UserController
func NewUserController(userService entity.UserService) *UserController {
	return &UserController{UserService: userService}
}

// InitRouter register handlers to gin.RouterGroup
func (u *UserController) InitRouter(g *gin.RouterGroup) {
	g.POST("/register", u.Register)
	g.POST("/login", u.Login)
	g.POST("/", u.Info)
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
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
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

	followerCount, err := u.UserService.GetFollowerCount(id)
	if err != nil {
		c.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	followCount, err := u.UserService.GetFollowCount(id)
	if err != nil {
		c.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	isFollow, err := u.UserService.IsFollow(userID, id)
	if err != nil {
		c.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, UserInfoResponse{
		Response:      core.ResponseOK(),
		ID:            userID,
		Name:          username,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	})
}
