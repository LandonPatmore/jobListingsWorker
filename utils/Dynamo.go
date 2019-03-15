package utils

import (
	"dataPullerWorker/types"
	"github.com/aws/aws-sdk-go/aws"
	session2 "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/landonp1203/goUtils/loggly"
)

var dynamoClient = createDynamoClient()

// Creates a dynamo client
func createDynamoClient() *dynamodb.DynamoDB {
	rKey := ReadAWSEnv()

	session, err := session2.NewSession(&aws.Config{
		Region: aws.String(rKey),
		// Credentials aren't here because we pass in ENV variables and the sdk auto detects them
	})

	if err != nil {
		loggly.Error(err)
	}

	// Create DynamoDB client
	client := dynamodb.New(session)

	return client
}

// Takes items in, marshals them, and then sends them to the database
func PutItem(job types.GithubJob) error {
	av, err := dynamodbattribute.MarshalMap(job)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Job-Listings"),
	}

	_, err = dynamoClient.PutItem(input)

	if err != nil {
		return err
	} else {
		loggly.Info(job)
	}

	return nil
}
