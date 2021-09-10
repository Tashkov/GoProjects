package handlers

import (
	"GoProjects/todolist-postgres/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /tasks
// Lists all tasks
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	models.DB.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// GET /tasks/id
// Gets a specific task by id
func GetTask() {

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

// UPDATE /tasks/id
// Updates a specific task by ID
func UpdateTask() {

}

// DELETE /tasks/id
// Deletes a specific task by ID
func DeleteTask() {

}
