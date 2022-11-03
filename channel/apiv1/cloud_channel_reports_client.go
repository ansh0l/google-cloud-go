// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package channel

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	channelpb "cloud.google.com/go/channel/apiv1/channelpb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newCloudChannelReportsClientHook clientHook

// CloudChannelReportsCallOptions contains the retry settings for each method of CloudChannelReportsClient.
type CloudChannelReportsCallOptions struct {
	RunReportJob       []gax.CallOption
	FetchReportResults []gax.CallOption
	ListReports        []gax.CallOption
	CancelOperation    []gax.CallOption
	DeleteOperation    []gax.CallOption
	GetOperation       []gax.CallOption
	ListOperations     []gax.CallOption
}

func defaultCloudChannelReportsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudchannel.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudchannel.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://cloudchannel.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCloudChannelReportsCallOptions() *CloudChannelReportsCallOptions {
	return &CloudChannelReportsCallOptions{
		RunReportJob:       []gax.CallOption{},
		FetchReportResults: []gax.CallOption{},
		ListReports:        []gax.CallOption{},
		CancelOperation:    []gax.CallOption{},
		DeleteOperation:    []gax.CallOption{},
		GetOperation:       []gax.CallOption{},
		ListOperations:     []gax.CallOption{},
	}
}

// internalCloudChannelReportsClient is an interface that defines the methods available from Cloud Channel API.
type internalCloudChannelReportsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	RunReportJob(context.Context, *channelpb.RunReportJobRequest, ...gax.CallOption) (*RunReportJobOperation, error)
	RunReportJobOperation(name string) *RunReportJobOperation
	FetchReportResults(context.Context, *channelpb.FetchReportResultsRequest, ...gax.CallOption) *RowIterator
	ListReports(context.Context, *channelpb.ListReportsRequest, ...gax.CallOption) *ReportIterator
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// CloudChannelReportsClient is a client for interacting with Cloud Channel API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// CloudChannelReportsService lets Google Cloud resellers and
// distributors retrieve and combine a variety of data in Cloud Channel for
// multiple products (Google Cloud Platform (GCP), Google Voice, and
// Google Workspace.)
type CloudChannelReportsClient struct {
	// The internal transport-dependent client.
	internalClient internalCloudChannelReportsClient

	// The call options for this service.
	CallOptions *CloudChannelReportsCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CloudChannelReportsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CloudChannelReportsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CloudChannelReportsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// RunReportJob begins generation of data for a given report. The report
// identifier is a UID (for example, 613bf59q).
//
// Possible error codes:
//
//	PERMISSION_DENIED: The user doesn’t have access to this report.
//
//	INVALID_ARGUMENT: Required request parameters are missing
//	or invalid.
//
//	NOT_FOUND: The report identifier was not found.
//
//	INTERNAL: Any non-user error related to a technical issue
//	in the backend. Contact Cloud Channel support.
//
//	UNKNOWN: Any non-user error related to a technical issue
//	in the backend. Contact Cloud Channel support.
//
// Return value:
// The ID of a long-running operation.
//
// To get the results of the operation, call the GetOperation method of
// CloudChannelOperationsService. The Operation metadata contains an
// instance of OperationMetadata.
//
// To get the results of report generation, call
// CloudChannelReportsService.FetchReportResults with the
// RunReportJobResponse.report_job.
func (c *CloudChannelReportsClient) RunReportJob(ctx context.Context, req *channelpb.RunReportJobRequest, opts ...gax.CallOption) (*RunReportJobOperation, error) {
	return c.internalClient.RunReportJob(ctx, req, opts...)
}

// RunReportJobOperation returns a new RunReportJobOperation from a given name.
// The name must be that of a previously created RunReportJobOperation, possibly from a different process.
func (c *CloudChannelReportsClient) RunReportJobOperation(name string) *RunReportJobOperation {
	return c.internalClient.RunReportJobOperation(name)
}

// FetchReportResults retrieves data generated by CloudChannelReportsService.RunReportJob.
func (c *CloudChannelReportsClient) FetchReportResults(ctx context.Context, req *channelpb.FetchReportResultsRequest, opts ...gax.CallOption) *RowIterator {
	return c.internalClient.FetchReportResults(ctx, req, opts...)
}

// ListReports lists the reports that RunReportJob can run. These reports include an ID,
// a description, and the list of columns that will be in the result.
func (c *CloudChannelReportsClient) ListReports(ctx context.Context, req *channelpb.ListReportsRequest, opts ...gax.CallOption) *ReportIterator {
	return c.internalClient.ListReports(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *CloudChannelReportsClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// DeleteOperation is a utility method from google.longrunning.Operations.
func (c *CloudChannelReportsClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *CloudChannelReportsClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *CloudChannelReportsClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// cloudChannelReportsGRPCClient is a client for interacting with Cloud Channel API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type cloudChannelReportsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CloudChannelReportsClient
	CallOptions **CloudChannelReportsCallOptions

	// The gRPC API client.
	cloudChannelReportsClient channelpb.CloudChannelReportsServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCloudChannelReportsClient creates a new cloud channel reports service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// CloudChannelReportsService lets Google Cloud resellers and
// distributors retrieve and combine a variety of data in Cloud Channel for
// multiple products (Google Cloud Platform (GCP), Google Voice, and
// Google Workspace.)
func NewCloudChannelReportsClient(ctx context.Context, opts ...option.ClientOption) (*CloudChannelReportsClient, error) {
	clientOpts := defaultCloudChannelReportsGRPCClientOptions()
	if newCloudChannelReportsClientHook != nil {
		hookOpts, err := newCloudChannelReportsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CloudChannelReportsClient{CallOptions: defaultCloudChannelReportsCallOptions()}

	c := &cloudChannelReportsGRPCClient{
		connPool:                  connPool,
		disableDeadlines:          disableDeadlines,
		cloudChannelReportsClient: channelpb.NewCloudChannelReportsServiceClient(connPool),
		CallOptions:               &client.CallOptions,
		operationsClient:          longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *cloudChannelReportsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *cloudChannelReportsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *cloudChannelReportsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *cloudChannelReportsGRPCClient) RunReportJob(ctx context.Context, req *channelpb.RunReportJobRequest, opts ...gax.CallOption) (*RunReportJobOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RunReportJob[0:len((*c.CallOptions).RunReportJob):len((*c.CallOptions).RunReportJob)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.cloudChannelReportsClient.RunReportJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &RunReportJobOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *cloudChannelReportsGRPCClient) FetchReportResults(ctx context.Context, req *channelpb.FetchReportResultsRequest, opts ...gax.CallOption) *RowIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "report_job", url.QueryEscape(req.GetReportJob())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).FetchReportResults[0:len((*c.CallOptions).FetchReportResults):len((*c.CallOptions).FetchReportResults)], opts...)
	it := &RowIterator{}
	req = proto.Clone(req).(*channelpb.FetchReportResultsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*channelpb.Row, string, error) {
		resp := &channelpb.FetchReportResultsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.cloudChannelReportsClient.FetchReportResults(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetRows(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *cloudChannelReportsGRPCClient) ListReports(ctx context.Context, req *channelpb.ListReportsRequest, opts ...gax.CallOption) *ReportIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListReports[0:len((*c.CallOptions).ListReports):len((*c.CallOptions).ListReports)], opts...)
	it := &ReportIterator{}
	req = proto.Clone(req).(*channelpb.ListReportsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*channelpb.Report, string, error) {
		resp := &channelpb.ListReportsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.cloudChannelReportsClient.ListReports(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetReports(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *cloudChannelReportsGRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.CancelOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *cloudChannelReportsGRPCClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteOperation[0:len((*c.CallOptions).DeleteOperation):len((*c.CallOptions).DeleteOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.DeleteOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *cloudChannelReportsGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *cloudChannelReportsGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// RunReportJobOperation manages a long-running operation from RunReportJob.
type RunReportJobOperation struct {
	lro *longrunning.Operation
}

// RunReportJobOperation returns a new RunReportJobOperation from a given name.
// The name must be that of a previously created RunReportJobOperation, possibly from a different process.
func (c *cloudChannelReportsGRPCClient) RunReportJobOperation(name string) *RunReportJobOperation {
	return &RunReportJobOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *RunReportJobOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*channelpb.RunReportJobResponse, error) {
	var resp channelpb.RunReportJobResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *RunReportJobOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*channelpb.RunReportJobResponse, error) {
	var resp channelpb.RunReportJobResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *RunReportJobOperation) Metadata() (*channelpb.OperationMetadata, error) {
	var meta channelpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *RunReportJobOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *RunReportJobOperation) Name() string {
	return op.lro.Name()
}

// OperationIterator manages a stream of *longrunningpb.Operation.
type OperationIterator struct {
	items    []*longrunningpb.Operation
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*longrunningpb.Operation, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *OperationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *OperationIterator) Next() (*longrunningpb.Operation, error) {
	var item *longrunningpb.Operation
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *OperationIterator) bufLen() int {
	return len(it.items)
}

func (it *OperationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ReportIterator manages a stream of *channelpb.Report.
type ReportIterator struct {
	items    []*channelpb.Report
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*channelpb.Report, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ReportIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ReportIterator) Next() (*channelpb.Report, error) {
	var item *channelpb.Report
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ReportIterator) bufLen() int {
	return len(it.items)
}

func (it *ReportIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// RowIterator manages a stream of *channelpb.Row.
type RowIterator struct {
	items    []*channelpb.Row
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*channelpb.Row, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *RowIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *RowIterator) Next() (*channelpb.Row, error) {
	var item *channelpb.Row
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *RowIterator) bufLen() int {
	return len(it.items)
}

func (it *RowIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
