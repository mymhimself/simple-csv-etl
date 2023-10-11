package handlers

import "github.com/labstack/echo/v4"

type Route struct {
	Name         string
	Method       string
	Path         string
	Handler      func(echo.Context) error
	Middleware   []echo.MiddlewareFunc
	Descriptions string
}
