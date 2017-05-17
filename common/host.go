package common

import (
	"github.com/labstack/echo"
)

type (
	// Host pointer to echo, to manage hosts of server
	Host struct {
		Echo *echo.Echo
	}
)
