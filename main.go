package main

import (
	"tech-shelf/controller"
	"tech-shelf/db"
	"tech-shelf/repository"
	"tech-shelf/router"
	"tech-shelf/usecase"
	"tech-shelf/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	bookUsecase := usecase.NewBookUsecase()
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	bookController := controller.NewBookController(bookUsecase)
	e := router.NewRouter(userController, taskController, bookController)
	e.Logger.Fatal(e.Start(":8080"))
}
