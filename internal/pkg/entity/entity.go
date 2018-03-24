package entity

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Tweet struct {
	EntityId  	string  `json:"entity_id"`
	Creator 	string  `json:"expanded_url"`
	Content     string  `json:"url"`
}


func Save() {

}
//sess, err := session.NewSession(&aws.Config{
//Region: aws.String("us-west-2")},
//)

// Create DynamoDB client
//svc := dynamodb.New(sess)