package main

import (
	"strings"
)

func actionFriendHandler(url string) bool {
	requestBody := map[string]interface{}{
		"guest1": "FIXME",
	}
	postDataProfiler(url, requestBody)
	return true
}

func getFriends(userID string) []string {
	results := getDataProfiler(userID, "http://localhost:8081/friends/"+userID)

	dest := strings.Split(results, ":")
	s := strings.ReplaceAll(dest[1], ",", "")
	a := strings.Split(s[:len(s)-2], "\"")

	for i := 0; i < len(a); i++ {
		if a[i] == "" {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}

	return a[1:]
}
