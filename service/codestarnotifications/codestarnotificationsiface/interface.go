// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codestarnotificationsiface provides an interface to enable mocking the AWS CodeStar Notifications service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codestarnotificationsiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/codestarnotifications"
)

// CodeStarNotificationsAPI provides an interface to enable mocking the
// codestarnotifications.CodeStarNotifications service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CodeStar Notifications.
//    func myFunc(svc codestarnotificationsiface.CodeStarNotificationsAPI) bool {
//        // Make svc.CreateNotificationRule request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := codestarnotifications.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCodeStarNotificationsClient struct {
//        codestarnotificationsiface.CodeStarNotificationsAPI
//    }
//    func (m *mockCodeStarNotificationsClient) CreateNotificationRule(input *codestarnotifications.CreateNotificationRuleInput) (*codestarnotifications.CreateNotificationRuleOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCodeStarNotificationsClient{}
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
type CodeStarNotificationsAPI interface {
	CreateNotificationRule(*codestarnotifications.CreateNotificationRuleInput) (*codestarnotifications.CreateNotificationRuleOutput, error)
	CreateNotificationRuleWithContext(aws.Context, *codestarnotifications.CreateNotificationRuleInput, ...request.Option) (*codestarnotifications.CreateNotificationRuleOutput, error)
	CreateNotificationRuleRequest(*codestarnotifications.CreateNotificationRuleInput) (*request.Request, *codestarnotifications.CreateNotificationRuleOutput)

	DeleteNotificationRule(*codestarnotifications.DeleteNotificationRuleInput) (*codestarnotifications.DeleteNotificationRuleOutput, error)
	DeleteNotificationRuleWithContext(aws.Context, *codestarnotifications.DeleteNotificationRuleInput, ...request.Option) (*codestarnotifications.DeleteNotificationRuleOutput, error)
	DeleteNotificationRuleRequest(*codestarnotifications.DeleteNotificationRuleInput) (*request.Request, *codestarnotifications.DeleteNotificationRuleOutput)

	DeleteTarget(*codestarnotifications.DeleteTargetInput) (*codestarnotifications.DeleteTargetOutput, error)
	DeleteTargetWithContext(aws.Context, *codestarnotifications.DeleteTargetInput, ...request.Option) (*codestarnotifications.DeleteTargetOutput, error)
	DeleteTargetRequest(*codestarnotifications.DeleteTargetInput) (*request.Request, *codestarnotifications.DeleteTargetOutput)

	DescribeNotificationRule(*codestarnotifications.DescribeNotificationRuleInput) (*codestarnotifications.DescribeNotificationRuleOutput, error)
	DescribeNotificationRuleWithContext(aws.Context, *codestarnotifications.DescribeNotificationRuleInput, ...request.Option) (*codestarnotifications.DescribeNotificationRuleOutput, error)
	DescribeNotificationRuleRequest(*codestarnotifications.DescribeNotificationRuleInput) (*request.Request, *codestarnotifications.DescribeNotificationRuleOutput)

	ListEventTypes(*codestarnotifications.ListEventTypesInput) (*codestarnotifications.ListEventTypesOutput, error)
	ListEventTypesWithContext(aws.Context, *codestarnotifications.ListEventTypesInput, ...request.Option) (*codestarnotifications.ListEventTypesOutput, error)
	ListEventTypesRequest(*codestarnotifications.ListEventTypesInput) (*request.Request, *codestarnotifications.ListEventTypesOutput)

	ListEventTypesPages(*codestarnotifications.ListEventTypesInput, func(*codestarnotifications.ListEventTypesOutput, bool) bool) error
	ListEventTypesPagesWithContext(aws.Context, *codestarnotifications.ListEventTypesInput, func(*codestarnotifications.ListEventTypesOutput, bool) bool, ...request.Option) error

	ListNotificationRules(*codestarnotifications.ListNotificationRulesInput) (*codestarnotifications.ListNotificationRulesOutput, error)
	ListNotificationRulesWithContext(aws.Context, *codestarnotifications.ListNotificationRulesInput, ...request.Option) (*codestarnotifications.ListNotificationRulesOutput, error)
	ListNotificationRulesRequest(*codestarnotifications.ListNotificationRulesInput) (*request.Request, *codestarnotifications.ListNotificationRulesOutput)

	ListNotificationRulesPages(*codestarnotifications.ListNotificationRulesInput, func(*codestarnotifications.ListNotificationRulesOutput, bool) bool) error
	ListNotificationRulesPagesWithContext(aws.Context, *codestarnotifications.ListNotificationRulesInput, func(*codestarnotifications.ListNotificationRulesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*codestarnotifications.ListTagsForResourceInput) (*codestarnotifications.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *codestarnotifications.ListTagsForResourceInput, ...request.Option) (*codestarnotifications.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*codestarnotifications.ListTagsForResourceInput) (*request.Request, *codestarnotifications.ListTagsForResourceOutput)

	ListTargets(*codestarnotifications.ListTargetsInput) (*codestarnotifications.ListTargetsOutput, error)
	ListTargetsWithContext(aws.Context, *codestarnotifications.ListTargetsInput, ...request.Option) (*codestarnotifications.ListTargetsOutput, error)
	ListTargetsRequest(*codestarnotifications.ListTargetsInput) (*request.Request, *codestarnotifications.ListTargetsOutput)

	ListTargetsPages(*codestarnotifications.ListTargetsInput, func(*codestarnotifications.ListTargetsOutput, bool) bool) error
	ListTargetsPagesWithContext(aws.Context, *codestarnotifications.ListTargetsInput, func(*codestarnotifications.ListTargetsOutput, bool) bool, ...request.Option) error

	Subscribe(*codestarnotifications.SubscribeInput) (*codestarnotifications.SubscribeOutput, error)
	SubscribeWithContext(aws.Context, *codestarnotifications.SubscribeInput, ...request.Option) (*codestarnotifications.SubscribeOutput, error)
	SubscribeRequest(*codestarnotifications.SubscribeInput) (*request.Request, *codestarnotifications.SubscribeOutput)

	TagResource(*codestarnotifications.TagResourceInput) (*codestarnotifications.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *codestarnotifications.TagResourceInput, ...request.Option) (*codestarnotifications.TagResourceOutput, error)
	TagResourceRequest(*codestarnotifications.TagResourceInput) (*request.Request, *codestarnotifications.TagResourceOutput)

	Unsubscribe(*codestarnotifications.UnsubscribeInput) (*codestarnotifications.UnsubscribeOutput, error)
	UnsubscribeWithContext(aws.Context, *codestarnotifications.UnsubscribeInput, ...request.Option) (*codestarnotifications.UnsubscribeOutput, error)
	UnsubscribeRequest(*codestarnotifications.UnsubscribeInput) (*request.Request, *codestarnotifications.UnsubscribeOutput)

	UntagResource(*codestarnotifications.UntagResourceInput) (*codestarnotifications.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *codestarnotifications.UntagResourceInput, ...request.Option) (*codestarnotifications.UntagResourceOutput, error)
	UntagResourceRequest(*codestarnotifications.UntagResourceInput) (*request.Request, *codestarnotifications.UntagResourceOutput)

	UpdateNotificationRule(*codestarnotifications.UpdateNotificationRuleInput) (*codestarnotifications.UpdateNotificationRuleOutput, error)
	UpdateNotificationRuleWithContext(aws.Context, *codestarnotifications.UpdateNotificationRuleInput, ...request.Option) (*codestarnotifications.UpdateNotificationRuleOutput, error)
	UpdateNotificationRuleRequest(*codestarnotifications.UpdateNotificationRuleInput) (*request.Request, *codestarnotifications.UpdateNotificationRuleOutput)
}

var _ CodeStarNotificationsAPI = (*codestarnotifications.CodeStarNotifications)(nil)
