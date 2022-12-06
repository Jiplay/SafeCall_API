package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func postDescription(c *gin.Context) {
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

func postFullName(c *gin.Context) {
	userID := c.Param("userID")
	description := c.Param("data")

	resp := postProfileHandler("FullName", userID, description)
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

func postPhoneNB(c *gin.Context) {
	userID := c.Param("userID")
	description := c.Param("data")

	resp := postProfileHandler("PhoneNB", userID, description)
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

func postEmail(c *gin.Context) {
	userID := c.Param("userID")
	email := c.Param("data")

	resp := postProfileHandler("Email", userID, email)
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

func getUserProfile(c *gin.Context) {
	userID := c.Param("userID")

	resp := getProfileHandler(userID)
	c.JSON(200, gin.H{
		"profile": resp,
	})
}

func SearchNameEndpoint(c *gin.Context) {
	userID := c.Param("userID")

	resp := searchName(userID)
	final := strings.Split(resp, ",")
	c.JSON(200, gin.H{
		"suggestions": final,
	})
}
