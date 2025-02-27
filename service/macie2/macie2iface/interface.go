// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package macie2iface provides an interface to enable mocking the Amazon Macie 2 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package macie2iface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/macie2"
)

// Macie2API provides an interface to enable mocking the
// macie2.Macie2 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Macie 2.
//    func myFunc(svc macie2iface.Macie2API) bool {
//        // Make svc.AcceptInvitation request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := macie2.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMacie2Client struct {
//        macie2iface.Macie2API
//    }
//    func (m *mockMacie2Client) AcceptInvitation(input *macie2.AcceptInvitationInput) (*macie2.AcceptInvitationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMacie2Client{}
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
type Macie2API interface {
	AcceptInvitation(*macie2.AcceptInvitationInput) (*macie2.AcceptInvitationOutput, error)
	AcceptInvitationWithContext(aws.Context, *macie2.AcceptInvitationInput, ...request.Option) (*macie2.AcceptInvitationOutput, error)
	AcceptInvitationRequest(*macie2.AcceptInvitationInput) (*request.Request, *macie2.AcceptInvitationOutput)

	BatchGetCustomDataIdentifiers(*macie2.BatchGetCustomDataIdentifiersInput) (*macie2.BatchGetCustomDataIdentifiersOutput, error)
	BatchGetCustomDataIdentifiersWithContext(aws.Context, *macie2.BatchGetCustomDataIdentifiersInput, ...request.Option) (*macie2.BatchGetCustomDataIdentifiersOutput, error)
	BatchGetCustomDataIdentifiersRequest(*macie2.BatchGetCustomDataIdentifiersInput) (*request.Request, *macie2.BatchGetCustomDataIdentifiersOutput)

	CreateClassificationJob(*macie2.CreateClassificationJobInput) (*macie2.CreateClassificationJobOutput, error)
	CreateClassificationJobWithContext(aws.Context, *macie2.CreateClassificationJobInput, ...request.Option) (*macie2.CreateClassificationJobOutput, error)
	CreateClassificationJobRequest(*macie2.CreateClassificationJobInput) (*request.Request, *macie2.CreateClassificationJobOutput)

	CreateCustomDataIdentifier(*macie2.CreateCustomDataIdentifierInput) (*macie2.CreateCustomDataIdentifierOutput, error)
	CreateCustomDataIdentifierWithContext(aws.Context, *macie2.CreateCustomDataIdentifierInput, ...request.Option) (*macie2.CreateCustomDataIdentifierOutput, error)
	CreateCustomDataIdentifierRequest(*macie2.CreateCustomDataIdentifierInput) (*request.Request, *macie2.CreateCustomDataIdentifierOutput)

	CreateFindingsFilter(*macie2.CreateFindingsFilterInput) (*macie2.CreateFindingsFilterOutput, error)
	CreateFindingsFilterWithContext(aws.Context, *macie2.CreateFindingsFilterInput, ...request.Option) (*macie2.CreateFindingsFilterOutput, error)
	CreateFindingsFilterRequest(*macie2.CreateFindingsFilterInput) (*request.Request, *macie2.CreateFindingsFilterOutput)

	CreateInvitations(*macie2.CreateInvitationsInput) (*macie2.CreateInvitationsOutput, error)
	CreateInvitationsWithContext(aws.Context, *macie2.CreateInvitationsInput, ...request.Option) (*macie2.CreateInvitationsOutput, error)
	CreateInvitationsRequest(*macie2.CreateInvitationsInput) (*request.Request, *macie2.CreateInvitationsOutput)

	CreateMember(*macie2.CreateMemberInput) (*macie2.CreateMemberOutput, error)
	CreateMemberWithContext(aws.Context, *macie2.CreateMemberInput, ...request.Option) (*macie2.CreateMemberOutput, error)
	CreateMemberRequest(*macie2.CreateMemberInput) (*request.Request, *macie2.CreateMemberOutput)

	CreateSampleFindings(*macie2.CreateSampleFindingsInput) (*macie2.CreateSampleFindingsOutput, error)
	CreateSampleFindingsWithContext(aws.Context, *macie2.CreateSampleFindingsInput, ...request.Option) (*macie2.CreateSampleFindingsOutput, error)
	CreateSampleFindingsRequest(*macie2.CreateSampleFindingsInput) (*request.Request, *macie2.CreateSampleFindingsOutput)

	DeclineInvitations(*macie2.DeclineInvitationsInput) (*macie2.DeclineInvitationsOutput, error)
	DeclineInvitationsWithContext(aws.Context, *macie2.DeclineInvitationsInput, ...request.Option) (*macie2.DeclineInvitationsOutput, error)
	DeclineInvitationsRequest(*macie2.DeclineInvitationsInput) (*request.Request, *macie2.DeclineInvitationsOutput)

	DeleteCustomDataIdentifier(*macie2.DeleteCustomDataIdentifierInput) (*macie2.DeleteCustomDataIdentifierOutput, error)
	DeleteCustomDataIdentifierWithContext(aws.Context, *macie2.DeleteCustomDataIdentifierInput, ...request.Option) (*macie2.DeleteCustomDataIdentifierOutput, error)
	DeleteCustomDataIdentifierRequest(*macie2.DeleteCustomDataIdentifierInput) (*request.Request, *macie2.DeleteCustomDataIdentifierOutput)

	DeleteFindingsFilter(*macie2.DeleteFindingsFilterInput) (*macie2.DeleteFindingsFilterOutput, error)
	DeleteFindingsFilterWithContext(aws.Context, *macie2.DeleteFindingsFilterInput, ...request.Option) (*macie2.DeleteFindingsFilterOutput, error)
	DeleteFindingsFilterRequest(*macie2.DeleteFindingsFilterInput) (*request.Request, *macie2.DeleteFindingsFilterOutput)

	DeleteInvitations(*macie2.DeleteInvitationsInput) (*macie2.DeleteInvitationsOutput, error)
	DeleteInvitationsWithContext(aws.Context, *macie2.DeleteInvitationsInput, ...request.Option) (*macie2.DeleteInvitationsOutput, error)
	DeleteInvitationsRequest(*macie2.DeleteInvitationsInput) (*request.Request, *macie2.DeleteInvitationsOutput)

	DeleteMember(*macie2.DeleteMemberInput) (*macie2.DeleteMemberOutput, error)
	DeleteMemberWithContext(aws.Context, *macie2.DeleteMemberInput, ...request.Option) (*macie2.DeleteMemberOutput, error)
	DeleteMemberRequest(*macie2.DeleteMemberInput) (*request.Request, *macie2.DeleteMemberOutput)

	DescribeBuckets(*macie2.DescribeBucketsInput) (*macie2.DescribeBucketsOutput, error)
	DescribeBucketsWithContext(aws.Context, *macie2.DescribeBucketsInput, ...request.Option) (*macie2.DescribeBucketsOutput, error)
	DescribeBucketsRequest(*macie2.DescribeBucketsInput) (*request.Request, *macie2.DescribeBucketsOutput)

	DescribeBucketsPages(*macie2.DescribeBucketsInput, func(*macie2.DescribeBucketsOutput, bool) bool) error
	DescribeBucketsPagesWithContext(aws.Context, *macie2.DescribeBucketsInput, func(*macie2.DescribeBucketsOutput, bool) bool, ...request.Option) error

	DescribeClassificationJob(*macie2.DescribeClassificationJobInput) (*macie2.DescribeClassificationJobOutput, error)
	DescribeClassificationJobWithContext(aws.Context, *macie2.DescribeClassificationJobInput, ...request.Option) (*macie2.DescribeClassificationJobOutput, error)
	DescribeClassificationJobRequest(*macie2.DescribeClassificationJobInput) (*request.Request, *macie2.DescribeClassificationJobOutput)

	DescribeOrganizationConfiguration(*macie2.DescribeOrganizationConfigurationInput) (*macie2.DescribeOrganizationConfigurationOutput, error)
	DescribeOrganizationConfigurationWithContext(aws.Context, *macie2.DescribeOrganizationConfigurationInput, ...request.Option) (*macie2.DescribeOrganizationConfigurationOutput, error)
	DescribeOrganizationConfigurationRequest(*macie2.DescribeOrganizationConfigurationInput) (*request.Request, *macie2.DescribeOrganizationConfigurationOutput)

	DisableMacie(*macie2.DisableMacieInput) (*macie2.DisableMacieOutput, error)
	DisableMacieWithContext(aws.Context, *macie2.DisableMacieInput, ...request.Option) (*macie2.DisableMacieOutput, error)
	DisableMacieRequest(*macie2.DisableMacieInput) (*request.Request, *macie2.DisableMacieOutput)

	DisableOrganizationAdminAccount(*macie2.DisableOrganizationAdminAccountInput) (*macie2.DisableOrganizationAdminAccountOutput, error)
	DisableOrganizationAdminAccountWithContext(aws.Context, *macie2.DisableOrganizationAdminAccountInput, ...request.Option) (*macie2.DisableOrganizationAdminAccountOutput, error)
	DisableOrganizationAdminAccountRequest(*macie2.DisableOrganizationAdminAccountInput) (*request.Request, *macie2.DisableOrganizationAdminAccountOutput)

	DisassociateFromAdministratorAccount(*macie2.DisassociateFromAdministratorAccountInput) (*macie2.DisassociateFromAdministratorAccountOutput, error)
	DisassociateFromAdministratorAccountWithContext(aws.Context, *macie2.DisassociateFromAdministratorAccountInput, ...request.Option) (*macie2.DisassociateFromAdministratorAccountOutput, error)
	DisassociateFromAdministratorAccountRequest(*macie2.DisassociateFromAdministratorAccountInput) (*request.Request, *macie2.DisassociateFromAdministratorAccountOutput)

	DisassociateFromMasterAccount(*macie2.DisassociateFromMasterAccountInput) (*macie2.DisassociateFromMasterAccountOutput, error)
	DisassociateFromMasterAccountWithContext(aws.Context, *macie2.DisassociateFromMasterAccountInput, ...request.Option) (*macie2.DisassociateFromMasterAccountOutput, error)
	DisassociateFromMasterAccountRequest(*macie2.DisassociateFromMasterAccountInput) (*request.Request, *macie2.DisassociateFromMasterAccountOutput)

	DisassociateMember(*macie2.DisassociateMemberInput) (*macie2.DisassociateMemberOutput, error)
	DisassociateMemberWithContext(aws.Context, *macie2.DisassociateMemberInput, ...request.Option) (*macie2.DisassociateMemberOutput, error)
	DisassociateMemberRequest(*macie2.DisassociateMemberInput) (*request.Request, *macie2.DisassociateMemberOutput)

	EnableMacie(*macie2.EnableMacieInput) (*macie2.EnableMacieOutput, error)
	EnableMacieWithContext(aws.Context, *macie2.EnableMacieInput, ...request.Option) (*macie2.EnableMacieOutput, error)
	EnableMacieRequest(*macie2.EnableMacieInput) (*request.Request, *macie2.EnableMacieOutput)

	EnableOrganizationAdminAccount(*macie2.EnableOrganizationAdminAccountInput) (*macie2.EnableOrganizationAdminAccountOutput, error)
	EnableOrganizationAdminAccountWithContext(aws.Context, *macie2.EnableOrganizationAdminAccountInput, ...request.Option) (*macie2.EnableOrganizationAdminAccountOutput, error)
	EnableOrganizationAdminAccountRequest(*macie2.EnableOrganizationAdminAccountInput) (*request.Request, *macie2.EnableOrganizationAdminAccountOutput)

	GetAdministratorAccount(*macie2.GetAdministratorAccountInput) (*macie2.GetAdministratorAccountOutput, error)
	GetAdministratorAccountWithContext(aws.Context, *macie2.GetAdministratorAccountInput, ...request.Option) (*macie2.GetAdministratorAccountOutput, error)
	GetAdministratorAccountRequest(*macie2.GetAdministratorAccountInput) (*request.Request, *macie2.GetAdministratorAccountOutput)

	GetBucketStatistics(*macie2.GetBucketStatisticsInput) (*macie2.GetBucketStatisticsOutput, error)
	GetBucketStatisticsWithContext(aws.Context, *macie2.GetBucketStatisticsInput, ...request.Option) (*macie2.GetBucketStatisticsOutput, error)
	GetBucketStatisticsRequest(*macie2.GetBucketStatisticsInput) (*request.Request, *macie2.GetBucketStatisticsOutput)

	GetClassificationExportConfiguration(*macie2.GetClassificationExportConfigurationInput) (*macie2.GetClassificationExportConfigurationOutput, error)
	GetClassificationExportConfigurationWithContext(aws.Context, *macie2.GetClassificationExportConfigurationInput, ...request.Option) (*macie2.GetClassificationExportConfigurationOutput, error)
	GetClassificationExportConfigurationRequest(*macie2.GetClassificationExportConfigurationInput) (*request.Request, *macie2.GetClassificationExportConfigurationOutput)

	GetCustomDataIdentifier(*macie2.GetCustomDataIdentifierInput) (*macie2.GetCustomDataIdentifierOutput, error)
	GetCustomDataIdentifierWithContext(aws.Context, *macie2.GetCustomDataIdentifierInput, ...request.Option) (*macie2.GetCustomDataIdentifierOutput, error)
	GetCustomDataIdentifierRequest(*macie2.GetCustomDataIdentifierInput) (*request.Request, *macie2.GetCustomDataIdentifierOutput)

	GetFindingStatistics(*macie2.GetFindingStatisticsInput) (*macie2.GetFindingStatisticsOutput, error)
	GetFindingStatisticsWithContext(aws.Context, *macie2.GetFindingStatisticsInput, ...request.Option) (*macie2.GetFindingStatisticsOutput, error)
	GetFindingStatisticsRequest(*macie2.GetFindingStatisticsInput) (*request.Request, *macie2.GetFindingStatisticsOutput)

	GetFindings(*macie2.GetFindingsInput) (*macie2.GetFindingsOutput, error)
	GetFindingsWithContext(aws.Context, *macie2.GetFindingsInput, ...request.Option) (*macie2.GetFindingsOutput, error)
	GetFindingsRequest(*macie2.GetFindingsInput) (*request.Request, *macie2.GetFindingsOutput)

	GetFindingsFilter(*macie2.GetFindingsFilterInput) (*macie2.GetFindingsFilterOutput, error)
	GetFindingsFilterWithContext(aws.Context, *macie2.GetFindingsFilterInput, ...request.Option) (*macie2.GetFindingsFilterOutput, error)
	GetFindingsFilterRequest(*macie2.GetFindingsFilterInput) (*request.Request, *macie2.GetFindingsFilterOutput)

	GetFindingsPublicationConfiguration(*macie2.GetFindingsPublicationConfigurationInput) (*macie2.GetFindingsPublicationConfigurationOutput, error)
	GetFindingsPublicationConfigurationWithContext(aws.Context, *macie2.GetFindingsPublicationConfigurationInput, ...request.Option) (*macie2.GetFindingsPublicationConfigurationOutput, error)
	GetFindingsPublicationConfigurationRequest(*macie2.GetFindingsPublicationConfigurationInput) (*request.Request, *macie2.GetFindingsPublicationConfigurationOutput)

	GetInvitationsCount(*macie2.GetInvitationsCountInput) (*macie2.GetInvitationsCountOutput, error)
	GetInvitationsCountWithContext(aws.Context, *macie2.GetInvitationsCountInput, ...request.Option) (*macie2.GetInvitationsCountOutput, error)
	GetInvitationsCountRequest(*macie2.GetInvitationsCountInput) (*request.Request, *macie2.GetInvitationsCountOutput)

	GetMacieSession(*macie2.GetMacieSessionInput) (*macie2.GetMacieSessionOutput, error)
	GetMacieSessionWithContext(aws.Context, *macie2.GetMacieSessionInput, ...request.Option) (*macie2.GetMacieSessionOutput, error)
	GetMacieSessionRequest(*macie2.GetMacieSessionInput) (*request.Request, *macie2.GetMacieSessionOutput)

	GetMasterAccount(*macie2.GetMasterAccountInput) (*macie2.GetMasterAccountOutput, error)
	GetMasterAccountWithContext(aws.Context, *macie2.GetMasterAccountInput, ...request.Option) (*macie2.GetMasterAccountOutput, error)
	GetMasterAccountRequest(*macie2.GetMasterAccountInput) (*request.Request, *macie2.GetMasterAccountOutput)

	GetMember(*macie2.GetMemberInput) (*macie2.GetMemberOutput, error)
	GetMemberWithContext(aws.Context, *macie2.GetMemberInput, ...request.Option) (*macie2.GetMemberOutput, error)
	GetMemberRequest(*macie2.GetMemberInput) (*request.Request, *macie2.GetMemberOutput)

	GetUsageStatistics(*macie2.GetUsageStatisticsInput) (*macie2.GetUsageStatisticsOutput, error)
	GetUsageStatisticsWithContext(aws.Context, *macie2.GetUsageStatisticsInput, ...request.Option) (*macie2.GetUsageStatisticsOutput, error)
	GetUsageStatisticsRequest(*macie2.GetUsageStatisticsInput) (*request.Request, *macie2.GetUsageStatisticsOutput)

	GetUsageStatisticsPages(*macie2.GetUsageStatisticsInput, func(*macie2.GetUsageStatisticsOutput, bool) bool) error
	GetUsageStatisticsPagesWithContext(aws.Context, *macie2.GetUsageStatisticsInput, func(*macie2.GetUsageStatisticsOutput, bool) bool, ...request.Option) error

	GetUsageTotals(*macie2.GetUsageTotalsInput) (*macie2.GetUsageTotalsOutput, error)
	GetUsageTotalsWithContext(aws.Context, *macie2.GetUsageTotalsInput, ...request.Option) (*macie2.GetUsageTotalsOutput, error)
	GetUsageTotalsRequest(*macie2.GetUsageTotalsInput) (*request.Request, *macie2.GetUsageTotalsOutput)

	ListClassificationJobs(*macie2.ListClassificationJobsInput) (*macie2.ListClassificationJobsOutput, error)
	ListClassificationJobsWithContext(aws.Context, *macie2.ListClassificationJobsInput, ...request.Option) (*macie2.ListClassificationJobsOutput, error)
	ListClassificationJobsRequest(*macie2.ListClassificationJobsInput) (*request.Request, *macie2.ListClassificationJobsOutput)

	ListClassificationJobsPages(*macie2.ListClassificationJobsInput, func(*macie2.ListClassificationJobsOutput, bool) bool) error
	ListClassificationJobsPagesWithContext(aws.Context, *macie2.ListClassificationJobsInput, func(*macie2.ListClassificationJobsOutput, bool) bool, ...request.Option) error

	ListCustomDataIdentifiers(*macie2.ListCustomDataIdentifiersInput) (*macie2.ListCustomDataIdentifiersOutput, error)
	ListCustomDataIdentifiersWithContext(aws.Context, *macie2.ListCustomDataIdentifiersInput, ...request.Option) (*macie2.ListCustomDataIdentifiersOutput, error)
	ListCustomDataIdentifiersRequest(*macie2.ListCustomDataIdentifiersInput) (*request.Request, *macie2.ListCustomDataIdentifiersOutput)

	ListCustomDataIdentifiersPages(*macie2.ListCustomDataIdentifiersInput, func(*macie2.ListCustomDataIdentifiersOutput, bool) bool) error
	ListCustomDataIdentifiersPagesWithContext(aws.Context, *macie2.ListCustomDataIdentifiersInput, func(*macie2.ListCustomDataIdentifiersOutput, bool) bool, ...request.Option) error

	ListFindings(*macie2.ListFindingsInput) (*macie2.ListFindingsOutput, error)
	ListFindingsWithContext(aws.Context, *macie2.ListFindingsInput, ...request.Option) (*macie2.ListFindingsOutput, error)
	ListFindingsRequest(*macie2.ListFindingsInput) (*request.Request, *macie2.ListFindingsOutput)

	ListFindingsPages(*macie2.ListFindingsInput, func(*macie2.ListFindingsOutput, bool) bool) error
	ListFindingsPagesWithContext(aws.Context, *macie2.ListFindingsInput, func(*macie2.ListFindingsOutput, bool) bool, ...request.Option) error

	ListFindingsFilters(*macie2.ListFindingsFiltersInput) (*macie2.ListFindingsFiltersOutput, error)
	ListFindingsFiltersWithContext(aws.Context, *macie2.ListFindingsFiltersInput, ...request.Option) (*macie2.ListFindingsFiltersOutput, error)
	ListFindingsFiltersRequest(*macie2.ListFindingsFiltersInput) (*request.Request, *macie2.ListFindingsFiltersOutput)

	ListFindingsFiltersPages(*macie2.ListFindingsFiltersInput, func(*macie2.ListFindingsFiltersOutput, bool) bool) error
	ListFindingsFiltersPagesWithContext(aws.Context, *macie2.ListFindingsFiltersInput, func(*macie2.ListFindingsFiltersOutput, bool) bool, ...request.Option) error

	ListInvitations(*macie2.ListInvitationsInput) (*macie2.ListInvitationsOutput, error)
	ListInvitationsWithContext(aws.Context, *macie2.ListInvitationsInput, ...request.Option) (*macie2.ListInvitationsOutput, error)
	ListInvitationsRequest(*macie2.ListInvitationsInput) (*request.Request, *macie2.ListInvitationsOutput)

	ListInvitationsPages(*macie2.ListInvitationsInput, func(*macie2.ListInvitationsOutput, bool) bool) error
	ListInvitationsPagesWithContext(aws.Context, *macie2.ListInvitationsInput, func(*macie2.ListInvitationsOutput, bool) bool, ...request.Option) error

	ListManagedDataIdentifiers(*macie2.ListManagedDataIdentifiersInput) (*macie2.ListManagedDataIdentifiersOutput, error)
	ListManagedDataIdentifiersWithContext(aws.Context, *macie2.ListManagedDataIdentifiersInput, ...request.Option) (*macie2.ListManagedDataIdentifiersOutput, error)
	ListManagedDataIdentifiersRequest(*macie2.ListManagedDataIdentifiersInput) (*request.Request, *macie2.ListManagedDataIdentifiersOutput)

	ListMembers(*macie2.ListMembersInput) (*macie2.ListMembersOutput, error)
	ListMembersWithContext(aws.Context, *macie2.ListMembersInput, ...request.Option) (*macie2.ListMembersOutput, error)
	ListMembersRequest(*macie2.ListMembersInput) (*request.Request, *macie2.ListMembersOutput)

	ListMembersPages(*macie2.ListMembersInput, func(*macie2.ListMembersOutput, bool) bool) error
	ListMembersPagesWithContext(aws.Context, *macie2.ListMembersInput, func(*macie2.ListMembersOutput, bool) bool, ...request.Option) error

	ListOrganizationAdminAccounts(*macie2.ListOrganizationAdminAccountsInput) (*macie2.ListOrganizationAdminAccountsOutput, error)
	ListOrganizationAdminAccountsWithContext(aws.Context, *macie2.ListOrganizationAdminAccountsInput, ...request.Option) (*macie2.ListOrganizationAdminAccountsOutput, error)
	ListOrganizationAdminAccountsRequest(*macie2.ListOrganizationAdminAccountsInput) (*request.Request, *macie2.ListOrganizationAdminAccountsOutput)

	ListOrganizationAdminAccountsPages(*macie2.ListOrganizationAdminAccountsInput, func(*macie2.ListOrganizationAdminAccountsOutput, bool) bool) error
	ListOrganizationAdminAccountsPagesWithContext(aws.Context, *macie2.ListOrganizationAdminAccountsInput, func(*macie2.ListOrganizationAdminAccountsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*macie2.ListTagsForResourceInput) (*macie2.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *macie2.ListTagsForResourceInput, ...request.Option) (*macie2.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*macie2.ListTagsForResourceInput) (*request.Request, *macie2.ListTagsForResourceOutput)

	PutClassificationExportConfiguration(*macie2.PutClassificationExportConfigurationInput) (*macie2.PutClassificationExportConfigurationOutput, error)
	PutClassificationExportConfigurationWithContext(aws.Context, *macie2.PutClassificationExportConfigurationInput, ...request.Option) (*macie2.PutClassificationExportConfigurationOutput, error)
	PutClassificationExportConfigurationRequest(*macie2.PutClassificationExportConfigurationInput) (*request.Request, *macie2.PutClassificationExportConfigurationOutput)

	PutFindingsPublicationConfiguration(*macie2.PutFindingsPublicationConfigurationInput) (*macie2.PutFindingsPublicationConfigurationOutput, error)
	PutFindingsPublicationConfigurationWithContext(aws.Context, *macie2.PutFindingsPublicationConfigurationInput, ...request.Option) (*macie2.PutFindingsPublicationConfigurationOutput, error)
	PutFindingsPublicationConfigurationRequest(*macie2.PutFindingsPublicationConfigurationInput) (*request.Request, *macie2.PutFindingsPublicationConfigurationOutput)

	SearchResources(*macie2.SearchResourcesInput) (*macie2.SearchResourcesOutput, error)
	SearchResourcesWithContext(aws.Context, *macie2.SearchResourcesInput, ...request.Option) (*macie2.SearchResourcesOutput, error)
	SearchResourcesRequest(*macie2.SearchResourcesInput) (*request.Request, *macie2.SearchResourcesOutput)

	SearchResourcesPages(*macie2.SearchResourcesInput, func(*macie2.SearchResourcesOutput, bool) bool) error
	SearchResourcesPagesWithContext(aws.Context, *macie2.SearchResourcesInput, func(*macie2.SearchResourcesOutput, bool) bool, ...request.Option) error

	TagResource(*macie2.TagResourceInput) (*macie2.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *macie2.TagResourceInput, ...request.Option) (*macie2.TagResourceOutput, error)
	TagResourceRequest(*macie2.TagResourceInput) (*request.Request, *macie2.TagResourceOutput)

	TestCustomDataIdentifier(*macie2.TestCustomDataIdentifierInput) (*macie2.TestCustomDataIdentifierOutput, error)
	TestCustomDataIdentifierWithContext(aws.Context, *macie2.TestCustomDataIdentifierInput, ...request.Option) (*macie2.TestCustomDataIdentifierOutput, error)
	TestCustomDataIdentifierRequest(*macie2.TestCustomDataIdentifierInput) (*request.Request, *macie2.TestCustomDataIdentifierOutput)

	UntagResource(*macie2.UntagResourceInput) (*macie2.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *macie2.UntagResourceInput, ...request.Option) (*macie2.UntagResourceOutput, error)
	UntagResourceRequest(*macie2.UntagResourceInput) (*request.Request, *macie2.UntagResourceOutput)

	UpdateClassificationJob(*macie2.UpdateClassificationJobInput) (*macie2.UpdateClassificationJobOutput, error)
	UpdateClassificationJobWithContext(aws.Context, *macie2.UpdateClassificationJobInput, ...request.Option) (*macie2.UpdateClassificationJobOutput, error)
	UpdateClassificationJobRequest(*macie2.UpdateClassificationJobInput) (*request.Request, *macie2.UpdateClassificationJobOutput)

	UpdateFindingsFilter(*macie2.UpdateFindingsFilterInput) (*macie2.UpdateFindingsFilterOutput, error)
	UpdateFindingsFilterWithContext(aws.Context, *macie2.UpdateFindingsFilterInput, ...request.Option) (*macie2.UpdateFindingsFilterOutput, error)
	UpdateFindingsFilterRequest(*macie2.UpdateFindingsFilterInput) (*request.Request, *macie2.UpdateFindingsFilterOutput)

	UpdateMacieSession(*macie2.UpdateMacieSessionInput) (*macie2.UpdateMacieSessionOutput, error)
	UpdateMacieSessionWithContext(aws.Context, *macie2.UpdateMacieSessionInput, ...request.Option) (*macie2.UpdateMacieSessionOutput, error)
	UpdateMacieSessionRequest(*macie2.UpdateMacieSessionInput) (*request.Request, *macie2.UpdateMacieSessionOutput)

	UpdateMemberSession(*macie2.UpdateMemberSessionInput) (*macie2.UpdateMemberSessionOutput, error)
	UpdateMemberSessionWithContext(aws.Context, *macie2.UpdateMemberSessionInput, ...request.Option) (*macie2.UpdateMemberSessionOutput, error)
	UpdateMemberSessionRequest(*macie2.UpdateMemberSessionInput) (*request.Request, *macie2.UpdateMemberSessionOutput)

	UpdateOrganizationConfiguration(*macie2.UpdateOrganizationConfigurationInput) (*macie2.UpdateOrganizationConfigurationOutput, error)
	UpdateOrganizationConfigurationWithContext(aws.Context, *macie2.UpdateOrganizationConfigurationInput, ...request.Option) (*macie2.UpdateOrganizationConfigurationOutput, error)
	UpdateOrganizationConfigurationRequest(*macie2.UpdateOrganizationConfigurationInput) (*request.Request, *macie2.UpdateOrganizationConfigurationOutput)
}

var _ Macie2API = (*macie2.Macie2)(nil)
