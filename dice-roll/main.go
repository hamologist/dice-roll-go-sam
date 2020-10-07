package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/hamologist/dice-roll/pkg/evaluator"
	"github.com/hamologist/dice-roll/pkg/model"

	"github.com/aws/aws-lambda-go/lambda"
)

func defaultError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       err.Error(),
	}, err
}

func handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rollPayload := model.RollPayload{}
	err := json.Unmarshal([]byte(event.Body), &rollPayload)
	if err != nil {
		return defaultError(err)
	}

	rollResponse, err := evaluator.EvaluateRoll(rollPayload)
	if err != nil {
		return defaultError(err)
	}

	response, err := json.Marshal(rollResponse)
	if err != nil {
		return defaultError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(response),
	}, nil
}

func main() {
	lambda.Start(handler)
}
