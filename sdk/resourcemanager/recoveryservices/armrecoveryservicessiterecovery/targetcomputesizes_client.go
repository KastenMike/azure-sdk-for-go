//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TargetComputeSizesClient contains the methods for the TargetComputeSizes group.
// Don't use this type directly, use NewTargetComputeSizesClient() instead.
type TargetComputeSizesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTargetComputeSizesClient creates a new instance of TargetComputeSizesClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTargetComputeSizesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TargetComputeSizesClient, error) {
	cl, err := arm.NewClient(moduleName+".TargetComputeSizesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TargetComputeSizesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByReplicationProtectedItemsPager - Lists the available target compute sizes for a replication protected item.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - protectionContainerName - protection container name.
//   - replicatedProtectedItemName - Replication protected item name.
//   - options - TargetComputeSizesClientListByReplicationProtectedItemsOptions contains the optional parameters for the TargetComputeSizesClient.NewListByReplicationProtectedItemsPager
//     method.
func (client *TargetComputeSizesClient) NewListByReplicationProtectedItemsPager(resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, replicatedProtectedItemName string, options *TargetComputeSizesClientListByReplicationProtectedItemsOptions) *runtime.Pager[TargetComputeSizesClientListByReplicationProtectedItemsResponse] {
	return runtime.NewPager(runtime.PagingHandler[TargetComputeSizesClientListByReplicationProtectedItemsResponse]{
		More: func(page TargetComputeSizesClientListByReplicationProtectedItemsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TargetComputeSizesClientListByReplicationProtectedItemsResponse) (TargetComputeSizesClientListByReplicationProtectedItemsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReplicationProtectedItemsCreateRequest(ctx, resourceName, resourceGroupName, fabricName, protectionContainerName, replicatedProtectedItemName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TargetComputeSizesClientListByReplicationProtectedItemsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TargetComputeSizesClientListByReplicationProtectedItemsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TargetComputeSizesClientListByReplicationProtectedItemsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReplicationProtectedItemsHandleResponse(resp)
		},
	})
}

// listByReplicationProtectedItemsCreateRequest creates the ListByReplicationProtectedItems request.
func (client *TargetComputeSizesClient) listByReplicationProtectedItemsCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, replicatedProtectedItemName string, options *TargetComputeSizesClientListByReplicationProtectedItemsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/targetComputeSizes"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if replicatedProtectedItemName == "" {
		return nil, errors.New("parameter replicatedProtectedItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicatedProtectedItemName}", url.PathEscape(replicatedProtectedItemName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationProtectedItemsHandleResponse handles the ListByReplicationProtectedItems response.
func (client *TargetComputeSizesClient) listByReplicationProtectedItemsHandleResponse(resp *http.Response) (TargetComputeSizesClientListByReplicationProtectedItemsResponse, error) {
	result := TargetComputeSizesClientListByReplicationProtectedItemsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TargetComputeSizeCollection); err != nil {
		return TargetComputeSizesClientListByReplicationProtectedItemsResponse{}, err
	}
	return result, nil
}
