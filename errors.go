package dynamoutils

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
)

//IsConditionalCheckFailed returns true if an error is an awserr.Error with code
//"CondtionalCheckFailedException"
func IsConditionFailed(err error) bool {
	if err == nil {
		return false
	}
	if awsErr, ok := err.(awserr.Error); ok {
		return awsErr.Code() == "ConditionalCheckFailedException"
	}
	return false
}
