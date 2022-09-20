package main

import (
	"strings"
)

type Profile struct {
	Name string `json:"name"`
	Nb   int    `json:"nb"`
}

type Weather struct {
	Temp    string `json:"temp"`
	Weather string `json:"weather"`
}

func loginHandler(id string, psw string) bool {

	response := account("login", id, psw)

	if strings.Contains(response, "success") == true {
		return true
	}
	return false
}

func registerHandler(id string, psw string) string {

	response := account("register", id, psw)

	if strings.Contains(response, "success") == true {
		return "200"
	}

	return response[11 : len(response)-2]
}
