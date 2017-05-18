package routers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mhhg/wms/common"
)

// Init Entry point for setting routes, return map of hosts with address
func Init() map[string]*common.Host {

	hosts := map[string]*common.Host{}

	// api
	api := echo.New()
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	hosts["api.localhost:8000"] = &common.Host{Echo: api}

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API")
	})

	return hosts
}
