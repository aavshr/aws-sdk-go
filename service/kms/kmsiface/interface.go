// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kmsiface provides an interface to enable mocking the AWS Key Management Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kmsiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/kms"
)

// KMSAPI provides an interface to enable mocking the
// kms.KMS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Key Management Service.
//    func myFunc(svc kmsiface.KMSAPI) bool {
//        // Make svc.CancelKeyDeletion request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kms.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKMSClient struct {
//        kmsiface.KMSAPI
//    }
//    func (m *mockKMSClient) CancelKeyDeletion(input *kms.CancelKeyDeletionInput) (*kms.CancelKeyDeletionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKMSClient{}
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
type KMSAPI interface {
	CancelKeyDeletion(*kms.CancelKeyDeletionInput) (*kms.CancelKeyDeletionOutput, error)
	CancelKeyDeletionWithContext(aws.Context, *kms.CancelKeyDeletionInput, ...request.Option) (*kms.CancelKeyDeletionOutput, error)
	CancelKeyDeletionRequest(*kms.CancelKeyDeletionInput) (*request.Request, *kms.CancelKeyDeletionOutput)

	ConnectCustomKeyStore(*kms.ConnectCustomKeyStoreInput) (*kms.ConnectCustomKeyStoreOutput, error)
	ConnectCustomKeyStoreWithContext(aws.Context, *kms.ConnectCustomKeyStoreInput, ...request.Option) (*kms.ConnectCustomKeyStoreOutput, error)
	ConnectCustomKeyStoreRequest(*kms.ConnectCustomKeyStoreInput) (*request.Request, *kms.ConnectCustomKeyStoreOutput)

	CreateAlias(*kms.CreateAliasInput) (*kms.CreateAliasOutput, error)
	CreateAliasWithContext(aws.Context, *kms.CreateAliasInput, ...request.Option) (*kms.CreateAliasOutput, error)
	CreateAliasRequest(*kms.CreateAliasInput) (*request.Request, *kms.CreateAliasOutput)

	CreateCustomKeyStore(*kms.CreateCustomKeyStoreInput) (*kms.CreateCustomKeyStoreOutput, error)
	CreateCustomKeyStoreWithContext(aws.Context, *kms.CreateCustomKeyStoreInput, ...request.Option) (*kms.CreateCustomKeyStoreOutput, error)
	CreateCustomKeyStoreRequest(*kms.CreateCustomKeyStoreInput) (*request.Request, *kms.CreateCustomKeyStoreOutput)

	CreateGrant(*kms.CreateGrantInput) (*kms.CreateGrantOutput, error)
	CreateGrantWithContext(aws.Context, *kms.CreateGrantInput, ...request.Option) (*kms.CreateGrantOutput, error)
	CreateGrantRequest(*kms.CreateGrantInput) (*request.Request, *kms.CreateGrantOutput)

	CreateKey(*kms.CreateKeyInput) (*kms.CreateKeyOutput, error)
	CreateKeyWithContext(aws.Context, *kms.CreateKeyInput, ...request.Option) (*kms.CreateKeyOutput, error)
	CreateKeyRequest(*kms.CreateKeyInput) (*request.Request, *kms.CreateKeyOutput)

	Decrypt(*kms.DecryptInput) (*kms.DecryptOutput, error)
	DecryptWithContext(aws.Context, *kms.DecryptInput, ...request.Option) (*kms.DecryptOutput, error)
	DecryptRequest(*kms.DecryptInput) (*request.Request, *kms.DecryptOutput)

	DeleteAlias(*kms.DeleteAliasInput) (*kms.DeleteAliasOutput, error)
	DeleteAliasWithContext(aws.Context, *kms.DeleteAliasInput, ...request.Option) (*kms.DeleteAliasOutput, error)
	DeleteAliasRequest(*kms.DeleteAliasInput) (*request.Request, *kms.DeleteAliasOutput)

	DeleteCustomKeyStore(*kms.DeleteCustomKeyStoreInput) (*kms.DeleteCustomKeyStoreOutput, error)
	DeleteCustomKeyStoreWithContext(aws.Context, *kms.DeleteCustomKeyStoreInput, ...request.Option) (*kms.DeleteCustomKeyStoreOutput, error)
	DeleteCustomKeyStoreRequest(*kms.DeleteCustomKeyStoreInput) (*request.Request, *kms.DeleteCustomKeyStoreOutput)

	DeleteImportedKeyMaterial(*kms.DeleteImportedKeyMaterialInput) (*kms.DeleteImportedKeyMaterialOutput, error)
	DeleteImportedKeyMaterialWithContext(aws.Context, *kms.DeleteImportedKeyMaterialInput, ...request.Option) (*kms.DeleteImportedKeyMaterialOutput, error)
	DeleteImportedKeyMaterialRequest(*kms.DeleteImportedKeyMaterialInput) (*request.Request, *kms.DeleteImportedKeyMaterialOutput)

	DescribeCustomKeyStores(*kms.DescribeCustomKeyStoresInput) (*kms.DescribeCustomKeyStoresOutput, error)
	DescribeCustomKeyStoresWithContext(aws.Context, *kms.DescribeCustomKeyStoresInput, ...request.Option) (*kms.DescribeCustomKeyStoresOutput, error)
	DescribeCustomKeyStoresRequest(*kms.DescribeCustomKeyStoresInput) (*request.Request, *kms.DescribeCustomKeyStoresOutput)

	DescribeKey(*kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error)
	DescribeKeyWithContext(aws.Context, *kms.DescribeKeyInput, ...request.Option) (*kms.DescribeKeyOutput, error)
	DescribeKeyRequest(*kms.DescribeKeyInput) (*request.Request, *kms.DescribeKeyOutput)

	DisableKey(*kms.DisableKeyInput) (*kms.DisableKeyOutput, error)
	DisableKeyWithContext(aws.Context, *kms.DisableKeyInput, ...request.Option) (*kms.DisableKeyOutput, error)
	DisableKeyRequest(*kms.DisableKeyInput) (*request.Request, *kms.DisableKeyOutput)

	DisableKeyRotation(*kms.DisableKeyRotationInput) (*kms.DisableKeyRotationOutput, error)
	DisableKeyRotationWithContext(aws.Context, *kms.DisableKeyRotationInput, ...request.Option) (*kms.DisableKeyRotationOutput, error)
	DisableKeyRotationRequest(*kms.DisableKeyRotationInput) (*request.Request, *kms.DisableKeyRotationOutput)

	DisconnectCustomKeyStore(*kms.DisconnectCustomKeyStoreInput) (*kms.DisconnectCustomKeyStoreOutput, error)
	DisconnectCustomKeyStoreWithContext(aws.Context, *kms.DisconnectCustomKeyStoreInput, ...request.Option) (*kms.DisconnectCustomKeyStoreOutput, error)
	DisconnectCustomKeyStoreRequest(*kms.DisconnectCustomKeyStoreInput) (*request.Request, *kms.DisconnectCustomKeyStoreOutput)

	EnableKey(*kms.EnableKeyInput) (*kms.EnableKeyOutput, error)
	EnableKeyWithContext(aws.Context, *kms.EnableKeyInput, ...request.Option) (*kms.EnableKeyOutput, error)
	EnableKeyRequest(*kms.EnableKeyInput) (*request.Request, *kms.EnableKeyOutput)

	EnableKeyRotation(*kms.EnableKeyRotationInput) (*kms.EnableKeyRotationOutput, error)
	EnableKeyRotationWithContext(aws.Context, *kms.EnableKeyRotationInput, ...request.Option) (*kms.EnableKeyRotationOutput, error)
	EnableKeyRotationRequest(*kms.EnableKeyRotationInput) (*request.Request, *kms.EnableKeyRotationOutput)

	Encrypt(*kms.EncryptInput) (*kms.EncryptOutput, error)
	EncryptWithContext(aws.Context, *kms.EncryptInput, ...request.Option) (*kms.EncryptOutput, error)
	EncryptRequest(*kms.EncryptInput) (*request.Request, *kms.EncryptOutput)

	GenerateDataKey(*kms.GenerateDataKeyInput) (*kms.GenerateDataKeyOutput, error)
	GenerateDataKeyWithContext(aws.Context, *kms.GenerateDataKeyInput, ...request.Option) (*kms.GenerateDataKeyOutput, error)
	GenerateDataKeyRequest(*kms.GenerateDataKeyInput) (*request.Request, *kms.GenerateDataKeyOutput)

	GenerateDataKeyPair(*kms.GenerateDataKeyPairInput) (*kms.GenerateDataKeyPairOutput, error)
	GenerateDataKeyPairWithContext(aws.Context, *kms.GenerateDataKeyPairInput, ...request.Option) (*kms.GenerateDataKeyPairOutput, error)
	GenerateDataKeyPairRequest(*kms.GenerateDataKeyPairInput) (*request.Request, *kms.GenerateDataKeyPairOutput)

	GenerateDataKeyPairWithoutPlaintext(*kms.GenerateDataKeyPairWithoutPlaintextInput) (*kms.GenerateDataKeyPairWithoutPlaintextOutput, error)
	GenerateDataKeyPairWithoutPlaintextWithContext(aws.Context, *kms.GenerateDataKeyPairWithoutPlaintextInput, ...request.Option) (*kms.GenerateDataKeyPairWithoutPlaintextOutput, error)
	GenerateDataKeyPairWithoutPlaintextRequest(*kms.GenerateDataKeyPairWithoutPlaintextInput) (*request.Request, *kms.GenerateDataKeyPairWithoutPlaintextOutput)

	GenerateDataKeyWithoutPlaintext(*kms.GenerateDataKeyWithoutPlaintextInput) (*kms.GenerateDataKeyWithoutPlaintextOutput, error)
	GenerateDataKeyWithoutPlaintextWithContext(aws.Context, *kms.GenerateDataKeyWithoutPlaintextInput, ...request.Option) (*kms.GenerateDataKeyWithoutPlaintextOutput, error)
	GenerateDataKeyWithoutPlaintextRequest(*kms.GenerateDataKeyWithoutPlaintextInput) (*request.Request, *kms.GenerateDataKeyWithoutPlaintextOutput)

	GenerateRandom(*kms.GenerateRandomInput) (*kms.GenerateRandomOutput, error)
	GenerateRandomWithContext(aws.Context, *kms.GenerateRandomInput, ...request.Option) (*kms.GenerateRandomOutput, error)
	GenerateRandomRequest(*kms.GenerateRandomInput) (*request.Request, *kms.GenerateRandomOutput)

	GetKeyPolicy(*kms.GetKeyPolicyInput) (*kms.GetKeyPolicyOutput, error)
	GetKeyPolicyWithContext(aws.Context, *kms.GetKeyPolicyInput, ...request.Option) (*kms.GetKeyPolicyOutput, error)
	GetKeyPolicyRequest(*kms.GetKeyPolicyInput) (*request.Request, *kms.GetKeyPolicyOutput)

	GetKeyRotationStatus(*kms.GetKeyRotationStatusInput) (*kms.GetKeyRotationStatusOutput, error)
	GetKeyRotationStatusWithContext(aws.Context, *kms.GetKeyRotationStatusInput, ...request.Option) (*kms.GetKeyRotationStatusOutput, error)
	GetKeyRotationStatusRequest(*kms.GetKeyRotationStatusInput) (*request.Request, *kms.GetKeyRotationStatusOutput)

	GetParametersForImport(*kms.GetParametersForImportInput) (*kms.GetParametersForImportOutput, error)
	GetParametersForImportWithContext(aws.Context, *kms.GetParametersForImportInput, ...request.Option) (*kms.GetParametersForImportOutput, error)
	GetParametersForImportRequest(*kms.GetParametersForImportInput) (*request.Request, *kms.GetParametersForImportOutput)

	GetPublicKey(*kms.GetPublicKeyInput) (*kms.GetPublicKeyOutput, error)
	GetPublicKeyWithContext(aws.Context, *kms.GetPublicKeyInput, ...request.Option) (*kms.GetPublicKeyOutput, error)
	GetPublicKeyRequest(*kms.GetPublicKeyInput) (*request.Request, *kms.GetPublicKeyOutput)

	ImportKeyMaterial(*kms.ImportKeyMaterialInput) (*kms.ImportKeyMaterialOutput, error)
	ImportKeyMaterialWithContext(aws.Context, *kms.ImportKeyMaterialInput, ...request.Option) (*kms.ImportKeyMaterialOutput, error)
	ImportKeyMaterialRequest(*kms.ImportKeyMaterialInput) (*request.Request, *kms.ImportKeyMaterialOutput)

	ListAliases(*kms.ListAliasesInput) (*kms.ListAliasesOutput, error)
	ListAliasesWithContext(aws.Context, *kms.ListAliasesInput, ...request.Option) (*kms.ListAliasesOutput, error)
	ListAliasesRequest(*kms.ListAliasesInput) (*request.Request, *kms.ListAliasesOutput)

	ListAliasesPages(*kms.ListAliasesInput, func(*kms.ListAliasesOutput, bool) bool) error
	ListAliasesPagesWithContext(aws.Context, *kms.ListAliasesInput, func(*kms.ListAliasesOutput, bool) bool, ...request.Option) error

	ListGrants(*kms.ListGrantsInput) (*kms.ListGrantsResponse, error)
	ListGrantsWithContext(aws.Context, *kms.ListGrantsInput, ...request.Option) (*kms.ListGrantsResponse, error)
	ListGrantsRequest(*kms.ListGrantsInput) (*request.Request, *kms.ListGrantsResponse)

	ListGrantsPages(*kms.ListGrantsInput, func(*kms.ListGrantsResponse, bool) bool) error
	ListGrantsPagesWithContext(aws.Context, *kms.ListGrantsInput, func(*kms.ListGrantsResponse, bool) bool, ...request.Option) error

	ListKeyPolicies(*kms.ListKeyPoliciesInput) (*kms.ListKeyPoliciesOutput, error)
	ListKeyPoliciesWithContext(aws.Context, *kms.ListKeyPoliciesInput, ...request.Option) (*kms.ListKeyPoliciesOutput, error)
	ListKeyPoliciesRequest(*kms.ListKeyPoliciesInput) (*request.Request, *kms.ListKeyPoliciesOutput)

	ListKeyPoliciesPages(*kms.ListKeyPoliciesInput, func(*kms.ListKeyPoliciesOutput, bool) bool) error
	ListKeyPoliciesPagesWithContext(aws.Context, *kms.ListKeyPoliciesInput, func(*kms.ListKeyPoliciesOutput, bool) bool, ...request.Option) error

	ListKeys(*kms.ListKeysInput) (*kms.ListKeysOutput, error)
	ListKeysWithContext(aws.Context, *kms.ListKeysInput, ...request.Option) (*kms.ListKeysOutput, error)
	ListKeysRequest(*kms.ListKeysInput) (*request.Request, *kms.ListKeysOutput)

	ListKeysPages(*kms.ListKeysInput, func(*kms.ListKeysOutput, bool) bool) error
	ListKeysPagesWithContext(aws.Context, *kms.ListKeysInput, func(*kms.ListKeysOutput, bool) bool, ...request.Option) error

	ListResourceTags(*kms.ListResourceTagsInput) (*kms.ListResourceTagsOutput, error)
	ListResourceTagsWithContext(aws.Context, *kms.ListResourceTagsInput, ...request.Option) (*kms.ListResourceTagsOutput, error)
	ListResourceTagsRequest(*kms.ListResourceTagsInput) (*request.Request, *kms.ListResourceTagsOutput)

	ListRetirableGrants(*kms.ListRetirableGrantsInput) (*kms.ListGrantsResponse, error)
	ListRetirableGrantsWithContext(aws.Context, *kms.ListRetirableGrantsInput, ...request.Option) (*kms.ListGrantsResponse, error)
	ListRetirableGrantsRequest(*kms.ListRetirableGrantsInput) (*request.Request, *kms.ListGrantsResponse)

	PutKeyPolicy(*kms.PutKeyPolicyInput) (*kms.PutKeyPolicyOutput, error)
	PutKeyPolicyWithContext(aws.Context, *kms.PutKeyPolicyInput, ...request.Option) (*kms.PutKeyPolicyOutput, error)
	PutKeyPolicyRequest(*kms.PutKeyPolicyInput) (*request.Request, *kms.PutKeyPolicyOutput)

	ReEncrypt(*kms.ReEncryptInput) (*kms.ReEncryptOutput, error)
	ReEncryptWithContext(aws.Context, *kms.ReEncryptInput, ...request.Option) (*kms.ReEncryptOutput, error)
	ReEncryptRequest(*kms.ReEncryptInput) (*request.Request, *kms.ReEncryptOutput)

	ReplicateKey(*kms.ReplicateKeyInput) (*kms.ReplicateKeyOutput, error)
	ReplicateKeyWithContext(aws.Context, *kms.ReplicateKeyInput, ...request.Option) (*kms.ReplicateKeyOutput, error)
	ReplicateKeyRequest(*kms.ReplicateKeyInput) (*request.Request, *kms.ReplicateKeyOutput)

	RetireGrant(*kms.RetireGrantInput) (*kms.RetireGrantOutput, error)
	RetireGrantWithContext(aws.Context, *kms.RetireGrantInput, ...request.Option) (*kms.RetireGrantOutput, error)
	RetireGrantRequest(*kms.RetireGrantInput) (*request.Request, *kms.RetireGrantOutput)

	RevokeGrant(*kms.RevokeGrantInput) (*kms.RevokeGrantOutput, error)
	RevokeGrantWithContext(aws.Context, *kms.RevokeGrantInput, ...request.Option) (*kms.RevokeGrantOutput, error)
	RevokeGrantRequest(*kms.RevokeGrantInput) (*request.Request, *kms.RevokeGrantOutput)

	ScheduleKeyDeletion(*kms.ScheduleKeyDeletionInput) (*kms.ScheduleKeyDeletionOutput, error)
	ScheduleKeyDeletionWithContext(aws.Context, *kms.ScheduleKeyDeletionInput, ...request.Option) (*kms.ScheduleKeyDeletionOutput, error)
	ScheduleKeyDeletionRequest(*kms.ScheduleKeyDeletionInput) (*request.Request, *kms.ScheduleKeyDeletionOutput)

	Sign(*kms.SignInput) (*kms.SignOutput, error)
	SignWithContext(aws.Context, *kms.SignInput, ...request.Option) (*kms.SignOutput, error)
	SignRequest(*kms.SignInput) (*request.Request, *kms.SignOutput)

	TagResource(*kms.TagResourceInput) (*kms.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *kms.TagResourceInput, ...request.Option) (*kms.TagResourceOutput, error)
	TagResourceRequest(*kms.TagResourceInput) (*request.Request, *kms.TagResourceOutput)

	UntagResource(*kms.UntagResourceInput) (*kms.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *kms.UntagResourceInput, ...request.Option) (*kms.UntagResourceOutput, error)
	UntagResourceRequest(*kms.UntagResourceInput) (*request.Request, *kms.UntagResourceOutput)

	UpdateAlias(*kms.UpdateAliasInput) (*kms.UpdateAliasOutput, error)
	UpdateAliasWithContext(aws.Context, *kms.UpdateAliasInput, ...request.Option) (*kms.UpdateAliasOutput, error)
	UpdateAliasRequest(*kms.UpdateAliasInput) (*request.Request, *kms.UpdateAliasOutput)

	UpdateCustomKeyStore(*kms.UpdateCustomKeyStoreInput) (*kms.UpdateCustomKeyStoreOutput, error)
	UpdateCustomKeyStoreWithContext(aws.Context, *kms.UpdateCustomKeyStoreInput, ...request.Option) (*kms.UpdateCustomKeyStoreOutput, error)
	UpdateCustomKeyStoreRequest(*kms.UpdateCustomKeyStoreInput) (*request.Request, *kms.UpdateCustomKeyStoreOutput)

	UpdateKeyDescription(*kms.UpdateKeyDescriptionInput) (*kms.UpdateKeyDescriptionOutput, error)
	UpdateKeyDescriptionWithContext(aws.Context, *kms.UpdateKeyDescriptionInput, ...request.Option) (*kms.UpdateKeyDescriptionOutput, error)
	UpdateKeyDescriptionRequest(*kms.UpdateKeyDescriptionInput) (*request.Request, *kms.UpdateKeyDescriptionOutput)

	UpdatePrimaryRegion(*kms.UpdatePrimaryRegionInput) (*kms.UpdatePrimaryRegionOutput, error)
	UpdatePrimaryRegionWithContext(aws.Context, *kms.UpdatePrimaryRegionInput, ...request.Option) (*kms.UpdatePrimaryRegionOutput, error)
	UpdatePrimaryRegionRequest(*kms.UpdatePrimaryRegionInput) (*request.Request, *kms.UpdatePrimaryRegionOutput)

	Verify(*kms.VerifyInput) (*kms.VerifyOutput, error)
	VerifyWithContext(aws.Context, *kms.VerifyInput, ...request.Option) (*kms.VerifyOutput, error)
	VerifyRequest(*kms.VerifyInput) (*request.Request, *kms.VerifyOutput)
}

var _ KMSAPI = (*kms.KMS)(nil)
