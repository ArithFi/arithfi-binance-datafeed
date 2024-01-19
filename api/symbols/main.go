package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
	"time"
)

func GenerateResponse(Body string, Code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: Body, StatusCode: Code}
}

func HandleRequest(_ context.Context, request events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
	return GenerateResponse(strconv.FormatInt(time.Now().Unix(), 10), 200), nil
}

func main() {
	lambda.Start(HandleRequest)
}
