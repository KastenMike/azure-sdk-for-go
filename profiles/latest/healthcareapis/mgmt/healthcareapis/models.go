// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package healthcareapis

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/healthcareapis/mgmt/2020-03-15/healthcareapis"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Kind = original.Kind

const (
	Fhir     Kind = original.Fhir
	FhirR4   Kind = original.FhirR4
	FhirStu3 Kind = original.FhirStu3
)

type ManagedServiceIdentityType = original.ManagedServiceIdentityType

const (
	None           ManagedServiceIdentityType = original.None
	SystemAssigned ManagedServiceIdentityType = original.SystemAssigned
)

type OperationResultStatus = original.OperationResultStatus

const (
	Canceled  OperationResultStatus = original.Canceled
	Failed    OperationResultStatus = original.Failed
	Requested OperationResultStatus = original.Requested
	Running   OperationResultStatus = original.Running
	Succeeded OperationResultStatus = original.Succeeded
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted      ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCanceled      ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreating      ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting      ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateDeprovisioned ProvisioningState = original.ProvisioningStateDeprovisioned
	ProvisioningStateFailed        ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateSucceeded     ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating      ProvisioningState = original.ProvisioningStateUpdating
	ProvisioningStateVerifying     ProvisioningState = original.ProvisioningStateVerifying
)

type ServiceNameUnavailabilityReason = original.ServiceNameUnavailabilityReason

const (
	AlreadyExists ServiceNameUnavailabilityReason = original.AlreadyExists
	Invalid       ServiceNameUnavailabilityReason = original.Invalid
)

type BaseClient = original.BaseClient
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type ErrorDetails = original.ErrorDetails
type ErrorDetailsInternal = original.ErrorDetailsInternal
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResultsClient = original.OperationResultsClient
type OperationResultsDescription = original.OperationResultsDescription
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type ResourceIdentity = original.ResourceIdentity
type ServiceAccessPolicyEntry = original.ServiceAccessPolicyEntry
type ServiceAuthenticationConfigurationInfo = original.ServiceAuthenticationConfigurationInfo
type ServiceCorsConfigurationInfo = original.ServiceCorsConfigurationInfo
type ServiceCosmosDbConfigurationInfo = original.ServiceCosmosDbConfigurationInfo
type ServiceExportConfigurationInfo = original.ServiceExportConfigurationInfo
type ServicesClient = original.ServicesClient
type ServicesCreateOrUpdateFuture = original.ServicesCreateOrUpdateFuture
type ServicesDeleteFuture = original.ServicesDeleteFuture
type ServicesDescription = original.ServicesDescription
type ServicesDescriptionListResult = original.ServicesDescriptionListResult
type ServicesDescriptionListResultIterator = original.ServicesDescriptionListResultIterator
type ServicesDescriptionListResultPage = original.ServicesDescriptionListResultPage
type ServicesNameAvailabilityInfo = original.ServicesNameAvailabilityInfo
type ServicesPatchDescription = original.ServicesPatchDescription
type ServicesProperties = original.ServicesProperties
type ServicesUpdateFuture = original.ServicesUpdateFuture
type SetObject = original.SetObject

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationResultsClient(subscriptionID string) OperationResultsClient {
	return original.NewOperationResultsClient(subscriptionID)
}
func NewOperationResultsClientWithBaseURI(baseURI string, subscriptionID string) OperationResultsClient {
	return original.NewOperationResultsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewServicesDescriptionListResultIterator(page ServicesDescriptionListResultPage) ServicesDescriptionListResultIterator {
	return original.NewServicesDescriptionListResultIterator(page)
}
func NewServicesDescriptionListResultPage(getNextPage func(context.Context, ServicesDescriptionListResult) (ServicesDescriptionListResult, error)) ServicesDescriptionListResultPage {
	return original.NewServicesDescriptionListResultPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return original.PossibleManagedServiceIdentityTypeValues()
}
func PossibleOperationResultStatusValues() []OperationResultStatus {
	return original.PossibleOperationResultStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleServiceNameUnavailabilityReasonValues() []ServiceNameUnavailabilityReason {
	return original.PossibleServiceNameUnavailabilityReasonValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
