package main

import (
	"GoProjects/todolist-postgres/handlers"
	"GoProjects/todolist-postgres/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDataBase()
	fmt.Println("Connect to DB")
	// Welcome page
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": "Welcome to your To Do List",
		})

	})

	// Task
	router.GET("/tasks", handlers.GetTasks)
	router.POST("/tasks", handlers.CreateTask)
	router.GET("tasks/:id", handlers.GetTask)
	router.PATCH("tasks/:id", handlers.UpdateTask)
	router.DELETE("tasks/:id", handlers.DeleteTask)

	router.Run()
}
