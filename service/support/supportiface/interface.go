// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package supportiface provides an interface to enable mocking the AWS Support service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package supportiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/support"
)

// SupportAPI provides an interface to enable mocking the
// support.Support service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Support.
//    func myFunc(svc supportiface.SupportAPI) bool {
//        // Make svc.AddAttachmentsToSet request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := support.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSupportClient struct {
//        supportiface.SupportAPI
//    }
//    func (m *mockSupportClient) AddAttachmentsToSet(input *support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSupportClient{}
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
type SupportAPI interface {
	AddAttachmentsToSet(*support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error)
	AddAttachmentsToSetWithContext(aws.Context, *support.AddAttachmentsToSetInput, ...request.Option) (*support.AddAttachmentsToSetOutput, error)
	AddAttachmentsToSetRequest(*support.AddAttachmentsToSetInput) (*request.Request, *support.AddAttachmentsToSetOutput)

	AddCommunicationToCase(*support.AddCommunicationToCaseInput) (*support.AddCommunicationToCaseOutput, error)
	AddCommunicationToCaseWithContext(aws.Context, *support.AddCommunicationToCaseInput, ...request.Option) (*support.AddCommunicationToCaseOutput, error)
	AddCommunicationToCaseRequest(*support.AddCommunicationToCaseInput) (*request.Request, *support.AddCommunicationToCaseOutput)

	CreateCase(*support.CreateCaseInput) (*support.CreateCaseOutput, error)
	CreateCaseWithContext(aws.Context, *support.CreateCaseInput, ...request.Option) (*support.CreateCaseOutput, error)
	CreateCaseRequest(*support.CreateCaseInput) (*request.Request, *support.CreateCaseOutput)

	DescribeAttachment(*support.DescribeAttachmentInput) (*support.DescribeAttachmentOutput, error)
	DescribeAttachmentWithContext(aws.Context, *support.DescribeAttachmentInput, ...request.Option) (*support.DescribeAttachmentOutput, error)
	DescribeAttachmentRequest(*support.DescribeAttachmentInput) (*request.Request, *support.DescribeAttachmentOutput)

	DescribeCases(*support.DescribeCasesInput) (*support.DescribeCasesOutput, error)
	DescribeCasesWithContext(aws.Context, *support.DescribeCasesInput, ...request.Option) (*support.DescribeCasesOutput, error)
	DescribeCasesRequest(*support.DescribeCasesInput) (*request.Request, *support.DescribeCasesOutput)

	DescribeCasesPages(*support.DescribeCasesInput, func(*support.DescribeCasesOutput, bool) bool) error
	DescribeCasesPagesWithContext(aws.Context, *support.DescribeCasesInput, func(*support.DescribeCasesOutput, bool) bool, ...request.Option) error

	DescribeCommunications(*support.DescribeCommunicationsInput) (*support.DescribeCommunicationsOutput, error)
	DescribeCommunicationsWithContext(aws.Context, *support.DescribeCommunicationsInput, ...request.Option) (*support.DescribeCommunicationsOutput, error)
	DescribeCommunicationsRequest(*support.DescribeCommunicationsInput) (*request.Request, *support.DescribeCommunicationsOutput)

	DescribeCommunicationsPages(*support.DescribeCommunicationsInput, func(*support.DescribeCommunicationsOutput, bool) bool) error
	DescribeCommunicationsPagesWithContext(aws.Context, *support.DescribeCommunicationsInput, func(*support.DescribeCommunicationsOutput, bool) bool, ...request.Option) error

	DescribeServices(*support.DescribeServicesInput) (*support.DescribeServicesOutput, error)
	DescribeServicesWithContext(aws.Context, *support.DescribeServicesInput, ...request.Option) (*support.DescribeServicesOutput, error)
	DescribeServicesRequest(*support.DescribeServicesInput) (*request.Request, *support.DescribeServicesOutput)

	DescribeSeverityLevels(*support.DescribeSeverityLevelsInput) (*support.DescribeSeverityLevelsOutput, error)
	DescribeSeverityLevelsWithContext(aws.Context, *support.DescribeSeverityLevelsInput, ...request.Option) (*support.DescribeSeverityLevelsOutput, error)
	DescribeSeverityLevelsRequest(*support.DescribeSeverityLevelsInput) (*request.Request, *support.DescribeSeverityLevelsOutput)

	DescribeTrustedAdvisorCheckRefreshStatuses(*support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error)
	DescribeTrustedAdvisorCheckRefreshStatusesWithContext(aws.Context, *support.DescribeTrustedAdvisorCheckRefreshStatusesInput, ...request.Option) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error)
	DescribeTrustedAdvisorCheckRefreshStatusesRequest(*support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*request.Request, *support.DescribeTrustedAdvisorCheckRefreshStatusesOutput)

	DescribeTrustedAdvisorCheckResult(*support.DescribeTrustedAdvisorCheckResultInput) (*support.DescribeTrustedAdvisorCheckResultOutput, error)
	DescribeTrustedAdvisorCheckResultWithContext(aws.Context, *support.DescribeTrustedAdvisorCheckResultInput, ...request.Option) (*support.DescribeTrustedAdvisorCheckResultOutput, error)
	DescribeTrustedAdvisorCheckResultRequest(*support.DescribeTrustedAdvisorCheckResultInput) (*request.Request, *support.DescribeTrustedAdvisorCheckResultOutput)

	DescribeTrustedAdvisorCheckSummaries(*support.DescribeTrustedAdvisorCheckSummariesInput) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error)
	DescribeTrustedAdvisorCheckSummariesWithContext(aws.Context, *support.DescribeTrustedAdvisorCheckSummariesInput, ...request.Option) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error)
	DescribeTrustedAdvisorCheckSummariesRequest(*support.DescribeTrustedAdvisorCheckSummariesInput) (*request.Request, *support.DescribeTrustedAdvisorCheckSummariesOutput)

	DescribeTrustedAdvisorChecks(*support.DescribeTrustedAdvisorChecksInput) (*support.DescribeTrustedAdvisorChecksOutput, error)
	DescribeTrustedAdvisorChecksWithContext(aws.Context, *support.DescribeTrustedAdvisorChecksInput, ...request.Option) (*support.DescribeTrustedAdvisorChecksOutput, error)
	DescribeTrustedAdvisorChecksRequest(*support.DescribeTrustedAdvisorChecksInput) (*request.Request, *support.DescribeTrustedAdvisorChecksOutput)

	RefreshTrustedAdvisorCheck(*support.RefreshTrustedAdvisorCheckInput) (*support.RefreshTrustedAdvisorCheckOutput, error)
	RefreshTrustedAdvisorCheckWithContext(aws.Context, *support.RefreshTrustedAdvisorCheckInput, ...request.Option) (*support.RefreshTrustedAdvisorCheckOutput, error)
	RefreshTrustedAdvisorCheckRequest(*support.RefreshTrustedAdvisorCheckInput) (*request.Request, *support.RefreshTrustedAdvisorCheckOutput)

	ResolveCase(*support.ResolveCaseInput) (*support.ResolveCaseOutput, error)
	ResolveCaseWithContext(aws.Context, *support.ResolveCaseInput, ...request.Option) (*support.ResolveCaseOutput, error)
	ResolveCaseRequest(*support.ResolveCaseInput) (*request.Request, *support.ResolveCaseOutput)
}

var _ SupportAPI = (*support.Support)(nil)
