package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostMessageStruct struct {
	Username   string `bson:"Username"`
	Friendname string `bson:"Friendname"`
	Message    string `bson:"Message"`
}

func GetConversations(c *gin.Context) {
	userID := c.Param("UserID")
	resp := GetConversation(userID)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func GetMessages(c *gin.Context) {
	userID := c.Param("UserID")
	friendID := c.Param("FriendID")
	resp := GetMessagesHandler(userID, friendID)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func PostMessage(c *gin.Context) {
	var data PostMessageStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	PostMessageHandler(data.Username, data.Friendname, data.Message)

	c.JSON(200, gin.H{
		"Success ": "True",
	})
}
