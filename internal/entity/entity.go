package entity

type Tweet struct {
	EntityId  	string  `json:"entity_id"`
	Creator 	string  `json:"expanded_url"`
	Content     string  `json:"url"`
}

// todo all this bs
func Save() {

}
//sess, err := session.NewSession(&aws.Config{
//Region: aws.String("us-west-2")},
//)

// Create DynamoDB client
//svc := dynamodb.New(sess)