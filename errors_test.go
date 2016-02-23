package dynamoutils

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestIsConditionFailed(t *testing.T) {
	key := &dynamodb.AttributeValue{S: aws.String("test_condition_failed")}
	val := FormatString(time.Now().Format(time.ANSIC))
	_, err := client.PutItem(&dynamodb.PutItemInput{
		TableName: table,
		Item: map[string]*dynamodb.AttributeValue{
			"test_id": key,
			"test_value": val,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if IsConditionFailed(err) {
		t.Error()
	}
	_, err = client.PutItem(&dynamodb.PutItemInput{
		TableName: table,
		Item: map[string]*dynamodb.AttributeValue{
			"test_id": key,
			"test_value": val,
		},
		ConditionExpression: aws.String("test_value <> :test_value"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":test_value": val,
		},
	})
	if !IsConditionFailed(err) {
		t.Error()
	}

}
