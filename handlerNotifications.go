package main

import "fmt"

func GetNotification(userID string) string {
	url := "http://profiler:8081/notification/" + userID
	res := getDataProfiler(userID, url)
	return res
}

func addNotificationHandler(UserID, Title, Content string, Status bool) string {
	url := "http://profiler:8081/AddNotification"

	requestBody := map[string]interface{}{
		"User":    UserID,
		"Title":   Title,
		"Content": Content,
		"Status":  fmt.Sprintf("%t", Status),
	}

	res := postDataProfiler(url, requestBody)
	return res
}

func delNotificationHandler(UserID, Title string) string {
	url := "http://profiler:8081/DelNotification"
	requestBody := map[string]interface{}{
		"User":  UserID,
		"Title": Title,
	}
	res := postDataProfiler(url, requestBody)
	return res
}
