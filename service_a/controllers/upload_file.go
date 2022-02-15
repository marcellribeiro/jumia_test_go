package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcellribeiro/jumia_test_go/service_a/providers"
	"github.com/marcellribeiro/jumia_test_go/service_a/use_cases"
	"log"
	"net/http"
	"time"
)

func UploadFile(c *gin.Context) {
	startTime := time.Now()
	log.Println("starting upload at ", startTime.Format("2006.01.02 15:04:05"))

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filePath := "uploads/" + file.Filename
	log.Println("new file uploaded: " + filePath)

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := use_cases.ProcessFile(filePath)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("sending to Service B")
	go providers.SendOrdersToServiceB(orders)

	elapsedTime := time.Since(startTime)
	log.Printf("Finished! This process took %s", elapsedTime)

	c.JSON(http.StatusOK, gin.H{"message": "The file was successfully uploaded"})
}
