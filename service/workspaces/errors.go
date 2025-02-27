// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// The user is not authorized to access a resource.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeInvalidParameterValuesException for service response error code
	// "InvalidParameterValuesException".
	//
	// One or more parameter values are not valid.
	ErrCodeInvalidParameterValuesException = "InvalidParameterValuesException"

	// ErrCodeInvalidResourceStateException for service response error code
	// "InvalidResourceStateException".
	//
	// The state of the resource is not valid for this operation.
	ErrCodeInvalidResourceStateException = "InvalidResourceStateException"

	// ErrCodeOperationInProgressException for service response error code
	// "OperationInProgressException".
	//
	// The properties of this WorkSpace are currently being modified. Try again
	// in a moment.
	ErrCodeOperationInProgressException = "OperationInProgressException"

	// ErrCodeOperationNotSupportedException for service response error code
	// "OperationNotSupportedException".
	//
	// This operation is not supported.
	ErrCodeOperationNotSupportedException = "OperationNotSupportedException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	//
	// The specified resource already exists.
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeResourceAssociatedException for service response error code
	// "ResourceAssociatedException".
	//
	// The resource is associated with a directory.
	ErrCodeResourceAssociatedException = "ResourceAssociatedException"

	// ErrCodeResourceCreationFailedException for service response error code
	// "ResourceCreationFailedException".
	//
	// The resource could not be created.
	ErrCodeResourceCreationFailedException = "ResourceCreationFailedException"

	// ErrCodeResourceLimitExceededException for service response error code
	// "ResourceLimitExceededException".
	//
	// Your resource limits have been exceeded.
	ErrCodeResourceLimitExceededException = "ResourceLimitExceededException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource could not be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceUnavailableException for service response error code
	// "ResourceUnavailableException".
	//
	// The specified resource is not available.
	ErrCodeResourceUnavailableException = "ResourceUnavailableException"

	// ErrCodeUnsupportedNetworkConfigurationException for service response error code
	// "UnsupportedNetworkConfigurationException".
	//
	// The configuration of this network is not supported for this operation, or
	// your network configuration conflicts with the Amazon WorkSpaces management
	// network IP range. For more information, see Configure a VPC for Amazon WorkSpaces
	// (https://docs.aws.amazon.com/workspaces/latest/adminguide/amazon-workspaces-vpc.html).
	ErrCodeUnsupportedNetworkConfigurationException = "UnsupportedNetworkConfigurationException"

	// ErrCodeUnsupportedWorkspaceConfigurationException for service response error code
	// "UnsupportedWorkspaceConfigurationException".
	//
	// The configuration of this WorkSpace is not supported for this operation.
	// For more information, see Required Configuration and Service Components for
	// WorkSpaces (https://docs.aws.amazon.com/workspaces/latest/adminguide/required-service-components.html).
	ErrCodeUnsupportedWorkspaceConfigurationException = "UnsupportedWorkspaceConfigurationException"

	// ErrCodeWorkspacesDefaultRoleNotFoundException for service response error code
	// "WorkspacesDefaultRoleNotFoundException".
	//
	// The workspaces_DefaultRole role could not be found. If this is the first
	// time you are registering a directory, you will need to create the workspaces_DefaultRole
	// role before you can register a directory. For more information, see Creating
	// the workspaces_DefaultRole Role (https://docs.aws.amazon.com/workspaces/latest/adminguide/workspaces-access-control.html#create-default-role).
	ErrCodeWorkspacesDefaultRoleNotFoundException = "WorkspacesDefaultRoleNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                      newErrorAccessDeniedException,
	"InvalidParameterValuesException":            newErrorInvalidParameterValuesException,
	"InvalidResourceStateException":              newErrorInvalidResourceStateException,
	"OperationInProgressException":               newErrorOperationInProgressException,
	"OperationNotSupportedException":             newErrorOperationNotSupportedException,
	"ResourceAlreadyExistsException":             newErrorResourceAlreadyExistsException,
	"ResourceAssociatedException":                newErrorResourceAssociatedException,
	"ResourceCreationFailedException":            newErrorResourceCreationFailedException,
	"ResourceLimitExceededException":             newErrorResourceLimitExceededException,
	"ResourceNotFoundException":                  newErrorResourceNotFoundException,
	"ResourceUnavailableException":               newErrorResourceUnavailableException,
	"UnsupportedNetworkConfigurationException":   newErrorUnsupportedNetworkConfigurationException,
	"UnsupportedWorkspaceConfigurationException": newErrorUnsupportedWorkspaceConfigurationException,
	"WorkspacesDefaultRoleNotFoundException":     newErrorWorkspacesDefaultRoleNotFoundException,
}
