// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53recoverycontrolconfig

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/client"
	"github.com/aavshr/aws-sdk-go/aws/client/metadata"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/aws/signer/v4"
	"github.com/aavshr/aws-sdk-go/private/protocol"
	"github.com/aavshr/aws-sdk-go/private/protocol/restjson"
)

// Route53RecoveryControlConfig provides the API operation methods for making requests to
// AWS Route53 Recovery Control Config. See this package's package overview docs
// for details on the service.
//
// Route53RecoveryControlConfig methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Route53RecoveryControlConfig struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "Route53 Recovery Control Config" // Name of service.
	EndpointsID = "route53-recovery-control-config" // ID to lookup a service endpoint with.
	ServiceID   = "Route53 Recovery Control Config" // ServiceID is a unique identifier of a specific service.
)

// New creates a new instance of the Route53RecoveryControlConfig client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     mySession := session.Must(session.NewSession())
//
//     // Create a Route53RecoveryControlConfig client from just a session.
//     svc := route53recoverycontrolconfig.New(mySession)
//
//     // Create a Route53RecoveryControlConfig client with additional configuration
//     svc := route53recoverycontrolconfig.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Route53RecoveryControlConfig {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "route53-recovery-control-config"
	}
	return newClient(*c.Config, c.Handlers, c.PartitionID, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, partitionID, endpoint, signingRegion, signingName string) *Route53RecoveryControlConfig {
	svc := &Route53RecoveryControlConfig{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				PartitionID:   partitionID,
				Endpoint:      endpoint,
				APIVersion:    "2020-11-02",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(
		protocol.NewUnmarshalErrorHandler(restjson.NewUnmarshalTypedError(exceptionFromCode)).NamedHandler(),
	)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Route53RecoveryControlConfig operation and runs any
// custom request initialization.
func (c *Route53RecoveryControlConfig) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
