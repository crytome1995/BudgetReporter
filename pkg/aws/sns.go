package aws

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SendSummary(username, message, region, arn string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Fatalf("Failed to create session for SNS due to %s", err)
	}
	svc := sns.New(sess)
	if svc != nil {
		result, err := svc.Publish(&sns.PublishInput{
			Message:  &message,
			TopicArn: &arn,
			MessageAttributes: map[string]*sns.MessageAttributeValue{
				"username": {
					DataType:    aws.String("String"),
					StringValue: aws.String(username),
				},
			},
		})
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(*result.MessageId)
	}

}
