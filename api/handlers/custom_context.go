package handlers

import (
	"errors"
	//"fmt"
	"sync"

	"app/models"

	"github.com/labstack/echo/v4"
)

// MemoryStore is used as a fake database for testing
type MemoryStore struct {
	tasks     map[string]*models.Task // similarly for the task structures
}

// CustomContext is used to pass the database connection and persistence functions to the handlers
type CustomContext struct {
	echo.Context

	sync.RWMutex

	ms *MemoryStore
}

// NewMemoryStore creates a MemoryStore instance and initializes the maps that it contains
func NewMemoryStore() (*MemoryStore, error) {
	var ms = new(MemoryStore)
	ms.tasks = make(map[string]*models.Task)
	return ms, nil
}

// SetDatabase sets the database connection that is passed to handlers
func (cc *CustomContext) SetDatabase(ms *MemoryStore) error {
	cc.ms = ms
	return nil
}


// StoreTask saves the task data for use in later handlers
func (cc *CustomContext) StoreTask(task *models.Task) error {
	cc.Lock()
	cc.ms.tasks[task.TaskID] = task
	cc.Unlock()
	return nil
}

// GetTaskFromID retrieves the task data given a task ID
func (cc *CustomContext) GetTaskFromID(taskID string) (*models.Task, error) {
	cc.RLock()
	task, exists := cc.ms.tasks[taskID]
	cc.RUnlock()
	if !exists {
		return nil, errors.New("bad task id")
	}
	return task, nil
}

// GetAllTasks gets a list of all tasks
func (cc *CustomContext) GetAllTasks() (tasks []*models.Task, err error) {
	//
	// this is not scalable, just for example
	//
	cc.RLock()
	for _, task := range cc.ms.tasks {
		tasks = append(tasks, task)
	}
	cc.RUnlock()

	return
}
