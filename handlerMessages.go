package main

import "fmt"

func GetConversation(userID string) string {
	url := "http://facteur:3000/get_all_conv/" + userID
	resp := getFromMessage(url)
	fmt.Println(resp)
	return resp
}

func GetMessagesHandler(userID, friendID string) string {
	url := "http://facteur:3000/conv/" + userID + "/" + friendID
	resp := getFromMessage(url)
	fmt.Println(resp)
	return resp
}

func PostMessageHandler(userID, friendID, message string) {
	url := "http://facteur:3000/send_message"
	postFacteur(url, userID, friendID, message)
	// fmt.Println(resp)
	// return resp
}
