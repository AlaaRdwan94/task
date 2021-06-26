package main

import (
	"github.com/InnoSoft/task/infrastructure/db"
	"github.com/InnoSoft/task/infrastructure/worker/api_worker"
	_userHandler "github.com/InnoSoft/task/user/handler"
	_userRepo "github.com/InnoSoft/task/user/repository"
	_userUsecase "github.com/InnoSoft/task/user/usecase"
	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v3"
)

func InitializeRouts(db *db.DB, router *gin.RouterGroup, client *redis.Client) {
	// user model
	userRepository := _userRepo.NewUserRepo(db.GormDB,client)
	userUsecase:= _userUsecase.NewUserUsecase(userRepository)
	_userHandler.NewUserHandler(router,userUsecase)
	api_worker.RunCron()
}