package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RequestIDFromContext(c echo.Context) string {
	rid, ok := c.Get(RequestID).(string)
	middleware.RequestID()
	if !ok {
		rid = c.RealIP() + uuid.NewString()
		c.Set(RequestID, rid)
		return rid
	}
	return rid
}
