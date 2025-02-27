// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rekognitioniface provides an interface to enable mocking the Amazon Rekognition service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rekognitioniface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/rekognition"
)

// RekognitionAPI provides an interface to enable mocking the
// rekognition.Rekognition service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Rekognition.
//    func myFunc(svc rekognitioniface.RekognitionAPI) bool {
//        // Make svc.CompareFaces request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := rekognition.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRekognitionClient struct {
//        rekognitioniface.RekognitionAPI
//    }
//    func (m *mockRekognitionClient) CompareFaces(input *rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRekognitionClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type RekognitionAPI interface {
	CompareFaces(*rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error)
	CompareFacesWithContext(aws.Context, *rekognition.CompareFacesInput, ...request.Option) (*rekognition.CompareFacesOutput, error)
	CompareFacesRequest(*rekognition.CompareFacesInput) (*request.Request, *rekognition.CompareFacesOutput)

	CreateCollection(*rekognition.CreateCollectionInput) (*rekognition.CreateCollectionOutput, error)
	CreateCollectionWithContext(aws.Context, *rekognition.CreateCollectionInput, ...request.Option) (*rekognition.CreateCollectionOutput, error)
	CreateCollectionRequest(*rekognition.CreateCollectionInput) (*request.Request, *rekognition.CreateCollectionOutput)

	CreateProject(*rekognition.CreateProjectInput) (*rekognition.CreateProjectOutput, error)
	CreateProjectWithContext(aws.Context, *rekognition.CreateProjectInput, ...request.Option) (*rekognition.CreateProjectOutput, error)
	CreateProjectRequest(*rekognition.CreateProjectInput) (*request.Request, *rekognition.CreateProjectOutput)

	CreateProjectVersion(*rekognition.CreateProjectVersionInput) (*rekognition.CreateProjectVersionOutput, error)
	CreateProjectVersionWithContext(aws.Context, *rekognition.CreateProjectVersionInput, ...request.Option) (*rekognition.CreateProjectVersionOutput, error)
	CreateProjectVersionRequest(*rekognition.CreateProjectVersionInput) (*request.Request, *rekognition.CreateProjectVersionOutput)

	CreateStreamProcessor(*rekognition.CreateStreamProcessorInput) (*rekognition.CreateStreamProcessorOutput, error)
	CreateStreamProcessorWithContext(aws.Context, *rekognition.CreateStreamProcessorInput, ...request.Option) (*rekognition.CreateStreamProcessorOutput, error)
	CreateStreamProcessorRequest(*rekognition.CreateStreamProcessorInput) (*request.Request, *rekognition.CreateStreamProcessorOutput)

	DeleteCollection(*rekognition.DeleteCollectionInput) (*rekognition.DeleteCollectionOutput, error)
	DeleteCollectionWithContext(aws.Context, *rekognition.DeleteCollectionInput, ...request.Option) (*rekognition.DeleteCollectionOutput, error)
	DeleteCollectionRequest(*rekognition.DeleteCollectionInput) (*request.Request, *rekognition.DeleteCollectionOutput)

	DeleteFaces(*rekognition.DeleteFacesInput) (*rekognition.DeleteFacesOutput, error)
	DeleteFacesWithContext(aws.Context, *rekognition.DeleteFacesInput, ...request.Option) (*rekognition.DeleteFacesOutput, error)
	DeleteFacesRequest(*rekognition.DeleteFacesInput) (*request.Request, *rekognition.DeleteFacesOutput)

	DeleteProject(*rekognition.DeleteProjectInput) (*rekognition.DeleteProjectOutput, error)
	DeleteProjectWithContext(aws.Context, *rekognition.DeleteProjectInput, ...request.Option) (*rekognition.DeleteProjectOutput, error)
	DeleteProjectRequest(*rekognition.DeleteProjectInput) (*request.Request, *rekognition.DeleteProjectOutput)

	DeleteProjectVersion(*rekognition.DeleteProjectVersionInput) (*rekognition.DeleteProjectVersionOutput, error)
	DeleteProjectVersionWithContext(aws.Context, *rekognition.DeleteProjectVersionInput, ...request.Option) (*rekognition.DeleteProjectVersionOutput, error)
	DeleteProjectVersionRequest(*rekognition.DeleteProjectVersionInput) (*request.Request, *rekognition.DeleteProjectVersionOutput)

	DeleteStreamProcessor(*rekognition.DeleteStreamProcessorInput) (*rekognition.DeleteStreamProcessorOutput, error)
	DeleteStreamProcessorWithContext(aws.Context, *rekognition.DeleteStreamProcessorInput, ...request.Option) (*rekognition.DeleteStreamProcessorOutput, error)
	DeleteStreamProcessorRequest(*rekognition.DeleteStreamProcessorInput) (*request.Request, *rekognition.DeleteStreamProcessorOutput)

	DescribeCollection(*rekognition.DescribeCollectionInput) (*rekognition.DescribeCollectionOutput, error)
	DescribeCollectionWithContext(aws.Context, *rekognition.DescribeCollectionInput, ...request.Option) (*rekognition.DescribeCollectionOutput, error)
	DescribeCollectionRequest(*rekognition.DescribeCollectionInput) (*request.Request, *rekognition.DescribeCollectionOutput)

	DescribeProjectVersions(*rekognition.DescribeProjectVersionsInput) (*rekognition.DescribeProjectVersionsOutput, error)
	DescribeProjectVersionsWithContext(aws.Context, *rekognition.DescribeProjectVersionsInput, ...request.Option) (*rekognition.DescribeProjectVersionsOutput, error)
	DescribeProjectVersionsRequest(*rekognition.DescribeProjectVersionsInput) (*request.Request, *rekognition.DescribeProjectVersionsOutput)

	DescribeProjectVersionsPages(*rekognition.DescribeProjectVersionsInput, func(*rekognition.DescribeProjectVersionsOutput, bool) bool) error
	DescribeProjectVersionsPagesWithContext(aws.Context, *rekognition.DescribeProjectVersionsInput, func(*rekognition.DescribeProjectVersionsOutput, bool) bool, ...request.Option) error

	DescribeProjects(*rekognition.DescribeProjectsInput) (*rekognition.DescribeProjectsOutput, error)
	DescribeProjectsWithContext(aws.Context, *rekognition.DescribeProjectsInput, ...request.Option) (*rekognition.DescribeProjectsOutput, error)
	DescribeProjectsRequest(*rekognition.DescribeProjectsInput) (*request.Request, *rekognition.DescribeProjectsOutput)

	DescribeProjectsPages(*rekognition.DescribeProjectsInput, func(*rekognition.DescribeProjectsOutput, bool) bool) error
	DescribeProjectsPagesWithContext(aws.Context, *rekognition.DescribeProjectsInput, func(*rekognition.DescribeProjectsOutput, bool) bool, ...request.Option) error

	DescribeStreamProcessor(*rekognition.DescribeStreamProcessorInput) (*rekognition.DescribeStreamProcessorOutput, error)
	DescribeStreamProcessorWithContext(aws.Context, *rekognition.DescribeStreamProcessorInput, ...request.Option) (*rekognition.DescribeStreamProcessorOutput, error)
	DescribeStreamProcessorRequest(*rekognition.DescribeStreamProcessorInput) (*request.Request, *rekognition.DescribeStreamProcessorOutput)

	DetectCustomLabels(*rekognition.DetectCustomLabelsInput) (*rekognition.DetectCustomLabelsOutput, error)
	DetectCustomLabelsWithContext(aws.Context, *rekognition.DetectCustomLabelsInput, ...request.Option) (*rekognition.DetectCustomLabelsOutput, error)
	DetectCustomLabelsRequest(*rekognition.DetectCustomLabelsInput) (*request.Request, *rekognition.DetectCustomLabelsOutput)

	DetectFaces(*rekognition.DetectFacesInput) (*rekognition.DetectFacesOutput, error)
	DetectFacesWithContext(aws.Context, *rekognition.DetectFacesInput, ...request.Option) (*rekognition.DetectFacesOutput, error)
	DetectFacesRequest(*rekognition.DetectFacesInput) (*request.Request, *rekognition.DetectFacesOutput)

	DetectLabels(*rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error)
	DetectLabelsWithContext(aws.Context, *rekognition.DetectLabelsInput, ...request.Option) (*rekognition.DetectLabelsOutput, error)
	DetectLabelsRequest(*rekognition.DetectLabelsInput) (*request.Request, *rekognition.DetectLabelsOutput)

	DetectModerationLabels(*rekognition.DetectModerationLabelsInput) (*rekognition.DetectModerationLabelsOutput, error)
	DetectModerationLabelsWithContext(aws.Context, *rekognition.DetectModerationLabelsInput, ...request.Option) (*rekognition.DetectModerationLabelsOutput, error)
	DetectModerationLabelsRequest(*rekognition.DetectModerationLabelsInput) (*request.Request, *rekognition.DetectModerationLabelsOutput)

	DetectProtectiveEquipment(*rekognition.DetectProtectiveEquipmentInput) (*rekognition.DetectProtectiveEquipmentOutput, error)
	DetectProtectiveEquipmentWithContext(aws.Context, *rekognition.DetectProtectiveEquipmentInput, ...request.Option) (*rekognition.DetectProtectiveEquipmentOutput, error)
	DetectProtectiveEquipmentRequest(*rekognition.DetectProtectiveEquipmentInput) (*request.Request, *rekognition.DetectProtectiveEquipmentOutput)

	DetectText(*rekognition.DetectTextInput) (*rekognition.DetectTextOutput, error)
	DetectTextWithContext(aws.Context, *rekognition.DetectTextInput, ...request.Option) (*rekognition.DetectTextOutput, error)
	DetectTextRequest(*rekognition.DetectTextInput) (*request.Request, *rekognition.DetectTextOutput)

	GetCelebrityInfo(*rekognition.GetCelebrityInfoInput) (*rekognition.GetCelebrityInfoOutput, error)
	GetCelebrityInfoWithContext(aws.Context, *rekognition.GetCelebrityInfoInput, ...request.Option) (*rekognition.GetCelebrityInfoOutput, error)
	GetCelebrityInfoRequest(*rekognition.GetCelebrityInfoInput) (*request.Request, *rekognition.GetCelebrityInfoOutput)

	GetCelebrityRecognition(*rekognition.GetCelebrityRecognitionInput) (*rekognition.GetCelebrityRecognitionOutput, error)
	GetCelebrityRecognitionWithContext(aws.Context, *rekognition.GetCelebrityRecognitionInput, ...request.Option) (*rekognition.GetCelebrityRecognitionOutput, error)
	GetCelebrityRecognitionRequest(*rekognition.GetCelebrityRecognitionInput) (*request.Request, *rekognition.GetCelebrityRecognitionOutput)

	GetCelebrityRecognitionPages(*rekognition.GetCelebrityRecognitionInput, func(*rekognition.GetCelebrityRecognitionOutput, bool) bool) error
	GetCelebrityRecognitionPagesWithContext(aws.Context, *rekognition.GetCelebrityRecognitionInput, func(*rekognition.GetCelebrityRecognitionOutput, bool) bool, ...request.Option) error

	GetContentModeration(*rekognition.GetContentModerationInput) (*rekognition.GetContentModerationOutput, error)
	GetContentModerationWithContext(aws.Context, *rekognition.GetContentModerationInput, ...request.Option) (*rekognition.GetContentModerationOutput, error)
	GetContentModerationRequest(*rekognition.GetContentModerationInput) (*request.Request, *rekognition.GetContentModerationOutput)

	GetContentModerationPages(*rekognition.GetContentModerationInput, func(*rekognition.GetContentModerationOutput, bool) bool) error
	GetContentModerationPagesWithContext(aws.Context, *rekognition.GetContentModerationInput, func(*rekognition.GetContentModerationOutput, bool) bool, ...request.Option) error

	GetFaceDetection(*rekognition.GetFaceDetectionInput) (*rekognition.GetFaceDetectionOutput, error)
	GetFaceDetectionWithContext(aws.Context, *rekognition.GetFaceDetectionInput, ...request.Option) (*rekognition.GetFaceDetectionOutput, error)
	GetFaceDetectionRequest(*rekognition.GetFaceDetectionInput) (*request.Request, *rekognition.GetFaceDetectionOutput)

	GetFaceDetectionPages(*rekognition.GetFaceDetectionInput, func(*rekognition.GetFaceDetectionOutput, bool) bool) error
	GetFaceDetectionPagesWithContext(aws.Context, *rekognition.GetFaceDetectionInput, func(*rekognition.GetFaceDetectionOutput, bool) bool, ...request.Option) error

	GetFaceSearch(*rekognition.GetFaceSearchInput) (*rekognition.GetFaceSearchOutput, error)
	GetFaceSearchWithContext(aws.Context, *rekognition.GetFaceSearchInput, ...request.Option) (*rekognition.GetFaceSearchOutput, error)
	GetFaceSearchRequest(*rekognition.GetFaceSearchInput) (*request.Request, *rekognition.GetFaceSearchOutput)

	GetFaceSearchPages(*rekognition.GetFaceSearchInput, func(*rekognition.GetFaceSearchOutput, bool) bool) error
	GetFaceSearchPagesWithContext(aws.Context, *rekognition.GetFaceSearchInput, func(*rekognition.GetFaceSearchOutput, bool) bool, ...request.Option) error

	GetLabelDetection(*rekognition.GetLabelDetectionInput) (*rekognition.GetLabelDetectionOutput, error)
	GetLabelDetectionWithContext(aws.Context, *rekognition.GetLabelDetectionInput, ...request.Option) (*rekognition.GetLabelDetectionOutput, error)
	GetLabelDetectionRequest(*rekognition.GetLabelDetectionInput) (*request.Request, *rekognition.GetLabelDetectionOutput)

	GetLabelDetectionPages(*rekognition.GetLabelDetectionInput, func(*rekognition.GetLabelDetectionOutput, bool) bool) error
	GetLabelDetectionPagesWithContext(aws.Context, *rekognition.GetLabelDetectionInput, func(*rekognition.GetLabelDetectionOutput, bool) bool, ...request.Option) error

	GetPersonTracking(*rekognition.GetPersonTrackingInput) (*rekognition.GetPersonTrackingOutput, error)
	GetPersonTrackingWithContext(aws.Context, *rekognition.GetPersonTrackingInput, ...request.Option) (*rekognition.GetPersonTrackingOutput, error)
	GetPersonTrackingRequest(*rekognition.GetPersonTrackingInput) (*request.Request, *rekognition.GetPersonTrackingOutput)

	GetPersonTrackingPages(*rekognition.GetPersonTrackingInput, func(*rekognition.GetPersonTrackingOutput, bool) bool) error
	GetPersonTrackingPagesWithContext(aws.Context, *rekognition.GetPersonTrackingInput, func(*rekognition.GetPersonTrackingOutput, bool) bool, ...request.Option) error

	GetSegmentDetection(*rekognition.GetSegmentDetectionInput) (*rekognition.GetSegmentDetectionOutput, error)
	GetSegmentDetectionWithContext(aws.Context, *rekognition.GetSegmentDetectionInput, ...request.Option) (*rekognition.GetSegmentDetectionOutput, error)
	GetSegmentDetectionRequest(*rekognition.GetSegmentDetectionInput) (*request.Request, *rekognition.GetSegmentDetectionOutput)

	GetSegmentDetectionPages(*rekognition.GetSegmentDetectionInput, func(*rekognition.GetSegmentDetectionOutput, bool) bool) error
	GetSegmentDetectionPagesWithContext(aws.Context, *rekognition.GetSegmentDetectionInput, func(*rekognition.GetSegmentDetectionOutput, bool) bool, ...request.Option) error

	GetTextDetection(*rekognition.GetTextDetectionInput) (*rekognition.GetTextDetectionOutput, error)
	GetTextDetectionWithContext(aws.Context, *rekognition.GetTextDetectionInput, ...request.Option) (*rekognition.GetTextDetectionOutput, error)
	GetTextDetectionRequest(*rekognition.GetTextDetectionInput) (*request.Request, *rekognition.GetTextDetectionOutput)

	GetTextDetectionPages(*rekognition.GetTextDetectionInput, func(*rekognition.GetTextDetectionOutput, bool) bool) error
	GetTextDetectionPagesWithContext(aws.Context, *rekognition.GetTextDetectionInput, func(*rekognition.GetTextDetectionOutput, bool) bool, ...request.Option) error

	IndexFaces(*rekognition.IndexFacesInput) (*rekognition.IndexFacesOutput, error)
	IndexFacesWithContext(aws.Context, *rekognition.IndexFacesInput, ...request.Option) (*rekognition.IndexFacesOutput, error)
	IndexFacesRequest(*rekognition.IndexFacesInput) (*request.Request, *rekognition.IndexFacesOutput)

	ListCollections(*rekognition.ListCollectionsInput) (*rekognition.ListCollectionsOutput, error)
	ListCollectionsWithContext(aws.Context, *rekognition.ListCollectionsInput, ...request.Option) (*rekognition.ListCollectionsOutput, error)
	ListCollectionsRequest(*rekognition.ListCollectionsInput) (*request.Request, *rekognition.ListCollectionsOutput)

	ListCollectionsPages(*rekognition.ListCollectionsInput, func(*rekognition.ListCollectionsOutput, bool) bool) error
	ListCollectionsPagesWithContext(aws.Context, *rekognition.ListCollectionsInput, func(*rekognition.ListCollectionsOutput, bool) bool, ...request.Option) error

	ListFaces(*rekognition.ListFacesInput) (*rekognition.ListFacesOutput, error)
	ListFacesWithContext(aws.Context, *rekognition.ListFacesInput, ...request.Option) (*rekognition.ListFacesOutput, error)
	ListFacesRequest(*rekognition.ListFacesInput) (*request.Request, *rekognition.ListFacesOutput)

	ListFacesPages(*rekognition.ListFacesInput, func(*rekognition.ListFacesOutput, bool) bool) error
	ListFacesPagesWithContext(aws.Context, *rekognition.ListFacesInput, func(*rekognition.ListFacesOutput, bool) bool, ...request.Option) error

	ListStreamProcessors(*rekognition.ListStreamProcessorsInput) (*rekognition.ListStreamProcessorsOutput, error)
	ListStreamProcessorsWithContext(aws.Context, *rekognition.ListStreamProcessorsInput, ...request.Option) (*rekognition.ListStreamProcessorsOutput, error)
	ListStreamProcessorsRequest(*rekognition.ListStreamProcessorsInput) (*request.Request, *rekognition.ListStreamProcessorsOutput)

	ListStreamProcessorsPages(*rekognition.ListStreamProcessorsInput, func(*rekognition.ListStreamProcessorsOutput, bool) bool) error
	ListStreamProcessorsPagesWithContext(aws.Context, *rekognition.ListStreamProcessorsInput, func(*rekognition.ListStreamProcessorsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*rekognition.ListTagsForResourceInput) (*rekognition.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *rekognition.ListTagsForResourceInput, ...request.Option) (*rekognition.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*rekognition.ListTagsForResourceInput) (*request.Request, *rekognition.ListTagsForResourceOutput)

	RecognizeCelebrities(*rekognition.RecognizeCelebritiesInput) (*rekognition.RecognizeCelebritiesOutput, error)
	RecognizeCelebritiesWithContext(aws.Context, *rekognition.RecognizeCelebritiesInput, ...request.Option) (*rekognition.RecognizeCelebritiesOutput, error)
	RecognizeCelebritiesRequest(*rekognition.RecognizeCelebritiesInput) (*request.Request, *rekognition.RecognizeCelebritiesOutput)

	SearchFaces(*rekognition.SearchFacesInput) (*rekognition.SearchFacesOutput, error)
	SearchFacesWithContext(aws.Context, *rekognition.SearchFacesInput, ...request.Option) (*rekognition.SearchFacesOutput, error)
	SearchFacesRequest(*rekognition.SearchFacesInput) (*request.Request, *rekognition.SearchFacesOutput)

	SearchFacesByImage(*rekognition.SearchFacesByImageInput) (*rekognition.SearchFacesByImageOutput, error)
	SearchFacesByImageWithContext(aws.Context, *rekognition.SearchFacesByImageInput, ...request.Option) (*rekognition.SearchFacesByImageOutput, error)
	SearchFacesByImageRequest(*rekognition.SearchFacesByImageInput) (*request.Request, *rekognition.SearchFacesByImageOutput)

	StartCelebrityRecognition(*rekognition.StartCelebrityRecognitionInput) (*rekognition.StartCelebrityRecognitionOutput, error)
	StartCelebrityRecognitionWithContext(aws.Context, *rekognition.StartCelebrityRecognitionInput, ...request.Option) (*rekognition.StartCelebrityRecognitionOutput, error)
	StartCelebrityRecognitionRequest(*rekognition.StartCelebrityRecognitionInput) (*request.Request, *rekognition.StartCelebrityRecognitionOutput)

	StartContentModeration(*rekognition.StartContentModerationInput) (*rekognition.StartContentModerationOutput, error)
	StartContentModerationWithContext(aws.Context, *rekognition.StartContentModerationInput, ...request.Option) (*rekognition.StartContentModerationOutput, error)
	StartContentModerationRequest(*rekognition.StartContentModerationInput) (*request.Request, *rekognition.StartContentModerationOutput)

	StartFaceDetection(*rekognition.StartFaceDetectionInput) (*rekognition.StartFaceDetectionOutput, error)
	StartFaceDetectionWithContext(aws.Context, *rekognition.StartFaceDetectionInput, ...request.Option) (*rekognition.StartFaceDetectionOutput, error)
	StartFaceDetectionRequest(*rekognition.StartFaceDetectionInput) (*request.Request, *rekognition.StartFaceDetectionOutput)

	StartFaceSearch(*rekognition.StartFaceSearchInput) (*rekognition.StartFaceSearchOutput, error)
	StartFaceSearchWithContext(aws.Context, *rekognition.StartFaceSearchInput, ...request.Option) (*rekognition.StartFaceSearchOutput, error)
	StartFaceSearchRequest(*rekognition.StartFaceSearchInput) (*request.Request, *rekognition.StartFaceSearchOutput)

	StartLabelDetection(*rekognition.StartLabelDetectionInput) (*rekognition.StartLabelDetectionOutput, error)
	StartLabelDetectionWithContext(aws.Context, *rekognition.StartLabelDetectionInput, ...request.Option) (*rekognition.StartLabelDetectionOutput, error)
	StartLabelDetectionRequest(*rekognition.StartLabelDetectionInput) (*request.Request, *rekognition.StartLabelDetectionOutput)

	StartPersonTracking(*rekognition.StartPersonTrackingInput) (*rekognition.StartPersonTrackingOutput, error)
	StartPersonTrackingWithContext(aws.Context, *rekognition.StartPersonTrackingInput, ...request.Option) (*rekognition.StartPersonTrackingOutput, error)
	StartPersonTrackingRequest(*rekognition.StartPersonTrackingInput) (*request.Request, *rekognition.StartPersonTrackingOutput)

	StartProjectVersion(*rekognition.StartProjectVersionInput) (*rekognition.StartProjectVersionOutput, error)
	StartProjectVersionWithContext(aws.Context, *rekognition.StartProjectVersionInput, ...request.Option) (*rekognition.StartProjectVersionOutput, error)
	StartProjectVersionRequest(*rekognition.StartProjectVersionInput) (*request.Request, *rekognition.StartProjectVersionOutput)

	StartSegmentDetection(*rekognition.StartSegmentDetectionInput) (*rekognition.StartSegmentDetectionOutput, error)
	StartSegmentDetectionWithContext(aws.Context, *rekognition.StartSegmentDetectionInput, ...request.Option) (*rekognition.StartSegmentDetectionOutput, error)
	StartSegmentDetectionRequest(*rekognition.StartSegmentDetectionInput) (*request.Request, *rekognition.StartSegmentDetectionOutput)

	StartStreamProcessor(*rekognition.StartStreamProcessorInput) (*rekognition.StartStreamProcessorOutput, error)
	StartStreamProcessorWithContext(aws.Context, *rekognition.StartStreamProcessorInput, ...request.Option) (*rekognition.StartStreamProcessorOutput, error)
	StartStreamProcessorRequest(*rekognition.StartStreamProcessorInput) (*request.Request, *rekognition.StartStreamProcessorOutput)

	StartTextDetection(*rekognition.StartTextDetectionInput) (*rekognition.StartTextDetectionOutput, error)
	StartTextDetectionWithContext(aws.Context, *rekognition.StartTextDetectionInput, ...request.Option) (*rekognition.StartTextDetectionOutput, error)
	StartTextDetectionRequest(*rekognition.StartTextDetectionInput) (*request.Request, *rekognition.StartTextDetectionOutput)

	StopProjectVersion(*rekognition.StopProjectVersionInput) (*rekognition.StopProjectVersionOutput, error)
	StopProjectVersionWithContext(aws.Context, *rekognition.StopProjectVersionInput, ...request.Option) (*rekognition.StopProjectVersionOutput, error)
	StopProjectVersionRequest(*rekognition.StopProjectVersionInput) (*request.Request, *rekognition.StopProjectVersionOutput)

	StopStreamProcessor(*rekognition.StopStreamProcessorInput) (*rekognition.StopStreamProcessorOutput, error)
	StopStreamProcessorWithContext(aws.Context, *rekognition.StopStreamProcessorInput, ...request.Option) (*rekognition.StopStreamProcessorOutput, error)
	StopStreamProcessorRequest(*rekognition.StopStreamProcessorInput) (*request.Request, *rekognition.StopStreamProcessorOutput)

	TagResource(*rekognition.TagResourceInput) (*rekognition.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *rekognition.TagResourceInput, ...request.Option) (*rekognition.TagResourceOutput, error)
	TagResourceRequest(*rekognition.TagResourceInput) (*request.Request, *rekognition.TagResourceOutput)

	UntagResource(*rekognition.UntagResourceInput) (*rekognition.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *rekognition.UntagResourceInput, ...request.Option) (*rekognition.UntagResourceOutput, error)
	UntagResourceRequest(*rekognition.UntagResourceInput) (*request.Request, *rekognition.UntagResourceOutput)

	WaitUntilProjectVersionRunning(*rekognition.DescribeProjectVersionsInput) error
	WaitUntilProjectVersionRunningWithContext(aws.Context, *rekognition.DescribeProjectVersionsInput, ...request.WaiterOption) error

	WaitUntilProjectVersionTrainingCompleted(*rekognition.DescribeProjectVersionsInput) error
	WaitUntilProjectVersionTrainingCompletedWithContext(aws.Context, *rekognition.DescribeProjectVersionsInput, ...request.WaiterOption) error
}

var _ RekognitionAPI = (*rekognition.Rekognition)(nil)
