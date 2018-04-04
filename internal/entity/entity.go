package entity

import (
	"fmt"
	"time"

	//"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/session"
	//"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Tweet struct {
	EntityId int64     `json:"entity_id"`
	Date     time.Time `json:"date"`
	Weight   int       `json:"weight"`
	Creator  string    `json:"creator"`
	Content  string    `json:"content"`
}

func init() {
	//sess, _ := session.NewSession(&aws.Config{
	//	Region: aws.String("us-west-2")},
	//)

	// Create DynamoDB client
	//svc := dynamodb.New(sess)
	//
	//input := &dynamodb.CreateTableInput{
	//	AttributeDefinitions: []*dynamodb.AttributeDefinition{
	//		{
	//			AttributeName: aws.String("creator"),
	//			AttributeType: aws.String("S"),
	//		},
	//		{
	//			AttributeName: aws.String("content"),
	//			AttributeType: aws.String("S"),
	//		},
	//	},
	//	KeySchema: []*dynamodb.KeySchemaElement{
	//		{
	//			AttributeName: aws.String("year"),
	//			KeyType:       aws.String("HASH"),
	//		},
	//		{
	//			AttributeName: aws.String("title"),
	//			KeyType:       aws.String("RANGE"),
	//		},
	//		{
	//			AttributeName: aws.String("entityId"),
	//			KeyType:       aws.String("N"),
	//		},
	//	},
	//	ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
	//		ReadCapacityUnits:  aws.Int64(10),
	//		WriteCapacityUnits: aws.Int64(10),
	//	},
	//	TableName: aws.String("Movies"),
	//}
}

// todo all this bs
func (t Tweet) Save() {
	// TODO
	fmt.Printf("Id:%v Date: %v Creator: %v Wieght: %v Content: %v\n", t.EntityId, t.Date, t.Creator, t.Weight, t.Content)
}

//sess, err := session.NewSession(&aws.Config{
//Region: aws.String("us-west-2")},
//)

// Create DynamoDB client
//svc := dynamodb.New(sess)
