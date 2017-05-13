package routers

import (
	"github.com/mhhg/wms/controllers"
)

func SetTaskRoutes(e echo.Echo) *echo.Echo {
	e.POST("/tasks")
	e.PUT("/tasks/:id")
}
