package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcellribeiro/jumia_test_go/service_b/models"
	"github.com/marcellribeiro/jumia_test_go/service_b/repositories"
	"github.com/marcellribeiro/jumia_test_go/service_b/use_cases"
	"log"
	"net/http"
)

func DailyCargo(c *gin.Context) {
	countries := getCountries()
	cargo := models.Cargo{}

	done := make(chan string)
	go processCargoFromCountries(countries, &cargo, done)
	log.Println(<-done)

	c.JSON(http.StatusOK, cargo)
}

func processCargoFromCountries(countries []string, cargo *models.Cargo, done chan string) {
	for _, country := range countries {
		orders := getOrders(country, -1)
		use_cases.ProcessCargo(cargo, orders, country)
	}
	done <- "finished"
}

func getOrders(country string, limit int) []models.Order {
	var orders = []models.Order{}
	if err := repositories.GetDB().Model(&models.Order{}).Where("country = ?", country).Find(&orders).Error; err != nil {
		log.Println(err)
	}
	return orders
}

func getCountries() []string {
	var countries []string
	if err := repositories.GetDB().Model(&models.Order{}).Distinct("country").Find(&countries).Error; err != nil {
		log.Println(err)
	}
	return countries
}
