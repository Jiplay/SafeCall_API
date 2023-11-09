package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ManageFriendStruct struct {
	UserID  string `bson:"UserID"`
	Friend  string `bson:"Friend"`
	Subject string `bson:"Subject"`
	Action  string `bson:"Action"`
}

type ReplyFriendStruct struct {
	UserID  string `bson:"UserID"`
	Friend  string `bson:"Friend"`
	Subject string `bson:"Subject"`
	Action  string `bson:"Action"`
}

func manageFriendEndpoint(c *gin.Context) {
	var data ManageFriendStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	url := "http://profiler:8081/friend"

	requestBody := map[string]interface{}{
		"UserID":  data.UserID,
		"Dest":    data.Friend,
		"Subject": data.Subject,
		"Action":  data.Action,
	}

	resp := actionFriendHandler(url, requestBody)
	if !resp {
		c.JSON(503, gin.H{
			"failed": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"success": "200",
		})
	}
}

func replyFriendEndpoint(c *gin.Context) {
	var data ReplyFriendStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := "http://profiler:8081/friendRequest"

	requestBody := map[string]interface{}{
		"UserID":  data.UserID,
		"Dest":    data.Friend,
		"Subject": data.Subject,
		"Action":  data.Action,
	}

	resp := actionFriendHandler(url, requestBody)
	if !resp {
		c.JSON(503, gin.H{
			"failed": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"success": "200",
		})
	}
}

func listFriends(c *gin.Context) {
	userID := c.Param("userID")

	friends := getFriends(userID)
	c.JSON(200, gin.H{
		"fetched": friends,
	})
}
