package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostFeedbackStruct struct {
	Username string `bson:"Username"`
	Message  string `bson:"Message"`
	Date     string `bson:"Date"`
}

type DelFeedbackStruct struct {
	Username string `bson:"Username"`
	Date     string `bson:"Date"`
}

type EditFeedbackStruct struct {
	Username string `bson:"Username"`
	Date     string `bson:"Date"`
	State    string `bson:"State"`
}

func NewFeedback(c *gin.Context) {
	var data PostFeedbackStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NewFeedbackHandler(data.Username, data.Date, data.Message)

	c.JSON(200, gin.H{
		"Success ": "True",
	})
}

func GetFeedback(c *gin.Context) {
	resp := GetFeedbackHandler()
	c.JSON(200, gin.H{
		"Success": resp,
	})
}

func EditFeedbackEndpoint(c *gin.Context) {
	var data EditFeedbackStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := EditFeedbackHandler(data.Username, data.Date, data.State)
	c.JSON(200, gin.H{
		"Success": resp,
	})
}

func DelFeedback(c *gin.Context) {
	var data DelFeedbackStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := DelFeedbackHandler(data.Username, data.Date)
	c.JSON(200, gin.H{
		"Success": resp,
	})
}
