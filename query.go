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

type Event struct {
	Guests    string `bson:"Guests"`
	Date      string `bson:"Date"`
	Subject   string `bson:"Subject"`
	Confirmed bool   `bson:"Confirmed"`
}

func AddUser(uri, login, psw, user, email string) bool {
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
		{Key: "email", Value: email},
		{Key: "psw", Value: psw},
		{Key: "code", Value: ""},
		{Key: "data", Value: user},
	})
	return true
}

func ProfilerRequest(url string) bool {
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

func GetUsers(uri, database string) []bson.M {
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
	loginCollection := quickstartDatabase.Collection(database)

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

func actionFriend(me, dest, action string) {

}

func editLoginInfo(uri, finder, new string, endpoint int) bool {
	src := ""
	dest := ""
	if endpoint == Password {
		src = "login"
		dest = "psw"
	} else if endpoint == Login {
		src = "email"
		dest = "code"
	} else if endpoint == Reset {
		src = "email"
		dest = "psw"
	}

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
	ProfileCollection := quickstartDatabase.Collection("loginInfo")

	filter := bson.D{{Key: src, Value: finder}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: dest, Value: new}}}}
	_, err = ProfileCollection.UpdateOne(ctx, filter, update)

	return err == nil
}

func getDataProfiler(userID, url string) string {
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

func postDataProfiler(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)

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
