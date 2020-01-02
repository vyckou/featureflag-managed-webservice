package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	configcat "gopkg.in/configcat/go-sdk.v1"
)

type contactv1 struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phonenumber"`
}

type contactv2 struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phonenumber"`
	Birthday    string `json:"bday"`
}

func list() (events.APIGatewayProxyResponse, error) {
	configCatAPIKey := os.Getenv("CONFIGCAT_APIKEY")
	client := configcat.NewClient(configCatAPIKey)

	isEnabled, _ := client.GetValue("enabled", false).(bool)
	isVersion2, _ := client.GetValue("version2", false).(bool)

	if !isEnabled {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotImplemented,
			Body:       "Not Implemented",
		}, nil
	}

	if isVersion2 {
		contact := contactv2{Name: "Jonas", PhoneNumber: "911", Birthday: "1990-11-11"}
		js, _ := json.Marshal(contact)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       string(js),
		}, nil
	}

	contact := contactv1{Name: "Jonas", PhoneNumber: "911"}
	js, _ := json.Marshal(contact)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func main() {
	lambda.Start(list)
}
