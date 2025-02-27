// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package elbiface provides an interface to enable mocking the Elastic Load Balancing service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package elbiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/elb"
)

// ELBAPI provides an interface to enable mocking the
// elb.ELB service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Elastic Load Balancing.
//    func myFunc(svc elbiface.ELBAPI) bool {
//        // Make svc.AddTags request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := elb.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockELBClient struct {
//        elbiface.ELBAPI
//    }
//    func (m *mockELBClient) AddTags(input *elb.AddTagsInput) (*elb.AddTagsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockELBClient{}
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
type ELBAPI interface {
	AddTags(*elb.AddTagsInput) (*elb.AddTagsOutput, error)
	AddTagsWithContext(aws.Context, *elb.AddTagsInput, ...request.Option) (*elb.AddTagsOutput, error)
	AddTagsRequest(*elb.AddTagsInput) (*request.Request, *elb.AddTagsOutput)

	ApplySecurityGroupsToLoadBalancer(*elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error)
	ApplySecurityGroupsToLoadBalancerWithContext(aws.Context, *elb.ApplySecurityGroupsToLoadBalancerInput, ...request.Option) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error)
	ApplySecurityGroupsToLoadBalancerRequest(*elb.ApplySecurityGroupsToLoadBalancerInput) (*request.Request, *elb.ApplySecurityGroupsToLoadBalancerOutput)

	AttachLoadBalancerToSubnets(*elb.AttachLoadBalancerToSubnetsInput) (*elb.AttachLoadBalancerToSubnetsOutput, error)
	AttachLoadBalancerToSubnetsWithContext(aws.Context, *elb.AttachLoadBalancerToSubnetsInput, ...request.Option) (*elb.AttachLoadBalancerToSubnetsOutput, error)
	AttachLoadBalancerToSubnetsRequest(*elb.AttachLoadBalancerToSubnetsInput) (*request.Request, *elb.AttachLoadBalancerToSubnetsOutput)

	ConfigureHealthCheck(*elb.ConfigureHealthCheckInput) (*elb.ConfigureHealthCheckOutput, error)
	ConfigureHealthCheckWithContext(aws.Context, *elb.ConfigureHealthCheckInput, ...request.Option) (*elb.ConfigureHealthCheckOutput, error)
	ConfigureHealthCheckRequest(*elb.ConfigureHealthCheckInput) (*request.Request, *elb.ConfigureHealthCheckOutput)

	CreateAppCookieStickinessPolicy(*elb.CreateAppCookieStickinessPolicyInput) (*elb.CreateAppCookieStickinessPolicyOutput, error)
	CreateAppCookieStickinessPolicyWithContext(aws.Context, *elb.CreateAppCookieStickinessPolicyInput, ...request.Option) (*elb.CreateAppCookieStickinessPolicyOutput, error)
	CreateAppCookieStickinessPolicyRequest(*elb.CreateAppCookieStickinessPolicyInput) (*request.Request, *elb.CreateAppCookieStickinessPolicyOutput)

	CreateLBCookieStickinessPolicy(*elb.CreateLBCookieStickinessPolicyInput) (*elb.CreateLBCookieStickinessPolicyOutput, error)
	CreateLBCookieStickinessPolicyWithContext(aws.Context, *elb.CreateLBCookieStickinessPolicyInput, ...request.Option) (*elb.CreateLBCookieStickinessPolicyOutput, error)
	CreateLBCookieStickinessPolicyRequest(*elb.CreateLBCookieStickinessPolicyInput) (*request.Request, *elb.CreateLBCookieStickinessPolicyOutput)

	CreateLoadBalancer(*elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error)
	CreateLoadBalancerWithContext(aws.Context, *elb.CreateLoadBalancerInput, ...request.Option) (*elb.CreateLoadBalancerOutput, error)
	CreateLoadBalancerRequest(*elb.CreateLoadBalancerInput) (*request.Request, *elb.CreateLoadBalancerOutput)

	CreateLoadBalancerListeners(*elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error)
	CreateLoadBalancerListenersWithContext(aws.Context, *elb.CreateLoadBalancerListenersInput, ...request.Option) (*elb.CreateLoadBalancerListenersOutput, error)
	CreateLoadBalancerListenersRequest(*elb.CreateLoadBalancerListenersInput) (*request.Request, *elb.CreateLoadBalancerListenersOutput)

	CreateLoadBalancerPolicy(*elb.CreateLoadBalancerPolicyInput) (*elb.CreateLoadBalancerPolicyOutput, error)
	CreateLoadBalancerPolicyWithContext(aws.Context, *elb.CreateLoadBalancerPolicyInput, ...request.Option) (*elb.CreateLoadBalancerPolicyOutput, error)
	CreateLoadBalancerPolicyRequest(*elb.CreateLoadBalancerPolicyInput) (*request.Request, *elb.CreateLoadBalancerPolicyOutput)

	DeleteLoadBalancer(*elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerWithContext(aws.Context, *elb.DeleteLoadBalancerInput, ...request.Option) (*elb.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerRequest(*elb.DeleteLoadBalancerInput) (*request.Request, *elb.DeleteLoadBalancerOutput)

	DeleteLoadBalancerListeners(*elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error)
	DeleteLoadBalancerListenersWithContext(aws.Context, *elb.DeleteLoadBalancerListenersInput, ...request.Option) (*elb.DeleteLoadBalancerListenersOutput, error)
	DeleteLoadBalancerListenersRequest(*elb.DeleteLoadBalancerListenersInput) (*request.Request, *elb.DeleteLoadBalancerListenersOutput)

	DeleteLoadBalancerPolicy(*elb.DeleteLoadBalancerPolicyInput) (*elb.DeleteLoadBalancerPolicyOutput, error)
	DeleteLoadBalancerPolicyWithContext(aws.Context, *elb.DeleteLoadBalancerPolicyInput, ...request.Option) (*elb.DeleteLoadBalancerPolicyOutput, error)
	DeleteLoadBalancerPolicyRequest(*elb.DeleteLoadBalancerPolicyInput) (*request.Request, *elb.DeleteLoadBalancerPolicyOutput)

	DeregisterInstancesFromLoadBalancer(*elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error)
	DeregisterInstancesFromLoadBalancerWithContext(aws.Context, *elb.DeregisterInstancesFromLoadBalancerInput, ...request.Option) (*elb.DeregisterInstancesFromLoadBalancerOutput, error)
	DeregisterInstancesFromLoadBalancerRequest(*elb.DeregisterInstancesFromLoadBalancerInput) (*request.Request, *elb.DeregisterInstancesFromLoadBalancerOutput)

	DescribeAccountLimits(*elb.DescribeAccountLimitsInput) (*elb.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsWithContext(aws.Context, *elb.DescribeAccountLimitsInput, ...request.Option) (*elb.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsRequest(*elb.DescribeAccountLimitsInput) (*request.Request, *elb.DescribeAccountLimitsOutput)

	DescribeInstanceHealth(*elb.DescribeInstanceHealthInput) (*elb.DescribeInstanceHealthOutput, error)
	DescribeInstanceHealthWithContext(aws.Context, *elb.DescribeInstanceHealthInput, ...request.Option) (*elb.DescribeInstanceHealthOutput, error)
	DescribeInstanceHealthRequest(*elb.DescribeInstanceHealthInput) (*request.Request, *elb.DescribeInstanceHealthOutput)

	DescribeLoadBalancerAttributes(*elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributesWithContext(aws.Context, *elb.DescribeLoadBalancerAttributesInput, ...request.Option) (*elb.DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributesRequest(*elb.DescribeLoadBalancerAttributesInput) (*request.Request, *elb.DescribeLoadBalancerAttributesOutput)

	DescribeLoadBalancerPolicies(*elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error)
	DescribeLoadBalancerPoliciesWithContext(aws.Context, *elb.DescribeLoadBalancerPoliciesInput, ...request.Option) (*elb.DescribeLoadBalancerPoliciesOutput, error)
	DescribeLoadBalancerPoliciesRequest(*elb.DescribeLoadBalancerPoliciesInput) (*request.Request, *elb.DescribeLoadBalancerPoliciesOutput)

	DescribeLoadBalancerPolicyTypes(*elb.DescribeLoadBalancerPolicyTypesInput) (*elb.DescribeLoadBalancerPolicyTypesOutput, error)
	DescribeLoadBalancerPolicyTypesWithContext(aws.Context, *elb.DescribeLoadBalancerPolicyTypesInput, ...request.Option) (*elb.DescribeLoadBalancerPolicyTypesOutput, error)
	DescribeLoadBalancerPolicyTypesRequest(*elb.DescribeLoadBalancerPolicyTypesInput) (*request.Request, *elb.DescribeLoadBalancerPolicyTypesOutput)

	DescribeLoadBalancers(*elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersWithContext(aws.Context, *elb.DescribeLoadBalancersInput, ...request.Option) (*elb.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersRequest(*elb.DescribeLoadBalancersInput) (*request.Request, *elb.DescribeLoadBalancersOutput)

	DescribeLoadBalancersPages(*elb.DescribeLoadBalancersInput, func(*elb.DescribeLoadBalancersOutput, bool) bool) error
	DescribeLoadBalancersPagesWithContext(aws.Context, *elb.DescribeLoadBalancersInput, func(*elb.DescribeLoadBalancersOutput, bool) bool, ...request.Option) error

	DescribeTags(*elb.DescribeTagsInput) (*elb.DescribeTagsOutput, error)
	DescribeTagsWithContext(aws.Context, *elb.DescribeTagsInput, ...request.Option) (*elb.DescribeTagsOutput, error)
	DescribeTagsRequest(*elb.DescribeTagsInput) (*request.Request, *elb.DescribeTagsOutput)

	DetachLoadBalancerFromSubnets(*elb.DetachLoadBalancerFromSubnetsInput) (*elb.DetachLoadBalancerFromSubnetsOutput, error)
	DetachLoadBalancerFromSubnetsWithContext(aws.Context, *elb.DetachLoadBalancerFromSubnetsInput, ...request.Option) (*elb.DetachLoadBalancerFromSubnetsOutput, error)
	DetachLoadBalancerFromSubnetsRequest(*elb.DetachLoadBalancerFromSubnetsInput) (*request.Request, *elb.DetachLoadBalancerFromSubnetsOutput)

	DisableAvailabilityZonesForLoadBalancer(*elb.DisableAvailabilityZonesForLoadBalancerInput) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error)
	DisableAvailabilityZonesForLoadBalancerWithContext(aws.Context, *elb.DisableAvailabilityZonesForLoadBalancerInput, ...request.Option) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error)
	DisableAvailabilityZonesForLoadBalancerRequest(*elb.DisableAvailabilityZonesForLoadBalancerInput) (*request.Request, *elb.DisableAvailabilityZonesForLoadBalancerOutput)

	EnableAvailabilityZonesForLoadBalancer(*elb.EnableAvailabilityZonesForLoadBalancerInput) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error)
	EnableAvailabilityZonesForLoadBalancerWithContext(aws.Context, *elb.EnableAvailabilityZonesForLoadBalancerInput, ...request.Option) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error)
	EnableAvailabilityZonesForLoadBalancerRequest(*elb.EnableAvailabilityZonesForLoadBalancerInput) (*request.Request, *elb.EnableAvailabilityZonesForLoadBalancerOutput)

	ModifyLoadBalancerAttributes(*elb.ModifyLoadBalancerAttributesInput) (*elb.ModifyLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributesWithContext(aws.Context, *elb.ModifyLoadBalancerAttributesInput, ...request.Option) (*elb.ModifyLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributesRequest(*elb.ModifyLoadBalancerAttributesInput) (*request.Request, *elb.ModifyLoadBalancerAttributesOutput)

	RegisterInstancesWithLoadBalancer(*elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error)
	RegisterInstancesWithLoadBalancerWithContext(aws.Context, *elb.RegisterInstancesWithLoadBalancerInput, ...request.Option) (*elb.RegisterInstancesWithLoadBalancerOutput, error)
	RegisterInstancesWithLoadBalancerRequest(*elb.RegisterInstancesWithLoadBalancerInput) (*request.Request, *elb.RegisterInstancesWithLoadBalancerOutput)

	RemoveTags(*elb.RemoveTagsInput) (*elb.RemoveTagsOutput, error)
	RemoveTagsWithContext(aws.Context, *elb.RemoveTagsInput, ...request.Option) (*elb.RemoveTagsOutput, error)
	RemoveTagsRequest(*elb.RemoveTagsInput) (*request.Request, *elb.RemoveTagsOutput)

	SetLoadBalancerListenerSSLCertificate(*elb.SetLoadBalancerListenerSSLCertificateInput) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error)
	SetLoadBalancerListenerSSLCertificateWithContext(aws.Context, *elb.SetLoadBalancerListenerSSLCertificateInput, ...request.Option) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error)
	SetLoadBalancerListenerSSLCertificateRequest(*elb.SetLoadBalancerListenerSSLCertificateInput) (*request.Request, *elb.SetLoadBalancerListenerSSLCertificateOutput)

	SetLoadBalancerPoliciesForBackendServer(*elb.SetLoadBalancerPoliciesForBackendServerInput) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error)
	SetLoadBalancerPoliciesForBackendServerWithContext(aws.Context, *elb.SetLoadBalancerPoliciesForBackendServerInput, ...request.Option) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error)
	SetLoadBalancerPoliciesForBackendServerRequest(*elb.SetLoadBalancerPoliciesForBackendServerInput) (*request.Request, *elb.SetLoadBalancerPoliciesForBackendServerOutput)

	SetLoadBalancerPoliciesOfListener(*elb.SetLoadBalancerPoliciesOfListenerInput) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error)
	SetLoadBalancerPoliciesOfListenerWithContext(aws.Context, *elb.SetLoadBalancerPoliciesOfListenerInput, ...request.Option) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error)
	SetLoadBalancerPoliciesOfListenerRequest(*elb.SetLoadBalancerPoliciesOfListenerInput) (*request.Request, *elb.SetLoadBalancerPoliciesOfListenerOutput)

	WaitUntilAnyInstanceInService(*elb.DescribeInstanceHealthInput) error
	WaitUntilAnyInstanceInServiceWithContext(aws.Context, *elb.DescribeInstanceHealthInput, ...request.WaiterOption) error

	WaitUntilInstanceDeregistered(*elb.DescribeInstanceHealthInput) error
	WaitUntilInstanceDeregisteredWithContext(aws.Context, *elb.DescribeInstanceHealthInput, ...request.WaiterOption) error

	WaitUntilInstanceInService(*elb.DescribeInstanceHealthInput) error
	WaitUntilInstanceInServiceWithContext(aws.Context, *elb.DescribeInstanceHealthInput, ...request.WaiterOption) error
}

var _ ELBAPI = (*elb.ELB)(nil)
