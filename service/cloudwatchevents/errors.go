// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchevents

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// There is concurrent modification on a rule, target, archive, or replay.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeIllegalStatusException for service response error code
	// "IllegalStatusException".
	//
	// An error occurred because a replay can be canceled only when the state is
	// Running or Starting.
	ErrCodeIllegalStatusException = "IllegalStatusException"

	// ErrCodeInternalException for service response error code
	// "InternalException".
	//
	// This exception occurs due to unexpected causes.
	ErrCodeInternalException = "InternalException"

	// ErrCodeInvalidEventPatternException for service response error code
	// "InvalidEventPatternException".
	//
	// The event pattern is not valid.
	ErrCodeInvalidEventPatternException = "InvalidEventPatternException"

	// ErrCodeInvalidStateException for service response error code
	// "InvalidStateException".
	//
	// The specified state is not a valid state for an event source.
	ErrCodeInvalidStateException = "InvalidStateException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The request failed because it attempted to create resource beyond the allowed
	// service quota.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeManagedRuleException for service response error code
	// "ManagedRuleException".
	//
	// This rule was created by an Amazon Web Services service on behalf of your
	// account. It is managed by that service. If you see this error in response
	// to DeleteRule or RemoveTargets, you can use the Force parameter in those
	// calls to delete the rule or remove targets from the rule. You cannot modify
	// these managed rules by using DisableRule, EnableRule, PutTargets, PutRule,
	// TagResource, or UntagResource.
	ErrCodeManagedRuleException = "ManagedRuleException"

	// ErrCodeOperationDisabledException for service response error code
	// "OperationDisabledException".
	//
	// The operation you are attempting is not available in this region.
	ErrCodeOperationDisabledException = "OperationDisabledException"

	// ErrCodePolicyLengthExceededException for service response error code
	// "PolicyLengthExceededException".
	//
	// The event bus policy is too long. For more information, see the limits.
	ErrCodePolicyLengthExceededException = "PolicyLengthExceededException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	//
	// The resource you are trying to create already exists.
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// An entity that you specified does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConcurrentModificationException": newErrorConcurrentModificationException,
	"IllegalStatusException":          newErrorIllegalStatusException,
	"InternalException":               newErrorInternalException,
	"InvalidEventPatternException":    newErrorInvalidEventPatternException,
	"InvalidStateException":           newErrorInvalidStateException,
	"LimitExceededException":          newErrorLimitExceededException,
	"ManagedRuleException":            newErrorManagedRuleException,
	"OperationDisabledException":      newErrorOperationDisabledException,
	"PolicyLengthExceededException":   newErrorPolicyLengthExceededException,
	"ResourceAlreadyExistsException":  newErrorResourceAlreadyExistsException,
	"ResourceNotFoundException":       newErrorResourceNotFoundException,
}
