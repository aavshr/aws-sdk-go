package sqs

import "github.com/aavshr/aws-sdk-go/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		setupChecksumValidation(r)
	}
}
