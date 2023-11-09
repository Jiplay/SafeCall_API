package main

import (
	"encoding/json"
	"fmt"
)

type Friends struct {
	Id      string `bson:"Id"`
	Subject string `bson:"Subject"`
	Active  bool   `bson:"Active"`
}

func actionFriendHandler(url string, body map[string]interface{}) bool {
	postDataProfiler(url, body)
	return true
}

func getFriends(userID string) []Friends {
	resp := getDataProfiler(userID, "http://localhost:8081/friends/"+userID)

	var friends []Friends
	err := json.Unmarshal([]byte(resp[12:len(resp)-1]), &friends)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return friends
}
