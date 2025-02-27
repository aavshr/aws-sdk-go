// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mediapackageiface provides an interface to enable mocking the AWS Elemental MediaPackage service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mediapackageiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/mediapackage"
)

// MediaPackageAPI provides an interface to enable mocking the
// mediapackage.MediaPackage service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Elemental MediaPackage.
//    func myFunc(svc mediapackageiface.MediaPackageAPI) bool {
//        // Make svc.ConfigureLogs request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := mediapackage.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMediaPackageClient struct {
//        mediapackageiface.MediaPackageAPI
//    }
//    func (m *mockMediaPackageClient) ConfigureLogs(input *mediapackage.ConfigureLogsInput) (*mediapackage.ConfigureLogsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMediaPackageClient{}
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
type MediaPackageAPI interface {
	ConfigureLogs(*mediapackage.ConfigureLogsInput) (*mediapackage.ConfigureLogsOutput, error)
	ConfigureLogsWithContext(aws.Context, *mediapackage.ConfigureLogsInput, ...request.Option) (*mediapackage.ConfigureLogsOutput, error)
	ConfigureLogsRequest(*mediapackage.ConfigureLogsInput) (*request.Request, *mediapackage.ConfigureLogsOutput)

	CreateChannel(*mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error)
	CreateChannelWithContext(aws.Context, *mediapackage.CreateChannelInput, ...request.Option) (*mediapackage.CreateChannelOutput, error)
	CreateChannelRequest(*mediapackage.CreateChannelInput) (*request.Request, *mediapackage.CreateChannelOutput)

	CreateHarvestJob(*mediapackage.CreateHarvestJobInput) (*mediapackage.CreateHarvestJobOutput, error)
	CreateHarvestJobWithContext(aws.Context, *mediapackage.CreateHarvestJobInput, ...request.Option) (*mediapackage.CreateHarvestJobOutput, error)
	CreateHarvestJobRequest(*mediapackage.CreateHarvestJobInput) (*request.Request, *mediapackage.CreateHarvestJobOutput)

	CreateOriginEndpoint(*mediapackage.CreateOriginEndpointInput) (*mediapackage.CreateOriginEndpointOutput, error)
	CreateOriginEndpointWithContext(aws.Context, *mediapackage.CreateOriginEndpointInput, ...request.Option) (*mediapackage.CreateOriginEndpointOutput, error)
	CreateOriginEndpointRequest(*mediapackage.CreateOriginEndpointInput) (*request.Request, *mediapackage.CreateOriginEndpointOutput)

	DeleteChannel(*mediapackage.DeleteChannelInput) (*mediapackage.DeleteChannelOutput, error)
	DeleteChannelWithContext(aws.Context, *mediapackage.DeleteChannelInput, ...request.Option) (*mediapackage.DeleteChannelOutput, error)
	DeleteChannelRequest(*mediapackage.DeleteChannelInput) (*request.Request, *mediapackage.DeleteChannelOutput)

	DeleteOriginEndpoint(*mediapackage.DeleteOriginEndpointInput) (*mediapackage.DeleteOriginEndpointOutput, error)
	DeleteOriginEndpointWithContext(aws.Context, *mediapackage.DeleteOriginEndpointInput, ...request.Option) (*mediapackage.DeleteOriginEndpointOutput, error)
	DeleteOriginEndpointRequest(*mediapackage.DeleteOriginEndpointInput) (*request.Request, *mediapackage.DeleteOriginEndpointOutput)

	DescribeChannel(*mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error)
	DescribeChannelWithContext(aws.Context, *mediapackage.DescribeChannelInput, ...request.Option) (*mediapackage.DescribeChannelOutput, error)
	DescribeChannelRequest(*mediapackage.DescribeChannelInput) (*request.Request, *mediapackage.DescribeChannelOutput)

	DescribeHarvestJob(*mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error)
	DescribeHarvestJobWithContext(aws.Context, *mediapackage.DescribeHarvestJobInput, ...request.Option) (*mediapackage.DescribeHarvestJobOutput, error)
	DescribeHarvestJobRequest(*mediapackage.DescribeHarvestJobInput) (*request.Request, *mediapackage.DescribeHarvestJobOutput)

	DescribeOriginEndpoint(*mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error)
	DescribeOriginEndpointWithContext(aws.Context, *mediapackage.DescribeOriginEndpointInput, ...request.Option) (*mediapackage.DescribeOriginEndpointOutput, error)
	DescribeOriginEndpointRequest(*mediapackage.DescribeOriginEndpointInput) (*request.Request, *mediapackage.DescribeOriginEndpointOutput)

	ListChannels(*mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error)
	ListChannelsWithContext(aws.Context, *mediapackage.ListChannelsInput, ...request.Option) (*mediapackage.ListChannelsOutput, error)
	ListChannelsRequest(*mediapackage.ListChannelsInput) (*request.Request, *mediapackage.ListChannelsOutput)

	ListChannelsPages(*mediapackage.ListChannelsInput, func(*mediapackage.ListChannelsOutput, bool) bool) error
	ListChannelsPagesWithContext(aws.Context, *mediapackage.ListChannelsInput, func(*mediapackage.ListChannelsOutput, bool) bool, ...request.Option) error

	ListHarvestJobs(*mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error)
	ListHarvestJobsWithContext(aws.Context, *mediapackage.ListHarvestJobsInput, ...request.Option) (*mediapackage.ListHarvestJobsOutput, error)
	ListHarvestJobsRequest(*mediapackage.ListHarvestJobsInput) (*request.Request, *mediapackage.ListHarvestJobsOutput)

	ListHarvestJobsPages(*mediapackage.ListHarvestJobsInput, func(*mediapackage.ListHarvestJobsOutput, bool) bool) error
	ListHarvestJobsPagesWithContext(aws.Context, *mediapackage.ListHarvestJobsInput, func(*mediapackage.ListHarvestJobsOutput, bool) bool, ...request.Option) error

	ListOriginEndpoints(*mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error)
	ListOriginEndpointsWithContext(aws.Context, *mediapackage.ListOriginEndpointsInput, ...request.Option) (*mediapackage.ListOriginEndpointsOutput, error)
	ListOriginEndpointsRequest(*mediapackage.ListOriginEndpointsInput) (*request.Request, *mediapackage.ListOriginEndpointsOutput)

	ListOriginEndpointsPages(*mediapackage.ListOriginEndpointsInput, func(*mediapackage.ListOriginEndpointsOutput, bool) bool) error
	ListOriginEndpointsPagesWithContext(aws.Context, *mediapackage.ListOriginEndpointsInput, func(*mediapackage.ListOriginEndpointsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *mediapackage.ListTagsForResourceInput, ...request.Option) (*mediapackage.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*mediapackage.ListTagsForResourceInput) (*request.Request, *mediapackage.ListTagsForResourceOutput)

	RotateChannelCredentials(*mediapackage.RotateChannelCredentialsInput) (*mediapackage.RotateChannelCredentialsOutput, error)
	RotateChannelCredentialsWithContext(aws.Context, *mediapackage.RotateChannelCredentialsInput, ...request.Option) (*mediapackage.RotateChannelCredentialsOutput, error)
	RotateChannelCredentialsRequest(*mediapackage.RotateChannelCredentialsInput) (*request.Request, *mediapackage.RotateChannelCredentialsOutput)

	RotateIngestEndpointCredentials(*mediapackage.RotateIngestEndpointCredentialsInput) (*mediapackage.RotateIngestEndpointCredentialsOutput, error)
	RotateIngestEndpointCredentialsWithContext(aws.Context, *mediapackage.RotateIngestEndpointCredentialsInput, ...request.Option) (*mediapackage.RotateIngestEndpointCredentialsOutput, error)
	RotateIngestEndpointCredentialsRequest(*mediapackage.RotateIngestEndpointCredentialsInput) (*request.Request, *mediapackage.RotateIngestEndpointCredentialsOutput)

	TagResource(*mediapackage.TagResourceInput) (*mediapackage.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *mediapackage.TagResourceInput, ...request.Option) (*mediapackage.TagResourceOutput, error)
	TagResourceRequest(*mediapackage.TagResourceInput) (*request.Request, *mediapackage.TagResourceOutput)

	UntagResource(*mediapackage.UntagResourceInput) (*mediapackage.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *mediapackage.UntagResourceInput, ...request.Option) (*mediapackage.UntagResourceOutput, error)
	UntagResourceRequest(*mediapackage.UntagResourceInput) (*request.Request, *mediapackage.UntagResourceOutput)

	UpdateChannel(*mediapackage.UpdateChannelInput) (*mediapackage.UpdateChannelOutput, error)
	UpdateChannelWithContext(aws.Context, *mediapackage.UpdateChannelInput, ...request.Option) (*mediapackage.UpdateChannelOutput, error)
	UpdateChannelRequest(*mediapackage.UpdateChannelInput) (*request.Request, *mediapackage.UpdateChannelOutput)

	UpdateOriginEndpoint(*mediapackage.UpdateOriginEndpointInput) (*mediapackage.UpdateOriginEndpointOutput, error)
	UpdateOriginEndpointWithContext(aws.Context, *mediapackage.UpdateOriginEndpointInput, ...request.Option) (*mediapackage.UpdateOriginEndpointOutput, error)
	UpdateOriginEndpointRequest(*mediapackage.UpdateOriginEndpointInput) (*request.Request, *mediapackage.UpdateOriginEndpointOutput)
}

var _ MediaPackageAPI = (*mediapackage.MediaPackage)(nil)
