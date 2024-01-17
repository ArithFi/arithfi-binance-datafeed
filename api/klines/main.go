package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func GenerateResponse(body string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Extract query parameters
	symbol := url.QueryEscape(request.QueryStringParameters["symbol"])
	interval := url.QueryEscape(request.QueryStringParameters["interval"])
	startTime := url.QueryEscape(request.QueryStringParameters["startTime"])
	endTime := url.QueryEscape(request.QueryStringParameters["endTime"])

	// Construct the Binance API URL
	apiUrl := fmt.Sprintf("https://fapi.binance.com/fapi/v1/klines?symbol=%s&interval=%s&startTime=%s&endTime=%s&limit=500", symbol, interval, startTime, endTime)

	// Create an HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return GenerateResponse(fmt.Sprintf(`{"error": "Error creating request: %s"}`, err.Error()), http.StatusInternalServerError), nil
	}

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return GenerateResponse(fmt.Sprintf(`{"error": "Error making request to Binance API: %s"}`, err.Error()), http.StatusInternalServerError), nil
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GenerateResponse(fmt.Sprintf(`{"error": "Error reading response body: %s"}`, err.Error()), http.StatusInternalServerError), nil
	}

	// Return the response from the Binance API
	return GenerateResponse(string(body), http.StatusOK), nil
}

func main() {
	lambda.Start(HandleRequest)
}
