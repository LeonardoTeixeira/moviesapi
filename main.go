package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Referência: https://medium.com/iq360/go-serverless-construindo-uma-api-usando-golang-e-aws-lambda-2a7d6a3019b9

func handleGetMovies(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	movies, err := GetMovies()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	js, err := json.Marshal(movies)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func handleCreateMovie(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var movie Movie
	err := json.Unmarshal([]byte(req.Body), &movie)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	err = CreateMovie(movie)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "{\"success\": true}",
	}, nil
}

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Path == "/movies" {
		if req.HTTPMethod == "GET" {
			return handleGetMovies(req)
		}
		if req.HTTPMethod == "POST" {
			return handleCreateMovie(req)
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       http.StatusText(http.StatusMethodNotAllowed),
	}, nil
}

func main() {
	lambda.Start(router)
}
