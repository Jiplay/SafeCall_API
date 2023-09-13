package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type PostDescriptionStruct struct {
	UserID string `bson:"UserID"`
	Data   string `bson:"Data"`
}

type PostFullNameStruct struct {
	UserID string `bson:"UserID"`
	Data   string `bson:"Data"`
}

type PostPhoneNBStruct struct {
	UserID string `bson:"UserID"`
	Data   string `bson:"Data"`
}

type PostEmailStruct struct {
	UserID string `bson:"UserID"`
	Data   string `bson:"Data"`
}

type PostProfilePicStruct struct {
	UserID string `bson:"UserID"`
	Data   string `bson:"Data"`
}

func postDescription(c *gin.Context) {
	var data PostDescriptionStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := postProfileHandler("Description", data.UserID, data.Data)
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
	var data PostFullNameStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := postProfileHandler("FullName", data.UserID, data.Data)
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
	var data PostPhoneNBStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := postProfileHandler("PhoneNB", data.UserID, data.Data)
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
	var data PostEmailStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := postProfileHandler("Email", data.UserID, data.Data)
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

func postProfilePic(c *gin.Context) {
	var data PostProfilePicStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := postProfileHandler("ProfilePic", data.UserID, data.Data)
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
