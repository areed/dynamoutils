//Provides utility functions for marshaling and unmarshaling values for storage
//in DynamoDB with the official client.
package dynamoutils

import (
	"errors"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//FormatString converts a string to an *AttributeValue.
func FormatString(s string) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{S: aws.String(s)}
}

//FormatString converts an AttributeValue to a string.
func ParseString(v *dynamodb.AttributeValue) string {
	return *v.S
}

//FormatInt converts an int to an *AttributeValue.
func FormatInt(i int) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{N: aws.String(strconv.FormatInt(int64(i), 10))}
}

//ParseInt converts an *AttributeValue to an int.
func ParseInt(v *dynamodb.AttributeValue) (int, error) {
	i, err := strconv.ParseInt(*v.N, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

//FormatTime converts a time to an *AttributeValue. The time is stored as an
//integer representing nanoseconds since epoch, and is not necessarily
//compatible with other serializations.
func FormatTime(t time.Time) *dynamodb.AttributeValue {
	var n int64
	if !t.IsZero() {
		n = t.UnixNano()
	}
	return &dynamodb.AttributeValue{N: aws.String(strconv.FormatInt(n, 10))}
}

//ParseTime converts an AttributeValue to a time, assuming it is stored in the
//format used in FormatTime.
func ParseTime(v *dynamodb.AttributeValue) (time.Time, error) {
	nanoseconds, err := strconv.ParseInt(*v.N, 10, 64)
	if err != nil {
		return time.Time{}, errors.New("N: " + *v.N)
	}
	if nanoseconds == 0 {
		return time.Time{}, nil
	}
	return time.Unix(0, nanoseconds), nil
}

//FormatDuration converts a time.Duration to an *AttributeValue number.
func FormatDuration(d time.Duration) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{N: aws.String(strconv.FormatInt(int64(d), 10))}
}

func ParseDuration(v *dynamodb.AttributeValue) (time.Duration, error) {
	i, err := strconv.ParseInt(*v.N, 10, 64)
	if err != nil {
		return 0, err
	}
	return time.Duration(i), nil
}
