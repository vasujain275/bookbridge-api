package util

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Response is the standard API response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Time    time.Time   `json:"timestamp"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
		Time:    time.Now(),
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(message string, err interface{}) Response {
	return Response{
		Success: false,
		Message: message,
		Error:   err,
		Time:    time.Now(),
	}
}

// SendSuccess sends a success response
func SendSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, NewSuccessResponse(message, data))
}

// SendError sends an error response
func SendError(c *gin.Context, code int, message string, err interface{}) {
	c.JSON(code, NewErrorResponse(message, err))
}

// SendBadRequest sends a bad request response
func SendBadRequest(c *gin.Context, message string, err interface{}) {
	SendError(c, http.StatusBadRequest, message, err)
}

// SendNotFound sends a not found response
func SendNotFound(c *gin.Context, message string) {
	SendError(c, http.StatusNotFound, message, nil)
}

// SendInternalServerError sends an internal server error response
func SendInternalServerError(c *gin.Context, err interface{}) {
	SendError(c, http.StatusInternalServerError, "Internal server error", err)
}

// SendUnauthorized sends an unauthorized response
func SendUnauthorized(c *gin.Context) {
	SendError(c, http.StatusUnauthorized, "Unauthorized", nil)
}

// SendCreated sends a created response
func SendCreated(c *gin.Context, message string, data interface{}) {
	SendSuccess(c, http.StatusCreated, message, data)
}

// SendOK sends an OK response
func SendOK(c *gin.Context, message string, data interface{}) {
	SendSuccess(c, http.StatusOK, message, data)
}

// SendNoContent sends a no content response
func SendNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
