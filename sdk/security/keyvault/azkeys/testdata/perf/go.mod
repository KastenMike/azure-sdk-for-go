module github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys/testdata/perf

go 1.18

replace github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys => ../..

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.9.0-beta.1
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.3.1
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.4.0
	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys v1.0.0
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/internal v1.0.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.1.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.0.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
