package dynamoutils

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var table = aws.String("utils-tests")

var client = dynamodb.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})

func TestString(t *testing.T) {
	cases := []string{
		"test",
		"50",
		"john@example.com",
	}
	for _, v := range cases {
		key := &dynamodb.AttributeValue{S: aws.String("test_string_" + v)}
		_, err := client.PutItem(&dynamodb.PutItemInput{
			TableName: table,
			Item: map[string]*dynamodb.AttributeValue{
				"test_id": key,
				"test_value": FormatString(v),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		resp, err := client.GetItem(&dynamodb.GetItemInput{
			TableName: table,
			Key: map[string]*dynamodb.AttributeValue{
				"test_id": key,
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if v != ParseString(resp.Item["test_value"]) {
			t.Errorf("%q -> %q", v, *resp.Item["test_value"].S)
		}
	}
}

func TestTime(t *testing.T) {
	key := &dynamodb.AttributeValue{S: aws.String("test_time_now")}
	now := time.Now()
	_, err := client.PutItem(&dynamodb.PutItemInput{
		TableName: table,
		Item: map[string]*dynamodb.AttributeValue{
			"test_id": key,
			"test_value": FormatTime(now),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.GetItem(&dynamodb.GetItemInput{
		TableName: table,
		Key: map[string]*dynamodb.AttributeValue{
			"test_id": key,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	when, err := ParseTime(resp.Item["test_value"])
	if err != nil {
		t.Fatal(err)
	}
	if !now.Equal(when) {
		t.Error("retreived time does not equal stored time")
	}
}
