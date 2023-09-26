package main

import (
	"tech-shelf/controller"
	"tech-shelf/repository"
	"tech-shelf/usecase"
	"tech-shelf/router"
	"tech-shelf/db"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}