// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package outposts

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have permission to perform this operation.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// Updating or deleting this resource can cause an inconsistent state.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An internal error has occurred.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// The specified request is not valid.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// You have exceeded a service quota.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// A parameter is not valid.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":         newErrorAccessDeniedException,
	"ConflictException":             newErrorConflictException,
	"InternalServerException":       newErrorInternalServerException,
	"NotFoundException":             newErrorNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"ValidationException":           newErrorValidationException,
}
