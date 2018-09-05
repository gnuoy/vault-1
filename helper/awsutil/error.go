package awsutil

import (
	awsRequest "github.com/aws/aws-sdk-go/aws/request"
	"github.com/hashicorp/vault/logical"
)

// CheckAWSError will examine an error and convert to a logical error if
// approrpiate. If no appropriate error is found, return nil
func CheckAWSError(err error) error {
	// IsErrorThrottle will check if the error returned is one that matches
	// known request limiting errors:
	// https://github.com/aws/aws-sdk-go/blob/488d634b5a699b9118ac2befb5135922b4a77210/aws/request/retryer.go#L35
	if awsRequest.IsErrorThrottle(err) {
		return logical.ErrUpstreamRateLimited
	}
	return nil
}
