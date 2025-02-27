//go:build !go1.6
// +build !go1.6

package request_test

import (
	"errors"

	"github.com/aavshr/aws-sdk-go/aws/awserr"
)

var errTimeout = awserr.New("foo", "bar", errors.New("net/http: request canceled Timeout"))
