package errorhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/mymhimself/simple-csv-reader/internal/gateway/handlers"
	"google.golang.org/grpc/status"
)

func DefaultErrorHandler(err error, c echo.Context) {
	w := handlers.GenericResponse{}
	w.HandleError(err)
	//todo get the status code of error
	c.JSON(code2Http[status.Code(err)], w)
}
