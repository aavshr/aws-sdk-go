// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package connectwisdomserviceiface provides an interface to enable mocking the Amazon Connect Wisdom Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package connectwisdomserviceiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/connectwisdomservice"
)

// ConnectWisdomServiceAPI provides an interface to enable mocking the
// connectwisdomservice.ConnectWisdomService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Connect Wisdom Service.
//    func myFunc(svc connectwisdomserviceiface.ConnectWisdomServiceAPI) bool {
//        // Make svc.CreateAssistant request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := connectwisdomservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockConnectWisdomServiceClient struct {
//        connectwisdomserviceiface.ConnectWisdomServiceAPI
//    }
//    func (m *mockConnectWisdomServiceClient) CreateAssistant(input *connectwisdomservice.CreateAssistantInput) (*connectwisdomservice.CreateAssistantOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockConnectWisdomServiceClient{}
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
type ConnectWisdomServiceAPI interface {
	CreateAssistant(*connectwisdomservice.CreateAssistantInput) (*connectwisdomservice.CreateAssistantOutput, error)
	CreateAssistantWithContext(aws.Context, *connectwisdomservice.CreateAssistantInput, ...request.Option) (*connectwisdomservice.CreateAssistantOutput, error)
	CreateAssistantRequest(*connectwisdomservice.CreateAssistantInput) (*request.Request, *connectwisdomservice.CreateAssistantOutput)

	CreateAssistantAssociation(*connectwisdomservice.CreateAssistantAssociationInput) (*connectwisdomservice.CreateAssistantAssociationOutput, error)
	CreateAssistantAssociationWithContext(aws.Context, *connectwisdomservice.CreateAssistantAssociationInput, ...request.Option) (*connectwisdomservice.CreateAssistantAssociationOutput, error)
	CreateAssistantAssociationRequest(*connectwisdomservice.CreateAssistantAssociationInput) (*request.Request, *connectwisdomservice.CreateAssistantAssociationOutput)

	CreateContent(*connectwisdomservice.CreateContentInput) (*connectwisdomservice.CreateContentOutput, error)
	CreateContentWithContext(aws.Context, *connectwisdomservice.CreateContentInput, ...request.Option) (*connectwisdomservice.CreateContentOutput, error)
	CreateContentRequest(*connectwisdomservice.CreateContentInput) (*request.Request, *connectwisdomservice.CreateContentOutput)

	CreateKnowledgeBase(*connectwisdomservice.CreateKnowledgeBaseInput) (*connectwisdomservice.CreateKnowledgeBaseOutput, error)
	CreateKnowledgeBaseWithContext(aws.Context, *connectwisdomservice.CreateKnowledgeBaseInput, ...request.Option) (*connectwisdomservice.CreateKnowledgeBaseOutput, error)
	CreateKnowledgeBaseRequest(*connectwisdomservice.CreateKnowledgeBaseInput) (*request.Request, *connectwisdomservice.CreateKnowledgeBaseOutput)

	CreateSession(*connectwisdomservice.CreateSessionInput) (*connectwisdomservice.CreateSessionOutput, error)
	CreateSessionWithContext(aws.Context, *connectwisdomservice.CreateSessionInput, ...request.Option) (*connectwisdomservice.CreateSessionOutput, error)
	CreateSessionRequest(*connectwisdomservice.CreateSessionInput) (*request.Request, *connectwisdomservice.CreateSessionOutput)

	DeleteAssistant(*connectwisdomservice.DeleteAssistantInput) (*connectwisdomservice.DeleteAssistantOutput, error)
	DeleteAssistantWithContext(aws.Context, *connectwisdomservice.DeleteAssistantInput, ...request.Option) (*connectwisdomservice.DeleteAssistantOutput, error)
	DeleteAssistantRequest(*connectwisdomservice.DeleteAssistantInput) (*request.Request, *connectwisdomservice.DeleteAssistantOutput)

	DeleteAssistantAssociation(*connectwisdomservice.DeleteAssistantAssociationInput) (*connectwisdomservice.DeleteAssistantAssociationOutput, error)
	DeleteAssistantAssociationWithContext(aws.Context, *connectwisdomservice.DeleteAssistantAssociationInput, ...request.Option) (*connectwisdomservice.DeleteAssistantAssociationOutput, error)
	DeleteAssistantAssociationRequest(*connectwisdomservice.DeleteAssistantAssociationInput) (*request.Request, *connectwisdomservice.DeleteAssistantAssociationOutput)

	DeleteContent(*connectwisdomservice.DeleteContentInput) (*connectwisdomservice.DeleteContentOutput, error)
	DeleteContentWithContext(aws.Context, *connectwisdomservice.DeleteContentInput, ...request.Option) (*connectwisdomservice.DeleteContentOutput, error)
	DeleteContentRequest(*connectwisdomservice.DeleteContentInput) (*request.Request, *connectwisdomservice.DeleteContentOutput)

	DeleteKnowledgeBase(*connectwisdomservice.DeleteKnowledgeBaseInput) (*connectwisdomservice.DeleteKnowledgeBaseOutput, error)
	DeleteKnowledgeBaseWithContext(aws.Context, *connectwisdomservice.DeleteKnowledgeBaseInput, ...request.Option) (*connectwisdomservice.DeleteKnowledgeBaseOutput, error)
	DeleteKnowledgeBaseRequest(*connectwisdomservice.DeleteKnowledgeBaseInput) (*request.Request, *connectwisdomservice.DeleteKnowledgeBaseOutput)

	GetAssistant(*connectwisdomservice.GetAssistantInput) (*connectwisdomservice.GetAssistantOutput, error)
	GetAssistantWithContext(aws.Context, *connectwisdomservice.GetAssistantInput, ...request.Option) (*connectwisdomservice.GetAssistantOutput, error)
	GetAssistantRequest(*connectwisdomservice.GetAssistantInput) (*request.Request, *connectwisdomservice.GetAssistantOutput)

	GetAssistantAssociation(*connectwisdomservice.GetAssistantAssociationInput) (*connectwisdomservice.GetAssistantAssociationOutput, error)
	GetAssistantAssociationWithContext(aws.Context, *connectwisdomservice.GetAssistantAssociationInput, ...request.Option) (*connectwisdomservice.GetAssistantAssociationOutput, error)
	GetAssistantAssociationRequest(*connectwisdomservice.GetAssistantAssociationInput) (*request.Request, *connectwisdomservice.GetAssistantAssociationOutput)

	GetContent(*connectwisdomservice.GetContentInput) (*connectwisdomservice.GetContentOutput, error)
	GetContentWithContext(aws.Context, *connectwisdomservice.GetContentInput, ...request.Option) (*connectwisdomservice.GetContentOutput, error)
	GetContentRequest(*connectwisdomservice.GetContentInput) (*request.Request, *connectwisdomservice.GetContentOutput)

	GetContentSummary(*connectwisdomservice.GetContentSummaryInput) (*connectwisdomservice.GetContentSummaryOutput, error)
	GetContentSummaryWithContext(aws.Context, *connectwisdomservice.GetContentSummaryInput, ...request.Option) (*connectwisdomservice.GetContentSummaryOutput, error)
	GetContentSummaryRequest(*connectwisdomservice.GetContentSummaryInput) (*request.Request, *connectwisdomservice.GetContentSummaryOutput)

	GetKnowledgeBase(*connectwisdomservice.GetKnowledgeBaseInput) (*connectwisdomservice.GetKnowledgeBaseOutput, error)
	GetKnowledgeBaseWithContext(aws.Context, *connectwisdomservice.GetKnowledgeBaseInput, ...request.Option) (*connectwisdomservice.GetKnowledgeBaseOutput, error)
	GetKnowledgeBaseRequest(*connectwisdomservice.GetKnowledgeBaseInput) (*request.Request, *connectwisdomservice.GetKnowledgeBaseOutput)

	GetRecommendations(*connectwisdomservice.GetRecommendationsInput) (*connectwisdomservice.GetRecommendationsOutput, error)
	GetRecommendationsWithContext(aws.Context, *connectwisdomservice.GetRecommendationsInput, ...request.Option) (*connectwisdomservice.GetRecommendationsOutput, error)
	GetRecommendationsRequest(*connectwisdomservice.GetRecommendationsInput) (*request.Request, *connectwisdomservice.GetRecommendationsOutput)

	GetSession(*connectwisdomservice.GetSessionInput) (*connectwisdomservice.GetSessionOutput, error)
	GetSessionWithContext(aws.Context, *connectwisdomservice.GetSessionInput, ...request.Option) (*connectwisdomservice.GetSessionOutput, error)
	GetSessionRequest(*connectwisdomservice.GetSessionInput) (*request.Request, *connectwisdomservice.GetSessionOutput)

	ListAssistantAssociations(*connectwisdomservice.ListAssistantAssociationsInput) (*connectwisdomservice.ListAssistantAssociationsOutput, error)
	ListAssistantAssociationsWithContext(aws.Context, *connectwisdomservice.ListAssistantAssociationsInput, ...request.Option) (*connectwisdomservice.ListAssistantAssociationsOutput, error)
	ListAssistantAssociationsRequest(*connectwisdomservice.ListAssistantAssociationsInput) (*request.Request, *connectwisdomservice.ListAssistantAssociationsOutput)

	ListAssistantAssociationsPages(*connectwisdomservice.ListAssistantAssociationsInput, func(*connectwisdomservice.ListAssistantAssociationsOutput, bool) bool) error
	ListAssistantAssociationsPagesWithContext(aws.Context, *connectwisdomservice.ListAssistantAssociationsInput, func(*connectwisdomservice.ListAssistantAssociationsOutput, bool) bool, ...request.Option) error

	ListAssistants(*connectwisdomservice.ListAssistantsInput) (*connectwisdomservice.ListAssistantsOutput, error)
	ListAssistantsWithContext(aws.Context, *connectwisdomservice.ListAssistantsInput, ...request.Option) (*connectwisdomservice.ListAssistantsOutput, error)
	ListAssistantsRequest(*connectwisdomservice.ListAssistantsInput) (*request.Request, *connectwisdomservice.ListAssistantsOutput)

	ListAssistantsPages(*connectwisdomservice.ListAssistantsInput, func(*connectwisdomservice.ListAssistantsOutput, bool) bool) error
	ListAssistantsPagesWithContext(aws.Context, *connectwisdomservice.ListAssistantsInput, func(*connectwisdomservice.ListAssistantsOutput, bool) bool, ...request.Option) error

	ListContents(*connectwisdomservice.ListContentsInput) (*connectwisdomservice.ListContentsOutput, error)
	ListContentsWithContext(aws.Context, *connectwisdomservice.ListContentsInput, ...request.Option) (*connectwisdomservice.ListContentsOutput, error)
	ListContentsRequest(*connectwisdomservice.ListContentsInput) (*request.Request, *connectwisdomservice.ListContentsOutput)

	ListContentsPages(*connectwisdomservice.ListContentsInput, func(*connectwisdomservice.ListContentsOutput, bool) bool) error
	ListContentsPagesWithContext(aws.Context, *connectwisdomservice.ListContentsInput, func(*connectwisdomservice.ListContentsOutput, bool) bool, ...request.Option) error

	ListKnowledgeBases(*connectwisdomservice.ListKnowledgeBasesInput) (*connectwisdomservice.ListKnowledgeBasesOutput, error)
	ListKnowledgeBasesWithContext(aws.Context, *connectwisdomservice.ListKnowledgeBasesInput, ...request.Option) (*connectwisdomservice.ListKnowledgeBasesOutput, error)
	ListKnowledgeBasesRequest(*connectwisdomservice.ListKnowledgeBasesInput) (*request.Request, *connectwisdomservice.ListKnowledgeBasesOutput)

	ListKnowledgeBasesPages(*connectwisdomservice.ListKnowledgeBasesInput, func(*connectwisdomservice.ListKnowledgeBasesOutput, bool) bool) error
	ListKnowledgeBasesPagesWithContext(aws.Context, *connectwisdomservice.ListKnowledgeBasesInput, func(*connectwisdomservice.ListKnowledgeBasesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*connectwisdomservice.ListTagsForResourceInput) (*connectwisdomservice.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *connectwisdomservice.ListTagsForResourceInput, ...request.Option) (*connectwisdomservice.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*connectwisdomservice.ListTagsForResourceInput) (*request.Request, *connectwisdomservice.ListTagsForResourceOutput)

	NotifyRecommendationsReceived(*connectwisdomservice.NotifyRecommendationsReceivedInput) (*connectwisdomservice.NotifyRecommendationsReceivedOutput, error)
	NotifyRecommendationsReceivedWithContext(aws.Context, *connectwisdomservice.NotifyRecommendationsReceivedInput, ...request.Option) (*connectwisdomservice.NotifyRecommendationsReceivedOutput, error)
	NotifyRecommendationsReceivedRequest(*connectwisdomservice.NotifyRecommendationsReceivedInput) (*request.Request, *connectwisdomservice.NotifyRecommendationsReceivedOutput)

	QueryAssistant(*connectwisdomservice.QueryAssistantInput) (*connectwisdomservice.QueryAssistantOutput, error)
	QueryAssistantWithContext(aws.Context, *connectwisdomservice.QueryAssistantInput, ...request.Option) (*connectwisdomservice.QueryAssistantOutput, error)
	QueryAssistantRequest(*connectwisdomservice.QueryAssistantInput) (*request.Request, *connectwisdomservice.QueryAssistantOutput)

	QueryAssistantPages(*connectwisdomservice.QueryAssistantInput, func(*connectwisdomservice.QueryAssistantOutput, bool) bool) error
	QueryAssistantPagesWithContext(aws.Context, *connectwisdomservice.QueryAssistantInput, func(*connectwisdomservice.QueryAssistantOutput, bool) bool, ...request.Option) error

	RemoveKnowledgeBaseTemplateUri(*connectwisdomservice.RemoveKnowledgeBaseTemplateUriInput) (*connectwisdomservice.RemoveKnowledgeBaseTemplateUriOutput, error)
	RemoveKnowledgeBaseTemplateUriWithContext(aws.Context, *connectwisdomservice.RemoveKnowledgeBaseTemplateUriInput, ...request.Option) (*connectwisdomservice.RemoveKnowledgeBaseTemplateUriOutput, error)
	RemoveKnowledgeBaseTemplateUriRequest(*connectwisdomservice.RemoveKnowledgeBaseTemplateUriInput) (*request.Request, *connectwisdomservice.RemoveKnowledgeBaseTemplateUriOutput)

	SearchContent(*connectwisdomservice.SearchContentInput) (*connectwisdomservice.SearchContentOutput, error)
	SearchContentWithContext(aws.Context, *connectwisdomservice.SearchContentInput, ...request.Option) (*connectwisdomservice.SearchContentOutput, error)
	SearchContentRequest(*connectwisdomservice.SearchContentInput) (*request.Request, *connectwisdomservice.SearchContentOutput)

	SearchContentPages(*connectwisdomservice.SearchContentInput, func(*connectwisdomservice.SearchContentOutput, bool) bool) error
	SearchContentPagesWithContext(aws.Context, *connectwisdomservice.SearchContentInput, func(*connectwisdomservice.SearchContentOutput, bool) bool, ...request.Option) error

	SearchSessions(*connectwisdomservice.SearchSessionsInput) (*connectwisdomservice.SearchSessionsOutput, error)
	SearchSessionsWithContext(aws.Context, *connectwisdomservice.SearchSessionsInput, ...request.Option) (*connectwisdomservice.SearchSessionsOutput, error)
	SearchSessionsRequest(*connectwisdomservice.SearchSessionsInput) (*request.Request, *connectwisdomservice.SearchSessionsOutput)

	SearchSessionsPages(*connectwisdomservice.SearchSessionsInput, func(*connectwisdomservice.SearchSessionsOutput, bool) bool) error
	SearchSessionsPagesWithContext(aws.Context, *connectwisdomservice.SearchSessionsInput, func(*connectwisdomservice.SearchSessionsOutput, bool) bool, ...request.Option) error

	StartContentUpload(*connectwisdomservice.StartContentUploadInput) (*connectwisdomservice.StartContentUploadOutput, error)
	StartContentUploadWithContext(aws.Context, *connectwisdomservice.StartContentUploadInput, ...request.Option) (*connectwisdomservice.StartContentUploadOutput, error)
	StartContentUploadRequest(*connectwisdomservice.StartContentUploadInput) (*request.Request, *connectwisdomservice.StartContentUploadOutput)

	TagResource(*connectwisdomservice.TagResourceInput) (*connectwisdomservice.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *connectwisdomservice.TagResourceInput, ...request.Option) (*connectwisdomservice.TagResourceOutput, error)
	TagResourceRequest(*connectwisdomservice.TagResourceInput) (*request.Request, *connectwisdomservice.TagResourceOutput)

	UntagResource(*connectwisdomservice.UntagResourceInput) (*connectwisdomservice.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *connectwisdomservice.UntagResourceInput, ...request.Option) (*connectwisdomservice.UntagResourceOutput, error)
	UntagResourceRequest(*connectwisdomservice.UntagResourceInput) (*request.Request, *connectwisdomservice.UntagResourceOutput)

	UpdateContent(*connectwisdomservice.UpdateContentInput) (*connectwisdomservice.UpdateContentOutput, error)
	UpdateContentWithContext(aws.Context, *connectwisdomservice.UpdateContentInput, ...request.Option) (*connectwisdomservice.UpdateContentOutput, error)
	UpdateContentRequest(*connectwisdomservice.UpdateContentInput) (*request.Request, *connectwisdomservice.UpdateContentOutput)

	UpdateKnowledgeBaseTemplateUri(*connectwisdomservice.UpdateKnowledgeBaseTemplateUriInput) (*connectwisdomservice.UpdateKnowledgeBaseTemplateUriOutput, error)
	UpdateKnowledgeBaseTemplateUriWithContext(aws.Context, *connectwisdomservice.UpdateKnowledgeBaseTemplateUriInput, ...request.Option) (*connectwisdomservice.UpdateKnowledgeBaseTemplateUriOutput, error)
	UpdateKnowledgeBaseTemplateUriRequest(*connectwisdomservice.UpdateKnowledgeBaseTemplateUriInput) (*request.Request, *connectwisdomservice.UpdateKnowledgeBaseTemplateUriOutput)
}

var _ ConnectWisdomServiceAPI = (*connectwisdomservice.ConnectWisdomService)(nil)
