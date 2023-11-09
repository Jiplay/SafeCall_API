package main

import (
	"net/http"

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

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	r.POST("/login", login)                   // TESTED
	r.GET("/profile/:userID", getUserProfile) // TESTED
	r.GET("/search/:userID", SearchNameEndpoint)

	r.POST("/forgetPassword", forgetPassword) // UNTESTABLE
	r.POST("/forgetPasswordCode", checkcode)  // UNTESTABLE
	r.POST("/setPassword", setPswEndpoint)
	r.POST("/editPassword", editPswEndpoint)

	r.POST("/register", register)                  // TESTED
	r.POST("/profileDescription", postDescription) // TESTED
	r.POST("/profileFullName", postFullName)       // TESTED
	r.POST("/profilePhoneNB", postPhoneNB)         // TESTED
	r.POST("/profileEmail", postEmail)             // TESTED
	r.POST("/profilePic", postProfilePic)
	r.POST("/delete", deleteUser) // TESTED

	r.POST("/manageFriend", manageFriendEndpoint) // TESTED
	r.POST("/replyFriend", replyFriendEndpoint)   // TESTED
	r.GET("/listFriends/:userID", listFriends)    // TESTED

	r.POST("/addEvent", addEventEndpoint)          // TESTED
	r.POST("/delEvent", delEventEndpoint)          // TESTED
	r.POST("/confirmEvent", confirmEvent)          // TESTED
	r.GET("/listEvent/:userID", listEventEndpoint) // TESTED

	r.POST("/AddNotification", addNotificationEndpoint) // FIXME Inform Front TESTED
	r.POST("/DelNotification", delNotificationEndpoint) // TESTED
	r.GET("/notification/:UserID", GetUserNotification) // TESTED

	r.POST("/sendMessage", PostMessage)
	r.GET("/conversations/:UserID", GetConversations)
	r.GET("/messages/:UserID/:FriendID", GetMessages)
	r.GET("/delRoom/:room", DelMessage)

	r.POST("/feedback", NewFeedback) // Tested
	r.POST("/editFeedback", EditFeedbackEndpoint)
	r.POST("/delFeedback", DelFeedback) // Tested
	r.GET("/feedback", GetFeedback)     // Tested

	r.POST("/report", NewReport)     // Tested
	r.POST("/delReport", DelReports) // Tested
	r.GET("/report", GetReports)     // Tested

	r.GET("/setupProfiler", SetupProfiler)
	r.GET("/tryCall", sendCall)

	r.Run()
}
