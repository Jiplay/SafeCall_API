package main

import (
	"fmt"
)

func GetNotification(userID string) string {
	url := "http://profiler:8081/notification/" + userID
	res := getDataProfiler(userID, url)
	return res
}

func addNotificationHandler(UserID, Title, Content string, Status bool) string {
	url := "http://profiler:8081/notification/" + UserID + "/" + Title + "/" + Content + "/" + fmt.Sprintf("%t", Status)
	res := postDataProfiler(url)
	return res
}

func delNotificationHandler(UserID, Title string) string {
	url := "http://profiler:8081/notification/" + UserID + "/" + Title
	res := postDataProfiler(url)
	return res
}
