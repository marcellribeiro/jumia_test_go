package use_cases

import (
	"github.com/joho/godotenv"
	"github.com/marcellribeiro/jumia_test_go/service_b/models"
	"log"
	"os"
	"strconv"
	"strings"
)

func ProcessCargo(cargo *models.Cargo, orders []models.Order, country string) {
	count := len(orders)
	log.Printf("Processing cargos (%v orders) from %v...\n", count, country)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	weightLimit := stringToFloat32(os.Getenv("WEIGHT_LIMIT_KG"))
	currentWeight := float32(0.0)
	var currentIds []int
	for idx, order := range orders {
		alreadyAdded := false
		nextOrder := idx + 1
		id := stringToInteger(order.ID)
		weight := stringToFloat32(order.Weight)
		newWeight := currentWeight + weight
		log.Printf("Processing order #%v\n", id)
		if newWeight > weightLimit {
			// send old values to new courier
			addCourier(currentIds, currentWeight, cargo)
			// create new values
			var currentIds []int
			currentIds = append(currentIds, id)
			currentWeight = weight
			alreadyAdded = true
		} else {
			// append values
			currentIds = append(currentIds, id)
			currentWeight = newWeight
		}

		//exception
		if !alreadyAdded && nextOrder == count {
			// send last values to new courier
			addCourier(currentIds, currentWeight, cargo)
		}
	}
	log.Println("Returning")
	return
}

func stringToInteger(text string) int {
	text = strings.Trim(text, " ")
	number, err := strconv.Atoi(text)
	if err != nil {
		return 0
	}
	return number
}

func stringToFloat32(text string) float32 {
	value, err := strconv.ParseFloat(text, 32)
	if err != nil {
		return 0.0
	}
	return float32(value)
}

func addCourier(ids []int, weight float32, cargo *models.Cargo) {
	newCourier := models.Courier{OrderIds: ids, TotalWeight: weight}
	cargo.Cargo = append(cargo.Cargo, newCourier)
	return
}
