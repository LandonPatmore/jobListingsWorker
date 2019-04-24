package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	session2 "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/landonp1203/goUtils/loggly"
	"jobListingsWorker/types"
)

var dynamoClient = createDynamoClient()

const TableName = "Job-Listings"

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
		TableName: aws.String(TableName),
	}

	_, err = dynamoClient.PutItem(input)

	if err != nil {
		loggly.Error(err)
		return err
	}

	loggly.Info(job)
	return nil
}

func GetAllItems() (items [] types.GithubJob, err error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
	}

	result, err := dynamoClient.Scan(params)

	if err != nil {
		loggly.Error(err)
		return nil, nil
	}

	var jobs [] types.GithubJob
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &jobs)

	if err != nil {
		loggly.Error(err)
		return nil, err
	}

	return jobs, nil
}

func GetRowCount() (count int64, err error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
	}

	result, err := dynamoClient.Scan(params)

	if err != nil {
		loggly.Error(err)
		return 0, nil
	}

	return *result.Count, nil
}
