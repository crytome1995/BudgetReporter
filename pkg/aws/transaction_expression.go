package aws

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

// BuildExpression username date
func BuildExpression(username, date string) (expression.Expression, error) {
	keyExp := expression.Key(userKey).Equal(expression.Value(username))
	filtDate := expression.Name(dateKey).Equal(expression.Value(date))
	expr, err := expression.NewBuilder().WithFilter(filtDate).WithKeyCondition(keyExp).Build()
	return expr, err
}
