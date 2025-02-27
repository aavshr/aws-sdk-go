// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package route53domainsiface provides an interface to enable mocking the Amazon Route 53 Domains service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package route53domainsiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/route53domains"
)

// Route53DomainsAPI provides an interface to enable mocking the
// route53domains.Route53Domains service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Route 53 Domains.
//    func myFunc(svc route53domainsiface.Route53DomainsAPI) bool {
//        // Make svc.AcceptDomainTransferFromAnotherAwsAccount request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := route53domains.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRoute53DomainsClient struct {
//        route53domainsiface.Route53DomainsAPI
//    }
//    func (m *mockRoute53DomainsClient) AcceptDomainTransferFromAnotherAwsAccount(input *route53domains.AcceptDomainTransferFromAnotherAwsAccountInput) (*route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRoute53DomainsClient{}
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
type Route53DomainsAPI interface {
	AcceptDomainTransferFromAnotherAwsAccount(*route53domains.AcceptDomainTransferFromAnotherAwsAccountInput) (*route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput, error)
	AcceptDomainTransferFromAnotherAwsAccountWithContext(aws.Context, *route53domains.AcceptDomainTransferFromAnotherAwsAccountInput, ...request.Option) (*route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput, error)
	AcceptDomainTransferFromAnotherAwsAccountRequest(*route53domains.AcceptDomainTransferFromAnotherAwsAccountInput) (*request.Request, *route53domains.AcceptDomainTransferFromAnotherAwsAccountOutput)

	CancelDomainTransferToAnotherAwsAccount(*route53domains.CancelDomainTransferToAnotherAwsAccountInput) (*route53domains.CancelDomainTransferToAnotherAwsAccountOutput, error)
	CancelDomainTransferToAnotherAwsAccountWithContext(aws.Context, *route53domains.CancelDomainTransferToAnotherAwsAccountInput, ...request.Option) (*route53domains.CancelDomainTransferToAnotherAwsAccountOutput, error)
	CancelDomainTransferToAnotherAwsAccountRequest(*route53domains.CancelDomainTransferToAnotherAwsAccountInput) (*request.Request, *route53domains.CancelDomainTransferToAnotherAwsAccountOutput)

	CheckDomainAvailability(*route53domains.CheckDomainAvailabilityInput) (*route53domains.CheckDomainAvailabilityOutput, error)
	CheckDomainAvailabilityWithContext(aws.Context, *route53domains.CheckDomainAvailabilityInput, ...request.Option) (*route53domains.CheckDomainAvailabilityOutput, error)
	CheckDomainAvailabilityRequest(*route53domains.CheckDomainAvailabilityInput) (*request.Request, *route53domains.CheckDomainAvailabilityOutput)

	CheckDomainTransferability(*route53domains.CheckDomainTransferabilityInput) (*route53domains.CheckDomainTransferabilityOutput, error)
	CheckDomainTransferabilityWithContext(aws.Context, *route53domains.CheckDomainTransferabilityInput, ...request.Option) (*route53domains.CheckDomainTransferabilityOutput, error)
	CheckDomainTransferabilityRequest(*route53domains.CheckDomainTransferabilityInput) (*request.Request, *route53domains.CheckDomainTransferabilityOutput)

	DeleteTagsForDomain(*route53domains.DeleteTagsForDomainInput) (*route53domains.DeleteTagsForDomainOutput, error)
	DeleteTagsForDomainWithContext(aws.Context, *route53domains.DeleteTagsForDomainInput, ...request.Option) (*route53domains.DeleteTagsForDomainOutput, error)
	DeleteTagsForDomainRequest(*route53domains.DeleteTagsForDomainInput) (*request.Request, *route53domains.DeleteTagsForDomainOutput)

	DisableDomainAutoRenew(*route53domains.DisableDomainAutoRenewInput) (*route53domains.DisableDomainAutoRenewOutput, error)
	DisableDomainAutoRenewWithContext(aws.Context, *route53domains.DisableDomainAutoRenewInput, ...request.Option) (*route53domains.DisableDomainAutoRenewOutput, error)
	DisableDomainAutoRenewRequest(*route53domains.DisableDomainAutoRenewInput) (*request.Request, *route53domains.DisableDomainAutoRenewOutput)

	DisableDomainTransferLock(*route53domains.DisableDomainTransferLockInput) (*route53domains.DisableDomainTransferLockOutput, error)
	DisableDomainTransferLockWithContext(aws.Context, *route53domains.DisableDomainTransferLockInput, ...request.Option) (*route53domains.DisableDomainTransferLockOutput, error)
	DisableDomainTransferLockRequest(*route53domains.DisableDomainTransferLockInput) (*request.Request, *route53domains.DisableDomainTransferLockOutput)

	EnableDomainAutoRenew(*route53domains.EnableDomainAutoRenewInput) (*route53domains.EnableDomainAutoRenewOutput, error)
	EnableDomainAutoRenewWithContext(aws.Context, *route53domains.EnableDomainAutoRenewInput, ...request.Option) (*route53domains.EnableDomainAutoRenewOutput, error)
	EnableDomainAutoRenewRequest(*route53domains.EnableDomainAutoRenewInput) (*request.Request, *route53domains.EnableDomainAutoRenewOutput)

	EnableDomainTransferLock(*route53domains.EnableDomainTransferLockInput) (*route53domains.EnableDomainTransferLockOutput, error)
	EnableDomainTransferLockWithContext(aws.Context, *route53domains.EnableDomainTransferLockInput, ...request.Option) (*route53domains.EnableDomainTransferLockOutput, error)
	EnableDomainTransferLockRequest(*route53domains.EnableDomainTransferLockInput) (*request.Request, *route53domains.EnableDomainTransferLockOutput)

	GetContactReachabilityStatus(*route53domains.GetContactReachabilityStatusInput) (*route53domains.GetContactReachabilityStatusOutput, error)
	GetContactReachabilityStatusWithContext(aws.Context, *route53domains.GetContactReachabilityStatusInput, ...request.Option) (*route53domains.GetContactReachabilityStatusOutput, error)
	GetContactReachabilityStatusRequest(*route53domains.GetContactReachabilityStatusInput) (*request.Request, *route53domains.GetContactReachabilityStatusOutput)

	GetDomainDetail(*route53domains.GetDomainDetailInput) (*route53domains.GetDomainDetailOutput, error)
	GetDomainDetailWithContext(aws.Context, *route53domains.GetDomainDetailInput, ...request.Option) (*route53domains.GetDomainDetailOutput, error)
	GetDomainDetailRequest(*route53domains.GetDomainDetailInput) (*request.Request, *route53domains.GetDomainDetailOutput)

	GetDomainSuggestions(*route53domains.GetDomainSuggestionsInput) (*route53domains.GetDomainSuggestionsOutput, error)
	GetDomainSuggestionsWithContext(aws.Context, *route53domains.GetDomainSuggestionsInput, ...request.Option) (*route53domains.GetDomainSuggestionsOutput, error)
	GetDomainSuggestionsRequest(*route53domains.GetDomainSuggestionsInput) (*request.Request, *route53domains.GetDomainSuggestionsOutput)

	GetOperationDetail(*route53domains.GetOperationDetailInput) (*route53domains.GetOperationDetailOutput, error)
	GetOperationDetailWithContext(aws.Context, *route53domains.GetOperationDetailInput, ...request.Option) (*route53domains.GetOperationDetailOutput, error)
	GetOperationDetailRequest(*route53domains.GetOperationDetailInput) (*request.Request, *route53domains.GetOperationDetailOutput)

	ListDomains(*route53domains.ListDomainsInput) (*route53domains.ListDomainsOutput, error)
	ListDomainsWithContext(aws.Context, *route53domains.ListDomainsInput, ...request.Option) (*route53domains.ListDomainsOutput, error)
	ListDomainsRequest(*route53domains.ListDomainsInput) (*request.Request, *route53domains.ListDomainsOutput)

	ListDomainsPages(*route53domains.ListDomainsInput, func(*route53domains.ListDomainsOutput, bool) bool) error
	ListDomainsPagesWithContext(aws.Context, *route53domains.ListDomainsInput, func(*route53domains.ListDomainsOutput, bool) bool, ...request.Option) error

	ListOperations(*route53domains.ListOperationsInput) (*route53domains.ListOperationsOutput, error)
	ListOperationsWithContext(aws.Context, *route53domains.ListOperationsInput, ...request.Option) (*route53domains.ListOperationsOutput, error)
	ListOperationsRequest(*route53domains.ListOperationsInput) (*request.Request, *route53domains.ListOperationsOutput)

	ListOperationsPages(*route53domains.ListOperationsInput, func(*route53domains.ListOperationsOutput, bool) bool) error
	ListOperationsPagesWithContext(aws.Context, *route53domains.ListOperationsInput, func(*route53domains.ListOperationsOutput, bool) bool, ...request.Option) error

	ListTagsForDomain(*route53domains.ListTagsForDomainInput) (*route53domains.ListTagsForDomainOutput, error)
	ListTagsForDomainWithContext(aws.Context, *route53domains.ListTagsForDomainInput, ...request.Option) (*route53domains.ListTagsForDomainOutput, error)
	ListTagsForDomainRequest(*route53domains.ListTagsForDomainInput) (*request.Request, *route53domains.ListTagsForDomainOutput)

	RegisterDomain(*route53domains.RegisterDomainInput) (*route53domains.RegisterDomainOutput, error)
	RegisterDomainWithContext(aws.Context, *route53domains.RegisterDomainInput, ...request.Option) (*route53domains.RegisterDomainOutput, error)
	RegisterDomainRequest(*route53domains.RegisterDomainInput) (*request.Request, *route53domains.RegisterDomainOutput)

	RejectDomainTransferFromAnotherAwsAccount(*route53domains.RejectDomainTransferFromAnotherAwsAccountInput) (*route53domains.RejectDomainTransferFromAnotherAwsAccountOutput, error)
	RejectDomainTransferFromAnotherAwsAccountWithContext(aws.Context, *route53domains.RejectDomainTransferFromAnotherAwsAccountInput, ...request.Option) (*route53domains.RejectDomainTransferFromAnotherAwsAccountOutput, error)
	RejectDomainTransferFromAnotherAwsAccountRequest(*route53domains.RejectDomainTransferFromAnotherAwsAccountInput) (*request.Request, *route53domains.RejectDomainTransferFromAnotherAwsAccountOutput)

	RenewDomain(*route53domains.RenewDomainInput) (*route53domains.RenewDomainOutput, error)
	RenewDomainWithContext(aws.Context, *route53domains.RenewDomainInput, ...request.Option) (*route53domains.RenewDomainOutput, error)
	RenewDomainRequest(*route53domains.RenewDomainInput) (*request.Request, *route53domains.RenewDomainOutput)

	ResendContactReachabilityEmail(*route53domains.ResendContactReachabilityEmailInput) (*route53domains.ResendContactReachabilityEmailOutput, error)
	ResendContactReachabilityEmailWithContext(aws.Context, *route53domains.ResendContactReachabilityEmailInput, ...request.Option) (*route53domains.ResendContactReachabilityEmailOutput, error)
	ResendContactReachabilityEmailRequest(*route53domains.ResendContactReachabilityEmailInput) (*request.Request, *route53domains.ResendContactReachabilityEmailOutput)

	RetrieveDomainAuthCode(*route53domains.RetrieveDomainAuthCodeInput) (*route53domains.RetrieveDomainAuthCodeOutput, error)
	RetrieveDomainAuthCodeWithContext(aws.Context, *route53domains.RetrieveDomainAuthCodeInput, ...request.Option) (*route53domains.RetrieveDomainAuthCodeOutput, error)
	RetrieveDomainAuthCodeRequest(*route53domains.RetrieveDomainAuthCodeInput) (*request.Request, *route53domains.RetrieveDomainAuthCodeOutput)

	TransferDomain(*route53domains.TransferDomainInput) (*route53domains.TransferDomainOutput, error)
	TransferDomainWithContext(aws.Context, *route53domains.TransferDomainInput, ...request.Option) (*route53domains.TransferDomainOutput, error)
	TransferDomainRequest(*route53domains.TransferDomainInput) (*request.Request, *route53domains.TransferDomainOutput)

	TransferDomainToAnotherAwsAccount(*route53domains.TransferDomainToAnotherAwsAccountInput) (*route53domains.TransferDomainToAnotherAwsAccountOutput, error)
	TransferDomainToAnotherAwsAccountWithContext(aws.Context, *route53domains.TransferDomainToAnotherAwsAccountInput, ...request.Option) (*route53domains.TransferDomainToAnotherAwsAccountOutput, error)
	TransferDomainToAnotherAwsAccountRequest(*route53domains.TransferDomainToAnotherAwsAccountInput) (*request.Request, *route53domains.TransferDomainToAnotherAwsAccountOutput)

	UpdateDomainContact(*route53domains.UpdateDomainContactInput) (*route53domains.UpdateDomainContactOutput, error)
	UpdateDomainContactWithContext(aws.Context, *route53domains.UpdateDomainContactInput, ...request.Option) (*route53domains.UpdateDomainContactOutput, error)
	UpdateDomainContactRequest(*route53domains.UpdateDomainContactInput) (*request.Request, *route53domains.UpdateDomainContactOutput)

	UpdateDomainContactPrivacy(*route53domains.UpdateDomainContactPrivacyInput) (*route53domains.UpdateDomainContactPrivacyOutput, error)
	UpdateDomainContactPrivacyWithContext(aws.Context, *route53domains.UpdateDomainContactPrivacyInput, ...request.Option) (*route53domains.UpdateDomainContactPrivacyOutput, error)
	UpdateDomainContactPrivacyRequest(*route53domains.UpdateDomainContactPrivacyInput) (*request.Request, *route53domains.UpdateDomainContactPrivacyOutput)

	UpdateDomainNameservers(*route53domains.UpdateDomainNameserversInput) (*route53domains.UpdateDomainNameserversOutput, error)
	UpdateDomainNameserversWithContext(aws.Context, *route53domains.UpdateDomainNameserversInput, ...request.Option) (*route53domains.UpdateDomainNameserversOutput, error)
	UpdateDomainNameserversRequest(*route53domains.UpdateDomainNameserversInput) (*request.Request, *route53domains.UpdateDomainNameserversOutput)

	UpdateTagsForDomain(*route53domains.UpdateTagsForDomainInput) (*route53domains.UpdateTagsForDomainOutput, error)
	UpdateTagsForDomainWithContext(aws.Context, *route53domains.UpdateTagsForDomainInput, ...request.Option) (*route53domains.UpdateTagsForDomainOutput, error)
	UpdateTagsForDomainRequest(*route53domains.UpdateTagsForDomainInput) (*request.Request, *route53domains.UpdateTagsForDomainOutput)

	ViewBilling(*route53domains.ViewBillingInput) (*route53domains.ViewBillingOutput, error)
	ViewBillingWithContext(aws.Context, *route53domains.ViewBillingInput, ...request.Option) (*route53domains.ViewBillingOutput, error)
	ViewBillingRequest(*route53domains.ViewBillingInput) (*request.Request, *route53domains.ViewBillingOutput)
}

var _ Route53DomainsAPI = (*route53domains.Route53Domains)(nil)
