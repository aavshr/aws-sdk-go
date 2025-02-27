// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53recoverycluster

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/client"
	"github.com/aavshr/aws-sdk-go/aws/client/metadata"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/aws/signer/v4"
	"github.com/aavshr/aws-sdk-go/private/protocol"
	"github.com/aavshr/aws-sdk-go/private/protocol/jsonrpc"
)

// Route53RecoveryCluster provides the API operation methods for making requests to
// Route53 Recovery Cluster. See this package's package overview docs
// for details on the service.
//
// Route53RecoveryCluster methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Route53RecoveryCluster struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "Route53 Recovery Cluster" // Name of service.
	EndpointsID = "route53-recovery-cluster" // ID to lookup a service endpoint with.
	ServiceID   = "Route53 Recovery Cluster" // ServiceID is a unique identifier of a specific service.
)

// New creates a new instance of the Route53RecoveryCluster client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     mySession := session.Must(session.NewSession())
//
//     // Create a Route53RecoveryCluster client from just a session.
//     svc := route53recoverycluster.New(mySession)
//
//     // Create a Route53RecoveryCluster client with additional configuration
//     svc := route53recoverycluster.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Route53RecoveryCluster {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "route53-recovery-cluster"
	}
	return newClient(*c.Config, c.Handlers, c.PartitionID, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, partitionID, endpoint, signingRegion, signingName string) *Route53RecoveryCluster {
	svc := &Route53RecoveryCluster{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				PartitionID:   partitionID,
				Endpoint:      endpoint,
				APIVersion:    "2019-12-02",
				JSONVersion:   "1.0",
				TargetPrefix:  "ToggleCustomerAPI",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(
		protocol.NewUnmarshalErrorHandler(jsonrpc.NewUnmarshalTypedError(exceptionFromCode)).NamedHandler(),
	)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Route53RecoveryCluster operation and runs any
// custom request initialization.
func (c *Route53RecoveryCluster) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
