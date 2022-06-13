package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhihaop/ticktok/follow/controller"
	"github.com/zhihaop/ticktok/follow/repository"
	"github.com/zhihaop/ticktok/follow/service"
	"github.com/zhihaop/ticktok/publish/controller"
	"github.com/zhihaop/ticktok/publish/repository"
	"github.com/zhihaop/ticktok/publish/service"
	"github.com/zhihaop/ticktok/user/controller"
	"github.com/zhihaop/ticktok/user/repository"
	"github.com/zhihaop/ticktok/user/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	// using sqlite as database
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// create a gin engine
	engine := gin.Default()

	// TODO initialize other domains
	// initialize repositories
	userRepository := user_repository.NewUserRepository(db)
	followRepository := follow_repository.NewFollowRepository(db)
	publishRepository := publish_repository.NewPublishRepository(db)

	// initialize services
	userService := user_service.NewUserService(userRepository, followRepository)
	followService := follow_service.NewFollowService(followRepository, userRepository)
	publishService := publish_service.NewPublishService(publishRepository, userService)

	// initialize controllers
	userController := user_controller.NewUserController(userService, followService)
	followController := follow_controller.NewFollowController(userService, followService)
	publishController := publish_controller.NewPublishController(publishService, userService)

	// TODO initialize other routers
	// initialize routers
	userController.InitRouter(engine.Group("/douyin/user"))
	followController.InitRouter(engine.Group("/douyin/relation"))
	publishController.InitRouter(engine.Group("/douyin/publish"))

	// listen on 0.0.0.0:8080
	if err := engine.Run(); err != nil {
		log.Fatalln(err)
	}
}
