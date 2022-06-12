package controller

import "github.com/gin-gonic/gin"

// Controller interface of Ticktok project
type Controller interface {
	InitRouter(g *gin.RouterGroup)
}
