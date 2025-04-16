package handlers

import (
	"net/http"
	"go-backend/sheets"
	"github.com/gin-gonic/gin"
)

type FormData struct {
	ClientName       string  `json:"clientName"`
	Date             string  `json:"date"`
	Volume           float64 `json:"volume"`
	Vintage          string  `json:"vintage"`
	Technology       string  `json:"technology"`
	Country          string  `json:"country"`
	Price            float64 `json:"price"`
	Comments         string  `json:"comments"`
}

func SubmitData(c *gin.Context) {
	var data FormData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := sheets.AppendToSheet(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write to sheet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetData(c *gin.Context) {
	records, err := sheets.ReadSheet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read data"})
		return
	}
	c.JSON(http.StatusOK, records)
}
