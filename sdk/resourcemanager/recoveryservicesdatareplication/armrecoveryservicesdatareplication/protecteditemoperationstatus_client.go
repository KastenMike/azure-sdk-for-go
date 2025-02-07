//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesdatareplication

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

// ProtectedItemOperationStatusClient contains the methods for the ProtectedItemOperationStatus group.
// Don't use this type directly, use NewProtectedItemOperationStatusClient() instead.
type ProtectedItemOperationStatusClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProtectedItemOperationStatusClient creates a new instance of ProtectedItemOperationStatusClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProtectedItemOperationStatusClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProtectedItemOperationStatusClient, error) {
	cl, err := arm.NewClient(moduleName+".ProtectedItemOperationStatusClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProtectedItemOperationStatusClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Tracks the results of an asynchronous operation on the protected item.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-16-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The vault name.
//   - protectedItemName - The protected item name.
//   - operationID - The ID of an ongoing async operation.
//   - options - ProtectedItemOperationStatusClientGetOptions contains the optional parameters for the ProtectedItemOperationStatusClient.Get
//     method.
func (client *ProtectedItemOperationStatusClient) Get(ctx context.Context, resourceGroupName string, vaultName string, protectedItemName string, operationID string, options *ProtectedItemOperationStatusClientGetOptions) (ProtectedItemOperationStatusClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, protectedItemName, operationID, options)
	if err != nil {
		return ProtectedItemOperationStatusClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectedItemOperationStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProtectedItemOperationStatusClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProtectedItemOperationStatusClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, protectedItemName string, operationID string, options *ProtectedItemOperationStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults/{vaultName}/protectedItems/{protectedItemName}/operations/{operationId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if protectedItemName == "" {
		return nil, errors.New("parameter protectedItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectedItemName}", url.PathEscape(protectedItemName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProtectedItemOperationStatusClient) getHandleResponse(resp *http.Response) (ProtectedItemOperationStatusClientGetResponse, error) {
	result := ProtectedItemOperationStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatus); err != nil {
		return ProtectedItemOperationStatusClientGetResponse{}, err
	}
	return result, nil
}
