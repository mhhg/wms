package main

import (
	"github.com/labstack/echo"
	"github.com/mhhg/wms/routers"
)

func main() {
	// Server
	e := echo.New()

	// Register routes
	hosts := routers.Init()

	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})
	e.Logger.Fatal(e.Start(":1323"))
}
