// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connectcontactlens

import (
	"fmt"

	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/awsutil"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/private/protocol"
)

const opListRealtimeContactAnalysisSegments = "ListRealtimeContactAnalysisSegments"

// ListRealtimeContactAnalysisSegmentsRequest generates a "aws/request.Request" representing the
// client's request for the ListRealtimeContactAnalysisSegments operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ListRealtimeContactAnalysisSegments for more information on using the ListRealtimeContactAnalysisSegments
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ListRealtimeContactAnalysisSegmentsRequest method.
//    req, resp := client.ListRealtimeContactAnalysisSegmentsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/connect-contact-lens-2020-08-21/ListRealtimeContactAnalysisSegments
func (c *ConnectContactLens) ListRealtimeContactAnalysisSegmentsRequest(input *ListRealtimeContactAnalysisSegmentsInput) (req *request.Request, output *ListRealtimeContactAnalysisSegmentsOutput) {
	op := &request.Operation{
		Name:       opListRealtimeContactAnalysisSegments,
		HTTPMethod: "POST",
		HTTPPath:   "/realtime-contact-analysis/analysis-segments",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListRealtimeContactAnalysisSegmentsInput{}
	}

	output = &ListRealtimeContactAnalysisSegmentsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListRealtimeContactAnalysisSegments API operation for Amazon Connect Contact Lens.
//
// Provides a list of analysis segments for a real-time analysis session.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Connect Contact Lens's
// API operation ListRealtimeContactAnalysisSegments for usage and error information.
//
// Returned Error Types:
//   * InvalidRequestException
//   The request is not valid.
//
//   * AccessDeniedException
//   You do not have sufficient access to perform this action.
//
//   * ResourceNotFoundException
//   The specified resource was not found.
//
//   * InternalServiceException
//   Request processing failed due to an error or failure with the service.
//
//   * ThrottlingException
//   The throttling limit has been exceeded.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/connect-contact-lens-2020-08-21/ListRealtimeContactAnalysisSegments
func (c *ConnectContactLens) ListRealtimeContactAnalysisSegments(input *ListRealtimeContactAnalysisSegmentsInput) (*ListRealtimeContactAnalysisSegmentsOutput, error) {
	req, out := c.ListRealtimeContactAnalysisSegmentsRequest(input)
	return out, req.Send()
}

// ListRealtimeContactAnalysisSegmentsWithContext is the same as ListRealtimeContactAnalysisSegments with the addition of
// the ability to pass a context and additional request options.
//
// See ListRealtimeContactAnalysisSegments for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ConnectContactLens) ListRealtimeContactAnalysisSegmentsWithContext(ctx aws.Context, input *ListRealtimeContactAnalysisSegmentsInput, opts ...request.Option) (*ListRealtimeContactAnalysisSegmentsOutput, error) {
	req, out := c.ListRealtimeContactAnalysisSegmentsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// ListRealtimeContactAnalysisSegmentsPages iterates over the pages of a ListRealtimeContactAnalysisSegments operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See ListRealtimeContactAnalysisSegments method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a ListRealtimeContactAnalysisSegments operation.
//    pageNum := 0
//    err := client.ListRealtimeContactAnalysisSegmentsPages(params,
//        func(page *connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *ConnectContactLens) ListRealtimeContactAnalysisSegmentsPages(input *ListRealtimeContactAnalysisSegmentsInput, fn func(*ListRealtimeContactAnalysisSegmentsOutput, bool) bool) error {
	return c.ListRealtimeContactAnalysisSegmentsPagesWithContext(aws.BackgroundContext(), input, fn)
}

// ListRealtimeContactAnalysisSegmentsPagesWithContext same as ListRealtimeContactAnalysisSegmentsPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ConnectContactLens) ListRealtimeContactAnalysisSegmentsPagesWithContext(ctx aws.Context, input *ListRealtimeContactAnalysisSegmentsInput, fn func(*ListRealtimeContactAnalysisSegmentsOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListRealtimeContactAnalysisSegmentsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListRealtimeContactAnalysisSegmentsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	for p.Next() {
		if !fn(p.Page().(*ListRealtimeContactAnalysisSegmentsOutput), !p.HasNextPage()) {
			break
		}
	}

	return p.Err()
}

// You do not have sufficient access to perform this action.
type AccessDeniedException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s AccessDeniedException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s AccessDeniedException) GoString() string {
	return s.String()
}

func newErrorAccessDeniedException(v protocol.ResponseMetadata) error {
	return &AccessDeniedException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *AccessDeniedException) Code() string {
	return "AccessDeniedException"
}

// Message returns the exception's message.
func (s *AccessDeniedException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *AccessDeniedException) OrigErr() error {
	return nil
}

func (s *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *AccessDeniedException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *AccessDeniedException) RequestID() string {
	return s.RespMetadata.RequestID
}

// Provides the category rules that are used to automatically categorize contacts
// based on uttered keywords and phrases.
type Categories struct {
	_ struct{} `type:"structure"`

	// The category rules that have been matched in the analyzed segment.
	//
	// MatchedCategories is a required field
	MatchedCategories []*string `type:"list" required:"true"`

	// The category rule that was matched and when it occurred in the transcript.
	//
	// MatchedDetails is a required field
	MatchedDetails map[string]*CategoryDetails `type:"map" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Categories) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Categories) GoString() string {
	return s.String()
}

// SetMatchedCategories sets the MatchedCategories field's value.
func (s *Categories) SetMatchedCategories(v []*string) *Categories {
	s.MatchedCategories = v
	return s
}

// SetMatchedDetails sets the MatchedDetails field's value.
func (s *Categories) SetMatchedDetails(v map[string]*CategoryDetails) *Categories {
	s.MatchedDetails = v
	return s
}

// Provides information about the category rule that was matched.
type CategoryDetails struct {
	_ struct{} `type:"structure"`

	// The section of audio where the category rule was detected.
	//
	// PointsOfInterest is a required field
	PointsOfInterest []*PointOfInterest `type:"list" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s CategoryDetails) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s CategoryDetails) GoString() string {
	return s.String()
}

// SetPointsOfInterest sets the PointsOfInterest field's value.
func (s *CategoryDetails) SetPointsOfInterest(v []*PointOfInterest) *CategoryDetails {
	s.PointsOfInterest = v
	return s
}

// For characters that were detected as issues, where they occur in the transcript.
type CharacterOffsets struct {
	_ struct{} `type:"structure"`

	// The beginning of the issue.
	//
	// BeginOffsetChar is a required field
	BeginOffsetChar *int64 `type:"integer" required:"true"`

	// The end of the issue.
	//
	// EndOffsetChar is a required field
	EndOffsetChar *int64 `type:"integer" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s CharacterOffsets) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s CharacterOffsets) GoString() string {
	return s.String()
}

// SetBeginOffsetChar sets the BeginOffsetChar field's value.
func (s *CharacterOffsets) SetBeginOffsetChar(v int64) *CharacterOffsets {
	s.BeginOffsetChar = &v
	return s
}

// SetEndOffsetChar sets the EndOffsetChar field's value.
func (s *CharacterOffsets) SetEndOffsetChar(v int64) *CharacterOffsets {
	s.EndOffsetChar = &v
	return s
}

// Request processing failed due to an error or failure with the service.
type InternalServiceException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InternalServiceException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InternalServiceException) GoString() string {
	return s.String()
}

func newErrorInternalServiceException(v protocol.ResponseMetadata) error {
	return &InternalServiceException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InternalServiceException) Code() string {
	return "InternalServiceException"
}

// Message returns the exception's message.
func (s *InternalServiceException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InternalServiceException) OrigErr() error {
	return nil
}

func (s *InternalServiceException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InternalServiceException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InternalServiceException) RequestID() string {
	return s.RespMetadata.RequestID
}

// The request is not valid.
type InvalidRequestException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidRequestException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidRequestException) GoString() string {
	return s.String()
}

func newErrorInvalidRequestException(v protocol.ResponseMetadata) error {
	return &InvalidRequestException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InvalidRequestException) Code() string {
	return "InvalidRequestException"
}

// Message returns the exception's message.
func (s *InvalidRequestException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InvalidRequestException) OrigErr() error {
	return nil
}

func (s *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InvalidRequestException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InvalidRequestException) RequestID() string {
	return s.RespMetadata.RequestID
}

// Potential issues that are detected based on an artificial intelligence analysis
// of each turn in the conversation.
type IssueDetected struct {
	_ struct{} `type:"structure"`

	// The offset for when the issue was detected in the segment.
	//
	// CharacterOffsets is a required field
	CharacterOffsets *CharacterOffsets `type:"structure" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s IssueDetected) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s IssueDetected) GoString() string {
	return s.String()
}

// SetCharacterOffsets sets the CharacterOffsets field's value.
func (s *IssueDetected) SetCharacterOffsets(v *CharacterOffsets) *IssueDetected {
	s.CharacterOffsets = v
	return s
}

type ListRealtimeContactAnalysisSegmentsInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the contact.
	//
	// ContactId is a required field
	ContactId *string `min:"1" type:"string" required:"true"`

	// The identifier of the Amazon Connect instance.
	//
	// InstanceId is a required field
	InstanceId *string `min:"1" type:"string" required:"true"`

	// The maximimum number of results to return per page.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ListRealtimeContactAnalysisSegmentsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ListRealtimeContactAnalysisSegmentsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRealtimeContactAnalysisSegmentsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListRealtimeContactAnalysisSegmentsInput"}
	if s.ContactId == nil {
		invalidParams.Add(request.NewErrParamRequired("ContactId"))
	}
	if s.ContactId != nil && len(*s.ContactId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ContactId", 1))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("InstanceId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetContactId sets the ContactId field's value.
func (s *ListRealtimeContactAnalysisSegmentsInput) SetContactId(v string) *ListRealtimeContactAnalysisSegmentsInput {
	s.ContactId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ListRealtimeContactAnalysisSegmentsInput) SetInstanceId(v string) *ListRealtimeContactAnalysisSegmentsInput {
	s.InstanceId = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListRealtimeContactAnalysisSegmentsInput) SetMaxResults(v int64) *ListRealtimeContactAnalysisSegmentsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListRealtimeContactAnalysisSegmentsInput) SetNextToken(v string) *ListRealtimeContactAnalysisSegmentsInput {
	s.NextToken = &v
	return s
}

type ListRealtimeContactAnalysisSegmentsOutput struct {
	_ struct{} `type:"structure"`

	// If there are additional results, this is the token for the next set of results.
	// If response includes nextToken there are two possible scenarios:
	//
	//    * There are more segments so another call is required to get them.
	//
	//    * There are no more segments at this time, but more may be available later
	//    (real-time analysis is in progress) so the client should call the operation
	//    again to get new segments.
	//
	// If response does not include nextToken, the analysis is completed (successfully
	// or failed) and there are no more segments to retrieve.
	NextToken *string `min:"1" type:"string"`

	// An analyzed transcript or category.
	//
	// Segments is a required field
	Segments []*RealtimeContactAnalysisSegment `type:"list" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ListRealtimeContactAnalysisSegmentsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ListRealtimeContactAnalysisSegmentsOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *ListRealtimeContactAnalysisSegmentsOutput) SetNextToken(v string) *ListRealtimeContactAnalysisSegmentsOutput {
	s.NextToken = &v
	return s
}

// SetSegments sets the Segments field's value.
func (s *ListRealtimeContactAnalysisSegmentsOutput) SetSegments(v []*RealtimeContactAnalysisSegment) *ListRealtimeContactAnalysisSegmentsOutput {
	s.Segments = v
	return s
}

// The section of the contact audio where that category rule was detected.
type PointOfInterest struct {
	_ struct{} `type:"structure"`

	// The beginning offset in milliseconds where the category rule was detected.
	//
	// BeginOffsetMillis is a required field
	BeginOffsetMillis *int64 `type:"integer" required:"true"`

	// The ending offset in milliseconds where the category rule was detected.
	//
	// EndOffsetMillis is a required field
	EndOffsetMillis *int64 `type:"integer" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PointOfInterest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PointOfInterest) GoString() string {
	return s.String()
}

// SetBeginOffsetMillis sets the BeginOffsetMillis field's value.
func (s *PointOfInterest) SetBeginOffsetMillis(v int64) *PointOfInterest {
	s.BeginOffsetMillis = &v
	return s
}

// SetEndOffsetMillis sets the EndOffsetMillis field's value.
func (s *PointOfInterest) SetEndOffsetMillis(v int64) *PointOfInterest {
	s.EndOffsetMillis = &v
	return s
}

// An analyzed segment for a real-time analysis session.
type RealtimeContactAnalysisSegment struct {
	_ struct{} `type:"structure"`

	// The matched category rules.
	Categories *Categories `type:"structure"`

	// The analyzed transcript.
	Transcript *Transcript `type:"structure"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s RealtimeContactAnalysisSegment) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s RealtimeContactAnalysisSegment) GoString() string {
	return s.String()
}

// SetCategories sets the Categories field's value.
func (s *RealtimeContactAnalysisSegment) SetCategories(v *Categories) *RealtimeContactAnalysisSegment {
	s.Categories = v
	return s
}

// SetTranscript sets the Transcript field's value.
func (s *RealtimeContactAnalysisSegment) SetTranscript(v *Transcript) *RealtimeContactAnalysisSegment {
	s.Transcript = v
	return s
}

// The specified resource was not found.
type ResourceNotFoundException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ResourceNotFoundException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ResourceNotFoundException) GoString() string {
	return s.String()
}

func newErrorResourceNotFoundException(v protocol.ResponseMetadata) error {
	return &ResourceNotFoundException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ResourceNotFoundException) Code() string {
	return "ResourceNotFoundException"
}

// Message returns the exception's message.
func (s *ResourceNotFoundException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ResourceNotFoundException) OrigErr() error {
	return nil
}

func (s *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ResourceNotFoundException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ResourceNotFoundException) RequestID() string {
	return s.RespMetadata.RequestID
}

// The throttling limit has been exceeded.
type ThrottlingException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ThrottlingException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ThrottlingException) GoString() string {
	return s.String()
}

func newErrorThrottlingException(v protocol.ResponseMetadata) error {
	return &ThrottlingException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ThrottlingException) Code() string {
	return "ThrottlingException"
}

// Message returns the exception's message.
func (s *ThrottlingException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ThrottlingException) OrigErr() error {
	return nil
}

func (s *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ThrottlingException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ThrottlingException) RequestID() string {
	return s.RespMetadata.RequestID
}

// A list of messages in the session.
type Transcript struct {
	_ struct{} `type:"structure"`

	// The beginning offset in the contact for this transcript.
	//
	// BeginOffsetMillis is a required field
	BeginOffsetMillis *int64 `type:"integer" required:"true"`

	// The content of the transcript.
	//
	// Content is a required field
	Content *string `min:"1" type:"string" required:"true"`

	// The end offset in the contact for this transcript.
	//
	// EndOffsetMillis is a required field
	EndOffsetMillis *int64 `type:"integer" required:"true"`

	// The identifier of the transcript.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// List of positions where issues were detected on the transcript.
	IssuesDetected []*IssueDetected `type:"list"`

	// The identifier of the participant.
	//
	// ParticipantId is a required field
	ParticipantId *string `min:"1" type:"string" required:"true"`

	// The role of participant. For example, is it a customer, agent, or system.
	//
	// ParticipantRole is a required field
	ParticipantRole *string `min:"1" type:"string" required:"true"`

	// The sentiment of the detected for this piece of transcript.
	//
	// Sentiment is a required field
	Sentiment *string `type:"string" required:"true" enum:"SentimentValue"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Transcript) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Transcript) GoString() string {
	return s.String()
}

// SetBeginOffsetMillis sets the BeginOffsetMillis field's value.
func (s *Transcript) SetBeginOffsetMillis(v int64) *Transcript {
	s.BeginOffsetMillis = &v
	return s
}

// SetContent sets the Content field's value.
func (s *Transcript) SetContent(v string) *Transcript {
	s.Content = &v
	return s
}

// SetEndOffsetMillis sets the EndOffsetMillis field's value.
func (s *Transcript) SetEndOffsetMillis(v int64) *Transcript {
	s.EndOffsetMillis = &v
	return s
}

// SetId sets the Id field's value.
func (s *Transcript) SetId(v string) *Transcript {
	s.Id = &v
	return s
}

// SetIssuesDetected sets the IssuesDetected field's value.
func (s *Transcript) SetIssuesDetected(v []*IssueDetected) *Transcript {
	s.IssuesDetected = v
	return s
}

// SetParticipantId sets the ParticipantId field's value.
func (s *Transcript) SetParticipantId(v string) *Transcript {
	s.ParticipantId = &v
	return s
}

// SetParticipantRole sets the ParticipantRole field's value.
func (s *Transcript) SetParticipantRole(v string) *Transcript {
	s.ParticipantRole = &v
	return s
}

// SetSentiment sets the Sentiment field's value.
func (s *Transcript) SetSentiment(v string) *Transcript {
	s.Sentiment = &v
	return s
}

const (
	// SentimentValuePositive is a SentimentValue enum value
	SentimentValuePositive = "POSITIVE"

	// SentimentValueNeutral is a SentimentValue enum value
	SentimentValueNeutral = "NEUTRAL"

	// SentimentValueNegative is a SentimentValue enum value
	SentimentValueNegative = "NEGATIVE"
)

// SentimentValue_Values returns all elements of the SentimentValue enum
func SentimentValue_Values() []string {
	return []string{
		SentimentValuePositive,
		SentimentValueNeutral,
		SentimentValueNegative,
	}
}
