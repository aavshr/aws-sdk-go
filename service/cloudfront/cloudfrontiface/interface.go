// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudfrontiface provides an interface to enable mocking the Amazon CloudFront service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudfrontiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/cloudfront"
)

// CloudFrontAPI provides an interface to enable mocking the
// cloudfront.CloudFront service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudFront.
//    func myFunc(svc cloudfrontiface.CloudFrontAPI) bool {
//        // Make svc.AssociateAlias request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloudfront.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudFrontClient struct {
//        cloudfrontiface.CloudFrontAPI
//    }
//    func (m *mockCloudFrontClient) AssociateAlias(input *cloudfront.AssociateAliasInput) (*cloudfront.AssociateAliasOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudFrontClient{}
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
type CloudFrontAPI interface {
	AssociateAlias(*cloudfront.AssociateAliasInput) (*cloudfront.AssociateAliasOutput, error)
	AssociateAliasWithContext(aws.Context, *cloudfront.AssociateAliasInput, ...request.Option) (*cloudfront.AssociateAliasOutput, error)
	AssociateAliasRequest(*cloudfront.AssociateAliasInput) (*request.Request, *cloudfront.AssociateAliasOutput)

	CreateCachePolicy(*cloudfront.CreateCachePolicyInput) (*cloudfront.CreateCachePolicyOutput, error)
	CreateCachePolicyWithContext(aws.Context, *cloudfront.CreateCachePolicyInput, ...request.Option) (*cloudfront.CreateCachePolicyOutput, error)
	CreateCachePolicyRequest(*cloudfront.CreateCachePolicyInput) (*request.Request, *cloudfront.CreateCachePolicyOutput)

	CreateCloudFrontOriginAccessIdentity(*cloudfront.CreateCloudFrontOriginAccessIdentityInput) (*cloudfront.CreateCloudFrontOriginAccessIdentityOutput, error)
	CreateCloudFrontOriginAccessIdentityWithContext(aws.Context, *cloudfront.CreateCloudFrontOriginAccessIdentityInput, ...request.Option) (*cloudfront.CreateCloudFrontOriginAccessIdentityOutput, error)
	CreateCloudFrontOriginAccessIdentityRequest(*cloudfront.CreateCloudFrontOriginAccessIdentityInput) (*request.Request, *cloudfront.CreateCloudFrontOriginAccessIdentityOutput)

	CreateDistribution(*cloudfront.CreateDistributionInput) (*cloudfront.CreateDistributionOutput, error)
	CreateDistributionWithContext(aws.Context, *cloudfront.CreateDistributionInput, ...request.Option) (*cloudfront.CreateDistributionOutput, error)
	CreateDistributionRequest(*cloudfront.CreateDistributionInput) (*request.Request, *cloudfront.CreateDistributionOutput)

	CreateDistributionWithTags(*cloudfront.CreateDistributionWithTagsInput) (*cloudfront.CreateDistributionWithTagsOutput, error)
	CreateDistributionWithTagsWithContext(aws.Context, *cloudfront.CreateDistributionWithTagsInput, ...request.Option) (*cloudfront.CreateDistributionWithTagsOutput, error)
	CreateDistributionWithTagsRequest(*cloudfront.CreateDistributionWithTagsInput) (*request.Request, *cloudfront.CreateDistributionWithTagsOutput)

	CreateFieldLevelEncryptionConfig(*cloudfront.CreateFieldLevelEncryptionConfigInput) (*cloudfront.CreateFieldLevelEncryptionConfigOutput, error)
	CreateFieldLevelEncryptionConfigWithContext(aws.Context, *cloudfront.CreateFieldLevelEncryptionConfigInput, ...request.Option) (*cloudfront.CreateFieldLevelEncryptionConfigOutput, error)
	CreateFieldLevelEncryptionConfigRequest(*cloudfront.CreateFieldLevelEncryptionConfigInput) (*request.Request, *cloudfront.CreateFieldLevelEncryptionConfigOutput)

	CreateFieldLevelEncryptionProfile(*cloudfront.CreateFieldLevelEncryptionProfileInput) (*cloudfront.CreateFieldLevelEncryptionProfileOutput, error)
	CreateFieldLevelEncryptionProfileWithContext(aws.Context, *cloudfront.CreateFieldLevelEncryptionProfileInput, ...request.Option) (*cloudfront.CreateFieldLevelEncryptionProfileOutput, error)
	CreateFieldLevelEncryptionProfileRequest(*cloudfront.CreateFieldLevelEncryptionProfileInput) (*request.Request, *cloudfront.CreateFieldLevelEncryptionProfileOutput)

	CreateFunction(*cloudfront.CreateFunctionInput) (*cloudfront.CreateFunctionOutput, error)
	CreateFunctionWithContext(aws.Context, *cloudfront.CreateFunctionInput, ...request.Option) (*cloudfront.CreateFunctionOutput, error)
	CreateFunctionRequest(*cloudfront.CreateFunctionInput) (*request.Request, *cloudfront.CreateFunctionOutput)

	CreateInvalidation(*cloudfront.CreateInvalidationInput) (*cloudfront.CreateInvalidationOutput, error)
	CreateInvalidationWithContext(aws.Context, *cloudfront.CreateInvalidationInput, ...request.Option) (*cloudfront.CreateInvalidationOutput, error)
	CreateInvalidationRequest(*cloudfront.CreateInvalidationInput) (*request.Request, *cloudfront.CreateInvalidationOutput)

	CreateKeyGroup(*cloudfront.CreateKeyGroupInput) (*cloudfront.CreateKeyGroupOutput, error)
	CreateKeyGroupWithContext(aws.Context, *cloudfront.CreateKeyGroupInput, ...request.Option) (*cloudfront.CreateKeyGroupOutput, error)
	CreateKeyGroupRequest(*cloudfront.CreateKeyGroupInput) (*request.Request, *cloudfront.CreateKeyGroupOutput)

	CreateMonitoringSubscription(*cloudfront.CreateMonitoringSubscriptionInput) (*cloudfront.CreateMonitoringSubscriptionOutput, error)
	CreateMonitoringSubscriptionWithContext(aws.Context, *cloudfront.CreateMonitoringSubscriptionInput, ...request.Option) (*cloudfront.CreateMonitoringSubscriptionOutput, error)
	CreateMonitoringSubscriptionRequest(*cloudfront.CreateMonitoringSubscriptionInput) (*request.Request, *cloudfront.CreateMonitoringSubscriptionOutput)

	CreateOriginRequestPolicy(*cloudfront.CreateOriginRequestPolicyInput) (*cloudfront.CreateOriginRequestPolicyOutput, error)
	CreateOriginRequestPolicyWithContext(aws.Context, *cloudfront.CreateOriginRequestPolicyInput, ...request.Option) (*cloudfront.CreateOriginRequestPolicyOutput, error)
	CreateOriginRequestPolicyRequest(*cloudfront.CreateOriginRequestPolicyInput) (*request.Request, *cloudfront.CreateOriginRequestPolicyOutput)

	CreatePublicKey(*cloudfront.CreatePublicKeyInput) (*cloudfront.CreatePublicKeyOutput, error)
	CreatePublicKeyWithContext(aws.Context, *cloudfront.CreatePublicKeyInput, ...request.Option) (*cloudfront.CreatePublicKeyOutput, error)
	CreatePublicKeyRequest(*cloudfront.CreatePublicKeyInput) (*request.Request, *cloudfront.CreatePublicKeyOutput)

	CreateRealtimeLogConfig(*cloudfront.CreateRealtimeLogConfigInput) (*cloudfront.CreateRealtimeLogConfigOutput, error)
	CreateRealtimeLogConfigWithContext(aws.Context, *cloudfront.CreateRealtimeLogConfigInput, ...request.Option) (*cloudfront.CreateRealtimeLogConfigOutput, error)
	CreateRealtimeLogConfigRequest(*cloudfront.CreateRealtimeLogConfigInput) (*request.Request, *cloudfront.CreateRealtimeLogConfigOutput)

	CreateStreamingDistribution(*cloudfront.CreateStreamingDistributionInput) (*cloudfront.CreateStreamingDistributionOutput, error)
	CreateStreamingDistributionWithContext(aws.Context, *cloudfront.CreateStreamingDistributionInput, ...request.Option) (*cloudfront.CreateStreamingDistributionOutput, error)
	CreateStreamingDistributionRequest(*cloudfront.CreateStreamingDistributionInput) (*request.Request, *cloudfront.CreateStreamingDistributionOutput)

	CreateStreamingDistributionWithTags(*cloudfront.CreateStreamingDistributionWithTagsInput) (*cloudfront.CreateStreamingDistributionWithTagsOutput, error)
	CreateStreamingDistributionWithTagsWithContext(aws.Context, *cloudfront.CreateStreamingDistributionWithTagsInput, ...request.Option) (*cloudfront.CreateStreamingDistributionWithTagsOutput, error)
	CreateStreamingDistributionWithTagsRequest(*cloudfront.CreateStreamingDistributionWithTagsInput) (*request.Request, *cloudfront.CreateStreamingDistributionWithTagsOutput)

	DeleteCachePolicy(*cloudfront.DeleteCachePolicyInput) (*cloudfront.DeleteCachePolicyOutput, error)
	DeleteCachePolicyWithContext(aws.Context, *cloudfront.DeleteCachePolicyInput, ...request.Option) (*cloudfront.DeleteCachePolicyOutput, error)
	DeleteCachePolicyRequest(*cloudfront.DeleteCachePolicyInput) (*request.Request, *cloudfront.DeleteCachePolicyOutput)

	DeleteCloudFrontOriginAccessIdentity(*cloudfront.DeleteCloudFrontOriginAccessIdentityInput) (*cloudfront.DeleteCloudFrontOriginAccessIdentityOutput, error)
	DeleteCloudFrontOriginAccessIdentityWithContext(aws.Context, *cloudfront.DeleteCloudFrontOriginAccessIdentityInput, ...request.Option) (*cloudfront.DeleteCloudFrontOriginAccessIdentityOutput, error)
	DeleteCloudFrontOriginAccessIdentityRequest(*cloudfront.DeleteCloudFrontOriginAccessIdentityInput) (*request.Request, *cloudfront.DeleteCloudFrontOriginAccessIdentityOutput)

	DeleteDistribution(*cloudfront.DeleteDistributionInput) (*cloudfront.DeleteDistributionOutput, error)
	DeleteDistributionWithContext(aws.Context, *cloudfront.DeleteDistributionInput, ...request.Option) (*cloudfront.DeleteDistributionOutput, error)
	DeleteDistributionRequest(*cloudfront.DeleteDistributionInput) (*request.Request, *cloudfront.DeleteDistributionOutput)

	DeleteFieldLevelEncryptionConfig(*cloudfront.DeleteFieldLevelEncryptionConfigInput) (*cloudfront.DeleteFieldLevelEncryptionConfigOutput, error)
	DeleteFieldLevelEncryptionConfigWithContext(aws.Context, *cloudfront.DeleteFieldLevelEncryptionConfigInput, ...request.Option) (*cloudfront.DeleteFieldLevelEncryptionConfigOutput, error)
	DeleteFieldLevelEncryptionConfigRequest(*cloudfront.DeleteFieldLevelEncryptionConfigInput) (*request.Request, *cloudfront.DeleteFieldLevelEncryptionConfigOutput)

	DeleteFieldLevelEncryptionProfile(*cloudfront.DeleteFieldLevelEncryptionProfileInput) (*cloudfront.DeleteFieldLevelEncryptionProfileOutput, error)
	DeleteFieldLevelEncryptionProfileWithContext(aws.Context, *cloudfront.DeleteFieldLevelEncryptionProfileInput, ...request.Option) (*cloudfront.DeleteFieldLevelEncryptionProfileOutput, error)
	DeleteFieldLevelEncryptionProfileRequest(*cloudfront.DeleteFieldLevelEncryptionProfileInput) (*request.Request, *cloudfront.DeleteFieldLevelEncryptionProfileOutput)

	DeleteFunction(*cloudfront.DeleteFunctionInput) (*cloudfront.DeleteFunctionOutput, error)
	DeleteFunctionWithContext(aws.Context, *cloudfront.DeleteFunctionInput, ...request.Option) (*cloudfront.DeleteFunctionOutput, error)
	DeleteFunctionRequest(*cloudfront.DeleteFunctionInput) (*request.Request, *cloudfront.DeleteFunctionOutput)

	DeleteKeyGroup(*cloudfront.DeleteKeyGroupInput) (*cloudfront.DeleteKeyGroupOutput, error)
	DeleteKeyGroupWithContext(aws.Context, *cloudfront.DeleteKeyGroupInput, ...request.Option) (*cloudfront.DeleteKeyGroupOutput, error)
	DeleteKeyGroupRequest(*cloudfront.DeleteKeyGroupInput) (*request.Request, *cloudfront.DeleteKeyGroupOutput)

	DeleteMonitoringSubscription(*cloudfront.DeleteMonitoringSubscriptionInput) (*cloudfront.DeleteMonitoringSubscriptionOutput, error)
	DeleteMonitoringSubscriptionWithContext(aws.Context, *cloudfront.DeleteMonitoringSubscriptionInput, ...request.Option) (*cloudfront.DeleteMonitoringSubscriptionOutput, error)
	DeleteMonitoringSubscriptionRequest(*cloudfront.DeleteMonitoringSubscriptionInput) (*request.Request, *cloudfront.DeleteMonitoringSubscriptionOutput)

	DeleteOriginRequestPolicy(*cloudfront.DeleteOriginRequestPolicyInput) (*cloudfront.DeleteOriginRequestPolicyOutput, error)
	DeleteOriginRequestPolicyWithContext(aws.Context, *cloudfront.DeleteOriginRequestPolicyInput, ...request.Option) (*cloudfront.DeleteOriginRequestPolicyOutput, error)
	DeleteOriginRequestPolicyRequest(*cloudfront.DeleteOriginRequestPolicyInput) (*request.Request, *cloudfront.DeleteOriginRequestPolicyOutput)

	DeletePublicKey(*cloudfront.DeletePublicKeyInput) (*cloudfront.DeletePublicKeyOutput, error)
	DeletePublicKeyWithContext(aws.Context, *cloudfront.DeletePublicKeyInput, ...request.Option) (*cloudfront.DeletePublicKeyOutput, error)
	DeletePublicKeyRequest(*cloudfront.DeletePublicKeyInput) (*request.Request, *cloudfront.DeletePublicKeyOutput)

	DeleteRealtimeLogConfig(*cloudfront.DeleteRealtimeLogConfigInput) (*cloudfront.DeleteRealtimeLogConfigOutput, error)
	DeleteRealtimeLogConfigWithContext(aws.Context, *cloudfront.DeleteRealtimeLogConfigInput, ...request.Option) (*cloudfront.DeleteRealtimeLogConfigOutput, error)
	DeleteRealtimeLogConfigRequest(*cloudfront.DeleteRealtimeLogConfigInput) (*request.Request, *cloudfront.DeleteRealtimeLogConfigOutput)

	DeleteStreamingDistribution(*cloudfront.DeleteStreamingDistributionInput) (*cloudfront.DeleteStreamingDistributionOutput, error)
	DeleteStreamingDistributionWithContext(aws.Context, *cloudfront.DeleteStreamingDistributionInput, ...request.Option) (*cloudfront.DeleteStreamingDistributionOutput, error)
	DeleteStreamingDistributionRequest(*cloudfront.DeleteStreamingDistributionInput) (*request.Request, *cloudfront.DeleteStreamingDistributionOutput)

	DescribeFunction(*cloudfront.DescribeFunctionInput) (*cloudfront.DescribeFunctionOutput, error)
	DescribeFunctionWithContext(aws.Context, *cloudfront.DescribeFunctionInput, ...request.Option) (*cloudfront.DescribeFunctionOutput, error)
	DescribeFunctionRequest(*cloudfront.DescribeFunctionInput) (*request.Request, *cloudfront.DescribeFunctionOutput)

	GetCachePolicy(*cloudfront.GetCachePolicyInput) (*cloudfront.GetCachePolicyOutput, error)
	GetCachePolicyWithContext(aws.Context, *cloudfront.GetCachePolicyInput, ...request.Option) (*cloudfront.GetCachePolicyOutput, error)
	GetCachePolicyRequest(*cloudfront.GetCachePolicyInput) (*request.Request, *cloudfront.GetCachePolicyOutput)

	GetCachePolicyConfig(*cloudfront.GetCachePolicyConfigInput) (*cloudfront.GetCachePolicyConfigOutput, error)
	GetCachePolicyConfigWithContext(aws.Context, *cloudfront.GetCachePolicyConfigInput, ...request.Option) (*cloudfront.GetCachePolicyConfigOutput, error)
	GetCachePolicyConfigRequest(*cloudfront.GetCachePolicyConfigInput) (*request.Request, *cloudfront.GetCachePolicyConfigOutput)

	GetCloudFrontOriginAccessIdentity(*cloudfront.GetCloudFrontOriginAccessIdentityInput) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error)
	GetCloudFrontOriginAccessIdentityWithContext(aws.Context, *cloudfront.GetCloudFrontOriginAccessIdentityInput, ...request.Option) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error)
	GetCloudFrontOriginAccessIdentityRequest(*cloudfront.GetCloudFrontOriginAccessIdentityInput) (*request.Request, *cloudfront.GetCloudFrontOriginAccessIdentityOutput)

	GetCloudFrontOriginAccessIdentityConfig(*cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error)
	GetCloudFrontOriginAccessIdentityConfigWithContext(aws.Context, *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput, ...request.Option) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error)
	GetCloudFrontOriginAccessIdentityConfigRequest(*cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) (*request.Request, *cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput)

	GetDistribution(*cloudfront.GetDistributionInput) (*cloudfront.GetDistributionOutput, error)
	GetDistributionWithContext(aws.Context, *cloudfront.GetDistributionInput, ...request.Option) (*cloudfront.GetDistributionOutput, error)
	GetDistributionRequest(*cloudfront.GetDistributionInput) (*request.Request, *cloudfront.GetDistributionOutput)

	GetDistributionConfig(*cloudfront.GetDistributionConfigInput) (*cloudfront.GetDistributionConfigOutput, error)
	GetDistributionConfigWithContext(aws.Context, *cloudfront.GetDistributionConfigInput, ...request.Option) (*cloudfront.GetDistributionConfigOutput, error)
	GetDistributionConfigRequest(*cloudfront.GetDistributionConfigInput) (*request.Request, *cloudfront.GetDistributionConfigOutput)

	GetFieldLevelEncryption(*cloudfront.GetFieldLevelEncryptionInput) (*cloudfront.GetFieldLevelEncryptionOutput, error)
	GetFieldLevelEncryptionWithContext(aws.Context, *cloudfront.GetFieldLevelEncryptionInput, ...request.Option) (*cloudfront.GetFieldLevelEncryptionOutput, error)
	GetFieldLevelEncryptionRequest(*cloudfront.GetFieldLevelEncryptionInput) (*request.Request, *cloudfront.GetFieldLevelEncryptionOutput)

	GetFieldLevelEncryptionConfig(*cloudfront.GetFieldLevelEncryptionConfigInput) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error)
	GetFieldLevelEncryptionConfigWithContext(aws.Context, *cloudfront.GetFieldLevelEncryptionConfigInput, ...request.Option) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error)
	GetFieldLevelEncryptionConfigRequest(*cloudfront.GetFieldLevelEncryptionConfigInput) (*request.Request, *cloudfront.GetFieldLevelEncryptionConfigOutput)

	GetFieldLevelEncryptionProfile(*cloudfront.GetFieldLevelEncryptionProfileInput) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error)
	GetFieldLevelEncryptionProfileWithContext(aws.Context, *cloudfront.GetFieldLevelEncryptionProfileInput, ...request.Option) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error)
	GetFieldLevelEncryptionProfileRequest(*cloudfront.GetFieldLevelEncryptionProfileInput) (*request.Request, *cloudfront.GetFieldLevelEncryptionProfileOutput)

	GetFieldLevelEncryptionProfileConfig(*cloudfront.GetFieldLevelEncryptionProfileConfigInput) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error)
	GetFieldLevelEncryptionProfileConfigWithContext(aws.Context, *cloudfront.GetFieldLevelEncryptionProfileConfigInput, ...request.Option) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error)
	GetFieldLevelEncryptionProfileConfigRequest(*cloudfront.GetFieldLevelEncryptionProfileConfigInput) (*request.Request, *cloudfront.GetFieldLevelEncryptionProfileConfigOutput)

	GetFunction(*cloudfront.GetFunctionInput) (*cloudfront.GetFunctionOutput, error)
	GetFunctionWithContext(aws.Context, *cloudfront.GetFunctionInput, ...request.Option) (*cloudfront.GetFunctionOutput, error)
	GetFunctionRequest(*cloudfront.GetFunctionInput) (*request.Request, *cloudfront.GetFunctionOutput)

	GetInvalidation(*cloudfront.GetInvalidationInput) (*cloudfront.GetInvalidationOutput, error)
	GetInvalidationWithContext(aws.Context, *cloudfront.GetInvalidationInput, ...request.Option) (*cloudfront.GetInvalidationOutput, error)
	GetInvalidationRequest(*cloudfront.GetInvalidationInput) (*request.Request, *cloudfront.GetInvalidationOutput)

	GetKeyGroup(*cloudfront.GetKeyGroupInput) (*cloudfront.GetKeyGroupOutput, error)
	GetKeyGroupWithContext(aws.Context, *cloudfront.GetKeyGroupInput, ...request.Option) (*cloudfront.GetKeyGroupOutput, error)
	GetKeyGroupRequest(*cloudfront.GetKeyGroupInput) (*request.Request, *cloudfront.GetKeyGroupOutput)

	GetKeyGroupConfig(*cloudfront.GetKeyGroupConfigInput) (*cloudfront.GetKeyGroupConfigOutput, error)
	GetKeyGroupConfigWithContext(aws.Context, *cloudfront.GetKeyGroupConfigInput, ...request.Option) (*cloudfront.GetKeyGroupConfigOutput, error)
	GetKeyGroupConfigRequest(*cloudfront.GetKeyGroupConfigInput) (*request.Request, *cloudfront.GetKeyGroupConfigOutput)

	GetMonitoringSubscription(*cloudfront.GetMonitoringSubscriptionInput) (*cloudfront.GetMonitoringSubscriptionOutput, error)
	GetMonitoringSubscriptionWithContext(aws.Context, *cloudfront.GetMonitoringSubscriptionInput, ...request.Option) (*cloudfront.GetMonitoringSubscriptionOutput, error)
	GetMonitoringSubscriptionRequest(*cloudfront.GetMonitoringSubscriptionInput) (*request.Request, *cloudfront.GetMonitoringSubscriptionOutput)

	GetOriginRequestPolicy(*cloudfront.GetOriginRequestPolicyInput) (*cloudfront.GetOriginRequestPolicyOutput, error)
	GetOriginRequestPolicyWithContext(aws.Context, *cloudfront.GetOriginRequestPolicyInput, ...request.Option) (*cloudfront.GetOriginRequestPolicyOutput, error)
	GetOriginRequestPolicyRequest(*cloudfront.GetOriginRequestPolicyInput) (*request.Request, *cloudfront.GetOriginRequestPolicyOutput)

	GetOriginRequestPolicyConfig(*cloudfront.GetOriginRequestPolicyConfigInput) (*cloudfront.GetOriginRequestPolicyConfigOutput, error)
	GetOriginRequestPolicyConfigWithContext(aws.Context, *cloudfront.GetOriginRequestPolicyConfigInput, ...request.Option) (*cloudfront.GetOriginRequestPolicyConfigOutput, error)
	GetOriginRequestPolicyConfigRequest(*cloudfront.GetOriginRequestPolicyConfigInput) (*request.Request, *cloudfront.GetOriginRequestPolicyConfigOutput)

	GetPublicKey(*cloudfront.GetPublicKeyInput) (*cloudfront.GetPublicKeyOutput, error)
	GetPublicKeyWithContext(aws.Context, *cloudfront.GetPublicKeyInput, ...request.Option) (*cloudfront.GetPublicKeyOutput, error)
	GetPublicKeyRequest(*cloudfront.GetPublicKeyInput) (*request.Request, *cloudfront.GetPublicKeyOutput)

	GetPublicKeyConfig(*cloudfront.GetPublicKeyConfigInput) (*cloudfront.GetPublicKeyConfigOutput, error)
	GetPublicKeyConfigWithContext(aws.Context, *cloudfront.GetPublicKeyConfigInput, ...request.Option) (*cloudfront.GetPublicKeyConfigOutput, error)
	GetPublicKeyConfigRequest(*cloudfront.GetPublicKeyConfigInput) (*request.Request, *cloudfront.GetPublicKeyConfigOutput)

	GetRealtimeLogConfig(*cloudfront.GetRealtimeLogConfigInput) (*cloudfront.GetRealtimeLogConfigOutput, error)
	GetRealtimeLogConfigWithContext(aws.Context, *cloudfront.GetRealtimeLogConfigInput, ...request.Option) (*cloudfront.GetRealtimeLogConfigOutput, error)
	GetRealtimeLogConfigRequest(*cloudfront.GetRealtimeLogConfigInput) (*request.Request, *cloudfront.GetRealtimeLogConfigOutput)

	GetStreamingDistribution(*cloudfront.GetStreamingDistributionInput) (*cloudfront.GetStreamingDistributionOutput, error)
	GetStreamingDistributionWithContext(aws.Context, *cloudfront.GetStreamingDistributionInput, ...request.Option) (*cloudfront.GetStreamingDistributionOutput, error)
	GetStreamingDistributionRequest(*cloudfront.GetStreamingDistributionInput) (*request.Request, *cloudfront.GetStreamingDistributionOutput)

	GetStreamingDistributionConfig(*cloudfront.GetStreamingDistributionConfigInput) (*cloudfront.GetStreamingDistributionConfigOutput, error)
	GetStreamingDistributionConfigWithContext(aws.Context, *cloudfront.GetStreamingDistributionConfigInput, ...request.Option) (*cloudfront.GetStreamingDistributionConfigOutput, error)
	GetStreamingDistributionConfigRequest(*cloudfront.GetStreamingDistributionConfigInput) (*request.Request, *cloudfront.GetStreamingDistributionConfigOutput)

	ListCachePolicies(*cloudfront.ListCachePoliciesInput) (*cloudfront.ListCachePoliciesOutput, error)
	ListCachePoliciesWithContext(aws.Context, *cloudfront.ListCachePoliciesInput, ...request.Option) (*cloudfront.ListCachePoliciesOutput, error)
	ListCachePoliciesRequest(*cloudfront.ListCachePoliciesInput) (*request.Request, *cloudfront.ListCachePoliciesOutput)

	ListCloudFrontOriginAccessIdentities(*cloudfront.ListCloudFrontOriginAccessIdentitiesInput) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error)
	ListCloudFrontOriginAccessIdentitiesWithContext(aws.Context, *cloudfront.ListCloudFrontOriginAccessIdentitiesInput, ...request.Option) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error)
	ListCloudFrontOriginAccessIdentitiesRequest(*cloudfront.ListCloudFrontOriginAccessIdentitiesInput) (*request.Request, *cloudfront.ListCloudFrontOriginAccessIdentitiesOutput)

	ListCloudFrontOriginAccessIdentitiesPages(*cloudfront.ListCloudFrontOriginAccessIdentitiesInput, func(*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, bool) bool) error
	ListCloudFrontOriginAccessIdentitiesPagesWithContext(aws.Context, *cloudfront.ListCloudFrontOriginAccessIdentitiesInput, func(*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, bool) bool, ...request.Option) error

	ListConflictingAliases(*cloudfront.ListConflictingAliasesInput) (*cloudfront.ListConflictingAliasesOutput, error)
	ListConflictingAliasesWithContext(aws.Context, *cloudfront.ListConflictingAliasesInput, ...request.Option) (*cloudfront.ListConflictingAliasesOutput, error)
	ListConflictingAliasesRequest(*cloudfront.ListConflictingAliasesInput) (*request.Request, *cloudfront.ListConflictingAliasesOutput)

	ListDistributions(*cloudfront.ListDistributionsInput) (*cloudfront.ListDistributionsOutput, error)
	ListDistributionsWithContext(aws.Context, *cloudfront.ListDistributionsInput, ...request.Option) (*cloudfront.ListDistributionsOutput, error)
	ListDistributionsRequest(*cloudfront.ListDistributionsInput) (*request.Request, *cloudfront.ListDistributionsOutput)

	ListDistributionsPages(*cloudfront.ListDistributionsInput, func(*cloudfront.ListDistributionsOutput, bool) bool) error
	ListDistributionsPagesWithContext(aws.Context, *cloudfront.ListDistributionsInput, func(*cloudfront.ListDistributionsOutput, bool) bool, ...request.Option) error

	ListDistributionsByCachePolicyId(*cloudfront.ListDistributionsByCachePolicyIdInput) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error)
	ListDistributionsByCachePolicyIdWithContext(aws.Context, *cloudfront.ListDistributionsByCachePolicyIdInput, ...request.Option) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error)
	ListDistributionsByCachePolicyIdRequest(*cloudfront.ListDistributionsByCachePolicyIdInput) (*request.Request, *cloudfront.ListDistributionsByCachePolicyIdOutput)

	ListDistributionsByKeyGroup(*cloudfront.ListDistributionsByKeyGroupInput) (*cloudfront.ListDistributionsByKeyGroupOutput, error)
	ListDistributionsByKeyGroupWithContext(aws.Context, *cloudfront.ListDistributionsByKeyGroupInput, ...request.Option) (*cloudfront.ListDistributionsByKeyGroupOutput, error)
	ListDistributionsByKeyGroupRequest(*cloudfront.ListDistributionsByKeyGroupInput) (*request.Request, *cloudfront.ListDistributionsByKeyGroupOutput)

	ListDistributionsByOriginRequestPolicyId(*cloudfront.ListDistributionsByOriginRequestPolicyIdInput) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error)
	ListDistributionsByOriginRequestPolicyIdWithContext(aws.Context, *cloudfront.ListDistributionsByOriginRequestPolicyIdInput, ...request.Option) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error)
	ListDistributionsByOriginRequestPolicyIdRequest(*cloudfront.ListDistributionsByOriginRequestPolicyIdInput) (*request.Request, *cloudfront.ListDistributionsByOriginRequestPolicyIdOutput)

	ListDistributionsByRealtimeLogConfig(*cloudfront.ListDistributionsByRealtimeLogConfigInput) (*cloudfront.ListDistributionsByRealtimeLogConfigOutput, error)
	ListDistributionsByRealtimeLogConfigWithContext(aws.Context, *cloudfront.ListDistributionsByRealtimeLogConfigInput, ...request.Option) (*cloudfront.ListDistributionsByRealtimeLogConfigOutput, error)
	ListDistributionsByRealtimeLogConfigRequest(*cloudfront.ListDistributionsByRealtimeLogConfigInput) (*request.Request, *cloudfront.ListDistributionsByRealtimeLogConfigOutput)

	ListDistributionsByWebACLId(*cloudfront.ListDistributionsByWebACLIdInput) (*cloudfront.ListDistributionsByWebACLIdOutput, error)
	ListDistributionsByWebACLIdWithContext(aws.Context, *cloudfront.ListDistributionsByWebACLIdInput, ...request.Option) (*cloudfront.ListDistributionsByWebACLIdOutput, error)
	ListDistributionsByWebACLIdRequest(*cloudfront.ListDistributionsByWebACLIdInput) (*request.Request, *cloudfront.ListDistributionsByWebACLIdOutput)

	ListFieldLevelEncryptionConfigs(*cloudfront.ListFieldLevelEncryptionConfigsInput) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error)
	ListFieldLevelEncryptionConfigsWithContext(aws.Context, *cloudfront.ListFieldLevelEncryptionConfigsInput, ...request.Option) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error)
	ListFieldLevelEncryptionConfigsRequest(*cloudfront.ListFieldLevelEncryptionConfigsInput) (*request.Request, *cloudfront.ListFieldLevelEncryptionConfigsOutput)

	ListFieldLevelEncryptionProfiles(*cloudfront.ListFieldLevelEncryptionProfilesInput) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error)
	ListFieldLevelEncryptionProfilesWithContext(aws.Context, *cloudfront.ListFieldLevelEncryptionProfilesInput, ...request.Option) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error)
	ListFieldLevelEncryptionProfilesRequest(*cloudfront.ListFieldLevelEncryptionProfilesInput) (*request.Request, *cloudfront.ListFieldLevelEncryptionProfilesOutput)

	ListFunctions(*cloudfront.ListFunctionsInput) (*cloudfront.ListFunctionsOutput, error)
	ListFunctionsWithContext(aws.Context, *cloudfront.ListFunctionsInput, ...request.Option) (*cloudfront.ListFunctionsOutput, error)
	ListFunctionsRequest(*cloudfront.ListFunctionsInput) (*request.Request, *cloudfront.ListFunctionsOutput)

	ListInvalidations(*cloudfront.ListInvalidationsInput) (*cloudfront.ListInvalidationsOutput, error)
	ListInvalidationsWithContext(aws.Context, *cloudfront.ListInvalidationsInput, ...request.Option) (*cloudfront.ListInvalidationsOutput, error)
	ListInvalidationsRequest(*cloudfront.ListInvalidationsInput) (*request.Request, *cloudfront.ListInvalidationsOutput)

	ListInvalidationsPages(*cloudfront.ListInvalidationsInput, func(*cloudfront.ListInvalidationsOutput, bool) bool) error
	ListInvalidationsPagesWithContext(aws.Context, *cloudfront.ListInvalidationsInput, func(*cloudfront.ListInvalidationsOutput, bool) bool, ...request.Option) error

	ListKeyGroups(*cloudfront.ListKeyGroupsInput) (*cloudfront.ListKeyGroupsOutput, error)
	ListKeyGroupsWithContext(aws.Context, *cloudfront.ListKeyGroupsInput, ...request.Option) (*cloudfront.ListKeyGroupsOutput, error)
	ListKeyGroupsRequest(*cloudfront.ListKeyGroupsInput) (*request.Request, *cloudfront.ListKeyGroupsOutput)

	ListOriginRequestPolicies(*cloudfront.ListOriginRequestPoliciesInput) (*cloudfront.ListOriginRequestPoliciesOutput, error)
	ListOriginRequestPoliciesWithContext(aws.Context, *cloudfront.ListOriginRequestPoliciesInput, ...request.Option) (*cloudfront.ListOriginRequestPoliciesOutput, error)
	ListOriginRequestPoliciesRequest(*cloudfront.ListOriginRequestPoliciesInput) (*request.Request, *cloudfront.ListOriginRequestPoliciesOutput)

	ListPublicKeys(*cloudfront.ListPublicKeysInput) (*cloudfront.ListPublicKeysOutput, error)
	ListPublicKeysWithContext(aws.Context, *cloudfront.ListPublicKeysInput, ...request.Option) (*cloudfront.ListPublicKeysOutput, error)
	ListPublicKeysRequest(*cloudfront.ListPublicKeysInput) (*request.Request, *cloudfront.ListPublicKeysOutput)

	ListRealtimeLogConfigs(*cloudfront.ListRealtimeLogConfigsInput) (*cloudfront.ListRealtimeLogConfigsOutput, error)
	ListRealtimeLogConfigsWithContext(aws.Context, *cloudfront.ListRealtimeLogConfigsInput, ...request.Option) (*cloudfront.ListRealtimeLogConfigsOutput, error)
	ListRealtimeLogConfigsRequest(*cloudfront.ListRealtimeLogConfigsInput) (*request.Request, *cloudfront.ListRealtimeLogConfigsOutput)

	ListStreamingDistributions(*cloudfront.ListStreamingDistributionsInput) (*cloudfront.ListStreamingDistributionsOutput, error)
	ListStreamingDistributionsWithContext(aws.Context, *cloudfront.ListStreamingDistributionsInput, ...request.Option) (*cloudfront.ListStreamingDistributionsOutput, error)
	ListStreamingDistributionsRequest(*cloudfront.ListStreamingDistributionsInput) (*request.Request, *cloudfront.ListStreamingDistributionsOutput)

	ListStreamingDistributionsPages(*cloudfront.ListStreamingDistributionsInput, func(*cloudfront.ListStreamingDistributionsOutput, bool) bool) error
	ListStreamingDistributionsPagesWithContext(aws.Context, *cloudfront.ListStreamingDistributionsInput, func(*cloudfront.ListStreamingDistributionsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*cloudfront.ListTagsForResourceInput) (*cloudfront.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *cloudfront.ListTagsForResourceInput, ...request.Option) (*cloudfront.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*cloudfront.ListTagsForResourceInput) (*request.Request, *cloudfront.ListTagsForResourceOutput)

	PublishFunction(*cloudfront.PublishFunctionInput) (*cloudfront.PublishFunctionOutput, error)
	PublishFunctionWithContext(aws.Context, *cloudfront.PublishFunctionInput, ...request.Option) (*cloudfront.PublishFunctionOutput, error)
	PublishFunctionRequest(*cloudfront.PublishFunctionInput) (*request.Request, *cloudfront.PublishFunctionOutput)

	TagResource(*cloudfront.TagResourceInput) (*cloudfront.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *cloudfront.TagResourceInput, ...request.Option) (*cloudfront.TagResourceOutput, error)
	TagResourceRequest(*cloudfront.TagResourceInput) (*request.Request, *cloudfront.TagResourceOutput)

	TestFunction(*cloudfront.TestFunctionInput) (*cloudfront.TestFunctionOutput, error)
	TestFunctionWithContext(aws.Context, *cloudfront.TestFunctionInput, ...request.Option) (*cloudfront.TestFunctionOutput, error)
	TestFunctionRequest(*cloudfront.TestFunctionInput) (*request.Request, *cloudfront.TestFunctionOutput)

	UntagResource(*cloudfront.UntagResourceInput) (*cloudfront.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *cloudfront.UntagResourceInput, ...request.Option) (*cloudfront.UntagResourceOutput, error)
	UntagResourceRequest(*cloudfront.UntagResourceInput) (*request.Request, *cloudfront.UntagResourceOutput)

	UpdateCachePolicy(*cloudfront.UpdateCachePolicyInput) (*cloudfront.UpdateCachePolicyOutput, error)
	UpdateCachePolicyWithContext(aws.Context, *cloudfront.UpdateCachePolicyInput, ...request.Option) (*cloudfront.UpdateCachePolicyOutput, error)
	UpdateCachePolicyRequest(*cloudfront.UpdateCachePolicyInput) (*request.Request, *cloudfront.UpdateCachePolicyOutput)

	UpdateCloudFrontOriginAccessIdentity(*cloudfront.UpdateCloudFrontOriginAccessIdentityInput) (*cloudfront.UpdateCloudFrontOriginAccessIdentityOutput, error)
	UpdateCloudFrontOriginAccessIdentityWithContext(aws.Context, *cloudfront.UpdateCloudFrontOriginAccessIdentityInput, ...request.Option) (*cloudfront.UpdateCloudFrontOriginAccessIdentityOutput, error)
	UpdateCloudFrontOriginAccessIdentityRequest(*cloudfront.UpdateCloudFrontOriginAccessIdentityInput) (*request.Request, *cloudfront.UpdateCloudFrontOriginAccessIdentityOutput)

	UpdateDistribution(*cloudfront.UpdateDistributionInput) (*cloudfront.UpdateDistributionOutput, error)
	UpdateDistributionWithContext(aws.Context, *cloudfront.UpdateDistributionInput, ...request.Option) (*cloudfront.UpdateDistributionOutput, error)
	UpdateDistributionRequest(*cloudfront.UpdateDistributionInput) (*request.Request, *cloudfront.UpdateDistributionOutput)

	UpdateFieldLevelEncryptionConfig(*cloudfront.UpdateFieldLevelEncryptionConfigInput) (*cloudfront.UpdateFieldLevelEncryptionConfigOutput, error)
	UpdateFieldLevelEncryptionConfigWithContext(aws.Context, *cloudfront.UpdateFieldLevelEncryptionConfigInput, ...request.Option) (*cloudfront.UpdateFieldLevelEncryptionConfigOutput, error)
	UpdateFieldLevelEncryptionConfigRequest(*cloudfront.UpdateFieldLevelEncryptionConfigInput) (*request.Request, *cloudfront.UpdateFieldLevelEncryptionConfigOutput)

	UpdateFieldLevelEncryptionProfile(*cloudfront.UpdateFieldLevelEncryptionProfileInput) (*cloudfront.UpdateFieldLevelEncryptionProfileOutput, error)
	UpdateFieldLevelEncryptionProfileWithContext(aws.Context, *cloudfront.UpdateFieldLevelEncryptionProfileInput, ...request.Option) (*cloudfront.UpdateFieldLevelEncryptionProfileOutput, error)
	UpdateFieldLevelEncryptionProfileRequest(*cloudfront.UpdateFieldLevelEncryptionProfileInput) (*request.Request, *cloudfront.UpdateFieldLevelEncryptionProfileOutput)

	UpdateFunction(*cloudfront.UpdateFunctionInput) (*cloudfront.UpdateFunctionOutput, error)
	UpdateFunctionWithContext(aws.Context, *cloudfront.UpdateFunctionInput, ...request.Option) (*cloudfront.UpdateFunctionOutput, error)
	UpdateFunctionRequest(*cloudfront.UpdateFunctionInput) (*request.Request, *cloudfront.UpdateFunctionOutput)

	UpdateKeyGroup(*cloudfront.UpdateKeyGroupInput) (*cloudfront.UpdateKeyGroupOutput, error)
	UpdateKeyGroupWithContext(aws.Context, *cloudfront.UpdateKeyGroupInput, ...request.Option) (*cloudfront.UpdateKeyGroupOutput, error)
	UpdateKeyGroupRequest(*cloudfront.UpdateKeyGroupInput) (*request.Request, *cloudfront.UpdateKeyGroupOutput)

	UpdateOriginRequestPolicy(*cloudfront.UpdateOriginRequestPolicyInput) (*cloudfront.UpdateOriginRequestPolicyOutput, error)
	UpdateOriginRequestPolicyWithContext(aws.Context, *cloudfront.UpdateOriginRequestPolicyInput, ...request.Option) (*cloudfront.UpdateOriginRequestPolicyOutput, error)
	UpdateOriginRequestPolicyRequest(*cloudfront.UpdateOriginRequestPolicyInput) (*request.Request, *cloudfront.UpdateOriginRequestPolicyOutput)

	UpdatePublicKey(*cloudfront.UpdatePublicKeyInput) (*cloudfront.UpdatePublicKeyOutput, error)
	UpdatePublicKeyWithContext(aws.Context, *cloudfront.UpdatePublicKeyInput, ...request.Option) (*cloudfront.UpdatePublicKeyOutput, error)
	UpdatePublicKeyRequest(*cloudfront.UpdatePublicKeyInput) (*request.Request, *cloudfront.UpdatePublicKeyOutput)

	UpdateRealtimeLogConfig(*cloudfront.UpdateRealtimeLogConfigInput) (*cloudfront.UpdateRealtimeLogConfigOutput, error)
	UpdateRealtimeLogConfigWithContext(aws.Context, *cloudfront.UpdateRealtimeLogConfigInput, ...request.Option) (*cloudfront.UpdateRealtimeLogConfigOutput, error)
	UpdateRealtimeLogConfigRequest(*cloudfront.UpdateRealtimeLogConfigInput) (*request.Request, *cloudfront.UpdateRealtimeLogConfigOutput)

	UpdateStreamingDistribution(*cloudfront.UpdateStreamingDistributionInput) (*cloudfront.UpdateStreamingDistributionOutput, error)
	UpdateStreamingDistributionWithContext(aws.Context, *cloudfront.UpdateStreamingDistributionInput, ...request.Option) (*cloudfront.UpdateStreamingDistributionOutput, error)
	UpdateStreamingDistributionRequest(*cloudfront.UpdateStreamingDistributionInput) (*request.Request, *cloudfront.UpdateStreamingDistributionOutput)

	WaitUntilDistributionDeployed(*cloudfront.GetDistributionInput) error
	WaitUntilDistributionDeployedWithContext(aws.Context, *cloudfront.GetDistributionInput, ...request.WaiterOption) error

	WaitUntilInvalidationCompleted(*cloudfront.GetInvalidationInput) error
	WaitUntilInvalidationCompletedWithContext(aws.Context, *cloudfront.GetInvalidationInput, ...request.WaiterOption) error

	WaitUntilStreamingDistributionDeployed(*cloudfront.GetStreamingDistributionInput) error
	WaitUntilStreamingDistributionDeployedWithContext(aws.Context, *cloudfront.GetStreamingDistributionInput, ...request.WaiterOption) error
}

var _ CloudFrontAPI = (*cloudfront.CloudFront)(nil)
