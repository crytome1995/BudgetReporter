package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// attribute contants for transaction
const (
	userKey     = "user"
	dateKey     = "date"
	amountKey   = "amount"
	businessKey = "business"
	cardKey     = "card"
)

// transaction representation in dynamodb
type Transaction struct {
	User     string
	Uuid     string
	Amount   float64
	Business string
	Card     string
	Date     string
}

// comment
func QueryTransactionsByDate(tableName string, region string, username string, dates []string) []Transaction {
	var transactions []Transaction
	// setup the dynamodb session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Fatalf("%s", err)
	}
	svc := dynamodb.New(sess)
	if svc != nil {
		for r := 0; r < len(dates); {
			date := dates[r]
			log.Printf("Obtaining transactions for date %s", date)
			expr, err := BuildExpression(username, date)
			if err != nil {
				log.Fatalf("Failed to build expression because of: %s", err)
			}
			input := &dynamodb.QueryInput{
				ExpressionAttributeValues: expr.Values(),
				ExpressionAttributeNames:  expr.Names(),
				KeyConditionExpression:    expr.KeyCondition(),
				TableName:                 aws.String(tableName),
				FilterExpression:          expr.Filter(),
			}
			result, err := svc.Query(input)
			if err != nil {
				log.Fatalf("Got error building expression: %s", err)
			}
			transactionItem := Transaction{}
			if *result.Count > int64(0) {
				items := result.Items
				for i := 0; i < len(items); {
					err = dynamodbattribute.UnmarshalMap(items[i], &transactionItem)
					if err != nil {
						log.Fatalf("Failed to unmarshal Record, %v", err)
					}
					transactions = append(transactions, transactionItem)
					i++
				}
			}
			r++
		}
	}
	return transactions
}
