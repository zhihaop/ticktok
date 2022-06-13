package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zhihaop/ticktok/clip/controller"
	"github.com/zhihaop/ticktok/clip/repository"
	"github.com/zhihaop/ticktok/clip/service"
	"github.com/zhihaop/ticktok/user/controller"
	"github.com/zhihaop/ticktok/user/repository"
	"github.com/zhihaop/ticktok/user/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}

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
	followRepository := user_repository.NewFollowRepository(db)
	publishRepository := clip_repository.NewPublishRepository(db)

	// initialize services
	userService := user_service.NewUserService(userRepository, followRepository)
	followService := user_service.NewFollowService(followRepository, userRepository)
	publishService := clip_service.NewClipService(publishRepository, userService)

	// initialize controllers
	userController := user_controller.NewUserController(userService, followService)
	followController := user_controller.NewFollowController(userService, followService)
	publishController := clip_controller.NewPublishController(publishService, userService)
	feedController := clip_controller.NewFeedController(userService, publishService)

	// TODO initialize other routers
	// initialize routers
	userController.InitRouter(engine.Group("/douyin/user"))
	followController.InitRouter(engine.Group("/douyin/relation"))
	publishController.InitRouter(engine.Group("/douyin/publish"))
	feedController.InitRouter(engine.Group("/douyin/feed"))

	// static resources router
	engine.Static("/douyin/static/", "./resources")

	// listen on 0.0.0.0:8080
	if err := engine.Run(viper.GetString("server.address")); err != nil {
		log.Fatalln(err)
	}
}
