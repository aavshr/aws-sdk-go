//go:build example
// +build example

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/arn"
	"github.com/aavshr/aws-sdk-go/aws/session"
	"github.com/aavshr/aws-sdk-go/service/s3"
	"github.com/aavshr/aws-sdk-go/service/s3control"
)

const (
	bucketName  = "myBucketName"
	keyName     = "myKeyName"
	accountID   = "123456789012"
	accessPoint = "accesspointname"

	// vpcBucketEndpoint will be used by the SDK to resolve an endpoint, when making a call to
	// access `bucket` data using s3 interface endpoint. This endpoint may be mutated by the SDK,
	// as per the input provided to work with ARNs.
	vpcBucketEndpoint = "https://bucket.vpce-0xxxxxxx-xxx8xxg.s3.us-west-2.vpce.amazonaws.com"

	// vpcAccesspointEndpoint will be used by the SDK to resolve an endpoint, when making a call to
	// access `access-point` data using s3 interface endpoint. This endpoint may be mutated by the SDK,
	// as per the input provided to work with ARNs.
	vpcAccesspointEndpoint = "https://accesspoint.vpce-0xxxxxxx-xxx8xxg.s3.us-west-2.vpce.amazonaws.com"

	// vpcControlEndpoint will be used by the SDK to resolve an endpoint, when making a call to
	// access `control` data using s3 interface endpoint. This endpoint may be mutated by the SDK,
	// as per the input provided to work with ARNs.
	vpcControlEndpoint = "https://control.vpce-0xxxxxxx-xxx8xxg.s3.us-west-2.vpce.amazonaws.com"
)

func main() {
	sess := session.Must(session.NewSession())

	s3BucketSvc := s3.New(sess, &aws.Config{
		Endpoint: aws.String(vpcBucketEndpoint),
	})

	s3AccesspointSvc := s3.New(sess, &aws.Config{
		Endpoint: aws.String(vpcAccesspointEndpoint),
	})

	s3ControlSvc := s3control.New(sess, &aws.Config{
		Endpoint: aws.String(vpcControlEndpoint),
	})

	// Create an S3 Bucket
	fmt.Println("create s3 bucket")
	_, err := s3BucketSvc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		panic(fmt.Errorf("failed to create bucket: %v", err))
	}

	// Wait for S3 Bucket to Exist
	fmt.Println("wait for s3 bucket to exist")
	err = s3BucketSvc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		panic(fmt.Sprintf("bucket failed to materialize: %v", err))
	}

	// Create an Access Point referring to the bucket
	fmt.Println("create an access point")
	_, err = s3ControlSvc.CreateAccessPoint(&s3control.CreateAccessPointInput{
		AccountId: aws.String(accountID),
		Bucket:    aws.String(bucketName),
		Name:      aws.String(accessPoint),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create access point: %v", err))
	}

	// Use the SDK's ARN builder to create an ARN for the Access Point.
	apARN := arn.ARN{
		Partition: "aws",
		Service:   "s3",
		Region:    aws.StringValue(sess.Config.Region),
		AccountID: accountID,
		Resource:  "accesspoint/" + accessPoint,
	}

	// And Use Access Point ARN where bucket parameters are accepted
	fmt.Println("get object using access point")
	getObjectOutput, err := s3AccesspointSvc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(apARN.String()),
		Key:    aws.String("somekey"),
	})
	if err != nil {
		panic(fmt.Sprintf("failed get object request: %v", err))
	}

	_, err = ioutil.ReadAll(getObjectOutput.Body)
	if err != nil {
		panic(fmt.Sprintf("failed to read object body: %v", err))
	}
}
