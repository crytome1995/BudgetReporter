package main

import (
	"fmt"
	"log"
	"time"

	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/crytome1995/BudgetReporter/pkg/aws"
	c "github.com/crytome1995/BudgetReporter/pkg/config"
	"github.com/crytome1995/BudgetReporter/pkg/date"
	"github.com/spf13/viper"
)

type UserNameEvent struct {
	Username string `json:"userName"`
}

func reportBudget(userName string) {
	// location timezone
	loc, err := time.LoadLocation("America/Panama")
	if err != nil {
		panic(err)
	}
	timeUTC := time.Now()
	timeCentral := timeUTC.In(loc)
	// set name of config file
	viper.SetConfigName("config")
	// set location of config.yaml
	viper.AddConfigPath(".")
	// allow for viper to look at env vars
	viper.AutomaticEnv()
	// set file type
	// loading of configuration
	viper.SetConfigType("yml")
	var configuration c.Configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)

	}
	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	dates := date.GenerateDateRange(timeCentral)
	transactions := aws.QueryTransactionsByDate(configuration.Database.TableName, configuration.Aws.DynamoRegion, userName, dates)
	fmt.Printf("Transactions: %v\n", transactions)
	dateSum := date.Sum(transactions)
	aws.SendSummary(userName, fmt.Sprintf("Total spent this week $%.2f", dateSum), configuration.Aws.SnsRegion, configuration.Aws.SnsArn)
}

func HandleRequest(ctx context.Context, username UserNameEvent) (string, error) {
	reportBudget(username.Username)
	return fmt.Sprintf("Ran report for user %s!", username.Username), nil
}
func main() {
	lambda.Start(HandleRequest)
}
