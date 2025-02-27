//go:build go1.8
// +build go1.8

package query

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/aavshr/aws-sdk-go/aws/awserr"
	"github.com/aavshr/aws-sdk-go/aws/request"
)

func TestUnmarshalError(t *testing.T) {
	cases := map[string]struct {
		Request   *request.Request
		Code, Msg string
		ReqID     string
		Status    int
	}{
		"ErrorResponse": {
			Request: &request.Request{
				HTTPResponse: &http.Response{
					StatusCode: 400,
					Header:     http.Header{},
					Body: ioutil.NopCloser(strings.NewReader(
						`<ErrorResponse>
							<Error>
								<Code>codeAbc</Code><Message>msg123</Message>
							</Error>
							<RequestId>reqID123</RequestId>
						</ErrorResponse>`)),
				},
			},
			Code: "codeAbc", Msg: "msg123",
			Status: 400, ReqID: "reqID123",
		},
		"ServiceUnavailableException": {
			Request: &request.Request{
				HTTPResponse: &http.Response{
					StatusCode: 502,
					Header:     http.Header{},
					Body: ioutil.NopCloser(strings.NewReader(
						`<ServiceUnavailableException>
							<Something>else</Something>
						</ServiceUnavailableException>`)),
				},
			},
			Code:   "ServiceUnavailableException",
			Msg:    "service is unavailable",
			Status: 502,
		},
		"unknown tag": {
			Request: &request.Request{
				HTTPResponse: &http.Response{
					StatusCode: 400,
					Header:     http.Header{},
					Body: ioutil.NopCloser(strings.NewReader(
						`<Hello>
							<World>.</World>
						</Hello>`)),
				},
			},
			Code:   request.ErrCodeSerialization,
			Msg:    "failed to unmarshal error message",
			Status: 400,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			r := c.Request
			UnmarshalError(r)
			if r.Error == nil {
				t.Fatalf("expect error, got none")
			}

			aerr := r.Error.(awserr.RequestFailure)
			if e, a := c.Code, aerr.Code(); e != a {
				t.Errorf("expect %v code, got %v", e, a)
			}
			if e, a := c.Msg, aerr.Message(); e != a {
				t.Errorf("expect %q message, got %q", e, a)
			}
			if e, a := c.ReqID, aerr.RequestID(); e != a {
				t.Errorf("expect %v request ID, got %v", e, a)
			}
			if e, a := c.Status, aerr.StatusCode(); e != a {
				t.Errorf("expect %v status code, got %v", e, a)
			}
		})
	}
}
