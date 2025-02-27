// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"time"

	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
)

// WaitUntilTableExists uses the DynamoDB API operation
// DescribeTable to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *DynamoDB) WaitUntilTableExists(input *DescribeTableInput) error {
	return c.WaitUntilTableExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilTableExistsWithContext is an extended version of WaitUntilTableExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DynamoDB) WaitUntilTableExistsWithContext(ctx aws.Context, input *DescribeTableInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilTableExists",
		MaxAttempts: 25,
		Delay:       request.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Table.TableStatus",
				Expected: "ACTIVE",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeTableInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeTableRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilTableNotExists uses the DynamoDB API operation
// DescribeTable to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *DynamoDB) WaitUntilTableNotExists(input *DescribeTableInput) error {
	return c.WaitUntilTableNotExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilTableNotExistsWithContext is an extended version of WaitUntilTableNotExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DynamoDB) WaitUntilTableNotExistsWithContext(ctx aws.Context, input *DescribeTableInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilTableNotExists",
		MaxAttempts: 25,
		Delay:       request.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeTableInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeTableRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
