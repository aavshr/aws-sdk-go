// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssooidc

import (
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAuthorizationPendingException for service response error code
	// "AuthorizationPendingException".
	//
	// Indicates that a request to authorize a client with an access user session
	// token is pending.
	ErrCodeAuthorizationPendingException = "AuthorizationPendingException"

	// ErrCodeExpiredTokenException for service response error code
	// "ExpiredTokenException".
	//
	// Indicates that the token issued by the service is expired and is no longer
	// valid.
	ErrCodeExpiredTokenException = "ExpiredTokenException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// Indicates that an error from the service occurred while trying to process
	// a request.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeInvalidClientException for service response error code
	// "InvalidClientException".
	//
	// Indicates that the clientId or clientSecret in the request is invalid. For
	// example, this can occur when a client sends an incorrect clientId or an expired
	// clientSecret.
	ErrCodeInvalidClientException = "InvalidClientException"

	// ErrCodeInvalidClientMetadataException for service response error code
	// "InvalidClientMetadataException".
	//
	// Indicates that the client information sent in the request during registration
	// is invalid.
	ErrCodeInvalidClientMetadataException = "InvalidClientMetadataException"

	// ErrCodeInvalidGrantException for service response error code
	// "InvalidGrantException".
	//
	// Indicates that a request contains an invalid grant. This can occur if a client
	// makes a CreateToken request with an invalid grant type.
	ErrCodeInvalidGrantException = "InvalidGrantException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// Indicates that something is wrong with the input to the request. For example,
	// a required parameter might be missing or out of range.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeInvalidScopeException for service response error code
	// "InvalidScopeException".
	//
	// Indicates that the scope provided in the request is invalid.
	ErrCodeInvalidScopeException = "InvalidScopeException"

	// ErrCodeSlowDownException for service response error code
	// "SlowDownException".
	//
	// Indicates that the client is making the request too frequently and is more
	// than the service can handle.
	ErrCodeSlowDownException = "SlowDownException"

	// ErrCodeUnauthorizedClientException for service response error code
	// "UnauthorizedClientException".
	//
	// Indicates that the client is not currently authorized to make the request.
	// This can happen when a clientId is not issued for a public client.
	ErrCodeUnauthorizedClientException = "UnauthorizedClientException"

	// ErrCodeUnsupportedGrantTypeException for service response error code
	// "UnsupportedGrantTypeException".
	//
	// Indicates that the grant type in the request is not supported by the service.
	ErrCodeUnsupportedGrantTypeException = "UnsupportedGrantTypeException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":          newErrorAccessDeniedException,
	"AuthorizationPendingException":  newErrorAuthorizationPendingException,
	"ExpiredTokenException":          newErrorExpiredTokenException,
	"InternalServerException":        newErrorInternalServerException,
	"InvalidClientException":         newErrorInvalidClientException,
	"InvalidClientMetadataException": newErrorInvalidClientMetadataException,
	"InvalidGrantException":          newErrorInvalidGrantException,
	"InvalidRequestException":        newErrorInvalidRequestException,
	"InvalidScopeException":          newErrorInvalidScopeException,
	"SlowDownException":              newErrorSlowDownException,
	"UnauthorizedClientException":    newErrorUnauthorizedClientException,
	"UnsupportedGrantTypeException":  newErrorUnsupportedGrantTypeException,
}
