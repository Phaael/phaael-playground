package errors

import (
	"net/http"
)

type ApiError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Cause   string `json:"cause"`
	Status  int    `json:"status"`
}

type ApiErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Cause   string `json:"cause"`
	Status  int    `json:"status"`
}

func (e ApiError) toMap() map[string]interface{} {
	return map[string]interface{}{
		"message": e.Message,
		"error":   e.Error,
		"status":  e.Status,
		"cause":   e.Cause,
	}
}

// GetError Method
func GetError(status int, message string) ApiErrorResponse {
	switch status {
	case http.StatusBadRequest:
		return BadRequestApiError(message)
	case http.StatusNotFound:
		return NotFoundApiError(message)
	case http.StatusForbidden:
		return ForbiddenApiError(message)
	case http.StatusUnauthorized:
		return UnauthorizedApiError(message)
	default:
		return InternalServerApiError(message, nil)
	}
}

func NotFoundApiError(message string) ApiErrorResponse {
	return ApiErrorResponse{message, "not_found", "", http.StatusNotFound}
}

func BadRequestApiError(message string) ApiErrorResponse {
	return ApiErrorResponse{message, "bad_request", "", http.StatusBadRequest}
}

func InternalServerApiError(message string, err error) ApiErrorResponse {
	return ApiErrorResponse{message, "internal_server_error", "", http.StatusInternalServerError}
}

func ForbiddenApiError(message string) ApiErrorResponse {
	return ApiErrorResponse{message, "forbidden", "", http.StatusForbidden}
}

func UnauthorizedApiError(message string) ApiErrorResponse {
	return ApiErrorResponse{message, "unauthorized", "", http.StatusUnauthorized}
}
