package main

import (
	"GoProjects/todolist-postgres/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Welcome to your To Do List",
		})

		models.ConnectDataBase()
		fmt.Println("Connect to DB")
		
		// Task
		r.POST("")

	})

	r.Run()
}
