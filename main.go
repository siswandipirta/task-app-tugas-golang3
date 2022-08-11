package main

import (
	"fmt"
	"net/http"
	TasksController "project/controllers"
	Model "project/models"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect datebase")
	}
	db.AutoMigrate(&Model.Tasks{})

	router := httprouter.New()

	router.GET("/", TasksController.Index)
	router.GET("/create", TasksController.Create)
	router.POST("/create", TasksController.Create)

	// router.ServeFiles("/static/*filepath", http.Dir("assets"))

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
