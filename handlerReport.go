package main

type Report struct {
	Username string `bson:"Username"`
	Date     string `bson:"Date"`
	Message  string `bson:"Message"`
}

func NewReportHandler(user, date, message string) {
	url := getCredentials()
	report := Report{user, date, message}
	AddReport(url.Uri, report)
}

func GetReportHandler() []Report {
	url := getCredentials()
	resp, _ := GetReportsQuery(url.Uri)
	return resp
}

func DelReportHandler(user, date string) string {
	url := getCredentials()
	resp := DeleteReport(url.Uri, user, date)

	if !resp { // if resp == false {
		return "No reports found"
	}
	return "Reports correctly deleted"

}
