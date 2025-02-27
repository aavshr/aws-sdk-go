//go:build go1.7
// +build go1.7

package finspace

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/awstesting/unit"
)

func TestClientContentType(t *testing.T) {
	sess := unit.Session.Copy()

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			contentType := r.Header.Get("Content-Type")
			if e, a := contentType, "application/x-amz-json-1.1"; !strings.EqualFold(e, a) {
				t.Errorf("expect %v content-type, got %v", e, a)
			}
		},
	))
	defer server.Close()

	client := New(sess, &aws.Config{Endpoint: &server.URL})
	_, err := client.CreateEnvironment(&CreateEnvironmentInput{
		Name: aws.String("fake name"),
	})
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
}
