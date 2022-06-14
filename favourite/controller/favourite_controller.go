package favourite_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhihaop/ticktok/core/controller"
	"github.com/zhihaop/ticktok/entity"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const (
	actionFavourite     = "1"
	actionUndoFavourite = "2"
)

type favouriteControllerImpl struct {
	favouriteService entity.FavouriteService
	userService      entity.UserService
}

type favouriteListResponse struct {
	controller.Response
	VideoList []entity.ClipInfo `json:"video_list,omitempty"`
}

func (f *favouriteControllerImpl) InitRouter(g *gin.RouterGroup) {
	g.POST("/action/", f.Action)
	g.GET("/list/", f.List)
}

func NewFavouriteController(favouriteService entity.FavouriteService, userService entity.UserService) controller.Controller {
	return &favouriteControllerImpl{
		favouriteService: favouriteService,
		userService:      userService,
	}
}

func (f *favouriteControllerImpl) List(c *gin.Context) {
	token := c.Query("token")
	sUserID := c.Query("user_id")

	queryID, err := strconv.ParseInt(sUserID, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	userID, err := f.userService.GetUserID(token)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	favourite, err := f.favouriteService.ListFavourite(userID, queryID)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, favouriteListResponse{
		Response:  controller.ResponseOK(),
		VideoList: favourite,
	})
}

func (f *favouriteControllerImpl) Action(c *gin.Context) {
	token := c.Query("token")
	sVideoID := c.Query("video_id")
	sAction := c.Query("action_type")

	videoID, err := strconv.ParseInt(sVideoID, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	userID, err := f.userService.GetUserID(token)
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	switch sAction {
	case actionFavourite:
		err = f.favouriteService.Favourite(userID, videoID)
	case actionUndoFavourite:
		err = f.favouriteService.UndoFavourite(userID, videoID)
	default:
		c.JSON(http.StatusOK, controller.ResponseError(gorm.ErrInvalidData))
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, controller.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, controller.ResponseOK())
}
