package finspace

import (
	"strings"

	"github.com/aavshr/aws-sdk-go/aws/client"
	"github.com/aavshr/aws-sdk-go/aws/request"
)

func init() {
	initClient = func(c *client.Client) {
		c.Handlers.Build.PushBack(func(r *request.Request) {
			if strings.EqualFold(r.HTTPRequest.Header.Get("Content-Type"), "application/json") {
				r.HTTPRequest.Header.Set("Content-Type", "application/x-amz-json-1.1")
			}
		})
	}
}
