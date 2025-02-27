// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package detectiveiface provides an interface to enable mocking the Amazon Detective service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package detectiveiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/detective"
)

// DetectiveAPI provides an interface to enable mocking the
// detective.Detective service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Detective.
//    func myFunc(svc detectiveiface.DetectiveAPI) bool {
//        // Make svc.AcceptInvitation request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := detective.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDetectiveClient struct {
//        detectiveiface.DetectiveAPI
//    }
//    func (m *mockDetectiveClient) AcceptInvitation(input *detective.AcceptInvitationInput) (*detective.AcceptInvitationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDetectiveClient{}
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
type DetectiveAPI interface {
	AcceptInvitation(*detective.AcceptInvitationInput) (*detective.AcceptInvitationOutput, error)
	AcceptInvitationWithContext(aws.Context, *detective.AcceptInvitationInput, ...request.Option) (*detective.AcceptInvitationOutput, error)
	AcceptInvitationRequest(*detective.AcceptInvitationInput) (*request.Request, *detective.AcceptInvitationOutput)

	CreateGraph(*detective.CreateGraphInput) (*detective.CreateGraphOutput, error)
	CreateGraphWithContext(aws.Context, *detective.CreateGraphInput, ...request.Option) (*detective.CreateGraphOutput, error)
	CreateGraphRequest(*detective.CreateGraphInput) (*request.Request, *detective.CreateGraphOutput)

	CreateMembers(*detective.CreateMembersInput) (*detective.CreateMembersOutput, error)
	CreateMembersWithContext(aws.Context, *detective.CreateMembersInput, ...request.Option) (*detective.CreateMembersOutput, error)
	CreateMembersRequest(*detective.CreateMembersInput) (*request.Request, *detective.CreateMembersOutput)

	DeleteGraph(*detective.DeleteGraphInput) (*detective.DeleteGraphOutput, error)
	DeleteGraphWithContext(aws.Context, *detective.DeleteGraphInput, ...request.Option) (*detective.DeleteGraphOutput, error)
	DeleteGraphRequest(*detective.DeleteGraphInput) (*request.Request, *detective.DeleteGraphOutput)

	DeleteMembers(*detective.DeleteMembersInput) (*detective.DeleteMembersOutput, error)
	DeleteMembersWithContext(aws.Context, *detective.DeleteMembersInput, ...request.Option) (*detective.DeleteMembersOutput, error)
	DeleteMembersRequest(*detective.DeleteMembersInput) (*request.Request, *detective.DeleteMembersOutput)

	DisassociateMembership(*detective.DisassociateMembershipInput) (*detective.DisassociateMembershipOutput, error)
	DisassociateMembershipWithContext(aws.Context, *detective.DisassociateMembershipInput, ...request.Option) (*detective.DisassociateMembershipOutput, error)
	DisassociateMembershipRequest(*detective.DisassociateMembershipInput) (*request.Request, *detective.DisassociateMembershipOutput)

	GetMembers(*detective.GetMembersInput) (*detective.GetMembersOutput, error)
	GetMembersWithContext(aws.Context, *detective.GetMembersInput, ...request.Option) (*detective.GetMembersOutput, error)
	GetMembersRequest(*detective.GetMembersInput) (*request.Request, *detective.GetMembersOutput)

	ListGraphs(*detective.ListGraphsInput) (*detective.ListGraphsOutput, error)
	ListGraphsWithContext(aws.Context, *detective.ListGraphsInput, ...request.Option) (*detective.ListGraphsOutput, error)
	ListGraphsRequest(*detective.ListGraphsInput) (*request.Request, *detective.ListGraphsOutput)

	ListGraphsPages(*detective.ListGraphsInput, func(*detective.ListGraphsOutput, bool) bool) error
	ListGraphsPagesWithContext(aws.Context, *detective.ListGraphsInput, func(*detective.ListGraphsOutput, bool) bool, ...request.Option) error

	ListInvitations(*detective.ListInvitationsInput) (*detective.ListInvitationsOutput, error)
	ListInvitationsWithContext(aws.Context, *detective.ListInvitationsInput, ...request.Option) (*detective.ListInvitationsOutput, error)
	ListInvitationsRequest(*detective.ListInvitationsInput) (*request.Request, *detective.ListInvitationsOutput)

	ListInvitationsPages(*detective.ListInvitationsInput, func(*detective.ListInvitationsOutput, bool) bool) error
	ListInvitationsPagesWithContext(aws.Context, *detective.ListInvitationsInput, func(*detective.ListInvitationsOutput, bool) bool, ...request.Option) error

	ListMembers(*detective.ListMembersInput) (*detective.ListMembersOutput, error)
	ListMembersWithContext(aws.Context, *detective.ListMembersInput, ...request.Option) (*detective.ListMembersOutput, error)
	ListMembersRequest(*detective.ListMembersInput) (*request.Request, *detective.ListMembersOutput)

	ListMembersPages(*detective.ListMembersInput, func(*detective.ListMembersOutput, bool) bool) error
	ListMembersPagesWithContext(aws.Context, *detective.ListMembersInput, func(*detective.ListMembersOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*detective.ListTagsForResourceInput) (*detective.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *detective.ListTagsForResourceInput, ...request.Option) (*detective.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*detective.ListTagsForResourceInput) (*request.Request, *detective.ListTagsForResourceOutput)

	RejectInvitation(*detective.RejectInvitationInput) (*detective.RejectInvitationOutput, error)
	RejectInvitationWithContext(aws.Context, *detective.RejectInvitationInput, ...request.Option) (*detective.RejectInvitationOutput, error)
	RejectInvitationRequest(*detective.RejectInvitationInput) (*request.Request, *detective.RejectInvitationOutput)

	StartMonitoringMember(*detective.StartMonitoringMemberInput) (*detective.StartMonitoringMemberOutput, error)
	StartMonitoringMemberWithContext(aws.Context, *detective.StartMonitoringMemberInput, ...request.Option) (*detective.StartMonitoringMemberOutput, error)
	StartMonitoringMemberRequest(*detective.StartMonitoringMemberInput) (*request.Request, *detective.StartMonitoringMemberOutput)

	TagResource(*detective.TagResourceInput) (*detective.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *detective.TagResourceInput, ...request.Option) (*detective.TagResourceOutput, error)
	TagResourceRequest(*detective.TagResourceInput) (*request.Request, *detective.TagResourceOutput)

	UntagResource(*detective.UntagResourceInput) (*detective.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *detective.UntagResourceInput, ...request.Option) (*detective.UntagResourceOutput, error)
	UntagResourceRequest(*detective.UntagResourceInput) (*request.Request, *detective.UntagResourceOutput)
}

var _ DetectiveAPI = (*detective.Detective)(nil)
