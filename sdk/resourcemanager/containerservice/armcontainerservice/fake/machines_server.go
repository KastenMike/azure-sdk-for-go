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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
	"net/http"
	"net/url"
	"regexp"
)

// MachinesServer is a fake server for instances of the armcontainerservice.MachinesClient type.
type MachinesServer struct {
	// Get is the fake for method MachinesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, machineName string, options *armcontainerservice.MachinesClientGetOptions) (resp azfake.Responder[armcontainerservice.MachinesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method MachinesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, resourceName string, agentPoolName string, options *armcontainerservice.MachinesClientListOptions) (resp azfake.PagerResponder[armcontainerservice.MachinesClientListResponse])
}

// NewMachinesServerTransport creates a new instance of MachinesServerTransport with the provided implementation.
// The returned MachinesServerTransport instance is connected to an instance of armcontainerservice.MachinesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMachinesServerTransport(srv *MachinesServer) *MachinesServerTransport {
	return &MachinesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcontainerservice.MachinesClientListResponse]](),
	}
}

// MachinesServerTransport connects instances of armcontainerservice.MachinesClient to instances of MachinesServer.
// Don't use this type directly, use NewMachinesServerTransport instead.
type MachinesServerTransport struct {
	srv          *MachinesServer
	newListPager *tracker[azfake.PagerResponder[armcontainerservice.MachinesClientListResponse]]
}

// Do implements the policy.Transporter interface for MachinesServerTransport.
func (m *MachinesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "MachinesClient.Get":
		resp, err = m.dispatchGet(req)
	case "MachinesClient.NewListPager":
		resp, err = m.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *MachinesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools/(?P<agentPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	agentPoolNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("agentPoolName")])
	if err != nil {
		return nil, err
	}
	machineNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameUnescaped, resourceNameUnescaped, agentPoolNameUnescaped, machineNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Machine, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MachinesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.ContainerService/managedClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentPools/(?P<agentPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/machines`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		agentPoolNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("agentPoolName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListPager(resourceGroupNameUnescaped, resourceNameUnescaped, agentPoolNameUnescaped, nil)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcontainerservice.MachinesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}
