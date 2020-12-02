package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser endpoint
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "we will Implement getuser endpoint")
}

// CreateUser controller
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "we will Implement createuser endpoint")
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "we will Implement search user endpoint")
// }
