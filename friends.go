package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func manageFriendEndpoint(c *gin.Context) {
	userID := c.Param("userID")
	friend := c.Param("friend")
	action := c.Param("action")
	url := fmt.Sprintf("http://localhost:8081/friend/%s/%s/%s", userID, friend, action)

	resp := actionFriendHandler(url)
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
	userID := c.Param("userID")
	friend := c.Param("friend")
	action := c.Param("action")

	url := fmt.Sprintf("http://localhost:8081/friendRequest/%s/%s/%s", userID, friend, action)

	resp := actionFriendHandler(url)
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
	resp := getFriends(userID)

	if resp != "success" {
		c.JSON(503, gin.H{
			"failed": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"success": resp,
		})
	}
}
