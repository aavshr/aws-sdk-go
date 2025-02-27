// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacecommerceanalytics

import (
	"fmt"
	"time"

	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/awsutil"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const opGenerateDataSet = "GenerateDataSet"

// GenerateDataSetRequest generates a "aws/request.Request" representing the
// client's request for the GenerateDataSet operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GenerateDataSet for more information on using the GenerateDataSet
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GenerateDataSetRequest method.
//    req, resp := client.GenerateDataSetRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/marketplacecommerceanalytics-2015-07-01/GenerateDataSet
func (c *MarketplaceCommerceAnalytics) GenerateDataSetRequest(input *GenerateDataSetInput) (req *request.Request, output *GenerateDataSetOutput) {
	op := &request.Operation{
		Name:       opGenerateDataSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GenerateDataSetInput{}
	}

	output = &GenerateDataSetOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GenerateDataSet API operation for AWS Marketplace Commerce Analytics.
//
// Given a data set type and data set publication date, asynchronously publishes
// the requested data set to the specified S3 bucket and notifies the specified
// SNS topic once the data is available. Returns a unique request identifier
// that can be used to correlate requests with notifications from the SNS topic.
// Data sets will be published in comma-separated values (CSV) format with the
// file name {data_set_type}_YYYY-MM-DD.csv. If a file with the same name already
// exists (e.g. if the same data set is requested twice), the original file
// will be overwritten by the new file. Requires a Role with an attached permissions
// policy providing Allow permissions for the following actions: s3:PutObject,
// s3:GetBucketLocation, sns:GetTopicAttributes, sns:Publish, iam:GetRolePolicy.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Marketplace Commerce Analytics's
// API operation GenerateDataSet for usage and error information.
//
// Returned Error Types:
//   * Exception
//   This exception is thrown when an internal service error occurs.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/marketplacecommerceanalytics-2015-07-01/GenerateDataSet
func (c *MarketplaceCommerceAnalytics) GenerateDataSet(input *GenerateDataSetInput) (*GenerateDataSetOutput, error) {
	req, out := c.GenerateDataSetRequest(input)
	return out, req.Send()
}

// GenerateDataSetWithContext is the same as GenerateDataSet with the addition of
// the ability to pass a context and additional request options.
//
// See GenerateDataSet for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MarketplaceCommerceAnalytics) GenerateDataSetWithContext(ctx aws.Context, input *GenerateDataSetInput, opts ...request.Option) (*GenerateDataSetOutput, error) {
	req, out := c.GenerateDataSetRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStartSupportDataExport = "StartSupportDataExport"

// StartSupportDataExportRequest generates a "aws/request.Request" representing the
// client's request for the StartSupportDataExport operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See StartSupportDataExport for more information on using the StartSupportDataExport
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the StartSupportDataExportRequest method.
//    req, resp := client.StartSupportDataExportRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/marketplacecommerceanalytics-2015-07-01/StartSupportDataExport
func (c *MarketplaceCommerceAnalytics) StartSupportDataExportRequest(input *StartSupportDataExportInput) (req *request.Request, output *StartSupportDataExportOutput) {
	op := &request.Operation{
		Name:       opStartSupportDataExport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartSupportDataExportInput{}
	}

	output = &StartSupportDataExportOutput{}
	req = c.newRequest(op, input, output)
	return
}

// StartSupportDataExport API operation for AWS Marketplace Commerce Analytics.
//
// Given a data set type and a from date, asynchronously publishes the requested
// customer support data to the specified S3 bucket and notifies the specified
// SNS topic once the data is available. Returns a unique request identifier
// that can be used to correlate requests with notifications from the SNS topic.
// Data sets will be published in comma-separated values (CSV) format with the
// file name {data_set_type}_YYYY-MM-DD'T'HH-mm-ss'Z'.csv. If a file with the
// same name already exists (e.g. if the same data set is requested twice),
// the original file will be overwritten by the new file. Requires a Role with
// an attached permissions policy providing Allow permissions for the following
// actions: s3:PutObject, s3:GetBucketLocation, sns:GetTopicAttributes, sns:Publish,
// iam:GetRolePolicy.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Marketplace Commerce Analytics's
// API operation StartSupportDataExport for usage and error information.
//
// Returned Error Types:
//   * Exception
//   This exception is thrown when an internal service error occurs.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/marketplacecommerceanalytics-2015-07-01/StartSupportDataExport
func (c *MarketplaceCommerceAnalytics) StartSupportDataExport(input *StartSupportDataExportInput) (*StartSupportDataExportOutput, error) {
	req, out := c.StartSupportDataExportRequest(input)
	return out, req.Send()
}

// StartSupportDataExportWithContext is the same as StartSupportDataExport with the addition of
// the ability to pass a context and additional request options.
//
// See StartSupportDataExport for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MarketplaceCommerceAnalytics) StartSupportDataExportWithContext(ctx aws.Context, input *StartSupportDataExportInput, opts ...request.Option) (*StartSupportDataExportOutput, error) {
	req, out := c.StartSupportDataExportRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// This exception is thrown when an internal service error occurs.
type Exception struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	// This message describes details of the error.
	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Exception) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Exception) GoString() string {
	return s.String()
}

func newErrorException(v protocol.ResponseMetadata) error {
	return &Exception{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *Exception) Code() string {
	return "MarketplaceCommerceAnalyticsException"
}

// Message returns the exception's message.
func (s *Exception) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *Exception) OrigErr() error {
	return nil
}

func (s *Exception) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *Exception) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *Exception) RequestID() string {
	return s.RespMetadata.RequestID
}

// Container for the parameters to the GenerateDataSet operation.
type GenerateDataSetInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Key-value pairs which will be returned, unmodified, in the Amazon
	// SNS notification message and the data set metadata file. These key-value
	// pairs can be used to correlated responses with tracking information from
	// other systems.
	CustomerDefinedValues map[string]*string `locationName:"customerDefinedValues" min:"1" type:"map"`

	// The date a data set was published. For daily data sets, provide a date with
	// day-level granularity for the desired day. For monthly data sets except those
	// with prefix disbursed_amount, provide a date with month-level granularity
	// for the desired month (the day value will be ignored). For data sets with
	// prefix disbursed_amount, provide a date with day-level granularity for the
	// desired day. For these data sets we will look backwards in time over the
	// range of 31 days until the first data set is found (the latest one).
	//
	// DataSetPublicationDate is a required field
	DataSetPublicationDate *time.Time `locationName:"dataSetPublicationDate" type:"timestamp" required:"true"`

	// The desired data set type.
	//
	//    * customer_subscriber_hourly_monthly_subscriptions From 2017-09-15 to
	//    present: Available daily by 24:00 UTC.
	//
	//    * customer_subscriber_annual_subscriptions From 2017-09-15 to present:
	//    Available daily by 24:00 UTC.
	//
	//    * daily_business_usage_by_instance_type From 2017-09-15 to present: Available
	//    daily by 24:00 UTC.
	//
	//    * daily_business_fees From 2017-09-15 to present: Available daily by 24:00
	//    UTC.
	//
	//    * daily_business_free_trial_conversions From 2017-09-15 to present: Available
	//    daily by 24:00 UTC.
	//
	//    * daily_business_new_instances From 2017-09-15 to present: Available daily
	//    by 24:00 UTC.
	//
	//    * daily_business_new_product_subscribers From 2017-09-15 to present: Available
	//    daily by 24:00 UTC.
	//
	//    * daily_business_canceled_product_subscribers From 2017-09-15 to present:
	//    Available daily by 24:00 UTC.
	//
	//    * monthly_revenue_billing_and_revenue_data From 2017-09-15 to present:
	//    Available monthly on the 15th day of the month by 24:00 UTC. Data includes
	//    metered transactions (e.g. hourly) from one month prior.
	//
	//    * monthly_revenue_annual_subscriptions From 2017-09-15 to present: Available
	//    monthly on the 15th day of the month by 24:00 UTC. Data includes up-front
	//    software charges (e.g. annual) from one month prior.
	//
	//    * monthly_revenue_field_demonstration_usage From 2018-03-15 to present:
	//    Available monthly on the 15th day of the month by 24:00 UTC.
	//
	//    * monthly_revenue_flexible_payment_schedule From 2018-11-15 to present:
	//    Available monthly on the 15th day of the month by 24:00 UTC.
	//
	//    * disbursed_amount_by_product From 2017-09-15 to present: Available every
	//    30 days by 24:00 UTC.
	//
	//    * disbursed_amount_by_instance_hours From 2017-09-15 to present: Available
	//    every 30 days by 24:00 UTC.
	//
	//    * disbursed_amount_by_customer_geo From 2017-09-15 to present: Available
	//    every 30 days by 24:00 UTC.
	//
	//    * disbursed_amount_by_age_of_uncollected_funds From 2017-09-15 to present:
	//    Available every 30 days by 24:00 UTC.
	//
	//    * disbursed_amount_by_age_of_disbursed_funds From 2017-09-15 to present:
	//    Available every 30 days by 24:00 UTC.
	//
	//    * disbursed_amount_by_age_of_past_due_funds From 2018-04-07 to present:
	//    Available every 30 days by 24:00 UTC.
	//
	//    * disbursed_amount_by_uncollected_funds_breakdown From 2019-10-04 to present:
	//    Available every 30 days by 24:00 UTC.
	//
	//    * sales_compensation_billed_revenue From 2017-09-15 to present: Available
	//    monthly on the 15th day of the month by 24:00 UTC. Data includes metered
	//    transactions (e.g. hourly) from one month prior, and up-front software
	//    charges (e.g. annual) from one month prior.
	//
	//    * us_sales_and_use_tax_records From 2017-09-15 to present: Available monthly
	//    on the 15th day of the month by 24:00 UTC.
	//
	//    * disbursed_amount_by_product_with_uncollected_funds This data set is
	//    deprecated. Download related reports from AMMP instead!
	//
	//    * customer_profile_by_industry This data set is deprecated. Download related
	//    reports from AMMP instead!
	//
	//    * customer_profile_by_revenue This data set is deprecated. Download related
	//    reports from AMMP instead!
	//
	//    * customer_profile_by_geography This data set is deprecated. Download
	//    related reports from AMMP instead!
	//
	// DataSetType is a required field
	DataSetType *string `locationName:"dataSetType" min:"1" type:"string" required:"true" enum:"DataSetType"`

	// The name (friendly name, not ARN) of the destination S3 bucket.
	//
	// DestinationS3BucketName is a required field
	DestinationS3BucketName *string `locationName:"destinationS3BucketName" min:"1" type:"string" required:"true"`

	// (Optional) The desired S3 prefix for the published data set, similar to a
	// directory path in standard file systems. For example, if given the bucket
	// name "mybucket" and the prefix "myprefix/mydatasets", the output file "outputfile"
	// would be published to "s3://mybucket/myprefix/mydatasets/outputfile". If
	// the prefix directory structure does not exist, it will be created. If no
	// prefix is provided, the data set will be published to the S3 bucket root.
	DestinationS3Prefix *string `locationName:"destinationS3Prefix" type:"string"`

	// The Amazon Resource Name (ARN) of the Role with an attached permissions policy
	// to interact with the provided AWS services.
	//
	// RoleNameArn is a required field
	RoleNameArn *string `locationName:"roleNameArn" min:"1" type:"string" required:"true"`

	// Amazon Resource Name (ARN) for the SNS Topic that will be notified when the
	// data set has been published or if an error has occurred.
	//
	// SnsTopicArn is a required field
	SnsTopicArn *string `locationName:"snsTopicArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GenerateDataSetInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GenerateDataSetInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GenerateDataSetInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GenerateDataSetInput"}
	if s.CustomerDefinedValues != nil && len(s.CustomerDefinedValues) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CustomerDefinedValues", 1))
	}
	if s.DataSetPublicationDate == nil {
		invalidParams.Add(request.NewErrParamRequired("DataSetPublicationDate"))
	}
	if s.DataSetType == nil {
		invalidParams.Add(request.NewErrParamRequired("DataSetType"))
	}
	if s.DataSetType != nil && len(*s.DataSetType) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("DataSetType", 1))
	}
	if s.DestinationS3BucketName == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationS3BucketName"))
	}
	if s.DestinationS3BucketName != nil && len(*s.DestinationS3BucketName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("DestinationS3BucketName", 1))
	}
	if s.RoleNameArn == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleNameArn"))
	}
	if s.RoleNameArn != nil && len(*s.RoleNameArn) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("RoleNameArn", 1))
	}
	if s.SnsTopicArn == nil {
		invalidParams.Add(request.NewErrParamRequired("SnsTopicArn"))
	}
	if s.SnsTopicArn != nil && len(*s.SnsTopicArn) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("SnsTopicArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCustomerDefinedValues sets the CustomerDefinedValues field's value.
func (s *GenerateDataSetInput) SetCustomerDefinedValues(v map[string]*string) *GenerateDataSetInput {
	s.CustomerDefinedValues = v
	return s
}

// SetDataSetPublicationDate sets the DataSetPublicationDate field's value.
func (s *GenerateDataSetInput) SetDataSetPublicationDate(v time.Time) *GenerateDataSetInput {
	s.DataSetPublicationDate = &v
	return s
}

// SetDataSetType sets the DataSetType field's value.
func (s *GenerateDataSetInput) SetDataSetType(v string) *GenerateDataSetInput {
	s.DataSetType = &v
	return s
}

// SetDestinationS3BucketName sets the DestinationS3BucketName field's value.
func (s *GenerateDataSetInput) SetDestinationS3BucketName(v string) *GenerateDataSetInput {
	s.DestinationS3BucketName = &v
	return s
}

// SetDestinationS3Prefix sets the DestinationS3Prefix field's value.
func (s *GenerateDataSetInput) SetDestinationS3Prefix(v string) *GenerateDataSetInput {
	s.DestinationS3Prefix = &v
	return s
}

// SetRoleNameArn sets the RoleNameArn field's value.
func (s *GenerateDataSetInput) SetRoleNameArn(v string) *GenerateDataSetInput {
	s.RoleNameArn = &v
	return s
}

// SetSnsTopicArn sets the SnsTopicArn field's value.
func (s *GenerateDataSetInput) SetSnsTopicArn(v string) *GenerateDataSetInput {
	s.SnsTopicArn = &v
	return s
}

// Container for the result of the GenerateDataSet operation.
type GenerateDataSetOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier representing a specific request to the GenerateDataSet
	// operation. This identifier can be used to correlate a request with notifications
	// from the SNS topic.
	DataSetRequestId *string `locationName:"dataSetRequestId" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GenerateDataSetOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GenerateDataSetOutput) GoString() string {
	return s.String()
}

// SetDataSetRequestId sets the DataSetRequestId field's value.
func (s *GenerateDataSetOutput) SetDataSetRequestId(v string) *GenerateDataSetOutput {
	s.DataSetRequestId = &v
	return s
}

// Container for the parameters to the StartSupportDataExport operation.
type StartSupportDataExportInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Key-value pairs which will be returned, unmodified, in the Amazon
	// SNS notification message and the data set metadata file.
	CustomerDefinedValues map[string]*string `locationName:"customerDefinedValues" min:"1" type:"map"`

	// Specifies the data set type to be written to the output csv file. The data
	// set types customer_support_contacts_data and test_customer_support_contacts_data
	// both result in a csv file containing the following fields: Product Id, Product
	// Code, Customer Guid, Subscription Guid, Subscription Start Date, Organization,
	// AWS Account Id, Given Name, Surname, Telephone Number, Email, Title, Country
	// Code, ZIP Code, Operation Type, and Operation Time.
	//
	//    * customer_support_contacts_data Customer support contact data. The data
	//    set will contain all changes (Creates, Updates, and Deletes) to customer
	//    support contact data from the date specified in the from_date parameter.
	//
	//    * test_customer_support_contacts_data An example data set containing static
	//    test data in the same format as customer_support_contacts_data
	//
	// DataSetType is a required field
	DataSetType *string `locationName:"dataSetType" min:"1" type:"string" required:"true" enum:"SupportDataSetType"`

	// The name (friendly name, not ARN) of the destination S3 bucket.
	//
	// DestinationS3BucketName is a required field
	DestinationS3BucketName *string `locationName:"destinationS3BucketName" min:"1" type:"string" required:"true"`

	// (Optional) The desired S3 prefix for the published data set, similar to a
	// directory path in standard file systems. For example, if given the bucket
	// name "mybucket" and the prefix "myprefix/mydatasets", the output file "outputfile"
	// would be published to "s3://mybucket/myprefix/mydatasets/outputfile". If
	// the prefix directory structure does not exist, it will be created. If no
	// prefix is provided, the data set will be published to the S3 bucket root.
	DestinationS3Prefix *string `locationName:"destinationS3Prefix" type:"string"`

	// The start date from which to retrieve the data set in UTC. This parameter
	// only affects the customer_support_contacts_data data set type.
	//
	// FromDate is a required field
	FromDate *time.Time `locationName:"fromDate" type:"timestamp" required:"true"`

	// The Amazon Resource Name (ARN) of the Role with an attached permissions policy
	// to interact with the provided AWS services.
	//
	// RoleNameArn is a required field
	RoleNameArn *string `locationName:"roleNameArn" min:"1" type:"string" required:"true"`

	// Amazon Resource Name (ARN) for the SNS Topic that will be notified when the
	// data set has been published or if an error has occurred.
	//
	// SnsTopicArn is a required field
	SnsTopicArn *string `locationName:"snsTopicArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s StartSupportDataExportInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s StartSupportDataExportInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSupportDataExportInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StartSupportDataExportInput"}
	if s.CustomerDefinedValues != nil && len(s.CustomerDefinedValues) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CustomerDefinedValues", 1))
	}
	if s.DataSetType == nil {
		invalidParams.Add(request.NewErrParamRequired("DataSetType"))
	}
	if s.DataSetType != nil && len(*s.DataSetType) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("DataSetType", 1))
	}
	if s.DestinationS3BucketName == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationS3BucketName"))
	}
	if s.DestinationS3BucketName != nil && len(*s.DestinationS3BucketName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("DestinationS3BucketName", 1))
	}
	if s.FromDate == nil {
		invalidParams.Add(request.NewErrParamRequired("FromDate"))
	}
	if s.RoleNameArn == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleNameArn"))
	}
	if s.RoleNameArn != nil && len(*s.RoleNameArn) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("RoleNameArn", 1))
	}
	if s.SnsTopicArn == nil {
		invalidParams.Add(request.NewErrParamRequired("SnsTopicArn"))
	}
	if s.SnsTopicArn != nil && len(*s.SnsTopicArn) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("SnsTopicArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCustomerDefinedValues sets the CustomerDefinedValues field's value.
func (s *StartSupportDataExportInput) SetCustomerDefinedValues(v map[string]*string) *StartSupportDataExportInput {
	s.CustomerDefinedValues = v
	return s
}

// SetDataSetType sets the DataSetType field's value.
func (s *StartSupportDataExportInput) SetDataSetType(v string) *StartSupportDataExportInput {
	s.DataSetType = &v
	return s
}

// SetDestinationS3BucketName sets the DestinationS3BucketName field's value.
func (s *StartSupportDataExportInput) SetDestinationS3BucketName(v string) *StartSupportDataExportInput {
	s.DestinationS3BucketName = &v
	return s
}

// SetDestinationS3Prefix sets the DestinationS3Prefix field's value.
func (s *StartSupportDataExportInput) SetDestinationS3Prefix(v string) *StartSupportDataExportInput {
	s.DestinationS3Prefix = &v
	return s
}

// SetFromDate sets the FromDate field's value.
func (s *StartSupportDataExportInput) SetFromDate(v time.Time) *StartSupportDataExportInput {
	s.FromDate = &v
	return s
}

// SetRoleNameArn sets the RoleNameArn field's value.
func (s *StartSupportDataExportInput) SetRoleNameArn(v string) *StartSupportDataExportInput {
	s.RoleNameArn = &v
	return s
}

// SetSnsTopicArn sets the SnsTopicArn field's value.
func (s *StartSupportDataExportInput) SetSnsTopicArn(v string) *StartSupportDataExportInput {
	s.SnsTopicArn = &v
	return s
}

// Container for the result of the StartSupportDataExport operation.
type StartSupportDataExportOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier representing a specific request to the StartSupportDataExport
	// operation. This identifier can be used to correlate a request with notifications
	// from the SNS topic.
	DataSetRequestId *string `locationName:"dataSetRequestId" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s StartSupportDataExportOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s StartSupportDataExportOutput) GoString() string {
	return s.String()
}

// SetDataSetRequestId sets the DataSetRequestId field's value.
func (s *StartSupportDataExportOutput) SetDataSetRequestId(v string) *StartSupportDataExportOutput {
	s.DataSetRequestId = &v
	return s
}

const (
	// DataSetTypeCustomerSubscriberHourlyMonthlySubscriptions is a DataSetType enum value
	DataSetTypeCustomerSubscriberHourlyMonthlySubscriptions = "customer_subscriber_hourly_monthly_subscriptions"

	// DataSetTypeCustomerSubscriberAnnualSubscriptions is a DataSetType enum value
	DataSetTypeCustomerSubscriberAnnualSubscriptions = "customer_subscriber_annual_subscriptions"

	// DataSetTypeDailyBusinessUsageByInstanceType is a DataSetType enum value
	DataSetTypeDailyBusinessUsageByInstanceType = "daily_business_usage_by_instance_type"

	// DataSetTypeDailyBusinessFees is a DataSetType enum value
	DataSetTypeDailyBusinessFees = "daily_business_fees"

	// DataSetTypeDailyBusinessFreeTrialConversions is a DataSetType enum value
	DataSetTypeDailyBusinessFreeTrialConversions = "daily_business_free_trial_conversions"

	// DataSetTypeDailyBusinessNewInstances is a DataSetType enum value
	DataSetTypeDailyBusinessNewInstances = "daily_business_new_instances"

	// DataSetTypeDailyBusinessNewProductSubscribers is a DataSetType enum value
	DataSetTypeDailyBusinessNewProductSubscribers = "daily_business_new_product_subscribers"

	// DataSetTypeDailyBusinessCanceledProductSubscribers is a DataSetType enum value
	DataSetTypeDailyBusinessCanceledProductSubscribers = "daily_business_canceled_product_subscribers"

	// DataSetTypeMonthlyRevenueBillingAndRevenueData is a DataSetType enum value
	DataSetTypeMonthlyRevenueBillingAndRevenueData = "monthly_revenue_billing_and_revenue_data"

	// DataSetTypeMonthlyRevenueAnnualSubscriptions is a DataSetType enum value
	DataSetTypeMonthlyRevenueAnnualSubscriptions = "monthly_revenue_annual_subscriptions"

	// DataSetTypeMonthlyRevenueFieldDemonstrationUsage is a DataSetType enum value
	DataSetTypeMonthlyRevenueFieldDemonstrationUsage = "monthly_revenue_field_demonstration_usage"

	// DataSetTypeMonthlyRevenueFlexiblePaymentSchedule is a DataSetType enum value
	DataSetTypeMonthlyRevenueFlexiblePaymentSchedule = "monthly_revenue_flexible_payment_schedule"

	// DataSetTypeDisbursedAmountByProduct is a DataSetType enum value
	DataSetTypeDisbursedAmountByProduct = "disbursed_amount_by_product"

	// DataSetTypeDisbursedAmountByProductWithUncollectedFunds is a DataSetType enum value
	DataSetTypeDisbursedAmountByProductWithUncollectedFunds = "disbursed_amount_by_product_with_uncollected_funds"

	// DataSetTypeDisbursedAmountByInstanceHours is a DataSetType enum value
	DataSetTypeDisbursedAmountByInstanceHours = "disbursed_amount_by_instance_hours"

	// DataSetTypeDisbursedAmountByCustomerGeo is a DataSetType enum value
	DataSetTypeDisbursedAmountByCustomerGeo = "disbursed_amount_by_customer_geo"

	// DataSetTypeDisbursedAmountByAgeOfUncollectedFunds is a DataSetType enum value
	DataSetTypeDisbursedAmountByAgeOfUncollectedFunds = "disbursed_amount_by_age_of_uncollected_funds"

	// DataSetTypeDisbursedAmountByAgeOfDisbursedFunds is a DataSetType enum value
	DataSetTypeDisbursedAmountByAgeOfDisbursedFunds = "disbursed_amount_by_age_of_disbursed_funds"

	// DataSetTypeDisbursedAmountByAgeOfPastDueFunds is a DataSetType enum value
	DataSetTypeDisbursedAmountByAgeOfPastDueFunds = "disbursed_amount_by_age_of_past_due_funds"

	// DataSetTypeDisbursedAmountByUncollectedFundsBreakdown is a DataSetType enum value
	DataSetTypeDisbursedAmountByUncollectedFundsBreakdown = "disbursed_amount_by_uncollected_funds_breakdown"

	// DataSetTypeCustomerProfileByIndustry is a DataSetType enum value
	DataSetTypeCustomerProfileByIndustry = "customer_profile_by_industry"

	// DataSetTypeCustomerProfileByRevenue is a DataSetType enum value
	DataSetTypeCustomerProfileByRevenue = "customer_profile_by_revenue"

	// DataSetTypeCustomerProfileByGeography is a DataSetType enum value
	DataSetTypeCustomerProfileByGeography = "customer_profile_by_geography"

	// DataSetTypeSalesCompensationBilledRevenue is a DataSetType enum value
	DataSetTypeSalesCompensationBilledRevenue = "sales_compensation_billed_revenue"

	// DataSetTypeUsSalesAndUseTaxRecords is a DataSetType enum value
	DataSetTypeUsSalesAndUseTaxRecords = "us_sales_and_use_tax_records"
)

// DataSetType_Values returns all elements of the DataSetType enum
func DataSetType_Values() []string {
	return []string{
		DataSetTypeCustomerSubscriberHourlyMonthlySubscriptions,
		DataSetTypeCustomerSubscriberAnnualSubscriptions,
		DataSetTypeDailyBusinessUsageByInstanceType,
		DataSetTypeDailyBusinessFees,
		DataSetTypeDailyBusinessFreeTrialConversions,
		DataSetTypeDailyBusinessNewInstances,
		DataSetTypeDailyBusinessNewProductSubscribers,
		DataSetTypeDailyBusinessCanceledProductSubscribers,
		DataSetTypeMonthlyRevenueBillingAndRevenueData,
		DataSetTypeMonthlyRevenueAnnualSubscriptions,
		DataSetTypeMonthlyRevenueFieldDemonstrationUsage,
		DataSetTypeMonthlyRevenueFlexiblePaymentSchedule,
		DataSetTypeDisbursedAmountByProduct,
		DataSetTypeDisbursedAmountByProductWithUncollectedFunds,
		DataSetTypeDisbursedAmountByInstanceHours,
		DataSetTypeDisbursedAmountByCustomerGeo,
		DataSetTypeDisbursedAmountByAgeOfUncollectedFunds,
		DataSetTypeDisbursedAmountByAgeOfDisbursedFunds,
		DataSetTypeDisbursedAmountByAgeOfPastDueFunds,
		DataSetTypeDisbursedAmountByUncollectedFundsBreakdown,
		DataSetTypeCustomerProfileByIndustry,
		DataSetTypeCustomerProfileByRevenue,
		DataSetTypeCustomerProfileByGeography,
		DataSetTypeSalesCompensationBilledRevenue,
		DataSetTypeUsSalesAndUseTaxRecords,
	}
}

const (
	// SupportDataSetTypeCustomerSupportContactsData is a SupportDataSetType enum value
	SupportDataSetTypeCustomerSupportContactsData = "customer_support_contacts_data"

	// SupportDataSetTypeTestCustomerSupportContactsData is a SupportDataSetType enum value
	SupportDataSetTypeTestCustomerSupportContactsData = "test_customer_support_contacts_data"
)

// SupportDataSetType_Values returns all elements of the SupportDataSetType enum
func SupportDataSetType_Values() []string {
	return []string{
		SupportDataSetTypeCustomerSupportContactsData,
		SupportDataSetTypeTestCustomerSupportContactsData,
	}
}
