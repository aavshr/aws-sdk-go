// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sso

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// Indicates that a problem occurred with the input to the request. For example,
	// a required parameter might be missing or out of range.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource doesn't exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// Indicates that the request is being made too frequently and is more than
	// what the server can handle.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// Indicates that the request is not authorized. This can happen due to an invalid
	// access token in the request.
	ErrCodeUnauthorizedException = "UnauthorizedException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InvalidRequestException":   newErrorInvalidRequestException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
	"TooManyRequestsException":  newErrorTooManyRequestsException,
	"UnauthorizedException":     newErrorUnauthorizedException,
}
