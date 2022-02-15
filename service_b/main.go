package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcellribeiro/jumia_test_go/service_b/controllers"
	"github.com/marcellribeiro/jumia_test_go/service_b/models"
	"github.com/marcellribeiro/jumia_test_go/service_b/repositories"
	"log"
)

func init() {
	if err := repositories.Init(); err != nil {
		log.Fatal(err)
	}
	err := repositories.GetDB().AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.POST("/order/new", controllers.NewOrder)
	router.GET("/cargo/daily", controllers.DailyCargo)
	router.Run(":8063")
}
