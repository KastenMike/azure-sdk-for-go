# Release History

## 1.2.0 (2023-10-27)
### Features Added

- New enum type `Operation` with values `OperationActive`, `OperationAddBegin`, `OperationAddComplete`, `OperationDeleteBegin`, `OperationDeleteComplete`
- New enum type `Status` with values `StatusActive`, `StatusDeleting`, `StatusFailed`, `StatusInProgress`
- New function `*ClientFactory.NewCreationSupportedClient() *CreationSupportedClient`
- New function `*ClientFactory.NewMonitoredSubscriptionsClient() *MonitoredSubscriptionsClient`
- New function `NewCreationSupportedClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CreationSupportedClient, error)`
- New function `*CreationSupportedClient.Get(context.Context, string, *CreationSupportedClientGetOptions) (CreationSupportedClientGetResponse, error)`
- New function `*CreationSupportedClient.NewListPager(string, *CreationSupportedClientListOptions) *runtime.Pager[CreationSupportedClientListResponse]`
- New function `NewMonitoredSubscriptionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*MonitoredSubscriptionsClient, error)`
- New function `*MonitoredSubscriptionsClient.BeginCreateorUpdate(context.Context, string, string, string, *MonitoredSubscriptionsClientBeginCreateorUpdateOptions) (*runtime.Poller[MonitoredSubscriptionsClientCreateorUpdateResponse], error)`
- New function `*MonitoredSubscriptionsClient.BeginDelete(context.Context, string, string, string, *MonitoredSubscriptionsClientBeginDeleteOptions) (*runtime.Poller[MonitoredSubscriptionsClientDeleteResponse], error)`
- New function `*MonitoredSubscriptionsClient.Get(context.Context, string, string, string, *MonitoredSubscriptionsClientGetOptions) (MonitoredSubscriptionsClientGetResponse, error)`
- New function `*MonitoredSubscriptionsClient.NewListPager(string, string, *MonitoredSubscriptionsClientListOptions) *runtime.Pager[MonitoredSubscriptionsClientListResponse]`
- New function `*MonitoredSubscriptionsClient.BeginUpdate(context.Context, string, string, string, *MonitoredSubscriptionsClientBeginUpdateOptions) (*runtime.Poller[MonitoredSubscriptionsClientUpdateResponse], error)`
- New struct `CreateResourceSupportedProperties`
- New struct `CreateResourceSupportedResponse`
- New struct `CreateResourceSupportedResponseList`
- New struct `MonitoredSubscription`
- New struct `MonitoredSubscriptionProperties`
- New struct `MonitoredSubscriptionPropertiesList`
- New struct `SubscriptionList`
- New field `Cspm` in struct `MonitorUpdateProperties`
- New field `Automuting` in struct `MonitoringTagRulesProperties`
- New field `Cspm` in struct `OrganizationProperties`


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-28)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).