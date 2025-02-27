// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeCodeValidationException for service response error code
	// "CodeValidationException".
	//
	// The user-provided application code (query) is not valid. This can be a simple
	// syntax error.
	ErrCodeCodeValidationException = "CodeValidationException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// Exception thrown as a result of concurrent modifications to an application.
	// This error can be the result of attempting to modify an application without
	// using the current application ID.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeInvalidApplicationConfigurationException for service response error code
	// "InvalidApplicationConfigurationException".
	//
	// The user-provided application configuration is not valid.
	ErrCodeInvalidApplicationConfigurationException = "InvalidApplicationConfigurationException"

	// ErrCodeInvalidArgumentException for service response error code
	// "InvalidArgumentException".
	//
	// The specified input parameter value is not valid.
	ErrCodeInvalidArgumentException = "InvalidArgumentException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The request JSON is not valid for the operation.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The number of allowed resources has been exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The application is not available for this operation.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Specified application can't be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceProvisionedThroughputExceededException for service response error code
	// "ResourceProvisionedThroughputExceededException".
	//
	// Discovery failed to get a record from the streaming source because of the
	// Kinesis Streams ProvisionedThroughputExceededException. For more information,
	// see GetRecords (http://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetRecords.html)
	// in the Amazon Kinesis Streams API Reference.
	ErrCodeResourceProvisionedThroughputExceededException = "ResourceProvisionedThroughputExceededException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service cannot complete the request.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// Application created with too many tags, or too many tags added to an application.
	// Note that the maximum number of application tags includes system tags. The
	// maximum number of user-defined application tags is 50.
	ErrCodeTooManyTagsException = "TooManyTagsException"

	// ErrCodeUnableToDetectSchemaException for service response error code
	// "UnableToDetectSchemaException".
	//
	// The data format is not valid. Kinesis Data Analytics cannot detect the schema
	// for the given streaming source.
	ErrCodeUnableToDetectSchemaException = "UnableToDetectSchemaException"

	// ErrCodeUnsupportedOperationException for service response error code
	// "UnsupportedOperationException".
	//
	// The request was rejected because a specified parameter is not supported or
	// a specified resource is not valid for this operation.
	ErrCodeUnsupportedOperationException = "UnsupportedOperationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"CodeValidationException":                        newErrorCodeValidationException,
	"ConcurrentModificationException":                newErrorConcurrentModificationException,
	"InvalidApplicationConfigurationException":       newErrorInvalidApplicationConfigurationException,
	"InvalidArgumentException":                       newErrorInvalidArgumentException,
	"InvalidRequestException":                        newErrorInvalidRequestException,
	"LimitExceededException":                         newErrorLimitExceededException,
	"ResourceInUseException":                         newErrorResourceInUseException,
	"ResourceNotFoundException":                      newErrorResourceNotFoundException,
	"ResourceProvisionedThroughputExceededException": newErrorResourceProvisionedThroughputExceededException,
	"ServiceUnavailableException":                    newErrorServiceUnavailableException,
	"TooManyTagsException":                           newErrorTooManyTagsException,
	"UnableToDetectSchemaException":                  newErrorUnableToDetectSchemaException,
	"UnsupportedOperationException":                  newErrorUnsupportedOperationException,
}
