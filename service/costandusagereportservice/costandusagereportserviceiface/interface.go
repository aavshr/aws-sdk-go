// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package costandusagereportserviceiface provides an interface to enable mocking the AWS Cost and Usage Report Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package costandusagereportserviceiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/costandusagereportservice"
)

// CostandUsageReportServiceAPI provides an interface to enable mocking the
// costandusagereportservice.CostandUsageReportService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Cost and Usage Report Service.
//    func myFunc(svc costandusagereportserviceiface.CostandUsageReportServiceAPI) bool {
//        // Make svc.DeleteReportDefinition request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := costandusagereportservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCostandUsageReportServiceClient struct {
//        costandusagereportserviceiface.CostandUsageReportServiceAPI
//    }
//    func (m *mockCostandUsageReportServiceClient) DeleteReportDefinition(input *costandusagereportservice.DeleteReportDefinitionInput) (*costandusagereportservice.DeleteReportDefinitionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCostandUsageReportServiceClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CostandUsageReportServiceAPI interface {
	DeleteReportDefinition(*costandusagereportservice.DeleteReportDefinitionInput) (*costandusagereportservice.DeleteReportDefinitionOutput, error)
	DeleteReportDefinitionWithContext(aws.Context, *costandusagereportservice.DeleteReportDefinitionInput, ...request.Option) (*costandusagereportservice.DeleteReportDefinitionOutput, error)
	DeleteReportDefinitionRequest(*costandusagereportservice.DeleteReportDefinitionInput) (*request.Request, *costandusagereportservice.DeleteReportDefinitionOutput)

	DescribeReportDefinitions(*costandusagereportservice.DescribeReportDefinitionsInput) (*costandusagereportservice.DescribeReportDefinitionsOutput, error)
	DescribeReportDefinitionsWithContext(aws.Context, *costandusagereportservice.DescribeReportDefinitionsInput, ...request.Option) (*costandusagereportservice.DescribeReportDefinitionsOutput, error)
	DescribeReportDefinitionsRequest(*costandusagereportservice.DescribeReportDefinitionsInput) (*request.Request, *costandusagereportservice.DescribeReportDefinitionsOutput)

	DescribeReportDefinitionsPages(*costandusagereportservice.DescribeReportDefinitionsInput, func(*costandusagereportservice.DescribeReportDefinitionsOutput, bool) bool) error
	DescribeReportDefinitionsPagesWithContext(aws.Context, *costandusagereportservice.DescribeReportDefinitionsInput, func(*costandusagereportservice.DescribeReportDefinitionsOutput, bool) bool, ...request.Option) error

	ModifyReportDefinition(*costandusagereportservice.ModifyReportDefinitionInput) (*costandusagereportservice.ModifyReportDefinitionOutput, error)
	ModifyReportDefinitionWithContext(aws.Context, *costandusagereportservice.ModifyReportDefinitionInput, ...request.Option) (*costandusagereportservice.ModifyReportDefinitionOutput, error)
	ModifyReportDefinitionRequest(*costandusagereportservice.ModifyReportDefinitionInput) (*request.Request, *costandusagereportservice.ModifyReportDefinitionOutput)

	PutReportDefinition(*costandusagereportservice.PutReportDefinitionInput) (*costandusagereportservice.PutReportDefinitionOutput, error)
	PutReportDefinitionWithContext(aws.Context, *costandusagereportservice.PutReportDefinitionInput, ...request.Option) (*costandusagereportservice.PutReportDefinitionOutput, error)
	PutReportDefinitionRequest(*costandusagereportservice.PutReportDefinitionInput) (*request.Request, *costandusagereportservice.PutReportDefinitionOutput)
}

var _ CostandUsageReportServiceAPI = (*costandusagereportservice.CostandUsageReportService)(nil)
