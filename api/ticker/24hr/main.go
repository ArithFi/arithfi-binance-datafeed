package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"os"
	"time"
)

type Response struct {
	ServerTime int64 `json:"serverTime"`
}

func GenerateResponse(Body string, Code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: Body, StatusCode: Code}
}

func HandleRequest(_ context.Context, request events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
	url := os.Getenv("UPSTASH_REDIS_REST_URL")
	response, err := http.NewRequest("GET", url+"/set/foo/bar", nil)
	if err != nil {
		fmt.Println("请求失败：", err)
	}
	token := os.Getenv("UPSTASH_REDIS_REST_TOKEN")
	response.Header.Set("Authorization", "Bearer "+token)
	defer response.Body.Close()

	data := Response{ServerTime: time.Now().Unix()}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return GenerateResponse(string(jsonData), 200), nil
}

func main() {
	lambda.Start(HandleRequest)
}
