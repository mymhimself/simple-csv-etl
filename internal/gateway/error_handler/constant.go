package errorhandler

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// The accpetable content type of request | For internal usage
const applicationJSON = "application/json"
const multipartForm = "multipart/form-data"

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
