package v1

import (
	apperr "true-kw/errors"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			// Обработка последней ошибки
			var err error = c.Errors.Last()
			status := apperr.ErrCodeToHTTPStatus(err)
			c.JSON(status, &HTTPError{Code: status, Message: apperr.Message(err)})
		}
	}
}
