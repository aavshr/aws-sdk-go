// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pricing_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/awserr"
	"github.com/aavshr/aws-sdk-go/aws/session"
	"github.com/aavshr/aws-sdk-go/service/pricing"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To retrieve a list of services and service codes
//
// Retrieves the service for the given Service Code.
func ExamplePricing_DescribeServices_shared00() {
	svc := pricing.New(session.New())
	input := &pricing.DescribeServicesInput{
		FormatVersion: aws.String("aws_v1"),
		MaxResults:    aws.Int64(1),
		ServiceCode:   aws.String("AmazonEC2"),
	}

	result, err := svc.DescribeServices(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case pricing.ErrCodeInternalErrorException:
				fmt.Println(pricing.ErrCodeInternalErrorException, aerr.Error())
			case pricing.ErrCodeInvalidParameterException:
				fmt.Println(pricing.ErrCodeInvalidParameterException, aerr.Error())
			case pricing.ErrCodeNotFoundException:
				fmt.Println(pricing.ErrCodeNotFoundException, aerr.Error())
			case pricing.ErrCodeInvalidNextTokenException:
				fmt.Println(pricing.ErrCodeInvalidNextTokenException, aerr.Error())
			case pricing.ErrCodeExpiredNextTokenException:
				fmt.Println(pricing.ErrCodeExpiredNextTokenException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To retrieve a list of attribute values
//
// This operation returns a list of values available for the given attribute.
func ExamplePricing_GetAttributeValues_shared00() {
	svc := pricing.New(session.New())
	input := &pricing.GetAttributeValuesInput{
		AttributeName: aws.String("volumeType"),
		MaxResults:    aws.Int64(2),
		ServiceCode:   aws.String("AmazonEC2"),
	}

	result, err := svc.GetAttributeValues(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case pricing.ErrCodeInternalErrorException:
				fmt.Println(pricing.ErrCodeInternalErrorException, aerr.Error())
			case pricing.ErrCodeInvalidParameterException:
				fmt.Println(pricing.ErrCodeInvalidParameterException, aerr.Error())
			case pricing.ErrCodeNotFoundException:
				fmt.Println(pricing.ErrCodeNotFoundException, aerr.Error())
			case pricing.ErrCodeInvalidNextTokenException:
				fmt.Println(pricing.ErrCodeInvalidNextTokenException, aerr.Error())
			case pricing.ErrCodeExpiredNextTokenException:
				fmt.Println(pricing.ErrCodeExpiredNextTokenException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
