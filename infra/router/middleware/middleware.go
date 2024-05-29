package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsHandler() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowCredentials: true,
			AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
			AllowOrigins:     []string{"*"},
		})
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusInternalServerError, "")
}
