package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgetPasswordStruct struct {
	Email string `bson:"Email"`
}

type CheckCodeStruct struct {
	Email string `bson:"Email"`
	Code  string `bson:"Code"`
}

type SetPasswordStruct struct {
	Email    string `bson:"Email"`
	Password string `bson:"Password"`
}

type EditPasswordStruct struct {
	UserID      string `bson:"UserID"`
	PasswordOld string `bson:"PasswordOld"`
	PasswordNew string `bson:"PasswordNew"`
}

type RegisterStruct struct {
	Login    string `bson:"Login"`
	Password string `bson:"Password"`
	Email    string `bson:"Email"`
}

type LoginStruct struct {
	Login    string `bson:"Login"`
	Password string `bson:"Password"`
}

type DeleteUserStruct struct {
	UserID string `bson:"UserID"`
}

func deleteUser(c *gin.Context) {
	var data DeleteUserStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := deleteUserData(data.UserID)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func login(c *gin.Context) {
	var data LoginStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// login := c.Param("login")
	// psw := c.Param("psw")

	resp := LoginHandler(data.Login, data.Password)

	if resp == "failed" {
		c.JSON(404, gin.H{
			"failed": "404",
		})
	} else {
		c.JSON(200, gin.H{
			"success": resp,
		})
	}
}

func register(c *gin.Context) {
	var data RegisterStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := RegisterHandler(data.Login, data.Password, data.Email)

	if resp != "200" {
		c.JSON(403, gin.H{
			"failed": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"success": "200",
		})
	}
}

func editPswEndpoint(c *gin.Context) {
	var data EditPasswordStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := PasswordHandler(data.UserID, data.PasswordOld, data.PasswordNew)

	if resp != "200" {
		c.JSON(403, gin.H{
			"failed": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"success": "200",
		})
	}
}

func forgetPassword(c *gin.Context) { // FIXME Update Front
	var data ForgetPasswordStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := ForgetPasswordHandler(data.Email)

	if resp != "200" {
		c.JSON(403, gin.H{
			"failed": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"success": "200",
		})
	}
}

func checkcode(c *gin.Context) {
	var data CheckCodeStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := checkCodeHandler(data.Email, data.Code)

	if !resp {
		c.JSON(403, gin.H{
			"failed": "404",
		})
	} else {
		c.JSON(200, gin.H{
			"success": "200",
		})
	}
}

func setPswEndpoint(c *gin.Context) {
	var data SetPasswordStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := setPasswordHandler(data.Email, data.Password)

	if resp != "200" {
		c.JSON(403, gin.H{
			"failed": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"success": "200",
		})
	}
}
