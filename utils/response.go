package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`            // Status response: success / error
	Message string      `json:"message,omitempty"` // Pesan singkat terkait response
	Data    interface{} `json:"data,omitempty"`    // Data yang dikembalikan
	Error   string      `json:"error,omitempty"`   // Pesan error (jika ada)
}

// SuccessResponse creates a standard success response
func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// ErrorResponse creates a standard error response
func ErrorResponse(c *gin.Context, statusCode int, message string, err string) {
	c.JSON(statusCode, Response{
		Status:  "error",
		Message: message,
		Error:   err,
	})
}
