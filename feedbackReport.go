package main

type Feedback struct {
	Username string `bson:"Username"`
	Date     string `bson:"Date"`
	Message  string `bson:"Message"`
}

func NewFeedbackHandler(user, date, message string) {
	url := getCredentials()
	feedback := Feedback{user, date, message}
	AddFeedback(url.Uri, feedback)
}

func GetFeedbackHandler() []Feedback {
	url := getCredentials()
	resp, _ := GetFeedbacks(url.Uri)
	return resp
}

func DelFeedbackHandler(user, date string) string {
	url := getCredentials()
	resp := DeleteFeedback(url.Uri, user, date)

	if !resp { // if resp == false {
		return "No feedback found"
	}
	return "Feedback correctly deleted"

}
