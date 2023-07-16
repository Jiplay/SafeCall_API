package main

import "fmt"

func GetNotification(userID string) string {
	url := "http://localhost:8081/notification/" + userID
	res := getDataProfiler(userID, url)
	return res
}

func addNotificationHandler(UserID, Title, Content string, Status bool) string {
	url := "http://localhost:8081/notification/" + UserID + "/" + Title + "/" + Content + "/" + fmt.Sprintf("%t", Status)

	requestBody := map[string]interface{}{
		"guest1":  "guest1",
		"guest2":  "guest2",
		"subject": "subject",
		"date":    "date",
	}

	res := postDataProfiler(url, requestBody)
	return res
}

func delNotificationHandler(UserID, Title string) string {
	url := "http://localhost:8081/notification/" + UserID + "/" + Title
	requestBody := map[string]interface{}{
		"guest1":  "guest1",
		"guest2":  "guest2",
		"subject": "subject",
		"date":    "date",
	}
	res := postDataProfiler(url, requestBody)
	return res
}
