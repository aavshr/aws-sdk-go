// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package restjsonserviceiface provides an interface to enable mocking the REST JSON Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package restjsonserviceiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/private/model/api/codegentest/service/restjsonservice"
)

// RESTJSONServiceAPI provides an interface to enable mocking the
// restjsonservice.RESTJSONService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // REST JSON Service.
//    func myFunc(svc restjsonserviceiface.RESTJSONServiceAPI) bool {
//        // Make svc.EmptyStream request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := restjsonservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRESTJSONServiceClient struct {
//        restjsonserviceiface.RESTJSONServiceAPI
//    }
//    func (m *mockRESTJSONServiceClient) EmptyStream(input *restjsonservice.EmptyStreamInput) (*restjsonservice.EmptyStreamOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRESTJSONServiceClient{}
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
type RESTJSONServiceAPI interface {
	EmptyStream(*restjsonservice.EmptyStreamInput) (*restjsonservice.EmptyStreamOutput, error)
	EmptyStreamWithContext(aws.Context, *restjsonservice.EmptyStreamInput, ...request.Option) (*restjsonservice.EmptyStreamOutput, error)
	EmptyStreamRequest(*restjsonservice.EmptyStreamInput) (*request.Request, *restjsonservice.EmptyStreamOutput)

	GetEventStream(*restjsonservice.GetEventStreamInput) (*restjsonservice.GetEventStreamOutput, error)
	GetEventStreamWithContext(aws.Context, *restjsonservice.GetEventStreamInput, ...request.Option) (*restjsonservice.GetEventStreamOutput, error)
	GetEventStreamRequest(*restjsonservice.GetEventStreamInput) (*request.Request, *restjsonservice.GetEventStreamOutput)

	OtherOperation(*restjsonservice.OtherOperationInput) (*restjsonservice.OtherOperationOutput, error)
	OtherOperationWithContext(aws.Context, *restjsonservice.OtherOperationInput, ...request.Option) (*restjsonservice.OtherOperationOutput, error)
	OtherOperationRequest(*restjsonservice.OtherOperationInput) (*request.Request, *restjsonservice.OtherOperationOutput)
}

var _ RESTJSONServiceAPI = (*restjsonservice.RESTJSONService)(nil)
