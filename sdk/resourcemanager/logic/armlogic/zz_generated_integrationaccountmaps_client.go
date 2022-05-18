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
	"strconv"
	"strings"
)

// IntegrationAccountMapsClient contains the methods for the IntegrationAccountMaps group.
// Don't use this type directly, use NewIntegrationAccountMapsClient() instead.
type IntegrationAccountMapsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIntegrationAccountMapsClient creates a new instance of IntegrationAccountMapsClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIntegrationAccountMapsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationAccountMapsClient, error) {
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
	client := &IntegrationAccountMapsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an integration account map. If the map is larger than 4 MB, you need to store the map
// in an Azure blob and use the blob's Shared Access Signature (SAS) URL as the 'contentLink'
// property value.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// mapName - The integration account map name.
// mapParam - The integration account map.
// options - IntegrationAccountMapsClientCreateOrUpdateOptions contains the optional parameters for the IntegrationAccountMapsClient.CreateOrUpdate
// method.
func (client *IntegrationAccountMapsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, mapParam IntegrationAccountMap, options *IntegrationAccountMapsClientCreateOrUpdateOptions) (IntegrationAccountMapsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, integrationAccountName, mapName, mapParam, options)
	if err != nil {
		return IntegrationAccountMapsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountMapsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IntegrationAccountMapsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationAccountMapsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, mapParam IntegrationAccountMap, options *IntegrationAccountMapsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if mapName == "" {
		return nil, errors.New("parameter mapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mapName}", url.PathEscape(mapName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, mapParam)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IntegrationAccountMapsClient) createOrUpdateHandleResponse(resp *http.Response) (IntegrationAccountMapsClientCreateOrUpdateResponse, error) {
	result := IntegrationAccountMapsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountMap); err != nil {
		return IntegrationAccountMapsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an integration account map.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// mapName - The integration account map name.
// options - IntegrationAccountMapsClientDeleteOptions contains the optional parameters for the IntegrationAccountMapsClient.Delete
// method.
func (client *IntegrationAccountMapsClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, options *IntegrationAccountMapsClientDeleteOptions) (IntegrationAccountMapsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, integrationAccountName, mapName, options)
	if err != nil {
		return IntegrationAccountMapsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountMapsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IntegrationAccountMapsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IntegrationAccountMapsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationAccountMapsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, options *IntegrationAccountMapsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if mapName == "" {
		return nil, errors.New("parameter mapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mapName}", url.PathEscape(mapName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an integration account map.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// mapName - The integration account map name.
// options - IntegrationAccountMapsClientGetOptions contains the optional parameters for the IntegrationAccountMapsClient.Get
// method.
func (client *IntegrationAccountMapsClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, options *IntegrationAccountMapsClientGetOptions) (IntegrationAccountMapsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, integrationAccountName, mapName, options)
	if err != nil {
		return IntegrationAccountMapsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountMapsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountMapsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationAccountMapsClient) getCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, options *IntegrationAccountMapsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if mapName == "" {
		return nil, errors.New("parameter mapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mapName}", url.PathEscape(mapName))
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

// getHandleResponse handles the Get response.
func (client *IntegrationAccountMapsClient) getHandleResponse(resp *http.Response) (IntegrationAccountMapsClientGetResponse, error) {
	result := IntegrationAccountMapsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountMap); err != nil {
		return IntegrationAccountMapsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of integration account maps.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// options - IntegrationAccountMapsClientListOptions contains the optional parameters for the IntegrationAccountMapsClient.List
// method.
func (client *IntegrationAccountMapsClient) NewListPager(resourceGroupName string, integrationAccountName string, options *IntegrationAccountMapsClientListOptions) *runtime.Pager[IntegrationAccountMapsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationAccountMapsClientListResponse]{
		More: func(page IntegrationAccountMapsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationAccountMapsClientListResponse) (IntegrationAccountMapsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IntegrationAccountMapsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IntegrationAccountMapsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IntegrationAccountMapsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *IntegrationAccountMapsClient) listCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountMapsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationAccountMapsClient) listHandleResponse(resp *http.Response) (IntegrationAccountMapsClientListResponse, error) {
	result := IntegrationAccountMapsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountMapListResult); err != nil {
		return IntegrationAccountMapsClientListResponse{}, err
	}
	return result, nil
}

// ListContentCallbackURL - Get the content callback url.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// mapName - The integration account map name.
// options - IntegrationAccountMapsClientListContentCallbackURLOptions contains the optional parameters for the IntegrationAccountMapsClient.ListContentCallbackURL
// method.
func (client *IntegrationAccountMapsClient) ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, listContentCallbackURL GetCallbackURLParameters, options *IntegrationAccountMapsClientListContentCallbackURLOptions) (IntegrationAccountMapsClientListContentCallbackURLResponse, error) {
	req, err := client.listContentCallbackURLCreateRequest(ctx, resourceGroupName, integrationAccountName, mapName, listContentCallbackURL, options)
	if err != nil {
		return IntegrationAccountMapsClientListContentCallbackURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountMapsClientListContentCallbackURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountMapsClientListContentCallbackURLResponse{}, runtime.NewResponseError(resp)
	}
	return client.listContentCallbackURLHandleResponse(resp)
}

// listContentCallbackURLCreateRequest creates the ListContentCallbackURL request.
func (client *IntegrationAccountMapsClient) listContentCallbackURLCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, mapName string, listContentCallbackURL GetCallbackURLParameters, options *IntegrationAccountMapsClientListContentCallbackURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}/listContentCallbackUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if mapName == "" {
		return nil, errors.New("parameter mapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mapName}", url.PathEscape(mapName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, listContentCallbackURL)
}

// listContentCallbackURLHandleResponse handles the ListContentCallbackURL response.
func (client *IntegrationAccountMapsClient) listContentCallbackURLHandleResponse(resp *http.Response) (IntegrationAccountMapsClientListContentCallbackURLResponse, error) {
	result := IntegrationAccountMapsClientListContentCallbackURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerCallbackURL); err != nil {
		return IntegrationAccountMapsClientListContentCallbackURLResponse{}, err
	}
	return result, nil
}
