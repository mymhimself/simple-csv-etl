package handlers

import (
	"net/http"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrContextIsNotSetOnTheHandler = status.Error(codes.Internal, "Context is not set on the handler")

type GenericResponse struct {
	Success  bool              `json:"success"`
	Messages []string          `json:"messages"`
	Data     interface{}       `json:"data"`
	Meta     map[string]string `json:"meta"`
	ctx      echo.Context      `json:"-"`
	err      error             `json:"-"`
}

// ─────────────────────────────────────────────────────────────────────────────
func NewResponse() *GenericResponse {
	return new(GenericResponse)
}

// ────────────────────────────────────────────────────────────────────────────────
func (r *GenericResponse) HandleError(err error) bool {
	if err == nil {
		// no error happened
		return false
	}
	// r.Error = err.Error()
	e, _ := status.FromError(err)

	// when a service is down, replace the error with a proper message and send an exception to Sentry
	if strings.Contains(err.Error(), "server misbehaving") || strings.Contains(err.Error(), "connection refused") {
		sentry.CaptureException(err)
		r.err = ErrServiceIsUnAvailable
	}

	r.Messages = append(r.Messages, e.Message())
	// r.Messages = append(r.Messages, err.Error())
	r.Success = false
	// error happened
	return true
}

// ────────────────────────────────────────────────────────────────────────────────
func (r *GenericResponse) ResponseJsonWithCode(c echo.Context, status int) error {
	return c.JSON(status, r)
}

// ─────────────────────────────────────────────────────────────────────────────
func (r *GenericResponse) ResponseJson(c echo.Context) error {
	r.Success = true
	return c.JSON(http.StatusOK, r)
}

// ─────────────────────────────────────────────────────────────────────────────
func (s *GenericResponse) AddMetadata(key, value string) *GenericResponse {
	if s.Meta == nil {
		s.Meta = make(map[string]string)
	}

	s.Meta[key] = value
	return s
}

// ─────────────────────────────────────────────────────────────────────────────
func (s *GenericResponse) JSON(d interface{}) *GenericResponse {
	s.Data = d
	return s
}

// ─────────────────────────────────────────────────────────────────────────────
func (s *GenericResponse) Write(ctx echo.Context) error {
	if s.err == nil {
		s.Success = true
		return ctx.JSON(http.StatusOK, s)
	}

	s.Success = false
	s.Messages = append(s.Messages, s.err.Error())
	return ctx.JSON(code2Http[status.Code(s.err)], s)
}

// ─────────────────────────────────────────────────────────────────────────────

// Err add error to the response
func (s *GenericResponse) Err(err error) *GenericResponse {
	s.err = err

	// when a service is down, replace the error with a proper message and send an exception to Sentry
	if strings.Contains(err.Error(), "server misbehaving") || strings.Contains(err.Error(), "connection refused") {
		sentry.CaptureException(err)
		s.err = ErrServiceIsUnAvailable
	}
	return s
}

// ─────────────────────────────────────────────────────────────────────────────
var code2Http = map[codes.Code]int{
	codes.Aborted:            http.StatusExpectationFailed,
	codes.AlreadyExists:      http.StatusConflict,
	codes.FailedPrecondition: http.StatusPreconditionFailed,
	codes.OutOfRange:         http.StatusRequestedRangeNotSatisfiable,
	codes.DataLoss:           http.StatusExpectationFailed,
	codes.Unknown:            http.StatusBadRequest,
	codes.InvalidArgument:    http.StatusFailedDependency,
	codes.NotFound:           http.StatusNotFound,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.Unauthenticated:    http.StatusUnauthorized,
	codes.Internal:           http.StatusInternalServerError,
	codes.ResourceExhausted:  http.StatusInternalServerError,
	codes.OK:                 http.StatusOK,
	codes.Unavailable:        http.StatusServiceUnavailable,
	codes.Unimplemented:      http.StatusNotImplemented,
}
