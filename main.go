package main

import (
	"github.com/gin-gonic/gin"
)

// This function is here for test purpose with Postman
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	// r.Use(CORS())

	r.GET("/login/:login/:psw", login)
	r.GET("/profile/:userID", getUserProfile)
	r.GET("/search/:userID", SearchNameEndpoint)

	r.POST("/forgetPassword/:email", forgetPassword)
	r.POST("/forgetPassword/:email/:code", checkcode)
	r.POST("/setPassword/:email/:new", setPswEndpoint)
	r.POST("/editPassword/:userID/:old/:new", editPswEndpoint)

	r.POST("/register/:login/:psw/:email", register)
	r.POST("/profileDescription/:userID/:data", postDescription)
	r.POST("/profileFullName/:userID/:data", postFullName)
	r.POST("/profilePhoneNB/:userID/:data", postPhoneNB)
	r.POST("/profileEmail/:userID/:data", postEmail)

	r.GET("/tryCall", sendCall)

	r.Run()
}

func login(c *gin.Context) {
	login := c.Param("login")
	psw := c.Param("psw")

	resp := LoginHandler(login, psw)

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
	login := c.Param("login")
	psw := c.Param("psw")
	email := c.Param("email")

	resp := RegisterHandler(login, psw, email)

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
	login := c.Param("userID")
	old := c.Param("old")
	new := c.Param("new")

	resp := PasswordHandler(login, old, new)

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

func forgetPassword(c *gin.Context) {
	email := c.Param("email")

	resp := ForgetPasswordHandler(email)

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
	email := c.Param("email")
	code := c.Param("code")

	resp := checkCodeHandler(email, code)

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
	email := c.Param("email")
	password := c.Param("new")

	resp := setPasswordHandler(email, password)

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
