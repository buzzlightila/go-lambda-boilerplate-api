package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	log "github.com/sirupsen/logrus"

	"net/http"
)

var (
	headers = map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Headers": "Content-Type",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "OPTIONS,GET",
	}
)

// Start
func main() {
	lambda.Start(router)
}

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.Path {
	case "/ping":
		return HandlerPing(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func HandlerPing(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := Ping(req)
	if err != nil {
		if re, ok := err.(*RequestError); ok {
			if re.ErrorCode() != 0 {
				return clientError(re.ErrorCode())
			}
		}
		return serverError(err)
	}
	return events.APIGatewayProxyResponse{
		Headers:    headers,
		StatusCode: http.StatusOK,
		//Body:       string(res),
	}, nil
}

func Ping(req events.APIGatewayProxyRequest) error {
	return nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Headers:    headers,
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	log.Error(err)
	return events.APIGatewayProxyResponse{
		Headers:    headers,
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}
