// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codepipelineiface provides an interface to enable mocking the AWS CodePipeline service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codepipelineiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/codepipeline"
)

// CodePipelineAPI provides an interface to enable mocking the
// codepipeline.CodePipeline service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CodePipeline.
//    func myFunc(svc codepipelineiface.CodePipelineAPI) bool {
//        // Make svc.AcknowledgeJob request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := codepipeline.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCodePipelineClient struct {
//        codepipelineiface.CodePipelineAPI
//    }
//    func (m *mockCodePipelineClient) AcknowledgeJob(input *codepipeline.AcknowledgeJobInput) (*codepipeline.AcknowledgeJobOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCodePipelineClient{}
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
type CodePipelineAPI interface {
	AcknowledgeJob(*codepipeline.AcknowledgeJobInput) (*codepipeline.AcknowledgeJobOutput, error)
	AcknowledgeJobWithContext(aws.Context, *codepipeline.AcknowledgeJobInput, ...request.Option) (*codepipeline.AcknowledgeJobOutput, error)
	AcknowledgeJobRequest(*codepipeline.AcknowledgeJobInput) (*request.Request, *codepipeline.AcknowledgeJobOutput)

	AcknowledgeThirdPartyJob(*codepipeline.AcknowledgeThirdPartyJobInput) (*codepipeline.AcknowledgeThirdPartyJobOutput, error)
	AcknowledgeThirdPartyJobWithContext(aws.Context, *codepipeline.AcknowledgeThirdPartyJobInput, ...request.Option) (*codepipeline.AcknowledgeThirdPartyJobOutput, error)
	AcknowledgeThirdPartyJobRequest(*codepipeline.AcknowledgeThirdPartyJobInput) (*request.Request, *codepipeline.AcknowledgeThirdPartyJobOutput)

	CreateCustomActionType(*codepipeline.CreateCustomActionTypeInput) (*codepipeline.CreateCustomActionTypeOutput, error)
	CreateCustomActionTypeWithContext(aws.Context, *codepipeline.CreateCustomActionTypeInput, ...request.Option) (*codepipeline.CreateCustomActionTypeOutput, error)
	CreateCustomActionTypeRequest(*codepipeline.CreateCustomActionTypeInput) (*request.Request, *codepipeline.CreateCustomActionTypeOutput)

	CreatePipeline(*codepipeline.CreatePipelineInput) (*codepipeline.CreatePipelineOutput, error)
	CreatePipelineWithContext(aws.Context, *codepipeline.CreatePipelineInput, ...request.Option) (*codepipeline.CreatePipelineOutput, error)
	CreatePipelineRequest(*codepipeline.CreatePipelineInput) (*request.Request, *codepipeline.CreatePipelineOutput)

	DeleteCustomActionType(*codepipeline.DeleteCustomActionTypeInput) (*codepipeline.DeleteCustomActionTypeOutput, error)
	DeleteCustomActionTypeWithContext(aws.Context, *codepipeline.DeleteCustomActionTypeInput, ...request.Option) (*codepipeline.DeleteCustomActionTypeOutput, error)
	DeleteCustomActionTypeRequest(*codepipeline.DeleteCustomActionTypeInput) (*request.Request, *codepipeline.DeleteCustomActionTypeOutput)

	DeletePipeline(*codepipeline.DeletePipelineInput) (*codepipeline.DeletePipelineOutput, error)
	DeletePipelineWithContext(aws.Context, *codepipeline.DeletePipelineInput, ...request.Option) (*codepipeline.DeletePipelineOutput, error)
	DeletePipelineRequest(*codepipeline.DeletePipelineInput) (*request.Request, *codepipeline.DeletePipelineOutput)

	DeleteWebhook(*codepipeline.DeleteWebhookInput) (*codepipeline.DeleteWebhookOutput, error)
	DeleteWebhookWithContext(aws.Context, *codepipeline.DeleteWebhookInput, ...request.Option) (*codepipeline.DeleteWebhookOutput, error)
	DeleteWebhookRequest(*codepipeline.DeleteWebhookInput) (*request.Request, *codepipeline.DeleteWebhookOutput)

	DeregisterWebhookWithThirdParty(*codepipeline.DeregisterWebhookWithThirdPartyInput) (*codepipeline.DeregisterWebhookWithThirdPartyOutput, error)
	DeregisterWebhookWithThirdPartyWithContext(aws.Context, *codepipeline.DeregisterWebhookWithThirdPartyInput, ...request.Option) (*codepipeline.DeregisterWebhookWithThirdPartyOutput, error)
	DeregisterWebhookWithThirdPartyRequest(*codepipeline.DeregisterWebhookWithThirdPartyInput) (*request.Request, *codepipeline.DeregisterWebhookWithThirdPartyOutput)

	DisableStageTransition(*codepipeline.DisableStageTransitionInput) (*codepipeline.DisableStageTransitionOutput, error)
	DisableStageTransitionWithContext(aws.Context, *codepipeline.DisableStageTransitionInput, ...request.Option) (*codepipeline.DisableStageTransitionOutput, error)
	DisableStageTransitionRequest(*codepipeline.DisableStageTransitionInput) (*request.Request, *codepipeline.DisableStageTransitionOutput)

	EnableStageTransition(*codepipeline.EnableStageTransitionInput) (*codepipeline.EnableStageTransitionOutput, error)
	EnableStageTransitionWithContext(aws.Context, *codepipeline.EnableStageTransitionInput, ...request.Option) (*codepipeline.EnableStageTransitionOutput, error)
	EnableStageTransitionRequest(*codepipeline.EnableStageTransitionInput) (*request.Request, *codepipeline.EnableStageTransitionOutput)

	GetActionType(*codepipeline.GetActionTypeInput) (*codepipeline.GetActionTypeOutput, error)
	GetActionTypeWithContext(aws.Context, *codepipeline.GetActionTypeInput, ...request.Option) (*codepipeline.GetActionTypeOutput, error)
	GetActionTypeRequest(*codepipeline.GetActionTypeInput) (*request.Request, *codepipeline.GetActionTypeOutput)

	GetJobDetails(*codepipeline.GetJobDetailsInput) (*codepipeline.GetJobDetailsOutput, error)
	GetJobDetailsWithContext(aws.Context, *codepipeline.GetJobDetailsInput, ...request.Option) (*codepipeline.GetJobDetailsOutput, error)
	GetJobDetailsRequest(*codepipeline.GetJobDetailsInput) (*request.Request, *codepipeline.GetJobDetailsOutput)

	GetPipeline(*codepipeline.GetPipelineInput) (*codepipeline.GetPipelineOutput, error)
	GetPipelineWithContext(aws.Context, *codepipeline.GetPipelineInput, ...request.Option) (*codepipeline.GetPipelineOutput, error)
	GetPipelineRequest(*codepipeline.GetPipelineInput) (*request.Request, *codepipeline.GetPipelineOutput)

	GetPipelineExecution(*codepipeline.GetPipelineExecutionInput) (*codepipeline.GetPipelineExecutionOutput, error)
	GetPipelineExecutionWithContext(aws.Context, *codepipeline.GetPipelineExecutionInput, ...request.Option) (*codepipeline.GetPipelineExecutionOutput, error)
	GetPipelineExecutionRequest(*codepipeline.GetPipelineExecutionInput) (*request.Request, *codepipeline.GetPipelineExecutionOutput)

	GetPipelineState(*codepipeline.GetPipelineStateInput) (*codepipeline.GetPipelineStateOutput, error)
	GetPipelineStateWithContext(aws.Context, *codepipeline.GetPipelineStateInput, ...request.Option) (*codepipeline.GetPipelineStateOutput, error)
	GetPipelineStateRequest(*codepipeline.GetPipelineStateInput) (*request.Request, *codepipeline.GetPipelineStateOutput)

	GetThirdPartyJobDetails(*codepipeline.GetThirdPartyJobDetailsInput) (*codepipeline.GetThirdPartyJobDetailsOutput, error)
	GetThirdPartyJobDetailsWithContext(aws.Context, *codepipeline.GetThirdPartyJobDetailsInput, ...request.Option) (*codepipeline.GetThirdPartyJobDetailsOutput, error)
	GetThirdPartyJobDetailsRequest(*codepipeline.GetThirdPartyJobDetailsInput) (*request.Request, *codepipeline.GetThirdPartyJobDetailsOutput)

	ListActionExecutions(*codepipeline.ListActionExecutionsInput) (*codepipeline.ListActionExecutionsOutput, error)
	ListActionExecutionsWithContext(aws.Context, *codepipeline.ListActionExecutionsInput, ...request.Option) (*codepipeline.ListActionExecutionsOutput, error)
	ListActionExecutionsRequest(*codepipeline.ListActionExecutionsInput) (*request.Request, *codepipeline.ListActionExecutionsOutput)

	ListActionExecutionsPages(*codepipeline.ListActionExecutionsInput, func(*codepipeline.ListActionExecutionsOutput, bool) bool) error
	ListActionExecutionsPagesWithContext(aws.Context, *codepipeline.ListActionExecutionsInput, func(*codepipeline.ListActionExecutionsOutput, bool) bool, ...request.Option) error

	ListActionTypes(*codepipeline.ListActionTypesInput) (*codepipeline.ListActionTypesOutput, error)
	ListActionTypesWithContext(aws.Context, *codepipeline.ListActionTypesInput, ...request.Option) (*codepipeline.ListActionTypesOutput, error)
	ListActionTypesRequest(*codepipeline.ListActionTypesInput) (*request.Request, *codepipeline.ListActionTypesOutput)

	ListActionTypesPages(*codepipeline.ListActionTypesInput, func(*codepipeline.ListActionTypesOutput, bool) bool) error
	ListActionTypesPagesWithContext(aws.Context, *codepipeline.ListActionTypesInput, func(*codepipeline.ListActionTypesOutput, bool) bool, ...request.Option) error

	ListPipelineExecutions(*codepipeline.ListPipelineExecutionsInput) (*codepipeline.ListPipelineExecutionsOutput, error)
	ListPipelineExecutionsWithContext(aws.Context, *codepipeline.ListPipelineExecutionsInput, ...request.Option) (*codepipeline.ListPipelineExecutionsOutput, error)
	ListPipelineExecutionsRequest(*codepipeline.ListPipelineExecutionsInput) (*request.Request, *codepipeline.ListPipelineExecutionsOutput)

	ListPipelineExecutionsPages(*codepipeline.ListPipelineExecutionsInput, func(*codepipeline.ListPipelineExecutionsOutput, bool) bool) error
	ListPipelineExecutionsPagesWithContext(aws.Context, *codepipeline.ListPipelineExecutionsInput, func(*codepipeline.ListPipelineExecutionsOutput, bool) bool, ...request.Option) error

	ListPipelines(*codepipeline.ListPipelinesInput) (*codepipeline.ListPipelinesOutput, error)
	ListPipelinesWithContext(aws.Context, *codepipeline.ListPipelinesInput, ...request.Option) (*codepipeline.ListPipelinesOutput, error)
	ListPipelinesRequest(*codepipeline.ListPipelinesInput) (*request.Request, *codepipeline.ListPipelinesOutput)

	ListPipelinesPages(*codepipeline.ListPipelinesInput, func(*codepipeline.ListPipelinesOutput, bool) bool) error
	ListPipelinesPagesWithContext(aws.Context, *codepipeline.ListPipelinesInput, func(*codepipeline.ListPipelinesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*codepipeline.ListTagsForResourceInput) (*codepipeline.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *codepipeline.ListTagsForResourceInput, ...request.Option) (*codepipeline.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*codepipeline.ListTagsForResourceInput) (*request.Request, *codepipeline.ListTagsForResourceOutput)

	ListTagsForResourcePages(*codepipeline.ListTagsForResourceInput, func(*codepipeline.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *codepipeline.ListTagsForResourceInput, func(*codepipeline.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	ListWebhooks(*codepipeline.ListWebhooksInput) (*codepipeline.ListWebhooksOutput, error)
	ListWebhooksWithContext(aws.Context, *codepipeline.ListWebhooksInput, ...request.Option) (*codepipeline.ListWebhooksOutput, error)
	ListWebhooksRequest(*codepipeline.ListWebhooksInput) (*request.Request, *codepipeline.ListWebhooksOutput)

	ListWebhooksPages(*codepipeline.ListWebhooksInput, func(*codepipeline.ListWebhooksOutput, bool) bool) error
	ListWebhooksPagesWithContext(aws.Context, *codepipeline.ListWebhooksInput, func(*codepipeline.ListWebhooksOutput, bool) bool, ...request.Option) error

	PollForJobs(*codepipeline.PollForJobsInput) (*codepipeline.PollForJobsOutput, error)
	PollForJobsWithContext(aws.Context, *codepipeline.PollForJobsInput, ...request.Option) (*codepipeline.PollForJobsOutput, error)
	PollForJobsRequest(*codepipeline.PollForJobsInput) (*request.Request, *codepipeline.PollForJobsOutput)

	PollForThirdPartyJobs(*codepipeline.PollForThirdPartyJobsInput) (*codepipeline.PollForThirdPartyJobsOutput, error)
	PollForThirdPartyJobsWithContext(aws.Context, *codepipeline.PollForThirdPartyJobsInput, ...request.Option) (*codepipeline.PollForThirdPartyJobsOutput, error)
	PollForThirdPartyJobsRequest(*codepipeline.PollForThirdPartyJobsInput) (*request.Request, *codepipeline.PollForThirdPartyJobsOutput)

	PutActionRevision(*codepipeline.PutActionRevisionInput) (*codepipeline.PutActionRevisionOutput, error)
	PutActionRevisionWithContext(aws.Context, *codepipeline.PutActionRevisionInput, ...request.Option) (*codepipeline.PutActionRevisionOutput, error)
	PutActionRevisionRequest(*codepipeline.PutActionRevisionInput) (*request.Request, *codepipeline.PutActionRevisionOutput)

	PutApprovalResult(*codepipeline.PutApprovalResultInput) (*codepipeline.PutApprovalResultOutput, error)
	PutApprovalResultWithContext(aws.Context, *codepipeline.PutApprovalResultInput, ...request.Option) (*codepipeline.PutApprovalResultOutput, error)
	PutApprovalResultRequest(*codepipeline.PutApprovalResultInput) (*request.Request, *codepipeline.PutApprovalResultOutput)

	PutJobFailureResult(*codepipeline.PutJobFailureResultInput) (*codepipeline.PutJobFailureResultOutput, error)
	PutJobFailureResultWithContext(aws.Context, *codepipeline.PutJobFailureResultInput, ...request.Option) (*codepipeline.PutJobFailureResultOutput, error)
	PutJobFailureResultRequest(*codepipeline.PutJobFailureResultInput) (*request.Request, *codepipeline.PutJobFailureResultOutput)

	PutJobSuccessResult(*codepipeline.PutJobSuccessResultInput) (*codepipeline.PutJobSuccessResultOutput, error)
	PutJobSuccessResultWithContext(aws.Context, *codepipeline.PutJobSuccessResultInput, ...request.Option) (*codepipeline.PutJobSuccessResultOutput, error)
	PutJobSuccessResultRequest(*codepipeline.PutJobSuccessResultInput) (*request.Request, *codepipeline.PutJobSuccessResultOutput)

	PutThirdPartyJobFailureResult(*codepipeline.PutThirdPartyJobFailureResultInput) (*codepipeline.PutThirdPartyJobFailureResultOutput, error)
	PutThirdPartyJobFailureResultWithContext(aws.Context, *codepipeline.PutThirdPartyJobFailureResultInput, ...request.Option) (*codepipeline.PutThirdPartyJobFailureResultOutput, error)
	PutThirdPartyJobFailureResultRequest(*codepipeline.PutThirdPartyJobFailureResultInput) (*request.Request, *codepipeline.PutThirdPartyJobFailureResultOutput)

	PutThirdPartyJobSuccessResult(*codepipeline.PutThirdPartyJobSuccessResultInput) (*codepipeline.PutThirdPartyJobSuccessResultOutput, error)
	PutThirdPartyJobSuccessResultWithContext(aws.Context, *codepipeline.PutThirdPartyJobSuccessResultInput, ...request.Option) (*codepipeline.PutThirdPartyJobSuccessResultOutput, error)
	PutThirdPartyJobSuccessResultRequest(*codepipeline.PutThirdPartyJobSuccessResultInput) (*request.Request, *codepipeline.PutThirdPartyJobSuccessResultOutput)

	PutWebhook(*codepipeline.PutWebhookInput) (*codepipeline.PutWebhookOutput, error)
	PutWebhookWithContext(aws.Context, *codepipeline.PutWebhookInput, ...request.Option) (*codepipeline.PutWebhookOutput, error)
	PutWebhookRequest(*codepipeline.PutWebhookInput) (*request.Request, *codepipeline.PutWebhookOutput)

	RegisterWebhookWithThirdParty(*codepipeline.RegisterWebhookWithThirdPartyInput) (*codepipeline.RegisterWebhookWithThirdPartyOutput, error)
	RegisterWebhookWithThirdPartyWithContext(aws.Context, *codepipeline.RegisterWebhookWithThirdPartyInput, ...request.Option) (*codepipeline.RegisterWebhookWithThirdPartyOutput, error)
	RegisterWebhookWithThirdPartyRequest(*codepipeline.RegisterWebhookWithThirdPartyInput) (*request.Request, *codepipeline.RegisterWebhookWithThirdPartyOutput)

	RetryStageExecution(*codepipeline.RetryStageExecutionInput) (*codepipeline.RetryStageExecutionOutput, error)
	RetryStageExecutionWithContext(aws.Context, *codepipeline.RetryStageExecutionInput, ...request.Option) (*codepipeline.RetryStageExecutionOutput, error)
	RetryStageExecutionRequest(*codepipeline.RetryStageExecutionInput) (*request.Request, *codepipeline.RetryStageExecutionOutput)

	StartPipelineExecution(*codepipeline.StartPipelineExecutionInput) (*codepipeline.StartPipelineExecutionOutput, error)
	StartPipelineExecutionWithContext(aws.Context, *codepipeline.StartPipelineExecutionInput, ...request.Option) (*codepipeline.StartPipelineExecutionOutput, error)
	StartPipelineExecutionRequest(*codepipeline.StartPipelineExecutionInput) (*request.Request, *codepipeline.StartPipelineExecutionOutput)

	StopPipelineExecution(*codepipeline.StopPipelineExecutionInput) (*codepipeline.StopPipelineExecutionOutput, error)
	StopPipelineExecutionWithContext(aws.Context, *codepipeline.StopPipelineExecutionInput, ...request.Option) (*codepipeline.StopPipelineExecutionOutput, error)
	StopPipelineExecutionRequest(*codepipeline.StopPipelineExecutionInput) (*request.Request, *codepipeline.StopPipelineExecutionOutput)

	TagResource(*codepipeline.TagResourceInput) (*codepipeline.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *codepipeline.TagResourceInput, ...request.Option) (*codepipeline.TagResourceOutput, error)
	TagResourceRequest(*codepipeline.TagResourceInput) (*request.Request, *codepipeline.TagResourceOutput)

	UntagResource(*codepipeline.UntagResourceInput) (*codepipeline.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *codepipeline.UntagResourceInput, ...request.Option) (*codepipeline.UntagResourceOutput, error)
	UntagResourceRequest(*codepipeline.UntagResourceInput) (*request.Request, *codepipeline.UntagResourceOutput)

	UpdateActionType(*codepipeline.UpdateActionTypeInput) (*codepipeline.UpdateActionTypeOutput, error)
	UpdateActionTypeWithContext(aws.Context, *codepipeline.UpdateActionTypeInput, ...request.Option) (*codepipeline.UpdateActionTypeOutput, error)
	UpdateActionTypeRequest(*codepipeline.UpdateActionTypeInput) (*request.Request, *codepipeline.UpdateActionTypeOutput)

	UpdatePipeline(*codepipeline.UpdatePipelineInput) (*codepipeline.UpdatePipelineOutput, error)
	UpdatePipelineWithContext(aws.Context, *codepipeline.UpdatePipelineInput, ...request.Option) (*codepipeline.UpdatePipelineOutput, error)
	UpdatePipelineRequest(*codepipeline.UpdatePipelineInput) (*request.Request, *codepipeline.UpdatePipelineOutput)
}

var _ CodePipelineAPI = (*codepipeline.CodePipeline)(nil)
