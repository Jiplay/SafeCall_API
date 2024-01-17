package main

import (
	"fmt"
	"log"
	"net/http"
	sync "sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func acmeFunc(c *gin.Context) {
	content := "FzevAlxqxAFFyS97EBQ0A9d754RkAv7XJUuTHJyazRQ.clTE2bdjTGolWDmWXCivTvIFXqCv6e-Fb8n5oZ-FA9c"

	// Répondre avec le contenu du fichier
	c.Data(http.StatusOK, "text/plain", []byte(content))

}

func main() {
	// Utiliser le mode Release pour la production
	gin.SetMode(gin.ReleaseMode)

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("a user connected")
		s.Emit("me", s.ID())
		return nil
	})
	log.Println("WebSocket server created successfully")
	server.OnDisconnect("/", func(s socketio.Conn, _ string) {
		log.Println("a user disconnected")
		s.Emit("/", "callEnded")
	})

	server.OnEvent("/", "callUser", func(s socketio.Conn, data map[string]interface{}) {
		log.Println("user to call:", data["userToCall"])
		targetSocketID := data["userToCall"].(string)
		server.BroadcastToNamespace(targetSocketID, "callUser", map[string]interface{}{
			"signal": data["signalData"],
			"from":   data["from"],
		})
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnEvent("/", "answerCall", func(s socketio.Conn, data map[string]interface{}) {
		log.Println("answer to:", data["to"])
		targetSocketID := data["to"].(string)
		server.BroadcastToNamespace(targetSocketID, "callAccepted", data["signal"])
	})

	// Créer un routeur Gin
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Remplacez par vos origines autorisées
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}
	r.Use(cors.New(config))

	// r.Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "https://safecall-web.vercel.app")
	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// 	if c.Request.Method == "OPTIONS" {
	// 		c.AbortWithStatus(http.StatusOK)
	// 		return
	// 	}

	// 	c.Next()
	// })
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	r.StaticFS("/public", http.Dir("../asset"))

	r.GET("/.well-known/acme-challenge/:data", acmeFunc)

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
	r.POST("/editReport", EditReportEndpoint)

	r.GET("/setupProfiler", SetupProfiler)
	r.GET("/tryCall", sendCall)

	// Configurer le serveur HTTPS
	portHTTPS := 443
	certFile := "certificates/cert.pem"
	keyFile := "certificates/privkey.pem"

	// Configurer le serveur HTTP
	portHTTP := 80

	var wg sync.WaitGroup

	// Lancer le serveur HTTPS dans une goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := r.RunTLS(fmt.Sprintf(":%d", portHTTPS), certFile, keyFile)
		if err != nil {
			log.Fatal("Erreur lors du démarrage du serveur HTTPS : ", err)
		}
	}()

	// Lancer le serveur HTTP dans une goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := http.ListenAndServe(fmt.Sprintf(":%d", portHTTP), r)
		if err != nil {
			log.Fatal("Erreur lors du démarrage du serveur HTTP : ", err)
		}
	}()
	// http.ListenAndServe(fmt.Sprintf(":%d", portHTTP), r)
	// Attendre que les serveurs se terminent
	wg.Wait()
}
