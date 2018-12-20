package manager

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CORSMiddleware cors middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

//HandlerFunc wrap handler to error
type HandlerFunc func(*gin.Context) (interface{}, error)

//WrapHandler wrap func
func WrapHandler(fc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := fc(c)
		if err != nil {
			log.Printf("fc(c) err : %v", err)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

//WrapWXHandler wx need string success
func WrapWXHandler(fc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := fc(c)
		if err != nil {
			log.Printf("WrapWXHandler Err; %v", err)
		}
		c.String(http.StatusOK, "success")
	}
}
