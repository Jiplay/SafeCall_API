package main

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
)

func weather(city string) string {

	url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&APPID=ab22018089b37f82cc084867cd1a3743&units=metric"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	return string(body)
}
