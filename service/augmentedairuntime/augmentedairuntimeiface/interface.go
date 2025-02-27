// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package augmentedairuntimeiface provides an interface to enable mocking the Amazon Augmented AI Runtime service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package augmentedairuntimeiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/augmentedairuntime"
)

// AugmentedAIRuntimeAPI provides an interface to enable mocking the
// augmentedairuntime.AugmentedAIRuntime service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Augmented AI Runtime.
//    func myFunc(svc augmentedairuntimeiface.AugmentedAIRuntimeAPI) bool {
//        // Make svc.DeleteHumanLoop request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := augmentedairuntime.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAugmentedAIRuntimeClient struct {
//        augmentedairuntimeiface.AugmentedAIRuntimeAPI
//    }
//    func (m *mockAugmentedAIRuntimeClient) DeleteHumanLoop(input *augmentedairuntime.DeleteHumanLoopInput) (*augmentedairuntime.DeleteHumanLoopOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAugmentedAIRuntimeClient{}
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
type AugmentedAIRuntimeAPI interface {
	DeleteHumanLoop(*augmentedairuntime.DeleteHumanLoopInput) (*augmentedairuntime.DeleteHumanLoopOutput, error)
	DeleteHumanLoopWithContext(aws.Context, *augmentedairuntime.DeleteHumanLoopInput, ...request.Option) (*augmentedairuntime.DeleteHumanLoopOutput, error)
	DeleteHumanLoopRequest(*augmentedairuntime.DeleteHumanLoopInput) (*request.Request, *augmentedairuntime.DeleteHumanLoopOutput)

	DescribeHumanLoop(*augmentedairuntime.DescribeHumanLoopInput) (*augmentedairuntime.DescribeHumanLoopOutput, error)
	DescribeHumanLoopWithContext(aws.Context, *augmentedairuntime.DescribeHumanLoopInput, ...request.Option) (*augmentedairuntime.DescribeHumanLoopOutput, error)
	DescribeHumanLoopRequest(*augmentedairuntime.DescribeHumanLoopInput) (*request.Request, *augmentedairuntime.DescribeHumanLoopOutput)

	ListHumanLoops(*augmentedairuntime.ListHumanLoopsInput) (*augmentedairuntime.ListHumanLoopsOutput, error)
	ListHumanLoopsWithContext(aws.Context, *augmentedairuntime.ListHumanLoopsInput, ...request.Option) (*augmentedairuntime.ListHumanLoopsOutput, error)
	ListHumanLoopsRequest(*augmentedairuntime.ListHumanLoopsInput) (*request.Request, *augmentedairuntime.ListHumanLoopsOutput)

	ListHumanLoopsPages(*augmentedairuntime.ListHumanLoopsInput, func(*augmentedairuntime.ListHumanLoopsOutput, bool) bool) error
	ListHumanLoopsPagesWithContext(aws.Context, *augmentedairuntime.ListHumanLoopsInput, func(*augmentedairuntime.ListHumanLoopsOutput, bool) bool, ...request.Option) error

	StartHumanLoop(*augmentedairuntime.StartHumanLoopInput) (*augmentedairuntime.StartHumanLoopOutput, error)
	StartHumanLoopWithContext(aws.Context, *augmentedairuntime.StartHumanLoopInput, ...request.Option) (*augmentedairuntime.StartHumanLoopOutput, error)
	StartHumanLoopRequest(*augmentedairuntime.StartHumanLoopInput) (*request.Request, *augmentedairuntime.StartHumanLoopOutput)

	StopHumanLoop(*augmentedairuntime.StopHumanLoopInput) (*augmentedairuntime.StopHumanLoopOutput, error)
	StopHumanLoopWithContext(aws.Context, *augmentedairuntime.StopHumanLoopInput, ...request.Option) (*augmentedairuntime.StopHumanLoopOutput, error)
	StopHumanLoopRequest(*augmentedairuntime.StopHumanLoopInput) (*request.Request, *augmentedairuntime.StopHumanLoopOutput)
}

var _ AugmentedAIRuntimeAPI = (*augmentedairuntime.AugmentedAIRuntime)(nil)
