package entity

import (
	"fmt"
	"os"
	//"os"
	"time"

	//"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/session"
	//"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	//"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/**
TODO play with inheritence if we get more than one table
 */
type Tweet struct {
	EntityId int64     `json:"entityId"`
	Symbol   string    `json:"symbol"`
	Date     time.Time `json:"date"`
	Weight   int       `json:"weight"`
	Creator  string    `json:"creator"`
	Content  string    `json:"content"`
}

var (
	tableName = "Tweet"
	service   *dynamodb.DynamoDB
)

func initConnection() {
	awsSession, err := session.NewSession()

	if err != nil {
		fmt.Println("Failed initiating AWS session\n")
		fmt.Printf("%v\n", err)
	}

	service = dynamodb.New(awsSession)
}

func init() {
	//Create DynamoDB client
	initConnection()
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("entityId"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("symbol"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("symbol"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("entityId"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := service.CreateTable(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceInUseException:
				fmt.Println(dynamodb.ErrCodeResourceInUseException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
				os.Exit(1) // todo a restart option
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	fmt.Printf("Verifying %v table", tableName)
	err = service.WaitUntilTableExists(&dynamodb.DescribeTableInput{
		TableName: &tableName,
	})

	if err != nil {
		fmt.Printf("Failed to verify %v table", tableName)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func (t Tweet) Save() {
	status, err := dynamodbattribute.MarshalMap(t)

	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      status,
		TableName: aws.String(tableName),
	}
	fmt.Printf("%v\n", input)

	_, err = service.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
	}
}
