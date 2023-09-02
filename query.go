package main

import (
	"bytes"
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

func getProfile(userID string) string {
	url := "http://localhost:8081/Profile" + "/" + userID

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
	url := "http://localhost:8081/search" + "/" + username

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

func postDataProfiler(url string, requestBody map[string]interface{}) string {

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return ""
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)

}

// getProfile clone
func getFromMessage(url string) [][]string {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var dat [][]string
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	return dat
}

func getAllConvQuery(url string) []string {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var dat []string
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	return dat
}

func postFacteur(url, userID, dest, message string) {

	requestBody := map[string]interface{}{
		"username":   userID,
		"friendname": dest,
		"message":    message,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func AddFeedback(uri string, feedback Feedback) bool {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return false
	}

	defer client.Disconnect(ctx)

	database := client.Database("Support")
	collection := database.Collection("Feedback")

	_, err = collection.InsertOne(ctx, feedback)
	if err != nil {
		fmt.Println("Failed to insert feedback:", err)
		return false
	}

	return true
}

func DeleteFeedback(uri string, username string, date string) bool {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return false
	}

	defer client.Disconnect(ctx)

	database := client.Database("Support")
	collection := database.Collection("Feedback")

	filter := bson.M{
		"Username": username,
		"Date":     date,
	}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println("Failed to delete feedback:", err)
		return false
	}

	if result.DeletedCount == 0 {
		fmt.Println("No feedback deleted")
		return false
	}

	return true
}

func GetFeedbacks(uri string) ([]Feedback, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer client.Disconnect(ctx)

	database := client.Database("Support")
	collection := database.Collection("Feedback")

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Failed to retrieve feedbacks:", err)
		return nil, err
	}

	var feedbacks []Feedback
	err = cur.All(ctx, &feedbacks)
	if err != nil {
		fmt.Println("Failed to decode feedbacks:", err)
		return nil, err
	}

	return feedbacks, nil
}

func AddReport(uri string, report Report) bool {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return false
	}

	defer client.Disconnect(ctx)

	database := client.Database("Support")
	collection := database.Collection("Report")

	_, err = collection.InsertOne(ctx, report)
	if err != nil {
		fmt.Println("Failed to insert feedback:", err)
		return false
	}

	return true
}

func GetReportsQuery(uri string) ([]Report, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer client.Disconnect(ctx)

	database := client.Database("Support")
	collection := database.Collection("Report")

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Failed to retrieve Reports:", err)
		return nil, err
	}

	var reports []Report
	err = cur.All(ctx, &reports)
	if err != nil {
		fmt.Println("Failed to decode Reports:", err)
		return nil, err
	}

	return reports, nil
}

func DeleteReport(uri string, username string, date string) bool {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return false
	}

	defer client.Disconnect(ctx)

	database := client.Database("Support")
	collection := database.Collection("Report")

	filter := bson.M{
		"Username": username,
		"Date":     date,
	}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println("Failed to delete report:", err)
		return false
	}

	if result.DeletedCount == 0 {
		fmt.Println("No report deleted")
		return false
	}

	return true
}
