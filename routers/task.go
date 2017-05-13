package routers

import (
	"github.com/mhhg/wms/controllers"
)

func SetTaskRoutes(e echo.Echo) *echo.Echo {
	e.POST("/tasks")       // create task
	e.PUT("/tasks/:id")    // update task
	e.GET("/tasks")        // get tasks
	e.GET("/tasks/:id")    // get task by id
	e.DELETE("/tasks/:id") // delete task

	return e
}
