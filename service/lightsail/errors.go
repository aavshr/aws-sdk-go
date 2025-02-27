// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// Lightsail throws this exception when the user cannot be authenticated or
	// uses invalid credentials to access a resource.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAccountSetupInProgressException for service response error code
	// "AccountSetupInProgressException".
	//
	// Lightsail throws this exception when an account is still in the setup in
	// progress state.
	ErrCodeAccountSetupInProgressException = "AccountSetupInProgressException"

	// ErrCodeInvalidInputException for service response error code
	// "InvalidInputException".
	//
	// Lightsail throws this exception when user input does not conform to the validation
	// rules of an input field.
	//
	// Domain and distribution APIs are only available in the N. Virginia (us-east-1)
	// AWS Region. Please set your AWS Region configuration to us-east-1 to create,
	// view, or edit these resources.
	ErrCodeInvalidInputException = "InvalidInputException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// Lightsail throws this exception when it cannot find a resource.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeOperationFailureException for service response error code
	// "OperationFailureException".
	//
	// Lightsail throws this exception when an operation fails to execute.
	ErrCodeOperationFailureException = "OperationFailureException"

	// ErrCodeServiceException for service response error code
	// "ServiceException".
	//
	// A general service exception.
	ErrCodeServiceException = "ServiceException"

	// ErrCodeUnauthenticatedException for service response error code
	// "UnauthenticatedException".
	//
	// Lightsail throws this exception when the user has not been authenticated.
	ErrCodeUnauthenticatedException = "UnauthenticatedException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":           newErrorAccessDeniedException,
	"AccountSetupInProgressException": newErrorAccountSetupInProgressException,
	"InvalidInputException":           newErrorInvalidInputException,
	"NotFoundException":               newErrorNotFoundException,
	"OperationFailureException":       newErrorOperationFailureException,
	"ServiceException":                newErrorServiceException,
	"UnauthenticatedException":        newErrorUnauthenticatedException,
}
