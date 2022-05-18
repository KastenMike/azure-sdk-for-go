//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// IntegrationServiceEnvironmentSKUsClient contains the methods for the IntegrationServiceEnvironmentSKUs group.
// Don't use this type directly, use NewIntegrationServiceEnvironmentSKUsClient() instead.
type IntegrationServiceEnvironmentSKUsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIntegrationServiceEnvironmentSKUsClient creates a new instance of IntegrationServiceEnvironmentSKUsClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIntegrationServiceEnvironmentSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationServiceEnvironmentSKUsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationServiceEnvironmentSKUsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Gets a list of integration service environment Skus.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroup - The resource group.
// integrationServiceEnvironmentName - The integration service environment name.
// options - IntegrationServiceEnvironmentSKUsClientListOptions contains the optional parameters for the IntegrationServiceEnvironmentSKUsClient.List
// method.
func (client *IntegrationServiceEnvironmentSKUsClient) NewListPager(resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentSKUsClientListOptions) *runtime.Pager[IntegrationServiceEnvironmentSKUsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationServiceEnvironmentSKUsClientListResponse]{
		More: func(page IntegrationServiceEnvironmentSKUsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationServiceEnvironmentSKUsClientListResponse) (IntegrationServiceEnvironmentSKUsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IntegrationServiceEnvironmentSKUsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IntegrationServiceEnvironmentSKUsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IntegrationServiceEnvironmentSKUsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *IntegrationServiceEnvironmentSKUsClient) listCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentSKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationServiceEnvironmentSKUsClient) listHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentSKUsClientListResponse, error) {
	result := IntegrationServiceEnvironmentSKUsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironmentSKUList); err != nil {
		return IntegrationServiceEnvironmentSKUsClientListResponse{}, err
	}
	return result, nil
}
