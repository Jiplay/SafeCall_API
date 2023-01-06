package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddUser(uri, login, psw, user string) bool {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return false
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("userData")
	loginCollection := quickstartDatabase.Collection("loginInfo")

	loginCollection.InsertOne(ctx, bson.D{
		{Key: "login", Value: login},
		{Key: "psw", Value: psw},
		{Key: "data", Value: user},
	})
	return true
}

func CreateProfile(uri, login string) bool {
	url := "http://profiler:8081/create/" + login

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		fmt.Println(err)
		return false
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	return true
}

func GetUsers(uri string) []bson.M {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("userData")
	loginCollection := quickstartDatabase.Collection("loginInfo")

	cursor, err := loginCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var users []bson.M
	if err = cursor.All(ctx, &users); err != nil {
		log.Fatal(err)
	}

	return users
}

func UpdateProfile(uri, endpoint, userID, data string) bool {
	url := "http://profiler:8081/" + endpoint + "/" + userID + "/" + data

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		fmt.Println(err)
		return false
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	return true
}

func getProfile(userID string) string {
	url := "http://profiler:8081/Profile" + "/" + userID

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return "false"
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "false"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "false"
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	return string(body)
}

func searchNameQuery(username string) string {
	url := "http://profiler:8081/search" + "/" + username

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return "false"
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "false"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "false"
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	return string(body)
}
