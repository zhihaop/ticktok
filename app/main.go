package main

import (
	"github.com/gin-gonic/gin"
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
	// initialize user's domain
	userRepository := _userRepository.NewUserRepository(db)
	userService := _userService.NewUserService(userRepository, nil)
	userController := _userController.NewUserController(userService)

	// TODO initialize other routers
	// initialize routers
	userController.InitRouter(engine.Group("/douyin/user"))

	// listen on 0.0.0.0:8080
	if err := engine.Run(); err != nil {
		log.Fatalln(err)
	}
}
