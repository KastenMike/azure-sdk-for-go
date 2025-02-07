//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ApplicationClient contains the methods for the ApplicationClient group.
// Don't use this type directly, use NewApplicationClient() instead.
type ApplicationClient struct {
	internal *arm.Client
}

// NewApplicationClient creates a new instance of ApplicationClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplicationClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationClient, error) {
	cl, err := arm.NewClient(moduleName+".ApplicationClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApplicationClient{
		internal: cl,
	}
	return client, nil
}

// NewListOperationsPager - Lists all of the available Microsoft.Solutions REST API operations.
//
// Generated from API version 2021-07-01
//   - options - ApplicationClientListOperationsOptions contains the optional parameters for the ApplicationClient.NewListOperationsPager
//     method.
func (client *ApplicationClient) NewListOperationsPager(options *ApplicationClientListOperationsOptions) *runtime.Pager[ApplicationClientListOperationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationClientListOperationsResponse]{
		More: func(page ApplicationClientListOperationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationClientListOperationsResponse) (ApplicationClientListOperationsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listOperationsCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationClientListOperationsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ApplicationClientListOperationsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationClientListOperationsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listOperationsHandleResponse(resp)
		},
	})
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *ApplicationClient) listOperationsCreateRequest(ctx context.Context, options *ApplicationClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Solutions/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *ApplicationClient) listOperationsHandleResponse(resp *http.Response) (ApplicationClientListOperationsResponse, error) {
	result := ApplicationClientListOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return ApplicationClientListOperationsResponse{}, err
	}
	return result, nil
}
