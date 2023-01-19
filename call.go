package main

import "github.com/gin-gonic/gin"

func sendCall(c *gin.Context) {
	a := c.Param("userID")
	b := c.Param("dest")

	resp := sendCallService(a, b)

	c.JSON(200, gin.H{
		"success": resp,
	})
}
