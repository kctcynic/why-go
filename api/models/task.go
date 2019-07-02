package models

import (
	//"errors"

	"github.com/google/uuid"
)

 

// Task holds state information
type Task struct {
	TaskID          string    `json:"task_id"`
	TaskName        string    `json:"task_name"`
	TaskState       string    `json:"task_state"`
}

func generateID() string {
	var id = uuid.New()

	return id.String()
}

func CreateTask(taskName string) (*Task, error) {

	var taskID = generateID()

	task := new(Task)

	task.TaskName = taskName
	task.TaskID = taskID
	task.TaskState = "created"

	return task, nil
}

func (task *Task) StartTask() error {
	task.TaskState = "started"
	return nil
}

func (task *Task) FinishTask() error {
	task.TaskState = "finished"
	return nil
}
