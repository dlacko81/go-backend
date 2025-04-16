package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/dlacko81/go-backend/sheets" // Replace with your actual repository path
)

func main() {
	r := gin.Default()

	// Endpoint to get data from Google Sheets
	r.GET("/api/data", func(c *gin.Context) {
		data, err := sheets.GetSheetData()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, data)
	})

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
