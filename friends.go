package main

import "github.com/gin-gonic/gin"

func manageFriendEndpoint(c *gin.Context) {
	userID := c.Param("userID")
	friend := c.Param("friend")
	action := c.Param("action")

	resp := actionFriendHandler(userID, friend, action)
	if resp != "success" {
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
	description := c.Param("data")

	resp := postProfileHandler("description", userID, description)
	if resp != "success" {
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

}
