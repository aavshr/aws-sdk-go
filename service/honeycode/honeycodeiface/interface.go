// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package honeycodeiface provides an interface to enable mocking the Amazon Honeycode service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package honeycodeiface

import (
	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/request"
	"github.com/aavshr/aws-sdk-go/service/honeycode"
)

// HoneycodeAPI provides an interface to enable mocking the
// honeycode.Honeycode service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Honeycode.
//    func myFunc(svc honeycodeiface.HoneycodeAPI) bool {
//        // Make svc.BatchCreateTableRows request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := honeycode.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockHoneycodeClient struct {
//        honeycodeiface.HoneycodeAPI
//    }
//    func (m *mockHoneycodeClient) BatchCreateTableRows(input *honeycode.BatchCreateTableRowsInput) (*honeycode.BatchCreateTableRowsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockHoneycodeClient{}
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
type HoneycodeAPI interface {
	BatchCreateTableRows(*honeycode.BatchCreateTableRowsInput) (*honeycode.BatchCreateTableRowsOutput, error)
	BatchCreateTableRowsWithContext(aws.Context, *honeycode.BatchCreateTableRowsInput, ...request.Option) (*honeycode.BatchCreateTableRowsOutput, error)
	BatchCreateTableRowsRequest(*honeycode.BatchCreateTableRowsInput) (*request.Request, *honeycode.BatchCreateTableRowsOutput)

	BatchDeleteTableRows(*honeycode.BatchDeleteTableRowsInput) (*honeycode.BatchDeleteTableRowsOutput, error)
	BatchDeleteTableRowsWithContext(aws.Context, *honeycode.BatchDeleteTableRowsInput, ...request.Option) (*honeycode.BatchDeleteTableRowsOutput, error)
	BatchDeleteTableRowsRequest(*honeycode.BatchDeleteTableRowsInput) (*request.Request, *honeycode.BatchDeleteTableRowsOutput)

	BatchUpdateTableRows(*honeycode.BatchUpdateTableRowsInput) (*honeycode.BatchUpdateTableRowsOutput, error)
	BatchUpdateTableRowsWithContext(aws.Context, *honeycode.BatchUpdateTableRowsInput, ...request.Option) (*honeycode.BatchUpdateTableRowsOutput, error)
	BatchUpdateTableRowsRequest(*honeycode.BatchUpdateTableRowsInput) (*request.Request, *honeycode.BatchUpdateTableRowsOutput)

	BatchUpsertTableRows(*honeycode.BatchUpsertTableRowsInput) (*honeycode.BatchUpsertTableRowsOutput, error)
	BatchUpsertTableRowsWithContext(aws.Context, *honeycode.BatchUpsertTableRowsInput, ...request.Option) (*honeycode.BatchUpsertTableRowsOutput, error)
	BatchUpsertTableRowsRequest(*honeycode.BatchUpsertTableRowsInput) (*request.Request, *honeycode.BatchUpsertTableRowsOutput)

	DescribeTableDataImportJob(*honeycode.DescribeTableDataImportJobInput) (*honeycode.DescribeTableDataImportJobOutput, error)
	DescribeTableDataImportJobWithContext(aws.Context, *honeycode.DescribeTableDataImportJobInput, ...request.Option) (*honeycode.DescribeTableDataImportJobOutput, error)
	DescribeTableDataImportJobRequest(*honeycode.DescribeTableDataImportJobInput) (*request.Request, *honeycode.DescribeTableDataImportJobOutput)

	GetScreenData(*honeycode.GetScreenDataInput) (*honeycode.GetScreenDataOutput, error)
	GetScreenDataWithContext(aws.Context, *honeycode.GetScreenDataInput, ...request.Option) (*honeycode.GetScreenDataOutput, error)
	GetScreenDataRequest(*honeycode.GetScreenDataInput) (*request.Request, *honeycode.GetScreenDataOutput)

	InvokeScreenAutomation(*honeycode.InvokeScreenAutomationInput) (*honeycode.InvokeScreenAutomationOutput, error)
	InvokeScreenAutomationWithContext(aws.Context, *honeycode.InvokeScreenAutomationInput, ...request.Option) (*honeycode.InvokeScreenAutomationOutput, error)
	InvokeScreenAutomationRequest(*honeycode.InvokeScreenAutomationInput) (*request.Request, *honeycode.InvokeScreenAutomationOutput)

	ListTableColumns(*honeycode.ListTableColumnsInput) (*honeycode.ListTableColumnsOutput, error)
	ListTableColumnsWithContext(aws.Context, *honeycode.ListTableColumnsInput, ...request.Option) (*honeycode.ListTableColumnsOutput, error)
	ListTableColumnsRequest(*honeycode.ListTableColumnsInput) (*request.Request, *honeycode.ListTableColumnsOutput)

	ListTableColumnsPages(*honeycode.ListTableColumnsInput, func(*honeycode.ListTableColumnsOutput, bool) bool) error
	ListTableColumnsPagesWithContext(aws.Context, *honeycode.ListTableColumnsInput, func(*honeycode.ListTableColumnsOutput, bool) bool, ...request.Option) error

	ListTableRows(*honeycode.ListTableRowsInput) (*honeycode.ListTableRowsOutput, error)
	ListTableRowsWithContext(aws.Context, *honeycode.ListTableRowsInput, ...request.Option) (*honeycode.ListTableRowsOutput, error)
	ListTableRowsRequest(*honeycode.ListTableRowsInput) (*request.Request, *honeycode.ListTableRowsOutput)

	ListTableRowsPages(*honeycode.ListTableRowsInput, func(*honeycode.ListTableRowsOutput, bool) bool) error
	ListTableRowsPagesWithContext(aws.Context, *honeycode.ListTableRowsInput, func(*honeycode.ListTableRowsOutput, bool) bool, ...request.Option) error

	ListTables(*honeycode.ListTablesInput) (*honeycode.ListTablesOutput, error)
	ListTablesWithContext(aws.Context, *honeycode.ListTablesInput, ...request.Option) (*honeycode.ListTablesOutput, error)
	ListTablesRequest(*honeycode.ListTablesInput) (*request.Request, *honeycode.ListTablesOutput)

	ListTablesPages(*honeycode.ListTablesInput, func(*honeycode.ListTablesOutput, bool) bool) error
	ListTablesPagesWithContext(aws.Context, *honeycode.ListTablesInput, func(*honeycode.ListTablesOutput, bool) bool, ...request.Option) error

	QueryTableRows(*honeycode.QueryTableRowsInput) (*honeycode.QueryTableRowsOutput, error)
	QueryTableRowsWithContext(aws.Context, *honeycode.QueryTableRowsInput, ...request.Option) (*honeycode.QueryTableRowsOutput, error)
	QueryTableRowsRequest(*honeycode.QueryTableRowsInput) (*request.Request, *honeycode.QueryTableRowsOutput)

	QueryTableRowsPages(*honeycode.QueryTableRowsInput, func(*honeycode.QueryTableRowsOutput, bool) bool) error
	QueryTableRowsPagesWithContext(aws.Context, *honeycode.QueryTableRowsInput, func(*honeycode.QueryTableRowsOutput, bool) bool, ...request.Option) error

	StartTableDataImportJob(*honeycode.StartTableDataImportJobInput) (*honeycode.StartTableDataImportJobOutput, error)
	StartTableDataImportJobWithContext(aws.Context, *honeycode.StartTableDataImportJobInput, ...request.Option) (*honeycode.StartTableDataImportJobOutput, error)
	StartTableDataImportJobRequest(*honeycode.StartTableDataImportJobInput) (*request.Request, *honeycode.StartTableDataImportJobOutput)
}

var _ HoneycodeAPI = (*honeycode.Honeycode)(nil)
