//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package shared

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/internal/errorinfo"
)

// NOTE: when adding a new context key type, it likely needs to be
// added to the deny-list of key types in ContextWithDeniedValues

// CtxWithHTTPHeaderKey is used as a context key for adding/retrieving http.Header.
type CtxWithHTTPHeaderKey struct{}

// CtxWithRetryOptionsKey is used as a context key for adding/retrieving RetryOptions.
type CtxWithRetryOptionsKey struct{}

// CtxWithCaptureResponse is used as a context key for retrieving the raw response.
type CtxWithCaptureResponse struct{}

// CtxWithTracingTracer is used as a context key for adding/retrieving tracing.Tracer.
type CtxWithTracingTracer struct{}

// CtxAPINameKey is used as a context key for adding/retrieving the API name.
type CtxAPINameKey struct{}

// Delay waits for the duration to elapse or the context to be cancelled.
func Delay(ctx context.Context, delay time.Duration) error {
	select {
	case <-time.After(delay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RetryAfter returns non-zero if the response contains a Retry-After header value.
func RetryAfter(resp *http.Response) time.Duration {
	if resp == nil {
		return 0
	}
	ra := resp.Header.Get(HeaderRetryAfter)
	if ra == "" {
		return 0
	}
	// retry-after values are expressed in either number of
	// seconds or an HTTP-date indicating when to try again
	if retryAfter, _ := strconv.Atoi(ra); retryAfter > 0 {
		return time.Duration(retryAfter) * time.Second
	} else if t, err := time.Parse(time.RFC1123, ra); err == nil {
		return time.Until(t)
	}
	return 0
}

// TypeOfT returns the type of the generic type param.
func TypeOfT[T any]() reflect.Type {
	// you can't, at present, obtain the type of
	// a type parameter, so this is the trick
	return reflect.TypeOf((*T)(nil)).Elem()
}

// TransportFunc is a helper to use a first-class func to satisfy the Transporter interface.
type TransportFunc func(*http.Request) (*http.Response, error)

// Do implements the Transporter interface for the TransportFunc type.
func (pf TransportFunc) Do(req *http.Request) (*http.Response, error) {
	return pf(req)
}

// ValidateModVer verifies that moduleVersion is a valid semver 2.0 string.
func ValidateModVer(moduleVersion string) error {
	modVerRegx := regexp.MustCompile(`^v\d+\.\d+\.\d+(?:-[a-zA-Z0-9_.-]+)?$`)
	if !modVerRegx.MatchString(moduleVersion) {
		return fmt.Errorf("malformed moduleVersion param value %s", moduleVersion)
	}
	return nil
}

// ExtractModuleName returns "module", "package.Client" from "module/package.Client" or
// "package", "package.Client" from "package.Client" when there's no "module/" prefix.
// If clientName is malformed, an error is returned.
func ExtractModuleName(clientName string) (string, string, error) {
	// uses unnamed capturing for "module", "package.Client", and "package"
	regex, err := regexp.Compile(`^(?:([a-z0-9]+)/)?(([a-z0-9]+)\.(?:[A-Za-z0-9]+))$`)
	if err != nil {
		return "", "", err
	}

	matches := regex.FindStringSubmatch(clientName)
	if len(matches) < 4 {
		return "", "", fmt.Errorf("malformed clientName %s", clientName)
	}

	// the first match is the entire string, the second is "module", the third is
	// "package.Client" and the fourth is "package".
	// if there was no "module/" prefix, the second match will be the empty string
	if matches[1] != "" {
		return matches[1], matches[2], nil
	}
	return matches[3], matches[2], nil
}

// ContextWithDeniedValues wraps an existing [context.Context], denying access to certain context values.
// Pipeline policies that create new requests to be sent down their own pipeline MUST wrap the caller's
// context with an instance of this type. This is to prevent context values from flowing across disjoint
// requests which can have unintended side-effects.
type ContextWithDeniedValues struct {
	context.Context
}

// Value implements part of the [context.Context] interface.
// It acts as a deny-list for certain context keys.
func (c *ContextWithDeniedValues) Value(key any) any {
	switch key.(type) {
	case CtxAPINameKey, CtxWithCaptureResponse, CtxWithHTTPHeaderKey, CtxWithRetryOptionsKey, CtxWithTracingTracer:
		return nil
	default:
		return c.Context.Value(key)
	}
}

// NonRetriableError marks the specified error as non-retriable.
func NonRetriableError(err error) error {
	return &nonRetriableError{err}
}

type nonRetriableError struct {
	error
}

func (p *nonRetriableError) Error() string {
	return p.error.Error()
}

func (*nonRetriableError) NonRetriable() {
	// marker method
}

func (p *nonRetriableError) Unwrap() error {
	return p.error
}

var _ errorinfo.NonRetriable = (*nonRetriableError)(nil)
