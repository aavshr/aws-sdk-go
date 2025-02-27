// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInsufficientCapacityException for service response error code
	// "InsufficientCapacityException".
	//
	// Returned if there is insufficient capacity to process this expedited request.
	// This error only applies to expedited retrievals and not to standard or bulk
	// retrievals.
	ErrCodeInsufficientCapacityException = "InsufficientCapacityException"

	// ErrCodeInvalidParameterValueException for service response error code
	// "InvalidParameterValueException".
	//
	// Returned if a parameter of the request is incorrectly specified.
	ErrCodeInvalidParameterValueException = "InvalidParameterValueException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Returned if the request results in a vault or account limit being exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMissingParameterValueException for service response error code
	// "MissingParameterValueException".
	//
	// Returned if a required header or parameter is missing from the request.
	ErrCodeMissingParameterValueException = "MissingParameterValueException"

	// ErrCodePolicyEnforcedException for service response error code
	// "PolicyEnforcedException".
	//
	// Returned if a retrieval job would exceed the current data policy's retrieval
	// rate limit. For more information about data retrieval policies,
	ErrCodePolicyEnforcedException = "PolicyEnforcedException"

	// ErrCodeRequestTimeoutException for service response error code
	// "RequestTimeoutException".
	//
	// Returned if, when uploading an archive, Amazon S3 Glacier times out while
	// receiving the upload.
	ErrCodeRequestTimeoutException = "RequestTimeoutException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Returned if the specified resource (such as a vault, upload ID, or job ID)
	// doesn't exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// Returned if the service cannot complete the request.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InsufficientCapacityException":  newErrorInsufficientCapacityException,
	"InvalidParameterValueException": newErrorInvalidParameterValueException,
	"LimitExceededException":         newErrorLimitExceededException,
	"MissingParameterValueException": newErrorMissingParameterValueException,
	"PolicyEnforcedException":        newErrorPolicyEnforcedException,
	"RequestTimeoutException":        newErrorRequestTimeoutException,
	"ResourceNotFoundException":      newErrorResourceNotFoundException,
	"ServiceUnavailableException":    newErrorServiceUnavailableException,
}
