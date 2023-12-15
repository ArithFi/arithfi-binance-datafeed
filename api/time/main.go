package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

type Response struct {
	ServerTime int64 `json:"serverTime"`
}

func GenerateResponse(Body string, Code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: Body, StatusCode: Code}
}

func HandleRequest(_ context.Context, request events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
	data := []Response{{ServerTime: time.Now().Unix()}}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return GenerateResponse(string(jsonData), 200), nil
}

func main() {
	lambda.Start(HandleRequest)
}
