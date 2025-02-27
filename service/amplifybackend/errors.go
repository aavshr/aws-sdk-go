// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplifybackend

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// An error returned if a request is not formed properly.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeGatewayTimeoutException for service response error code
	// "GatewayTimeoutException".
	//
	// An error returned if there's a temporary issue with the service.
	ErrCodeGatewayTimeoutException = "GatewayTimeoutException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// An error returned when a specific resource type is not found.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// An error that is returned when a limit of a specific type has been exceeded.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":      newErrorBadRequestException,
	"GatewayTimeoutException":  newErrorGatewayTimeoutException,
	"NotFoundException":        newErrorNotFoundException,
	"TooManyRequestsException": newErrorTooManyRequestsException,
}
