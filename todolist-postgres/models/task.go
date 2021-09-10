package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	TaskName        string `json:"task_name"`
	TaskDescription string `json:"task_description"`
	ID              uint   `json:"id"`
}
