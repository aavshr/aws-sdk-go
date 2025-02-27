// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package transferiface provides an interface to enable mocking the AWS Transfer Family service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package transferiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/transfer"
)

// TransferAPI provides an interface to enable mocking the
// transfer.Transfer service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Transfer Family.
//    func myFunc(svc transferiface.TransferAPI) bool {
//        // Make svc.CreateAccess request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := transfer.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockTransferClient struct {
//        transferiface.TransferAPI
//    }
//    func (m *mockTransferClient) CreateAccess(input *transfer.CreateAccessInput) (*transfer.CreateAccessOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockTransferClient{}
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
type TransferAPI interface {
	CreateAccess(*transfer.CreateAccessInput) (*transfer.CreateAccessOutput, error)
	CreateAccessWithContext(aws.Context, *transfer.CreateAccessInput, ...request.Option) (*transfer.CreateAccessOutput, error)
	CreateAccessRequest(*transfer.CreateAccessInput) (*request.Request, *transfer.CreateAccessOutput)

	CreateServer(*transfer.CreateServerInput) (*transfer.CreateServerOutput, error)
	CreateServerWithContext(aws.Context, *transfer.CreateServerInput, ...request.Option) (*transfer.CreateServerOutput, error)
	CreateServerRequest(*transfer.CreateServerInput) (*request.Request, *transfer.CreateServerOutput)

	CreateUser(*transfer.CreateUserInput) (*transfer.CreateUserOutput, error)
	CreateUserWithContext(aws.Context, *transfer.CreateUserInput, ...request.Option) (*transfer.CreateUserOutput, error)
	CreateUserRequest(*transfer.CreateUserInput) (*request.Request, *transfer.CreateUserOutput)

	CreateWorkflow(*transfer.CreateWorkflowInput) (*transfer.CreateWorkflowOutput, error)
	CreateWorkflowWithContext(aws.Context, *transfer.CreateWorkflowInput, ...request.Option) (*transfer.CreateWorkflowOutput, error)
	CreateWorkflowRequest(*transfer.CreateWorkflowInput) (*request.Request, *transfer.CreateWorkflowOutput)

	DeleteAccess(*transfer.DeleteAccessInput) (*transfer.DeleteAccessOutput, error)
	DeleteAccessWithContext(aws.Context, *transfer.DeleteAccessInput, ...request.Option) (*transfer.DeleteAccessOutput, error)
	DeleteAccessRequest(*transfer.DeleteAccessInput) (*request.Request, *transfer.DeleteAccessOutput)

	DeleteServer(*transfer.DeleteServerInput) (*transfer.DeleteServerOutput, error)
	DeleteServerWithContext(aws.Context, *transfer.DeleteServerInput, ...request.Option) (*transfer.DeleteServerOutput, error)
	DeleteServerRequest(*transfer.DeleteServerInput) (*request.Request, *transfer.DeleteServerOutput)

	DeleteSshPublicKey(*transfer.DeleteSshPublicKeyInput) (*transfer.DeleteSshPublicKeyOutput, error)
	DeleteSshPublicKeyWithContext(aws.Context, *transfer.DeleteSshPublicKeyInput, ...request.Option) (*transfer.DeleteSshPublicKeyOutput, error)
	DeleteSshPublicKeyRequest(*transfer.DeleteSshPublicKeyInput) (*request.Request, *transfer.DeleteSshPublicKeyOutput)

	DeleteUser(*transfer.DeleteUserInput) (*transfer.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *transfer.DeleteUserInput, ...request.Option) (*transfer.DeleteUserOutput, error)
	DeleteUserRequest(*transfer.DeleteUserInput) (*request.Request, *transfer.DeleteUserOutput)

	DeleteWorkflow(*transfer.DeleteWorkflowInput) (*transfer.DeleteWorkflowOutput, error)
	DeleteWorkflowWithContext(aws.Context, *transfer.DeleteWorkflowInput, ...request.Option) (*transfer.DeleteWorkflowOutput, error)
	DeleteWorkflowRequest(*transfer.DeleteWorkflowInput) (*request.Request, *transfer.DeleteWorkflowOutput)

	DescribeAccess(*transfer.DescribeAccessInput) (*transfer.DescribeAccessOutput, error)
	DescribeAccessWithContext(aws.Context, *transfer.DescribeAccessInput, ...request.Option) (*transfer.DescribeAccessOutput, error)
	DescribeAccessRequest(*transfer.DescribeAccessInput) (*request.Request, *transfer.DescribeAccessOutput)

	DescribeExecution(*transfer.DescribeExecutionInput) (*transfer.DescribeExecutionOutput, error)
	DescribeExecutionWithContext(aws.Context, *transfer.DescribeExecutionInput, ...request.Option) (*transfer.DescribeExecutionOutput, error)
	DescribeExecutionRequest(*transfer.DescribeExecutionInput) (*request.Request, *transfer.DescribeExecutionOutput)

	DescribeSecurityPolicy(*transfer.DescribeSecurityPolicyInput) (*transfer.DescribeSecurityPolicyOutput, error)
	DescribeSecurityPolicyWithContext(aws.Context, *transfer.DescribeSecurityPolicyInput, ...request.Option) (*transfer.DescribeSecurityPolicyOutput, error)
	DescribeSecurityPolicyRequest(*transfer.DescribeSecurityPolicyInput) (*request.Request, *transfer.DescribeSecurityPolicyOutput)

	DescribeServer(*transfer.DescribeServerInput) (*transfer.DescribeServerOutput, error)
	DescribeServerWithContext(aws.Context, *transfer.DescribeServerInput, ...request.Option) (*transfer.DescribeServerOutput, error)
	DescribeServerRequest(*transfer.DescribeServerInput) (*request.Request, *transfer.DescribeServerOutput)

	DescribeUser(*transfer.DescribeUserInput) (*transfer.DescribeUserOutput, error)
	DescribeUserWithContext(aws.Context, *transfer.DescribeUserInput, ...request.Option) (*transfer.DescribeUserOutput, error)
	DescribeUserRequest(*transfer.DescribeUserInput) (*request.Request, *transfer.DescribeUserOutput)

	DescribeWorkflow(*transfer.DescribeWorkflowInput) (*transfer.DescribeWorkflowOutput, error)
	DescribeWorkflowWithContext(aws.Context, *transfer.DescribeWorkflowInput, ...request.Option) (*transfer.DescribeWorkflowOutput, error)
	DescribeWorkflowRequest(*transfer.DescribeWorkflowInput) (*request.Request, *transfer.DescribeWorkflowOutput)

	ImportSshPublicKey(*transfer.ImportSshPublicKeyInput) (*transfer.ImportSshPublicKeyOutput, error)
	ImportSshPublicKeyWithContext(aws.Context, *transfer.ImportSshPublicKeyInput, ...request.Option) (*transfer.ImportSshPublicKeyOutput, error)
	ImportSshPublicKeyRequest(*transfer.ImportSshPublicKeyInput) (*request.Request, *transfer.ImportSshPublicKeyOutput)

	ListAccesses(*transfer.ListAccessesInput) (*transfer.ListAccessesOutput, error)
	ListAccessesWithContext(aws.Context, *transfer.ListAccessesInput, ...request.Option) (*transfer.ListAccessesOutput, error)
	ListAccessesRequest(*transfer.ListAccessesInput) (*request.Request, *transfer.ListAccessesOutput)

	ListAccessesPages(*transfer.ListAccessesInput, func(*transfer.ListAccessesOutput, bool) bool) error
	ListAccessesPagesWithContext(aws.Context, *transfer.ListAccessesInput, func(*transfer.ListAccessesOutput, bool) bool, ...request.Option) error

	ListExecutions(*transfer.ListExecutionsInput) (*transfer.ListExecutionsOutput, error)
	ListExecutionsWithContext(aws.Context, *transfer.ListExecutionsInput, ...request.Option) (*transfer.ListExecutionsOutput, error)
	ListExecutionsRequest(*transfer.ListExecutionsInput) (*request.Request, *transfer.ListExecutionsOutput)

	ListExecutionsPages(*transfer.ListExecutionsInput, func(*transfer.ListExecutionsOutput, bool) bool) error
	ListExecutionsPagesWithContext(aws.Context, *transfer.ListExecutionsInput, func(*transfer.ListExecutionsOutput, bool) bool, ...request.Option) error

	ListSecurityPolicies(*transfer.ListSecurityPoliciesInput) (*transfer.ListSecurityPoliciesOutput, error)
	ListSecurityPoliciesWithContext(aws.Context, *transfer.ListSecurityPoliciesInput, ...request.Option) (*transfer.ListSecurityPoliciesOutput, error)
	ListSecurityPoliciesRequest(*transfer.ListSecurityPoliciesInput) (*request.Request, *transfer.ListSecurityPoliciesOutput)

	ListSecurityPoliciesPages(*transfer.ListSecurityPoliciesInput, func(*transfer.ListSecurityPoliciesOutput, bool) bool) error
	ListSecurityPoliciesPagesWithContext(aws.Context, *transfer.ListSecurityPoliciesInput, func(*transfer.ListSecurityPoliciesOutput, bool) bool, ...request.Option) error

	ListServers(*transfer.ListServersInput) (*transfer.ListServersOutput, error)
	ListServersWithContext(aws.Context, *transfer.ListServersInput, ...request.Option) (*transfer.ListServersOutput, error)
	ListServersRequest(*transfer.ListServersInput) (*request.Request, *transfer.ListServersOutput)

	ListServersPages(*transfer.ListServersInput, func(*transfer.ListServersOutput, bool) bool) error
	ListServersPagesWithContext(aws.Context, *transfer.ListServersInput, func(*transfer.ListServersOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*transfer.ListTagsForResourceInput) (*transfer.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *transfer.ListTagsForResourceInput, ...request.Option) (*transfer.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*transfer.ListTagsForResourceInput) (*request.Request, *transfer.ListTagsForResourceOutput)

	ListTagsForResourcePages(*transfer.ListTagsForResourceInput, func(*transfer.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *transfer.ListTagsForResourceInput, func(*transfer.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	ListUsers(*transfer.ListUsersInput) (*transfer.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *transfer.ListUsersInput, ...request.Option) (*transfer.ListUsersOutput, error)
	ListUsersRequest(*transfer.ListUsersInput) (*request.Request, *transfer.ListUsersOutput)

	ListUsersPages(*transfer.ListUsersInput, func(*transfer.ListUsersOutput, bool) bool) error
	ListUsersPagesWithContext(aws.Context, *transfer.ListUsersInput, func(*transfer.ListUsersOutput, bool) bool, ...request.Option) error

	ListWorkflows(*transfer.ListWorkflowsInput) (*transfer.ListWorkflowsOutput, error)
	ListWorkflowsWithContext(aws.Context, *transfer.ListWorkflowsInput, ...request.Option) (*transfer.ListWorkflowsOutput, error)
	ListWorkflowsRequest(*transfer.ListWorkflowsInput) (*request.Request, *transfer.ListWorkflowsOutput)

	ListWorkflowsPages(*transfer.ListWorkflowsInput, func(*transfer.ListWorkflowsOutput, bool) bool) error
	ListWorkflowsPagesWithContext(aws.Context, *transfer.ListWorkflowsInput, func(*transfer.ListWorkflowsOutput, bool) bool, ...request.Option) error

	SendWorkflowStepState(*transfer.SendWorkflowStepStateInput) (*transfer.SendWorkflowStepStateOutput, error)
	SendWorkflowStepStateWithContext(aws.Context, *transfer.SendWorkflowStepStateInput, ...request.Option) (*transfer.SendWorkflowStepStateOutput, error)
	SendWorkflowStepStateRequest(*transfer.SendWorkflowStepStateInput) (*request.Request, *transfer.SendWorkflowStepStateOutput)

	StartServer(*transfer.StartServerInput) (*transfer.StartServerOutput, error)
	StartServerWithContext(aws.Context, *transfer.StartServerInput, ...request.Option) (*transfer.StartServerOutput, error)
	StartServerRequest(*transfer.StartServerInput) (*request.Request, *transfer.StartServerOutput)

	StopServer(*transfer.StopServerInput) (*transfer.StopServerOutput, error)
	StopServerWithContext(aws.Context, *transfer.StopServerInput, ...request.Option) (*transfer.StopServerOutput, error)
	StopServerRequest(*transfer.StopServerInput) (*request.Request, *transfer.StopServerOutput)

	TagResource(*transfer.TagResourceInput) (*transfer.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *transfer.TagResourceInput, ...request.Option) (*transfer.TagResourceOutput, error)
	TagResourceRequest(*transfer.TagResourceInput) (*request.Request, *transfer.TagResourceOutput)

	TestIdentityProvider(*transfer.TestIdentityProviderInput) (*transfer.TestIdentityProviderOutput, error)
	TestIdentityProviderWithContext(aws.Context, *transfer.TestIdentityProviderInput, ...request.Option) (*transfer.TestIdentityProviderOutput, error)
	TestIdentityProviderRequest(*transfer.TestIdentityProviderInput) (*request.Request, *transfer.TestIdentityProviderOutput)

	UntagResource(*transfer.UntagResourceInput) (*transfer.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *transfer.UntagResourceInput, ...request.Option) (*transfer.UntagResourceOutput, error)
	UntagResourceRequest(*transfer.UntagResourceInput) (*request.Request, *transfer.UntagResourceOutput)

	UpdateAccess(*transfer.UpdateAccessInput) (*transfer.UpdateAccessOutput, error)
	UpdateAccessWithContext(aws.Context, *transfer.UpdateAccessInput, ...request.Option) (*transfer.UpdateAccessOutput, error)
	UpdateAccessRequest(*transfer.UpdateAccessInput) (*request.Request, *transfer.UpdateAccessOutput)

	UpdateServer(*transfer.UpdateServerInput) (*transfer.UpdateServerOutput, error)
	UpdateServerWithContext(aws.Context, *transfer.UpdateServerInput, ...request.Option) (*transfer.UpdateServerOutput, error)
	UpdateServerRequest(*transfer.UpdateServerInput) (*request.Request, *transfer.UpdateServerOutput)

	UpdateUser(*transfer.UpdateUserInput) (*transfer.UpdateUserOutput, error)
	UpdateUserWithContext(aws.Context, *transfer.UpdateUserInput, ...request.Option) (*transfer.UpdateUserOutput, error)
	UpdateUserRequest(*transfer.UpdateUserInput) (*request.Request, *transfer.UpdateUserOutput)
}

var _ TransferAPI = (*transfer.Transfer)(nil)
