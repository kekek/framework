package handler

import (
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.POST("/ctl/token", Token)
}
