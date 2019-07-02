package handlers


import (
	"fmt"
	"net/http"

	"app/bindings"

	"github.com/labstack/echo/v4"
)

// FinishTask handles API requests on /finishTask
func FinishTask(c echo.Context) error {
	var msg *bindings.ErrorResult

	var params = new(bindings.FinishTaskParameters)
	if err := c.Bind(params); err != nil {
		fmt.Println(err)
		msg = new(bindings.ErrorResult)
		msg.ErrorMessage = "Invalid Parameters"
		return c.JSON(http.StatusOK, msg)
	}

	cc := c.(*CustomContext)

	task, err := cc.GetTaskFromID(params.TaskID)
	if err != nil {
		fmt.Println(err)
		msg = new(bindings.ErrorResult)
		msg.ErrorMessage = "Invalid Task ID"
		return c.JSON(http.StatusOK, msg)
	}

	var result = new(bindings.CreateTaskResult)

	result.Output = "Finished task " + task.TaskID

	return c.JSON(http.StatusOK, result)
}
