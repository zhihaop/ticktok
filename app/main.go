package main

import (
	"github.com/gin-gonic/gin"
	_followController "github.com/zhihaop/ticktok/follow/controller"
	_followRepository "github.com/zhihaop/ticktok/follow/repository"
	_followService "github.com/zhihaop/ticktok/follow/service"
	_userController "github.com/zhihaop/ticktok/user/controller"
	_userRepository "github.com/zhihaop/ticktok/user/repository"
	_userService "github.com/zhihaop/ticktok/user/service"
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
	userRepository := _userRepository.NewUserRepository(db)
	followRepository := _followRepository.NewFollowRepository(db)

	// initialize services
	userService := _userService.NewUserService(userRepository)
	followService := _followService.NewFollowService(followRepository, userRepository)

	// initialize controllers
	userController := _userController.NewUserController(userService, followService)
	followController := _followController.NewFollowController(userService, followService)

	// TODO initialize other routers
	// initialize routers
	userController.InitRouter(engine.Group("/douyin/user"))
	followController.InitRouter(engine.Group("/douyin/relation"))

	// listen on 0.0.0.0:8080
	if err := engine.Run(); err != nil {
		log.Fatalln(err)
	}
}
