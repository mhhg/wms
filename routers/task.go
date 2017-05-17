package routers

import (
	"github.com/labstack/echo"
	"github.com/mhhg/wms/controllers"
)

// SetTaskRoutes register routes with controllers
func SetTaskRoutes(e *echo.Echo) *echo.Echo {
	e.POST("/tasks", controllers.CreateTask)       // create task
	e.PUT("/tasks/:id", controllers.UpdateTask)    // update task
	e.GET("/tasks", controllers.GetTasks)          // get tasks
	e.GET("/tasks/:id", controllers.GetTaskByID)   // get task by id
	e.DELETE("/tasks/:id", controllers.DeleteTask) // delete task

	return e
}
