// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pi

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInternalServiceError for service response error code
	// "InternalServiceError".
	//
	// The request failed due to an unknown error.
	ErrCodeInternalServiceError = "InternalServiceError"

	// ErrCodeInvalidArgumentException for service response error code
	// "InvalidArgumentException".
	//
	// One of the arguments provided is invalid for this request.
	ErrCodeInvalidArgumentException = "InvalidArgumentException"

	// ErrCodeNotAuthorizedException for service response error code
	// "NotAuthorizedException".
	//
	// The user is not authorized to perform this request.
	ErrCodeNotAuthorizedException = "NotAuthorizedException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalServiceError":     newErrorInternalServiceError,
	"InvalidArgumentException": newErrorInvalidArgumentException,
	"NotAuthorizedException":   newErrorNotAuthorizedException,
}
