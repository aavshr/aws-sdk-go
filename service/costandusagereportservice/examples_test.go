// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costandusagereportservice_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/awserr"
	"github.com/aavshr/aws-sdk-go/aws/session"
	"github.com/aavshr/aws-sdk-go/service/costandusagereportservice"
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

// To delete the AWS Cost and Usage report named ExampleReport.
//
// The following example deletes the AWS Cost and Usage report named ExampleReport.
func ExampleCostandUsageReportService_DeleteReportDefinition_shared00() {
	svc := costandusagereportservice.New(session.New())
	input := &costandusagereportservice.DeleteReportDefinitionInput{
		ReportName: aws.String("ExampleReport"),
	}

	result, err := svc.DeleteReportDefinition(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case costandusagereportservice.ErrCodeInternalErrorException:
				fmt.Println(costandusagereportservice.ErrCodeInternalErrorException, aerr.Error())
			case costandusagereportservice.ErrCodeValidationException:
				fmt.Println(costandusagereportservice.ErrCodeValidationException, aerr.Error())
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

// To list the AWS Cost and Usage reports for the account.
//
// The following example lists the AWS Cost and Usage reports for the account.
func ExampleCostandUsageReportService_DescribeReportDefinitions_shared00() {
	svc := costandusagereportservice.New(session.New())
	input := &costandusagereportservice.DescribeReportDefinitionsInput{
		MaxResults: aws.Int64(5),
	}

	result, err := svc.DescribeReportDefinitions(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case costandusagereportservice.ErrCodeInternalErrorException:
				fmt.Println(costandusagereportservice.ErrCodeInternalErrorException, aerr.Error())
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

// To create a report named ExampleReport.
//
// The following example creates a AWS Cost and Usage report named ExampleReport.
func ExampleCostandUsageReportService_PutReportDefinition_shared00() {
	svc := costandusagereportservice.New(session.New())
	input := &costandusagereportservice.PutReportDefinitionInput{
		ReportDefinition: &costandusagereportservice.ReportDefinition{
			AdditionalArtifacts: []*string{
				aws.String("REDSHIFT"),
				aws.String("QUICKSIGHT"),
			},
			AdditionalSchemaElements: []*string{
				aws.String("RESOURCES"),
			},
			Compression: aws.String("ZIP"),
			Format:      aws.String("textORcsv"),
			ReportName:  aws.String("ExampleReport"),
			S3Bucket:    aws.String("example-s3-bucket"),
			S3Prefix:    aws.String("exampleprefix"),
			S3Region:    aws.String("us-east-1"),
			TimeUnit:    aws.String("DAILY"),
		},
	}

	result, err := svc.PutReportDefinition(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case costandusagereportservice.ErrCodeDuplicateReportNameException:
				fmt.Println(costandusagereportservice.ErrCodeDuplicateReportNameException, aerr.Error())
			case costandusagereportservice.ErrCodeReportLimitReachedException:
				fmt.Println(costandusagereportservice.ErrCodeReportLimitReachedException, aerr.Error())
			case costandusagereportservice.ErrCodeInternalErrorException:
				fmt.Println(costandusagereportservice.ErrCodeInternalErrorException, aerr.Error())
			case costandusagereportservice.ErrCodeValidationException:
				fmt.Println(costandusagereportservice.ErrCodeValidationException, aerr.Error())
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
