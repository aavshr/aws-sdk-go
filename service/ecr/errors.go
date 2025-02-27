// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeEmptyUploadException for service response error code
	// "EmptyUploadException".
	//
	// The specified layer upload does not contain any layer parts.
	ErrCodeEmptyUploadException = "EmptyUploadException"

	// ErrCodeImageAlreadyExistsException for service response error code
	// "ImageAlreadyExistsException".
	//
	// The specified image has already been pushed, and there were no changes to
	// the manifest or image tag after the last push.
	ErrCodeImageAlreadyExistsException = "ImageAlreadyExistsException"

	// ErrCodeImageDigestDoesNotMatchException for service response error code
	// "ImageDigestDoesNotMatchException".
	//
	// The specified image digest does not match the digest that Amazon ECR calculated
	// for the image.
	ErrCodeImageDigestDoesNotMatchException = "ImageDigestDoesNotMatchException"

	// ErrCodeImageNotFoundException for service response error code
	// "ImageNotFoundException".
	//
	// The image requested does not exist in the specified repository.
	ErrCodeImageNotFoundException = "ImageNotFoundException"

	// ErrCodeImageTagAlreadyExistsException for service response error code
	// "ImageTagAlreadyExistsException".
	//
	// The specified image is tagged with a tag that already exists. The repository
	// is configured for tag immutability.
	ErrCodeImageTagAlreadyExistsException = "ImageTagAlreadyExistsException"

	// ErrCodeInvalidLayerException for service response error code
	// "InvalidLayerException".
	//
	// The layer digest calculation performed by Amazon ECR upon receipt of the
	// image layer does not match the digest specified.
	ErrCodeInvalidLayerException = "InvalidLayerException"

	// ErrCodeInvalidLayerPartException for service response error code
	// "InvalidLayerPartException".
	//
	// The layer part size is not valid, or the first byte specified is not consecutive
	// to the last byte of a previous layer part upload.
	ErrCodeInvalidLayerPartException = "InvalidLayerPartException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// The specified parameter is invalid. Review the available parameters for the
	// API request.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidTagParameterException for service response error code
	// "InvalidTagParameterException".
	//
	// An invalid parameter has been specified. Tag keys can have a maximum character
	// length of 128 characters, and tag values can have a maximum length of 256
	// characters.
	ErrCodeInvalidTagParameterException = "InvalidTagParameterException"

	// ErrCodeKmsException for service response error code
	// "KmsException".
	//
	// The operation failed due to a KMS exception.
	ErrCodeKmsException = "KmsException"

	// ErrCodeLayerAlreadyExistsException for service response error code
	// "LayerAlreadyExistsException".
	//
	// The image layer already exists in the associated repository.
	ErrCodeLayerAlreadyExistsException = "LayerAlreadyExistsException"

	// ErrCodeLayerInaccessibleException for service response error code
	// "LayerInaccessibleException".
	//
	// The specified layer is not available because it is not associated with an
	// image. Unassociated image layers may be cleaned up at any time.
	ErrCodeLayerInaccessibleException = "LayerInaccessibleException"

	// ErrCodeLayerPartTooSmallException for service response error code
	// "LayerPartTooSmallException".
	//
	// Layer parts must be at least 5 MiB in size.
	ErrCodeLayerPartTooSmallException = "LayerPartTooSmallException"

	// ErrCodeLayersNotFoundException for service response error code
	// "LayersNotFoundException".
	//
	// The specified layers could not be found, or the specified layer is not valid
	// for this repository.
	ErrCodeLayersNotFoundException = "LayersNotFoundException"

	// ErrCodeLifecyclePolicyNotFoundException for service response error code
	// "LifecyclePolicyNotFoundException".
	//
	// The lifecycle policy could not be found, and no policy is set to the repository.
	ErrCodeLifecyclePolicyNotFoundException = "LifecyclePolicyNotFoundException"

	// ErrCodeLifecyclePolicyPreviewInProgressException for service response error code
	// "LifecyclePolicyPreviewInProgressException".
	//
	// The previous lifecycle policy preview request has not completed. Wait and
	// try again.
	ErrCodeLifecyclePolicyPreviewInProgressException = "LifecyclePolicyPreviewInProgressException"

	// ErrCodeLifecyclePolicyPreviewNotFoundException for service response error code
	// "LifecyclePolicyPreviewNotFoundException".
	//
	// There is no dry run for this repository.
	ErrCodeLifecyclePolicyPreviewNotFoundException = "LifecyclePolicyPreviewNotFoundException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The operation did not succeed because it would have exceeded a service limit
	// for your account. For more information, see Amazon ECR service quotas (https://docs.aws.amazon.com/AmazonECR/latest/userguide/service-quotas.html)
	// in the Amazon Elastic Container Registry User Guide.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeReferencedImagesNotFoundException for service response error code
	// "ReferencedImagesNotFoundException".
	//
	// The manifest list is referencing an image that does not exist.
	ErrCodeReferencedImagesNotFoundException = "ReferencedImagesNotFoundException"

	// ErrCodeRegistryPolicyNotFoundException for service response error code
	// "RegistryPolicyNotFoundException".
	//
	// The registry doesn't have an associated registry policy.
	ErrCodeRegistryPolicyNotFoundException = "RegistryPolicyNotFoundException"

	// ErrCodeRepositoryAlreadyExistsException for service response error code
	// "RepositoryAlreadyExistsException".
	//
	// The specified repository already exists in the specified registry.
	ErrCodeRepositoryAlreadyExistsException = "RepositoryAlreadyExistsException"

	// ErrCodeRepositoryNotEmptyException for service response error code
	// "RepositoryNotEmptyException".
	//
	// The specified repository contains images. To delete a repository that contains
	// images, you must force the deletion with the force parameter.
	ErrCodeRepositoryNotEmptyException = "RepositoryNotEmptyException"

	// ErrCodeRepositoryNotFoundException for service response error code
	// "RepositoryNotFoundException".
	//
	// The specified repository could not be found. Check the spelling of the specified
	// repository and ensure that you are performing operations on the correct registry.
	ErrCodeRepositoryNotFoundException = "RepositoryNotFoundException"

	// ErrCodeRepositoryPolicyNotFoundException for service response error code
	// "RepositoryPolicyNotFoundException".
	//
	// The specified repository and registry combination does not have an associated
	// repository policy.
	ErrCodeRepositoryPolicyNotFoundException = "RepositoryPolicyNotFoundException"

	// ErrCodeScanNotFoundException for service response error code
	// "ScanNotFoundException".
	//
	// The specified image scan could not be found. Ensure that image scanning is
	// enabled on the repository and try again.
	ErrCodeScanNotFoundException = "ScanNotFoundException"

	// ErrCodeServerException for service response error code
	// "ServerException".
	//
	// These errors are usually caused by a server-side issue.
	ErrCodeServerException = "ServerException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// The list of tags on the repository is over the limit. The maximum number
	// of tags that can be applied to a repository is 50.
	ErrCodeTooManyTagsException = "TooManyTagsException"

	// ErrCodeUnsupportedImageTypeException for service response error code
	// "UnsupportedImageTypeException".
	//
	// The image is of a type that cannot be scanned.
	ErrCodeUnsupportedImageTypeException = "UnsupportedImageTypeException"

	// ErrCodeUploadNotFoundException for service response error code
	// "UploadNotFoundException".
	//
	// The upload could not be found, or the specified upload ID is not valid for
	// this repository.
	ErrCodeUploadNotFoundException = "UploadNotFoundException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// There was an exception validating this request.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"EmptyUploadException":                      newErrorEmptyUploadException,
	"ImageAlreadyExistsException":               newErrorImageAlreadyExistsException,
	"ImageDigestDoesNotMatchException":          newErrorImageDigestDoesNotMatchException,
	"ImageNotFoundException":                    newErrorImageNotFoundException,
	"ImageTagAlreadyExistsException":            newErrorImageTagAlreadyExistsException,
	"InvalidLayerException":                     newErrorInvalidLayerException,
	"InvalidLayerPartException":                 newErrorInvalidLayerPartException,
	"InvalidParameterException":                 newErrorInvalidParameterException,
	"InvalidTagParameterException":              newErrorInvalidTagParameterException,
	"KmsException":                              newErrorKmsException,
	"LayerAlreadyExistsException":               newErrorLayerAlreadyExistsException,
	"LayerInaccessibleException":                newErrorLayerInaccessibleException,
	"LayerPartTooSmallException":                newErrorLayerPartTooSmallException,
	"LayersNotFoundException":                   newErrorLayersNotFoundException,
	"LifecyclePolicyNotFoundException":          newErrorLifecyclePolicyNotFoundException,
	"LifecyclePolicyPreviewInProgressException": newErrorLifecyclePolicyPreviewInProgressException,
	"LifecyclePolicyPreviewNotFoundException":   newErrorLifecyclePolicyPreviewNotFoundException,
	"LimitExceededException":                    newErrorLimitExceededException,
	"ReferencedImagesNotFoundException":         newErrorReferencedImagesNotFoundException,
	"RegistryPolicyNotFoundException":           newErrorRegistryPolicyNotFoundException,
	"RepositoryAlreadyExistsException":          newErrorRepositoryAlreadyExistsException,
	"RepositoryNotEmptyException":               newErrorRepositoryNotEmptyException,
	"RepositoryNotFoundException":               newErrorRepositoryNotFoundException,
	"RepositoryPolicyNotFoundException":         newErrorRepositoryPolicyNotFoundException,
	"ScanNotFoundException":                     newErrorScanNotFoundException,
	"ServerException":                           newErrorServerException,
	"TooManyTagsException":                      newErrorTooManyTagsException,
	"UnsupportedImageTypeException":             newErrorUnsupportedImageTypeException,
	"UploadNotFoundException":                   newErrorUploadNotFoundException,
	"ValidationException":                       newErrorValidationException,
}
