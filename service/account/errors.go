// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package account

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// The operation failed because the calling identity doesn't have the minimum
	// required permissions.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// The operation failed because of an error internal to Amazon Web Services.
	// Try your operation again later.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The operation failed because it specified a resource that can't be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// The operation failed because it was called too frequently and exceeded a
	// throttle limit.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The operation failed because one of the input parameters was invalid.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":     newErrorAccessDeniedException,
	"InternalServerException":   newErrorInternalServerException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
	"TooManyRequestsException":  newErrorTooManyRequestsException,
	"ValidationException":       newErrorValidationException,
}
