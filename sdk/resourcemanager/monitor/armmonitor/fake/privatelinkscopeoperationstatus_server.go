//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinkScopeOperationStatusServer is a fake server for instances of the armmonitor.PrivateLinkScopeOperationStatusClient type.
type PrivateLinkScopeOperationStatusServer struct {
	// Get is the fake for method PrivateLinkScopeOperationStatusClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, asyncOperationID string, resourceGroupName string, options *armmonitor.PrivateLinkScopeOperationStatusClientGetOptions) (resp azfake.Responder[armmonitor.PrivateLinkScopeOperationStatusClientGetResponse], errResp azfake.ErrorResponder)
}

// NewPrivateLinkScopeOperationStatusServerTransport creates a new instance of PrivateLinkScopeOperationStatusServerTransport with the provided implementation.
// The returned PrivateLinkScopeOperationStatusServerTransport instance is connected to an instance of armmonitor.PrivateLinkScopeOperationStatusClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinkScopeOperationStatusServerTransport(srv *PrivateLinkScopeOperationStatusServer) *PrivateLinkScopeOperationStatusServerTransport {
	return &PrivateLinkScopeOperationStatusServerTransport{srv: srv}
}

// PrivateLinkScopeOperationStatusServerTransport connects instances of armmonitor.PrivateLinkScopeOperationStatusClient to instances of PrivateLinkScopeOperationStatusServer.
// Don't use this type directly, use NewPrivateLinkScopeOperationStatusServerTransport instead.
type PrivateLinkScopeOperationStatusServerTransport struct {
	srv *PrivateLinkScopeOperationStatusServer
}

// Do implements the policy.Transporter interface for PrivateLinkScopeOperationStatusServerTransport.
func (p *PrivateLinkScopeOperationStatusServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateLinkScopeOperationStatusClient.Get":
		resp, err = p.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateLinkScopeOperationStatusServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/microsoft.insights/privateLinkScopeOperationStatuses/(?P<asyncOperationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	asyncOperationIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("asyncOperationId")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), asyncOperationIDUnescaped, resourceGroupNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
