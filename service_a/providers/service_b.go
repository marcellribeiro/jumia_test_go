package providers

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/marcellribeiro/jumia_test_go/service_a/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ServiceBResponse struct {
	ID     int  `json:"id"`
	Status bool `json:"status"`
}

func SendOrdersToServiceB(orders []models.Order) bool {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	serviceBUrl := os.Getenv("SERVICE_B_URL") + "/order/new"

	for idx, _ := range orders {
		sendOrder(&orders[idx], serviceBUrl)
	}
	return true
}

func sendOrder(order *models.Order, serviceBUrl string) bool {
	ordersJSON, err := json.Marshal(order)

	req, err := http.Post(serviceBUrl, "application/json",
		bytes.NewBuffer(ordersJSON))
	if err != nil {
		log.Println(err)
		return false
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return false
	}

	var responseObject ServiceBResponse
	json.Unmarshal([]byte(body), &responseObject)
	return responseObject.Status
}
