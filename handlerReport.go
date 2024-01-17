package main

type Report struct {
	Username string `bson:"Username"`
	Date     string `bson:"Date"`
	Message  string `bson:"Message"`
	Status   string `bson:"Status"`
}

func NewReportHandler(user, date, message string) {
	url := getCredentials()
	report := Report{user, date, message, "NEW"}
	AddReport(url.Uri, report)
}

func GetReportHandler() []Report {
	url := getCredentials()
	resp, _ := GetReportsQuery(url.Uri)
	return resp
}

func EditReportHandler(user, date, state string) bool {
	url := getCredentials()
	UpdateReport(url.Uri, user, date, state)
	return true
}

func DelReportHandler(user, date string) string {
	url := getCredentials()
	resp := DeleteReport(url.Uri, user, date)

	if !resp { // if resp == false {
		return "No reports found"
	}
	return "Reports correctly deleted"
}
