package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zhihaop/ticktok/clip/controller"
	"github.com/zhihaop/ticktok/clip/repository"
	"github.com/zhihaop/ticktok/clip/service"
	"github.com/zhihaop/ticktok/favourite/controller"
	"github.com/zhihaop/ticktok/favourite/repository"
	"github.com/zhihaop/ticktok/favourite/service"
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
	clipRepository := clip_repository.NewClipRepository(db)
	favouriteRepository := favourite_repository.NewFavouriteRepository(db)

	// initialize services
	userService := user_service.NewUserService(userRepository, followRepository)
	followService := user_service.NewFollowService(followRepository, userRepository)
	clipService := clip_service.NewClipService(clipRepository, userService)
	favouriteService := favourite_service.NewFavouriteService(favouriteRepository, clipService)

	// initialize controllers
	userController := user_controller.NewUserController(userService, followService)
	followController := user_controller.NewFollowController(userService, followService)
	publishController := clip_controller.NewPublishController(clipService, userService)
	feedController := clip_controller.NewFeedController(userService, clipService)
	favouriteController := favourite_controller.NewFavouriteController(favouriteService, userService)

	// TODO initialize other routers
	// initialize routers
	userController.InitRouter(engine.Group("/douyin/user"))
	followController.InitRouter(engine.Group("/douyin/relation"))
	publishController.InitRouter(engine.Group("/douyin/publish"))
	feedController.InitRouter(engine.Group("/douyin/feed"))
	favouriteController.InitRouter(engine.Group("/douyin/favorite"))

	// static resources router
	engine.Static("/douyin/static/", "./resources")

	// listen on 0.0.0.0:8080
	if err := engine.Run(viper.GetString("server.address")); err != nil {
		log.Fatalln(err)
	}
}
