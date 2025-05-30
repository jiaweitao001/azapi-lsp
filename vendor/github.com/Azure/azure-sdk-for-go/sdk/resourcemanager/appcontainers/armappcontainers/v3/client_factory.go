// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewAvailableWorkloadProfilesClient creates a new instance of AvailableWorkloadProfilesClient.
func (c *ClientFactory) NewAvailableWorkloadProfilesClient() *AvailableWorkloadProfilesClient {
	return &AvailableWorkloadProfilesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBillingMetersClient creates a new instance of BillingMetersClient.
func (c *ClientFactory) NewBillingMetersClient() *BillingMetersClient {
	return &BillingMetersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCertificatesClient creates a new instance of CertificatesClient.
func (c *ClientFactory) NewCertificatesClient() *CertificatesClient {
	return &CertificatesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConnectedEnvironmentsCertificatesClient creates a new instance of ConnectedEnvironmentsCertificatesClient.
func (c *ClientFactory) NewConnectedEnvironmentsCertificatesClient() *ConnectedEnvironmentsCertificatesClient {
	return &ConnectedEnvironmentsCertificatesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConnectedEnvironmentsClient creates a new instance of ConnectedEnvironmentsClient.
func (c *ClientFactory) NewConnectedEnvironmentsClient() *ConnectedEnvironmentsClient {
	return &ConnectedEnvironmentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConnectedEnvironmentsDaprComponentsClient creates a new instance of ConnectedEnvironmentsDaprComponentsClient.
func (c *ClientFactory) NewConnectedEnvironmentsDaprComponentsClient() *ConnectedEnvironmentsDaprComponentsClient {
	return &ConnectedEnvironmentsDaprComponentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConnectedEnvironmentsStoragesClient creates a new instance of ConnectedEnvironmentsStoragesClient.
func (c *ClientFactory) NewConnectedEnvironmentsStoragesClient() *ConnectedEnvironmentsStoragesClient {
	return &ConnectedEnvironmentsStoragesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewContainerAppsAPIClient creates a new instance of ContainerAppsAPIClient.
func (c *ClientFactory) NewContainerAppsAPIClient() *ContainerAppsAPIClient {
	return &ContainerAppsAPIClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewContainerAppsAuthConfigsClient creates a new instance of ContainerAppsAuthConfigsClient.
func (c *ClientFactory) NewContainerAppsAuthConfigsClient() *ContainerAppsAuthConfigsClient {
	return &ContainerAppsAuthConfigsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewContainerAppsClient creates a new instance of ContainerAppsClient.
func (c *ClientFactory) NewContainerAppsClient() *ContainerAppsClient {
	return &ContainerAppsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewContainerAppsDiagnosticsClient creates a new instance of ContainerAppsDiagnosticsClient.
func (c *ClientFactory) NewContainerAppsDiagnosticsClient() *ContainerAppsDiagnosticsClient {
	return &ContainerAppsDiagnosticsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewContainerAppsRevisionReplicasClient creates a new instance of ContainerAppsRevisionReplicasClient.
func (c *ClientFactory) NewContainerAppsRevisionReplicasClient() *ContainerAppsRevisionReplicasClient {
	return &ContainerAppsRevisionReplicasClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewContainerAppsRevisionsClient creates a new instance of ContainerAppsRevisionsClient.
func (c *ClientFactory) NewContainerAppsRevisionsClient() *ContainerAppsRevisionsClient {
	return &ContainerAppsRevisionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewContainerAppsSessionPoolsClient creates a new instance of ContainerAppsSessionPoolsClient.
func (c *ClientFactory) NewContainerAppsSessionPoolsClient() *ContainerAppsSessionPoolsClient {
	return &ContainerAppsSessionPoolsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewContainerAppsSourceControlsClient creates a new instance of ContainerAppsSourceControlsClient.
func (c *ClientFactory) NewContainerAppsSourceControlsClient() *ContainerAppsSourceControlsClient {
	return &ContainerAppsSourceControlsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDaprComponentsClient creates a new instance of DaprComponentsClient.
func (c *ClientFactory) NewDaprComponentsClient() *DaprComponentsClient {
	return &DaprComponentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewJavaComponentsClient creates a new instance of JavaComponentsClient.
func (c *ClientFactory) NewJavaComponentsClient() *JavaComponentsClient {
	return &JavaComponentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewJobsClient creates a new instance of JobsClient.
func (c *ClientFactory) NewJobsClient() *JobsClient {
	return &JobsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewJobsExecutionsClient creates a new instance of JobsExecutionsClient.
func (c *ClientFactory) NewJobsExecutionsClient() *JobsExecutionsClient {
	return &JobsExecutionsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewManagedCertificatesClient creates a new instance of ManagedCertificatesClient.
func (c *ClientFactory) NewManagedCertificatesClient() *ManagedCertificatesClient {
	return &ManagedCertificatesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewManagedEnvironmentDiagnosticsClient creates a new instance of ManagedEnvironmentDiagnosticsClient.
func (c *ClientFactory) NewManagedEnvironmentDiagnosticsClient() *ManagedEnvironmentDiagnosticsClient {
	return &ManagedEnvironmentDiagnosticsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewManagedEnvironmentUsagesClient creates a new instance of ManagedEnvironmentUsagesClient.
func (c *ClientFactory) NewManagedEnvironmentUsagesClient() *ManagedEnvironmentUsagesClient {
	return &ManagedEnvironmentUsagesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewManagedEnvironmentsClient creates a new instance of ManagedEnvironmentsClient.
func (c *ClientFactory) NewManagedEnvironmentsClient() *ManagedEnvironmentsClient {
	return &ManagedEnvironmentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewManagedEnvironmentsDiagnosticsClient creates a new instance of ManagedEnvironmentsDiagnosticsClient.
func (c *ClientFactory) NewManagedEnvironmentsDiagnosticsClient() *ManagedEnvironmentsDiagnosticsClient {
	return &ManagedEnvironmentsDiagnosticsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewManagedEnvironmentsStoragesClient creates a new instance of ManagedEnvironmentsStoragesClient.
func (c *ClientFactory) NewManagedEnvironmentsStoragesClient() *ManagedEnvironmentsStoragesClient {
	return &ManagedEnvironmentsStoragesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewNamespacesClient creates a new instance of NamespacesClient.
func (c *ClientFactory) NewNamespacesClient() *NamespacesClient {
	return &NamespacesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewUsagesClient creates a new instance of UsagesClient.
func (c *ClientFactory) NewUsagesClient() *UsagesClient {
	return &UsagesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
