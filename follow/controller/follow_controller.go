package follow_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhihaop/ticktok/core"
	"github.com/zhihaop/ticktok/entity"
	"net/http"
	"strconv"
)

const (
	ActionFollow   = "1"
	ActionUnfollow = "2"
)

// FollowController handles all the request mapping to '/douyin/relation'
type FollowController struct {
	UserService   entity.UserService
	FollowService entity.FollowService
}

type followListResponse struct {
	core.Response
	UserList []entity.UserInfo `json:"user_list"`
}

// NewFollowController creates an instance of FollowController
func NewFollowController(userService entity.UserService, followService entity.FollowService) *FollowController {
	return &FollowController{
		UserService:   userService,
		FollowService: followService,
	}
}

// InitRouter register handlers to gin.RouterGroup
func (u *FollowController) InitRouter(g *gin.RouterGroup) {
	g.POST("/action/", u.Action)
	g.GET("/follow/list/", u.ListFollow)
	g.GET("/follower/list/", u.ListFollower)
}

func (u *FollowController) Action(g *gin.Context) {
	token := g.Query("token")
	sFollowID := g.Query("to_user_id")
	sAction := g.Query("action_type")

	followID, err := strconv.ParseInt(sFollowID, 10, 64)
	if err != nil {
		g.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	followerID, err := u.UserService.GetUserID(token)
	if err != nil {
		g.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	switch sAction {
	case ActionFollow:
		err := u.FollowService.Follow(followerID, followID)
		if err != nil {
			g.JSON(http.StatusOK, core.ResponseError(err))
		}
	case ActionUnfollow:
		err := u.FollowService.UnFollow(followerID, followID)
		if err != nil {
			g.JSON(http.StatusOK, core.ResponseError(err))
		}
	default:
		g.JSON(http.StatusOK, core.ResponseError(core.ErrActionInValid))
		return
	}

	g.JSON(http.StatusOK, core.ResponseOK())
}

func (u *FollowController) ListFollow(g *gin.Context) {
	sUserID := g.Query("user_id")
	token := g.Query("token")

	userID, err := strconv.ParseInt(sUserID, 10, 64)
	if err != nil {
		g.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	id, err := u.UserService.GetUserID(token)
	if err != nil {
		g.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	if userID == 0 {
		userID = id
	}

	follows, err := u.FollowService.ListFollow(userID)
	if err != nil {
		g.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	g.JSON(http.StatusOK, &followListResponse{
		Response: core.ResponseOK(),
		UserList: follows,
	})
}

func (u *FollowController) ListFollower(g *gin.Context) {
	sUserID := g.Query("user_id")
	token := g.Query("token")

	userID, err := strconv.ParseInt(sUserID, 10, 64)
	if err != nil {
		g.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	id, err := u.UserService.GetUserID(token)
	if err != nil {
		g.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	if userID == 0 {
		userID = id
	}

	followers, err := u.FollowService.ListFollower(userID)
	if err != nil {
		g.JSON(http.StatusOK, core.ResponseError(err))
		return
	}

	g.JSON(http.StatusOK, &followListResponse{
		Response: core.ResponseOK(),
		UserList: followers,
	})
}
