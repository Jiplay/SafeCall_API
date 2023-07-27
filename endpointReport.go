package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostReportStruct struct {
	Username string `bson:"Username"`
	Message  string `bson:"Message"`
	Date     string `bson:"Message"`
}

type DelReportStruct struct {
	Username string `bson:"Username"`
	Date     string `bson:"Date"`
}

func NewReport(c *gin.Context) {
	var data PostFeedbackStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NewReportHandler(data.Username, data.Date, data.Message)

	c.JSON(200, gin.H{
		"Success ": "True",
	})
}

func GetReports(c *gin.Context) {
	resp := GetReportHandler()
	c.JSON(200, gin.H{
		"Success": resp,
	})
}

func DelReports(c *gin.Context) {
	var data DelFeedbackStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := DelReportHandler(data.Username, data.Date)
	c.JSON(200, gin.H{
		"Success": resp,
	})
}
