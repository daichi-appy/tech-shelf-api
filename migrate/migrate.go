package main

import (
	"fmt"
	"tech-shelf/db"
	"tech-shelf/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Connected")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}