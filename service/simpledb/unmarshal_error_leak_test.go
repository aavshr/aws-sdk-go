package simpledb

import (
	"net/http"
	"testing"

	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/awstesting"
)

func TestUnmarhsalErrorLeak(t *testing.T) {
	req := &request.Request{
		HTTPRequest: &http.Request{
			Header: make(http.Header),
			Body:   &awstesting.ReadCloser{Size: 2048},
		},
	}
	req.HTTPResponse = &http.Response{
		Body: &awstesting.ReadCloser{Size: 2048},
		Header: http.Header{
			"X-Amzn-Requestid": []string{"1"},
		},
		StatusCode: http.StatusOK,
	}

	reader := req.HTTPResponse.Body.(*awstesting.ReadCloser)
	unmarshalError(req)

	if req.Error == nil {
		t.Errorf("expect error, got nil")
	}
	if !reader.Closed {
		t.Errorf("expect closed, was not")
	}
	if e, a := 0, reader.Size; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}
