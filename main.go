package main

import (
	"net/http"
	"strconv"

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

type PostMessageStruct struct {
	Username   string `bson:"Username"`
	Friendname string `bson:"Friendname"`
	Message    string `bson:"Message"`
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
	r.POST("/delete/:userID", deleteUser)

	r.POST("/manageFriend/:userID/:friend/:action", manageFriendEndpoint)
	r.POST("/replyFriend/:userID/:friend/:action", replyFriendEndpoint)
	r.GET("/listFriends/:userID", listFriends)

	r.POST("/addEvent/", addEventEndpoint)
	r.POST("/delEvent/:guest1/:guest2/:date", delEventEndpoint)
	r.GET("/listEvent/:userID", listEventEndpoint)

	r.POST("/notification/:UserID/:Title/:Content/:Status", addNotificationEndpoint)
	r.POST("/notification/:UserID/:Title", delNotificationEndpoint)
	r.GET("/notification/:UserID", GetUserNotification)

	r.GET("/conversations/:UserID", GetConversations)
	r.GET("/messages/:UserID/:FriendID", GetMessages)
	r.POST("/sendMessage", PostMessage)

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

func deleteUser(c *gin.Context) {
	userID := c.Param("userID")
	resp := deleteUserData(userID)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

type PostAddEventStruct struct {
	Guest1  string `bson:"Guest1"`
	Guest2  string `bson:"Guest2"`
	Subject string `bson:"Subject"`
	Date    string `bson:"Date"`
}

func addEventEndpoint(c *gin.Context) {
	var data PostAddEventStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := addEventHandler(data.Guest1, data.Guest2, data.Subject, data.Date)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func delEventEndpoint(c *gin.Context) {
	guest1 := c.Param("guest1")
	guest2 := c.Param("guest2")
	date := c.Param("date")

	resp := delEventHandler(guest1, guest2, date)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func listEventEndpoint(c *gin.Context) {
	user := c.Param("userID")

	a := listEventHandler(user)

	c.JSON(200, gin.H{
		"Success ": a,
	})
}

func GetUserNotification(c *gin.Context) {
	userID := c.Param("UserID")
	resp := GetNotification(userID)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func addNotificationEndpoint(c *gin.Context) {
	UserID := c.Param("UserID")
	Title := c.Param("Title")
	Content := c.Param("Content")
	Status := c.Param("Status")

	res, err := strconv.ParseBool(Status)

	if err != nil {
		c.JSON(200, gin.H{
			"Success ": err,
		})
	}

	resp := addNotificationHandler(UserID, Title, Content, res)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func delNotificationEndpoint(c *gin.Context) {
	UserID := c.Param("UserID")
	Title := c.Param("Title")

	resp := delNotificationHandler(UserID, Title)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func GetConversations(c *gin.Context) {
	userID := c.Param("UserID")
	resp := GetConversation(userID)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func GetMessages(c *gin.Context) {
	userID := c.Param("UserID")
	friendID := c.Param("FriendID")
	resp := GetMessagesHandler(userID, friendID)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func PostMessage(c *gin.Context) {
	var data PostMessageStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	PostMessageHandler(data.Username, data.Friendname, data.Message)

	c.JSON(200, gin.H{
		"Success ": "True",
	})
}
