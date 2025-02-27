// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"time"

	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
)

// WaitUntilChannelCreated uses the MediaLive API operation
// DescribeChannel to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaLive) WaitUntilChannelCreated(input *DescribeChannelInput) error {
	return c.WaitUntilChannelCreatedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilChannelCreatedWithContext is an extended version of WaitUntilChannelCreated.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaLive) WaitUntilChannelCreatedWithContext(ctx aws.Context, input *DescribeChannelInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilChannelCreated",
		MaxAttempts: 5,
		Delay:       request.ConstantWaiterDelay(3 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "IDLE",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "CREATING",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "CREATE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeChannelInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeChannelRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilChannelDeleted uses the MediaLive API operation
// DescribeChannel to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaLive) WaitUntilChannelDeleted(input *DescribeChannelInput) error {
	return c.WaitUntilChannelDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilChannelDeletedWithContext is an extended version of WaitUntilChannelDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaLive) WaitUntilChannelDeletedWithContext(ctx aws.Context, input *DescribeChannelInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilChannelDeleted",
		MaxAttempts: 84,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "DELETED",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "DELETING",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeChannelInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeChannelRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilChannelRunning uses the MediaLive API operation
// DescribeChannel to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaLive) WaitUntilChannelRunning(input *DescribeChannelInput) error {
	return c.WaitUntilChannelRunningWithContext(aws.BackgroundContext(), input)
}

// WaitUntilChannelRunningWithContext is an extended version of WaitUntilChannelRunning.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaLive) WaitUntilChannelRunningWithContext(ctx aws.Context, input *DescribeChannelInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilChannelRunning",
		MaxAttempts: 120,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "RUNNING",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "STARTING",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeChannelInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeChannelRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilChannelStopped uses the MediaLive API operation
// DescribeChannel to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaLive) WaitUntilChannelStopped(input *DescribeChannelInput) error {
	return c.WaitUntilChannelStoppedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilChannelStoppedWithContext is an extended version of WaitUntilChannelStopped.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaLive) WaitUntilChannelStoppedWithContext(ctx aws.Context, input *DescribeChannelInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilChannelStopped",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "IDLE",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "STOPPING",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeChannelInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeChannelRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilInputAttached uses the MediaLive API operation
// DescribeInput to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaLive) WaitUntilInputAttached(input *DescribeInputInput) error {
	return c.WaitUntilInputAttachedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInputAttachedWithContext is an extended version of WaitUntilInputAttached.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaLive) WaitUntilInputAttachedWithContext(ctx aws.Context, input *DescribeInputInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInputAttached",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "ATTACHED",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "DETACHED",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeInputInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeInputRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilInputDeleted uses the MediaLive API operation
// DescribeInput to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaLive) WaitUntilInputDeleted(input *DescribeInputInput) error {
	return c.WaitUntilInputDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInputDeletedWithContext is an extended version of WaitUntilInputDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaLive) WaitUntilInputDeletedWithContext(ctx aws.Context, input *DescribeInputInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInputDeleted",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "DELETED",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "DELETING",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeInputInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeInputRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilInputDetached uses the MediaLive API operation
// DescribeInput to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaLive) WaitUntilInputDetached(input *DescribeInputInput) error {
	return c.WaitUntilInputDetachedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInputDetachedWithContext is an extended version of WaitUntilInputDetached.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaLive) WaitUntilInputDetachedWithContext(ctx aws.Context, input *DescribeInputInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInputDetached",
		MaxAttempts: 84,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "DETACHED",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "CREATING",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "ATTACHED",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeInputInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeInputRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilMultiplexCreated uses the MediaLive API operation
// DescribeMultiplex to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaLive) WaitUntilMultiplexCreated(input *DescribeMultiplexInput) error {
	return c.WaitUntilMultiplexCreatedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilMultiplexCreatedWithContext is an extended version of WaitUntilMultiplexCreated.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaLive) WaitUntilMultiplexCreatedWithContext(ctx aws.Context, input *DescribeMultiplexInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilMultiplexCreated",
		MaxAttempts: 5,
		Delay:       request.ConstantWaiterDelay(3 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "IDLE",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "CREATING",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "CREATE_FAILED",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeMultiplexInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeMultiplexRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilMultiplexDeleted uses the MediaLive API operation
// DescribeMultiplex to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaLive) WaitUntilMultiplexDeleted(input *DescribeMultiplexInput) error {
	return c.WaitUntilMultiplexDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilMultiplexDeletedWithContext is an extended version of WaitUntilMultiplexDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaLive) WaitUntilMultiplexDeletedWithContext(ctx aws.Context, input *DescribeMultiplexInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilMultiplexDeleted",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "DELETED",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "DELETING",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeMultiplexInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeMultiplexRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilMultiplexRunning uses the MediaLive API operation
// DescribeMultiplex to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaLive) WaitUntilMultiplexRunning(input *DescribeMultiplexInput) error {
	return c.WaitUntilMultiplexRunningWithContext(aws.BackgroundContext(), input)
}

// WaitUntilMultiplexRunningWithContext is an extended version of WaitUntilMultiplexRunning.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaLive) WaitUntilMultiplexRunningWithContext(ctx aws.Context, input *DescribeMultiplexInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilMultiplexRunning",
		MaxAttempts: 120,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "RUNNING",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "STARTING",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeMultiplexInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeMultiplexRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilMultiplexStopped uses the MediaLive API operation
// DescribeMultiplex to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *MediaLive) WaitUntilMultiplexStopped(input *DescribeMultiplexInput) error {
	return c.WaitUntilMultiplexStoppedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilMultiplexStoppedWithContext is an extended version of WaitUntilMultiplexStopped.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaLive) WaitUntilMultiplexStoppedWithContext(ctx aws.Context, input *DescribeMultiplexInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilMultiplexStopped",
		MaxAttempts: 28,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "IDLE",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "State",
				Expected: "STOPPING",
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 500,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeMultiplexInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeMultiplexRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
