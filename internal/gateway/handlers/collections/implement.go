package etl

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/mymhimself/simple-csv-reader/internal/gateway/handlers"
	"github.com/mymhimself/simple-csv-reader/internal/services/writer"
	"github.com/mymhimself/simple-csv-reader/pkg/constants"
)

// ────────────────────────────────────────────────────────────────────────────────
func (s *iCollections) List(c echo.Context) error {

	args := writer.ListParams{
		Collection: c.Param("collection"),
	}

	resp, err := s.service.List(c.Request().Context(), &args)
	if err != nil {
		return err
	}

	var w handlers.GenericResponse

	return w.JSON(resp).AddMetadata(constants.NextPageToken, "").Write(c)
}

// ─────────────────────────────────────────────────────────────────────────────
func (s *iCollections) HealthCheck(c echo.Context) error {
	w := handlers.GenericResponse{}
	return w.Write(c)
}

// ────────────────────────────────────────────────────────────────────────────────
func (s *iCollections) Register(e *echo.Echo) error {
	var routes = []handlers.Route{
		{
			Name:    "health-check",
			Method:  http.MethodGet,
			Path:    "etl/health-check",
			Handler: s.HealthCheck,
		},
		{
			Name:       "tutorial.list",
			Method:     http.MethodGet,
			Path:       "etl/collections/:collection",
			Handler:    s.List,
			Middleware: []echo.MiddlewareFunc{},
		},
	}

	for _, v := range routes {
		e.Add(v.Method, v.Path, v.Handler, v.Middleware...)
	}
	return nil
}
