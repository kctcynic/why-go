package handlers


import (
	"fmt"
	"net/http"

	"app/bindings"
	"app/models"

	"github.com/labstack/echo/v4"
)

// CreateTask handles API requests on /createTask
func CreateTask(c echo.Context) error {
	var msg *bindings.ErrorResult

	var params = new(bindings.CreateTaskParameters)
	if err := c.Bind(params); err != nil {
		fmt.Println(err)
		msg = new(bindings.ErrorResult)
		msg.ErrorMessage = "Invalid Parameters"
		return c.JSON(http.StatusOK, msg)
	}

	cc := c.(*CustomContext)

	task, err := models.CreateTask(params.TaskName)
	if err != nil {
		fmt.Println(err)
		msg = new(bindings.ErrorResult)
		msg.ErrorMessage = "Cannot create task"
		return c.JSON(http.StatusOK, msg)
	}

	cc.StoreTask(task)

	var result = new(bindings.CreateTaskResult)

	result.Output = "Created Task " + task.TaskID

	return c.JSON(http.StatusOK, result)
}
