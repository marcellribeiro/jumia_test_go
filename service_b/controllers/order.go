package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcellribeiro/jumia_test_go/service_b/models"
	"github.com/marcellribeiro/jumia_test_go/service_b/repositories"
	"log"
	"net/http"
)

type orderReceived struct {
	ID      string `json:"id" binding:"required"`
	Email   string `json:"email"  binding:"required"`
	Phone   string `json:"phone_number" binding:"required"`
	Weight  string `json:"parcel_weight" binding:"required"`
	Country string `json:"country" binding:"required"`
}

func NewOrder(c *gin.Context) {
	var json orderReceived

	if err := c.BindJSON(&json); err != nil {
		log.Println(err)
		c.AbortWithStatus(400)
		return
	}

	log.Printf("received: %v\n", json)

	order := models.Order{
		ID:      json.ID,
		Email:   json.Email,
		Phone:   json.Phone,
		Weight:  json.Weight,
		Country: json.Country,
	}
	c.JSON(http.StatusOK, gin.H{"id": json.ID, "status": true})

	repositories.GetDB().Create(&order)
}
