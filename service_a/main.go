package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcellribeiro/jumia_test_go/service_a/controllers"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.POST("/upload", controllers.UploadFile)
	router.Run(":8062")
}
