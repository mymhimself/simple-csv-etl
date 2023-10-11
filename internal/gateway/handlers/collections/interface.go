package etl

import echo "github.com/labstack/echo/v4"

type ICollections interface {
	Register(*echo.Echo) error
}
