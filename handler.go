package main

import (
// "strconv"
// "strings"
)

type Profile struct {
	Name string `json:"name"`
	Nb   int    `json:"nb"`
}

type Weather struct {
	Temp    string `json:"temp"`
	Weather string `json:"weather"`
}

// func demo(uid, nb string) Profile {
// 	nbAsStr, _ := strconv.Atoi(nb)
// 	userProfile := Profile{Name: uid, Nb: nbAsStr}

// 	return userProfile
// }

// func loginHandler(id, psw string) bool {

// resp := weather(city)
// parsing := strings.Split(resp, ":")
// weather := ""
// temp := ""

// for _, s := range parsing {
// 	if strings.Contains(s, "description") == true {
// 		a := strings.Split(s, ",")
// 		weather = a[0][1 : len(a[0])-1]
// 	}
// 	if strings.Contains(s, "feels_like") == true {
// 		a := strings.Split(s, ",")
// 		temp = a[0]
// 	}
// }

// dest := Weather{Temp: temp, Weather: weather}

// 	return true
// }
