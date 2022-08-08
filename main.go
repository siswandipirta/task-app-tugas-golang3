package main

import (
	"fmt"
	"net/http"
	TasksController "project/controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", TasksController.Index)
	router.GET("/create", TasksController.Create)

	router.ServeFiles("/static/*filepath", http.Dir("assets"))

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
