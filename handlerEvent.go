package main

import (
	"encoding/json"
	"fmt"
)

func addEventHandler(guest1, guest2, subject, date string) string {
	url := "http://localhost:8081/addEvent"

	requestBody := map[string]interface{}{
		"Guest1":  guest1,
		"Guest2":  guest2,
		"Subject": subject,
		"Date":    date,
	}

	res := postDataProfiler(url, requestBody)
	return res
}

func delEventHandler(guest1, guest2, date string) string {
	url := "http://localhost:8081/delEvent"

	requestBody := map[string]interface{}{
		"Guest1": guest1,
		"Guest2": guest2,
		"Date":   date,
	}

	resp := postDataProfiler(url, requestBody)
	return resp
}

func listEventHandler(userID string) []Event {
	resp := getDataProfiler(userID, "http://localhost:8081/listEvent/"+userID)
	var events []Event
	err := json.Unmarshal([]byte(resp[12:len(resp)-1]), &events)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return events
}
