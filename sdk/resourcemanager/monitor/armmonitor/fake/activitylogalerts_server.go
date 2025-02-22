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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"net/http"
	"net/url"
	"regexp"
)

// ActivityLogAlertsServer is a fake server for instances of the armmonitor.ActivityLogAlertsClient type.
type ActivityLogAlertsServer struct {
	// CreateOrUpdate is the fake for method ActivityLogAlertsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertRule armmonitor.ActivityLogAlertResource, options *armmonitor.ActivityLogAlertsClientCreateOrUpdateOptions) (resp azfake.Responder[armmonitor.ActivityLogAlertsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ActivityLogAlertsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, activityLogAlertName string, options *armmonitor.ActivityLogAlertsClientDeleteOptions) (resp azfake.Responder[armmonitor.ActivityLogAlertsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ActivityLogAlertsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, activityLogAlertName string, options *armmonitor.ActivityLogAlertsClientGetOptions) (resp azfake.Responder[armmonitor.ActivityLogAlertsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method ActivityLogAlertsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmonitor.ActivityLogAlertsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmonitor.ActivityLogAlertsClientListByResourceGroupResponse])

	// NewListBySubscriptionIDPager is the fake for method ActivityLogAlertsClient.NewListBySubscriptionIDPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionIDPager func(options *armmonitor.ActivityLogAlertsClientListBySubscriptionIDOptions) (resp azfake.PagerResponder[armmonitor.ActivityLogAlertsClientListBySubscriptionIDResponse])

	// Update is the fake for method ActivityLogAlertsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertRulePatch armmonitor.AlertRulePatchObject, options *armmonitor.ActivityLogAlertsClientUpdateOptions) (resp azfake.Responder[armmonitor.ActivityLogAlertsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewActivityLogAlertsServerTransport creates a new instance of ActivityLogAlertsServerTransport with the provided implementation.
// The returned ActivityLogAlertsServerTransport instance is connected to an instance of armmonitor.ActivityLogAlertsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewActivityLogAlertsServerTransport(srv *ActivityLogAlertsServer) *ActivityLogAlertsServerTransport {
	return &ActivityLogAlertsServerTransport{
		srv:                          srv,
		newListByResourceGroupPager:  newTracker[azfake.PagerResponder[armmonitor.ActivityLogAlertsClientListByResourceGroupResponse]](),
		newListBySubscriptionIDPager: newTracker[azfake.PagerResponder[armmonitor.ActivityLogAlertsClientListBySubscriptionIDResponse]](),
	}
}

// ActivityLogAlertsServerTransport connects instances of armmonitor.ActivityLogAlertsClient to instances of ActivityLogAlertsServer.
// Don't use this type directly, use NewActivityLogAlertsServerTransport instead.
type ActivityLogAlertsServerTransport struct {
	srv                          *ActivityLogAlertsServer
	newListByResourceGroupPager  *tracker[azfake.PagerResponder[armmonitor.ActivityLogAlertsClientListByResourceGroupResponse]]
	newListBySubscriptionIDPager *tracker[azfake.PagerResponder[armmonitor.ActivityLogAlertsClientListBySubscriptionIDResponse]]
}

// Do implements the policy.Transporter interface for ActivityLogAlertsServerTransport.
func (a *ActivityLogAlertsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ActivityLogAlertsClient.CreateOrUpdate":
		resp, err = a.dispatchCreateOrUpdate(req)
	case "ActivityLogAlertsClient.Delete":
		resp, err = a.dispatchDelete(req)
	case "ActivityLogAlertsClient.Get":
		resp, err = a.dispatchGet(req)
	case "ActivityLogAlertsClient.NewListByResourceGroupPager":
		resp, err = a.dispatchNewListByResourceGroupPager(req)
	case "ActivityLogAlertsClient.NewListBySubscriptionIDPager":
		resp, err = a.dispatchNewListBySubscriptionIDPager(req)
	case "ActivityLogAlertsClient.Update":
		resp, err = a.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ActivityLogAlertsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/activityLogAlerts/(?P<activityLogAlertName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmonitor.ActivityLogAlertResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	activityLogAlertNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("activityLogAlertName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.CreateOrUpdate(req.Context(), resourceGroupNameUnescaped, activityLogAlertNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ActivityLogAlertResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ActivityLogAlertsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/activityLogAlerts/(?P<activityLogAlertName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	activityLogAlertNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("activityLogAlertName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Delete(req.Context(), resourceGroupNameUnescaped, activityLogAlertNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ActivityLogAlertsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/activityLogAlerts/(?P<activityLogAlertName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	activityLogAlertNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("activityLogAlertName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameUnescaped, activityLogAlertNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ActivityLogAlertResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ActivityLogAlertsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := a.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/activityLogAlerts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByResourceGroupPager(resourceGroupNameUnescaped, nil)
		newListByResourceGroupPager = &resp
		a.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmonitor.ActivityLogAlertsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		a.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (a *ActivityLogAlertsServerTransport) dispatchNewListBySubscriptionIDPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListBySubscriptionIDPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionIDPager not implemented")}
	}
	newListBySubscriptionIDPager := a.newListBySubscriptionIDPager.get(req)
	if newListBySubscriptionIDPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/activityLogAlerts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListBySubscriptionIDPager(nil)
		newListBySubscriptionIDPager = &resp
		a.newListBySubscriptionIDPager.add(req, newListBySubscriptionIDPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionIDPager, req, func(page *armmonitor.ActivityLogAlertsClientListBySubscriptionIDResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionIDPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListBySubscriptionIDPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionIDPager) {
		a.newListBySubscriptionIDPager.remove(req)
	}
	return resp, nil
}

func (a *ActivityLogAlertsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/activityLogAlerts/(?P<activityLogAlertName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmonitor.AlertRulePatchObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	activityLogAlertNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("activityLogAlertName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Update(req.Context(), resourceGroupNameUnescaped, activityLogAlertNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ActivityLogAlertResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
