package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

//StartApplication init
func StartApplication() {
	mapUrls()
	router.Run()
}
