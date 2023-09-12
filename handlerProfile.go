package main

import (
	"strings"
)

type Profile struct {
	FullName    string
	Description string
	PhoneNb     string
	Email       string
	ProfilePic  string
}

// Capital letters function to export them
func NewProfile(fullName, description, phoneNB, email, profilePic string) Profile {
	product := Profile{fullName, description, phoneNB, email, profilePic}
	return product
}

func StringToProfile(input string) Profile {
	data := strings.Split(input, "\"")

	return Profile{data[5], data[9], data[13], data[17], data[21]}
}
