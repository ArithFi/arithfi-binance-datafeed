package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Config struct {
	SupportedResolutions   []string `json:"supported_resolutions"`
	SupportsSearch         bool     `json:"supports_search"`
	SupportsGroupRequest   bool     `json:"supports_group_request"`
	SupportsMarks          bool     `json:"supports_marks"`
	SupportsTimescaleMarks bool     `json:"supports_timescale_marks"`
	SupportsTime           bool     `json:"supports_time"`
}

func GenerateResponse(Body string, Code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: Body, StatusCode: Code}
}

func HandleRequest(_ context.Context, request events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
	config := Config{
		SupportsSearch:         true,
		SupportsGroupRequest:   false,
		SupportsMarks:          false,
		SupportsTimescaleMarks: false,
		SupportedResolutions:   []string{"1", "3", "5", "15", "30", "60", "120", "240", "360", "480", "720", "1D", "3D", "1W", "1M"},
		SupportsTime:           true,
	}
	jsonData, err := json.Marshal(config)
	if err != nil {
		return GenerateResponse(err.Error(), 500), nil
	}
	return GenerateResponse(string(jsonData), 200), nil
}

func main() {
	lambda.Start(HandleRequest)
}
