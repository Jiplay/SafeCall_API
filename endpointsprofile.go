package main

import (
	"github.com/gin-gonic/gin"
)

func postDescription(c *gin.Context) {
	userID := c.Param("userID")
	description := c.Param("data")

	resp := postProfileHandler("description", userID, description)
	if resp == "failed" {
		c.JSON(503, gin.H{
			"failed": "404",
		})
	} else {
		c.JSON(200, gin.H{
			"success": resp,
		})
	}
}

func postFullName(c *gin.Context) {
	userID := c.Param("userID")
	description := c.Param("data")

	resp := postProfileHandler("FullName", userID, description)
	if resp == "failed" {
		c.JSON(503, gin.H{
			"failed": "503",
		})
	} else {
		c.JSON(200, gin.H{
			"success": resp,
		})
	}
}
