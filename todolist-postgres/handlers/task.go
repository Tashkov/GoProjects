package handlers

import (
	"GoProjects/todolist-postgres/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /tasks
// Lists all tasks
func GetTasks(ctx *gin.Context) {
	var tasks []models.Task
	models.DB.Find(&tasks)

	ctx.JSON(http.StatusOK, gin.H{"data": tasks})
}

// GET /tasks/:id
// Gets a specific task by id
func GetTask(ctx *gin.Context) {
	var task models.Task

	if err := models.DB.Debug().Where("id = ?", ctx.Param("id")).First(&task).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": task})

}

// Schema to validate user input
type CreateTaskInput struct {
	TaskName        string `json:"task_name" binding:"required"`
	TaskDescription string `json:"task_description"`
}

// POST /tasks
// Creates a Task
func CreateTask(c *gin.Context) {
	// Validating user input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		TaskName:        input.TaskName,
		TaskDescription: input.TaskDescription}

	models.DB.Create(&task)

}

// Schema to validate user input
// When using PATCH method, include ID uint variable to resolve internal server error
type UpdateTaskInput struct {
	id              uint   `json:"id" binding:"required"`
	TaskName        string `json:"task_name" binding:"required"`
	TaskDescription string `json:"task_description"`
}

// UPDATE /tasks/:id
// Updates a specific task by ID
func UpdateTask(ctx *gin.Context) {
	var task models.Task
	if err := models.DB.Debug().Where("id = ?", ctx.Param("id")).First(&task).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task not found!"})
		return
	}

	var input UpdateTaskInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Debug().Model(&task).Updates(input)

	ctx.JSON(http.StatusOK, gin.H{"data": task})
}

// DELETE /tasks/:id
// Deletes a specific task by ID
func DeleteTask(ctx *gin.Context) {
	var task models.Task

	if err := models.DB.Debug().Where("id = ?", ctx.Param("id")).First(&task).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Task not found!"})
		return
	}

	models.DB.Debug().Delete(&task)

	ctx.JSON(http.StatusOK, gin.H{"data": "delete task"})
}
