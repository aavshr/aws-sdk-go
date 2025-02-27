// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// Indicates that an error occurred while processing the request and that the
	// request was not completed.
	ErrCodeInternalServerError = "InternalServerError"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// This exception occurs when there is an internal failure in the Amazon EMR
	// service.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// This exception occurs when there is something wrong with user input.
	ErrCodeInvalidRequestException = "InvalidRequestException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalServerError":     newErrorInternalServerError,
	"InternalServerException": newErrorInternalServerException,
	"InvalidRequestException": newErrorInvalidRequestException,
}
