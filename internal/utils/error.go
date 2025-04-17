package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	formatted := make(map[string]string)

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fieldErr := range ve {
			formatted[fieldErr.Field()] = fieldErr.Tag() // ex: Email -> "required"
		}
	} else {
		formatted["error"] = err.Error()
	}

	return formatted
}

func SendError(c *gin.Context, status int, message string, details any) {
	c.JSON(status, gin.H{
		"error":   message,
		"details": details,
	})
}
