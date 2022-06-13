package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhihaop/ticktok/core/controller"
	"github.com/zhihaop/ticktok/core/service"
	"github.com/zhihaop/ticktok/entity"
	"net/http"
	"strconv"
)

const (
	ActionFollow   = "1"
	ActionUnfollow = "2"
)

// followController handles all the request mapping to '/douyin/relation'
type followController struct {
	controller.Controller
	UserService   entity.UserService
	FollowService entity.FollowService
}

type followListResponse struct {
	controller.Response
	UserList []entity.UserInfo `json:"user_list"`
}

// NewFollowController creates an instance of followController
func NewFollowController(userService entity.UserService, followService entity.FollowService) controller.Controller {
	return &followController{
		UserService:   userService,
		FollowService: followService,
	}
}

// InitRouter register handlers to gin.RouterGroup
func (u *followController) InitRouter(g *gin.RouterGroup) {
	g.POST("/action/", u.Action)
	g.GET("/follow/list/", u.ListFollow)
	g.GET("/follower/list/", u.ListFollower)
}

func (u *followController) Action(g *gin.Context) {
	token := g.Query("token")
	sFollowID := g.Query("to_user_id")
	sAction := g.Query("action_type")

	followID, err := strconv.ParseInt(sFollowID, 10, 64)
	if err != nil {
		g.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	followerID, err := u.UserService.GetUserID(token)
	if err != nil {
		g.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	switch sAction {
	case ActionFollow:
		err := u.FollowService.Follow(followerID, followID)
		if err != nil {
			g.JSON(http.StatusOK, controller.ResponseError(err))
		}
	case ActionUnfollow:
		err := u.FollowService.UnFollow(followerID, followID)
		if err != nil {
			g.JSON(http.StatusOK, controller.ResponseError(err))
		}
	default:
		g.JSON(http.StatusOK, controller.ResponseError(service.ErrActionInValid))
		return
	}

	g.JSON(http.StatusOK, controller.ResponseOK())
}

func (u *followController) ListFollow(g *gin.Context) {
	sUserID := g.Query("user_id")
	token := g.Query("token")

	userID, err := strconv.ParseInt(sUserID, 10, 64)
	if err != nil {
		g.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	id, err := u.UserService.GetUserID(token)
	if err != nil {
		g.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	if userID == 0 {
		userID = id
	}

	follows, err := u.FollowService.ListFollow(id, userID)
	if err != nil {
		g.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	g.JSON(http.StatusOK, &followListResponse{
		Response: controller.ResponseOK(),
		UserList: follows,
	})
}

func (u *followController) ListFollower(g *gin.Context) {
	sUserID := g.Query("user_id")
	token := g.Query("token")

	userID, err := strconv.ParseInt(sUserID, 10, 64)
	if err != nil {
		g.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	id, err := u.UserService.GetUserID(token)
	if err != nil {
		g.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	if userID == 0 {
		userID = id
	}

	followers, err := u.FollowService.ListFollower(id, userID)
	if err != nil {
		g.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	g.JSON(http.StatusOK, &followListResponse{
		Response: controller.ResponseOK(),
		UserList: followers,
	})
}
