package main

func actionFriendHandler(url string) bool {
	// url := fmt.Sprintf("http://localhost:8081/friend/%s/%s/%s", me, dest, action)
	resp := ProfilerRequest(url)
	return resp
}

func getFriends(userID string) string {
	results := getDataProfiler(userID, "http://localhost:8081/friends/"+userID)
	return results
}
