// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// A bad request exception object.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeInternalServerErrorException for service response error code
	// "InternalServerErrorException".
	//
	// An internal server error exception object.
	ErrCodeInternalServerErrorException = "InternalServerErrorException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":          newErrorBadRequestException,
	"InternalServerErrorException": newErrorInternalServerErrorException,
}
