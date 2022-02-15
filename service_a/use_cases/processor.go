package use_cases

import (
	"encoding/json"
	"github.com/marcellribeiro/jumia_test_go/service_a/models"
	"log"
	"sync"
)

func ProcessFile(filePath string) ([]models.Order, error) {
	log.Println("Processing file " + filePath)
	orders, err := getOrders(filePath)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func getOrders(filePath string) ([]models.Order, error) {
	var orders []models.Order

	// Get results from csv importer
	results, _ := ReadCSV(filePath)

	var counter int
	var wg = &sync.WaitGroup{}

	requests := make(chan *models.Order, 20)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for {
				req := <-requests
				if req == nil {
					break
				}
				if req.ID != "" {
					// Create new order and add to order array
					orders = append(orders, models.Order{
						ID:      req.ID,
						Email:   req.Email,
						Phone:   req.Phone,
						Weight:  req.Weight,
						Country: req.Country,
					})
				}
			}
			wg.Done()
		}()
	}

Loop:
	for {
		select {
		case result, open := <-results.Results:
			if !open {
				break Loop
			}
			byt, _ := json.Marshal(result)
			var req models.Order
			if err := json.Unmarshal(byt, &req); err != nil {
				return nil, err
			}
			req.Country = GetCountryByPhone(req.Phone)
			requests <- &req
			counter++
		}
	}
	close(requests)

	log.Printf("%v orders were read\n", counter)
	wg.Wait()

	return orders, nil
}
