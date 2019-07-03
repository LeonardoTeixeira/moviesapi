package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/satori/go.uuid"
)

const AWS_REGION = "us-east-1"
const TABLE_NAME = "movies"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(AWS_REGION))

func GetMovies() ([]Movie, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String(TABLE_NAME),
	}
	result, err := db.Scan(input)
	if err != nil {
		return []Movie{}, err
	}
	if len(result.Items) == 0 {
		return []Movie{}, nil
	}

	var movies []Movie
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &movies)
	if err != nil {
		return []Movie{}, err
	}

	return movies, nil
}

func CreateMovie(movie Movie) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(fmt.Sprintf("%v", uuid)),
			},
			"name": {
				S: aws.String(movie.Name),
			},
			"image": {
				S: aws.String(movie.Image),
			},
		},
	}

	_, err = db.PutItem(input)
	return err
}
