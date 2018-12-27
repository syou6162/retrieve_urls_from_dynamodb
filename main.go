package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Example struct {
	Url string `json:"url"`
}

func main() {
	region := "us-east-1"
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(region)}))
	svc := dynamodb.New(sess)

	tableName := "examples-to-be-added"
	scanInput := dynamodb.ScanInput{
		TableName: &tableName,
	}

	output, err := svc.Scan(&scanInput)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, item := range output.Items {
		e := Example{}
		err = dynamodbattribute.UnmarshalMap(item, &e)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(e.Url)
	}
}
