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

	// initialize user's domain
	userRepository := _userRepository.NewUserRepository(db)
	userService := _userService.NewUserService(userRepository, nil)
	_userController.RegisterUserController(engine.Group("/douyin/user"), userService)

	// TODO initialize other domains

	// listen on 0.0.0.0:8080
	if err := engine.Run(); err != nil {
		log.Fatalln(err)
	}
}
