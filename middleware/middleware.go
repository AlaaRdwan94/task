package middleware

import (
	"github.com/InnoSoft/task/middleware/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

//SetMiddlewareAuthentication check for the validity of the authentication token provided
func SetMiddlewareAuthentication(next gin.HandlerFunc) gin.HandlerFunc {
	return func (c *gin.Context) {
		uid ,err := auth.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		c.Writer.Header().Set("uid",uid)
		next(c)
	}
}
