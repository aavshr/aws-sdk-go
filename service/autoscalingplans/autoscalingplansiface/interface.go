// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package autoscalingplansiface provides an interface to enable mocking the AWS Auto Scaling Plans service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package autoscalingplansiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/autoscalingplans"
)

// AutoScalingPlansAPI provides an interface to enable mocking the
// autoscalingplans.AutoScalingPlans service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Auto Scaling Plans.
//    func myFunc(svc autoscalingplansiface.AutoScalingPlansAPI) bool {
//        // Make svc.CreateScalingPlan request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := autoscalingplans.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAutoScalingPlansClient struct {
//        autoscalingplansiface.AutoScalingPlansAPI
//    }
//    func (m *mockAutoScalingPlansClient) CreateScalingPlan(input *autoscalingplans.CreateScalingPlanInput) (*autoscalingplans.CreateScalingPlanOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAutoScalingPlansClient{}
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
type AutoScalingPlansAPI interface {
	CreateScalingPlan(*autoscalingplans.CreateScalingPlanInput) (*autoscalingplans.CreateScalingPlanOutput, error)
	CreateScalingPlanWithContext(aws.Context, *autoscalingplans.CreateScalingPlanInput, ...request.Option) (*autoscalingplans.CreateScalingPlanOutput, error)
	CreateScalingPlanRequest(*autoscalingplans.CreateScalingPlanInput) (*request.Request, *autoscalingplans.CreateScalingPlanOutput)

	DeleteScalingPlan(*autoscalingplans.DeleteScalingPlanInput) (*autoscalingplans.DeleteScalingPlanOutput, error)
	DeleteScalingPlanWithContext(aws.Context, *autoscalingplans.DeleteScalingPlanInput, ...request.Option) (*autoscalingplans.DeleteScalingPlanOutput, error)
	DeleteScalingPlanRequest(*autoscalingplans.DeleteScalingPlanInput) (*request.Request, *autoscalingplans.DeleteScalingPlanOutput)

	DescribeScalingPlanResources(*autoscalingplans.DescribeScalingPlanResourcesInput) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error)
	DescribeScalingPlanResourcesWithContext(aws.Context, *autoscalingplans.DescribeScalingPlanResourcesInput, ...request.Option) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error)
	DescribeScalingPlanResourcesRequest(*autoscalingplans.DescribeScalingPlanResourcesInput) (*request.Request, *autoscalingplans.DescribeScalingPlanResourcesOutput)

	DescribeScalingPlans(*autoscalingplans.DescribeScalingPlansInput) (*autoscalingplans.DescribeScalingPlansOutput, error)
	DescribeScalingPlansWithContext(aws.Context, *autoscalingplans.DescribeScalingPlansInput, ...request.Option) (*autoscalingplans.DescribeScalingPlansOutput, error)
	DescribeScalingPlansRequest(*autoscalingplans.DescribeScalingPlansInput) (*request.Request, *autoscalingplans.DescribeScalingPlansOutput)

	GetScalingPlanResourceForecastData(*autoscalingplans.GetScalingPlanResourceForecastDataInput) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error)
	GetScalingPlanResourceForecastDataWithContext(aws.Context, *autoscalingplans.GetScalingPlanResourceForecastDataInput, ...request.Option) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error)
	GetScalingPlanResourceForecastDataRequest(*autoscalingplans.GetScalingPlanResourceForecastDataInput) (*request.Request, *autoscalingplans.GetScalingPlanResourceForecastDataOutput)

	UpdateScalingPlan(*autoscalingplans.UpdateScalingPlanInput) (*autoscalingplans.UpdateScalingPlanOutput, error)
	UpdateScalingPlanWithContext(aws.Context, *autoscalingplans.UpdateScalingPlanInput, ...request.Option) (*autoscalingplans.UpdateScalingPlanOutput, error)
	UpdateScalingPlanRequest(*autoscalingplans.UpdateScalingPlanInput) (*request.Request, *autoscalingplans.UpdateScalingPlanOutput)
}

var _ AutoScalingPlansAPI = (*autoscalingplans.AutoScalingPlans)(nil)
