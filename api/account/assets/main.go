package account_assets

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Balance struct {
	Asset string  `json:"asset"`
	Free  float32 `json:"free"`
	Copy  float32 `json:"copy"`
}

type Response struct {
	Balances []Balance `json:"balances"`
}

func GenerateResponse(Body string, Code int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: Body, StatusCode: Code}
}

func HandleRequest(_ context.Context, request events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
	balances := []Balance{
		{
			Asset: "ATF",
			Free:  10000.0,
			Copy:  200.0,
		},
	}

	response := Response{Balances: balances}

	jsonData, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return GenerateResponse(string(jsonData), 200), nil
}

func main() {
	lambda.Start(HandleRequest)
}
