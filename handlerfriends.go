package main

func actionFriendHandler(me, dest, action string) string {

	resp, err := ProfilerRequest(me, dest, action)

	if !err {
		return "Internal Error"
	}
	return "200"
}
