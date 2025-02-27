package corehandlers_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/awserr"
	"github.com/aavshr/aws-sdk-go/aws/client"
	"github.com/aavshr/aws-sdk-go/aws/client/metadata"
	"github.com/aavshr/aws-sdk-go/aws/corehandlers"
	"github.com/aavshr/aws-sdk-go/aws/credentials"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/awstesting"
	"github.com/aavshr/aws-sdk-go/awstesting/unit"
	"github.com/aavshr/aws-sdk-go/internal/sdktesting"
	"github.com/aavshr/aws-sdk-go/service/s3"
)

func TestValidateEndpointHandler(t *testing.T) {
	restoreEnvFn := sdktesting.StashEnv()
	defer restoreEnvFn()
	svc := awstesting.NewClient(aws.NewConfig().WithRegion("us-west-2"))
	svc.Handlers.Clear()
	svc.Handlers.Validate.PushBackNamed(corehandlers.ValidateEndpointHandler)

	req := svc.NewRequest(&request.Operation{Name: "Operation"}, nil, nil)
	err := req.Build()

	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestValidateEndpointHandlerErrorRegion(t *testing.T) {
	restoreEnvFn := sdktesting.StashEnv()
	defer restoreEnvFn()
	svc := awstesting.NewClient()
	svc.Handlers.Clear()
	svc.Handlers.Validate.PushBackNamed(corehandlers.ValidateEndpointHandler)

	req := svc.NewRequest(&request.Operation{Name: "Operation"}, nil, nil)
	err := req.Build()

	if err == nil {
		t.Errorf("expect error, got none")
	}
	if e, a := aws.ErrMissingRegion, err; e != a {
		t.Errorf("expect %v to be %v", e, a)
	}
}

type mockCredsProvider struct {
	expired        bool
	retrieveCalled bool
}

func (m *mockCredsProvider) Retrieve() (credentials.Value, error) {
	m.retrieveCalled = true
	return credentials.Value{
		AccessKeyID:     "AKID",
		SecretAccessKey: "SECRET",
		ProviderName:    "mockCredsProvider",
	}, nil
}

func (m *mockCredsProvider) IsExpired() bool {
	return m.expired
}

func TestAfterRetryRefreshCreds(t *testing.T) {
	restoreEnvFn := sdktesting.StashEnv()
	defer restoreEnvFn()

	credProvider := &mockCredsProvider{}

	sess := unit.Session.Copy(&aws.Config{
		Credentials: credentials.NewCredentials(credProvider),
		MaxRetries:  aws.Int(2),
	})
	clientInfo := metadata.ClientInfo{
		Endpoint:    "http://endpoint",
		SigningName: "",
	}
	svc := client.New(*sess.Config, clientInfo, sess.Handlers)

	svc.Handlers.Sign.PushBack(func(r *request.Request) {
		if !svc.Config.Credentials.IsExpired() {
			t.Errorf("expect credentials of of been expired before request attempt")
		}
		_, err := svc.Config.Credentials.Get()
		r.Error = err
	})

	var respID int
	resps := []struct {
		Resp *http.Response
		Err  error
	}{
		{
			Resp: &http.Response{
				StatusCode: 403,
				Header:     http.Header{},
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			},
			Err: awserr.New("ExpiredToken", "", nil),
		},
		{
			Resp: &http.Response{
				StatusCode: 403,
				Header:     http.Header{},
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			},
			Err: awserr.New("ExpiredToken", "", nil),
		},
		{
			Resp: &http.Response{
				StatusCode: 200,
				Header:     http.Header{},
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte{})),
			},
		},
	}
	svc.Handlers.Send.Clear()
	svc.Handlers.Send.PushBack(func(r *request.Request) {
		r.HTTPResponse = resps[respID].Resp
	})
	svc.Handlers.UnmarshalError.PushBack(func(r *request.Request) {
		r.Error = resps[respID].Err
	})
	svc.Handlers.CompleteAttempt.PushBack(func(r *request.Request) {
		respID++
	})

	if !svc.Config.Credentials.IsExpired() {
		t.Fatalf("expect to start out expired")
	}
	if credProvider.retrieveCalled {
		t.Fatalf("expect retrieve not yet called")
	}

	req := svc.NewRequest(&request.Operation{Name: "Operation"}, nil, nil)
	if err := req.Send(); err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if e, a := len(resps)-1, req.RetryCount; e != a {
		t.Errorf("expect %v retries, got %v", e, a)
	}
	if svc.Config.Credentials.IsExpired() {
		t.Errorf("expect credentials not to be expired")
	}
	if !credProvider.retrieveCalled {
		t.Errorf("expect retrieve to be called")
	}
}

func TestAfterRetryWithContextCanceled(t *testing.T) {
	c := awstesting.NewClient()

	req := c.NewRequest(&request.Operation{Name: "Operation"}, nil, nil)

	ctx := &awstesting.FakeContext{DoneCh: make(chan struct{})}
	req.SetContext(ctx)

	req.Error = fmt.Errorf("some error")
	req.Retryable = aws.Bool(true)
	req.HTTPResponse = &http.Response{
		StatusCode: 500,
	}

	close(ctx.DoneCh)
	ctx.Error = fmt.Errorf("context canceled")

	corehandlers.AfterRetryHandler.Fn(req)

	if req.Error == nil {
		t.Fatalf("expect error but didn't receive one")
	}

	aerr := req.Error.(awserr.Error)

	if e, a := request.CanceledErrorCode, aerr.Code(); e != a {
		t.Errorf("expect %q, error code got %q", e, a)
	}
}

func TestAfterRetryWithContext(t *testing.T) {
	c := awstesting.NewClient()

	req := c.NewRequest(&request.Operation{Name: "Operation"}, nil, nil)

	ctx := &awstesting.FakeContext{DoneCh: make(chan struct{})}
	req.SetContext(ctx)

	req.Error = fmt.Errorf("some error")
	req.Retryable = aws.Bool(true)
	req.HTTPResponse = &http.Response{
		StatusCode: 500,
	}

	corehandlers.AfterRetryHandler.Fn(req)

	if req.Error != nil {
		t.Fatalf("expect no error, got %v", req.Error)
	}
	if e, a := 1, req.RetryCount; e != a {
		t.Errorf("expect retry count to be %d, got %d", e, a)
	}
}

func TestSendWithContextCanceled(t *testing.T) {
	c := awstesting.NewClient(&aws.Config{
		SleepDelay: func(dur time.Duration) {
			t.Errorf("SleepDelay should not be called")
		},
	})

	req := c.NewRequest(&request.Operation{Name: "Operation"}, nil, nil)

	ctx := &awstesting.FakeContext{DoneCh: make(chan struct{})}
	req.SetContext(ctx)

	req.Error = fmt.Errorf("some error")
	req.Retryable = aws.Bool(true)
	req.HTTPResponse = &http.Response{
		StatusCode: 500,
	}

	close(ctx.DoneCh)
	ctx.Error = fmt.Errorf("context canceled")

	corehandlers.SendHandler.Fn(req)

	if req.Error == nil {
		t.Fatalf("expect error but didn't receive one")
	}

	aerr := req.Error.(awserr.Error)

	if e, a := request.CanceledErrorCode, aerr.Code(); e != a {
		t.Errorf("expect %q, error code got %q", e, a)
	}
}

type testSendHandlerTransport struct{}

func (t *testSendHandlerTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("mock error")
}

func TestSendHandlerError(t *testing.T) {
	svc := awstesting.NewClient(&aws.Config{
		HTTPClient: &http.Client{
			Transport: &testSendHandlerTransport{},
		},
	})
	svc.Handlers.Clear()
	svc.Handlers.Send.PushBackNamed(corehandlers.SendHandler)
	r := svc.NewRequest(&request.Operation{Name: "Operation"}, nil, nil)

	r.Send()

	if r.Error == nil {
		t.Errorf("expect error, got none")
	}
	if r.HTTPResponse == nil {
		t.Errorf("expect response, got none")
	}
}

func TestSendWithoutFollowRedirects(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/original":
			w.Header().Set("Location", "/redirected")
			w.WriteHeader(301)
		case "/redirected":
			t.Fatalf("expect not to redirect, but was")
		}
	}))
	defer server.Close()

	svc := awstesting.NewClient(&aws.Config{
		DisableSSL: aws.Bool(true),
		Endpoint:   aws.String(server.URL),
	})
	svc.Handlers.Clear()
	svc.Handlers.Send.PushBackNamed(corehandlers.SendHandler)

	r := svc.NewRequest(&request.Operation{
		Name:     "Operation",
		HTTPPath: "/original",
	}, nil, nil)
	r.DisableFollowRedirects = true

	err := r.Send()
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := 301, r.HTTPResponse.StatusCode; e != a {
		t.Errorf("expect %d status code, got %d", e, a)
	}
}

func TestValidateReqSigHandler(t *testing.T) {
	cases := []struct {
		Req    *request.Request
		Resign bool
	}{
		{
			Req: &request.Request{
				Config: aws.Config{Credentials: credentials.AnonymousCredentials},
				Time:   time.Now().Add(-15 * time.Minute),
			},
			Resign: false,
		},
		{
			Req: &request.Request{
				Time: time.Now().Add(-15 * time.Minute),
			},
			Resign: true,
		},
		{
			Req: &request.Request{
				Time: time.Now().Add(-1 * time.Minute),
			},
			Resign: false,
		},
	}

	for i, c := range cases {
		c.Req.HTTPRequest = &http.Request{URL: &url.URL{}}

		resigned := false
		c.Req.Handlers.Sign.PushBack(func(r *request.Request) {
			resigned = true
		})

		corehandlers.ValidateReqSigHandler.Fn(c.Req)

		if c.Req.Error != nil {
			t.Errorf("expect no error, got %v", c.Req.Error)
		}
		if e, a := c.Resign, resigned; e != a {
			t.Errorf("%d, expect %v to be %v", i, e, a)
		}
	}
}

func setupContentLengthTestServer(t *testing.T, hasContentLength bool, contentLength int64) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := r.Header["Content-Length"]
		if e, a := hasContentLength, ok; e != a {
			t.Errorf("expect %v to be %v", e, a)
		}
		if hasContentLength {
			if e, a := contentLength, r.ContentLength; e != a {
				t.Errorf("expect %v to be %v", e, a)
			}
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}
		r.Body.Close()

		authHeader := r.Header.Get("Authorization")
		if hasContentLength {
			if e, a := "content-length", authHeader; !strings.Contains(a, e) {
				t.Errorf("expect %v to be in %v", e, a)
			}
		} else {
			if e, a := "content-length", authHeader; strings.Contains(a, e) {
				t.Errorf("expect %v to not be in %v", e, a)
			}
		}

		if e, a := contentLength, int64(len(b)); e != a {
			t.Errorf("expect %v to be %v", e, a)
		}
	}))

	return server
}

func TestBuildContentLength_ZeroBody(t *testing.T) {
	server := setupContentLengthTestServer(t, false, 0)
	defer server.Close()

	svc := s3.New(unit.Session, &aws.Config{
		Endpoint:         aws.String(server.URL),
		S3ForcePathStyle: aws.Bool(true),
		DisableSSL:       aws.Bool(true),
	})
	_, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("bucketname"),
		Key:    aws.String("keyname"),
	})

	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestBuildContentLength_NegativeBody(t *testing.T) {
	server := setupContentLengthTestServer(t, false, 0)
	defer server.Close()

	svc := s3.New(unit.Session, &aws.Config{
		Endpoint:         aws.String(server.URL),
		S3ForcePathStyle: aws.Bool(true),
		DisableSSL:       aws.Bool(true),
	})
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("bucketname"),
		Key:    aws.String("keyname"),
	})

	req.HTTPRequest.Header.Set("Content-Length", "-1")

	if req.Error != nil {
		t.Errorf("expect no error, got %v", req.Error)
	}
}

func TestBuildContentLength_WithBody(t *testing.T) {
	server := setupContentLengthTestServer(t, true, 1024)
	defer server.Close()

	svc := s3.New(unit.Session, &aws.Config{
		Endpoint:         aws.String(server.URL),
		S3ForcePathStyle: aws.Bool(true),
		DisableSSL:       aws.Bool(true),
	})
	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("bucketname"),
		Key:    aws.String("keyname"),
		Body:   bytes.NewReader(make([]byte, 1024)),
	})

	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
