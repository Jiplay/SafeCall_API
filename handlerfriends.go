package main

import (
	"strings"
)

func actionFriendHandler(url string) bool {
	// url := fmt.Sprintf("http://localhost:8081/friend/%s/%s/%s", me, dest, action)
	resp := ProfilerRequest(url)
	return resp
}

func getFriends(userID string) []string {
	results := getDataProfiler(userID, "http://profiler:8081/friends/"+userID)

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
