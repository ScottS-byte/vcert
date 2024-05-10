// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package cloudproviders

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type CertificateProvisioningAWSOptionsInput struct {
	// Amazon Resource Name (ARN) uniquely identifying AWS certificate resource
	Arn  *string                                  `json:"arn"`
	Tags []*CertificateProvisioningTagOptionInput `json:"tags,omitempty"`
}

// GetArn returns CertificateProvisioningAWSOptionsInput.Arn, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningAWSOptionsInput) GetArn() *string { return v.Arn }

// GetTags returns CertificateProvisioningAWSOptionsInput.Tags, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningAWSOptionsInput) GetTags() []*CertificateProvisioningTagOptionInput {
	return v.Tags
}

type CertificateProvisioningAzureOptionsInput struct {
	// An Azure object-name is a user provided name for and must be unique within a key vault. The name must be a 1-127 character string, starting with a letter and containing only 0-9, a-z, A-Z, and -.
	Name       *string                                  `json:"name"`
	Enabled    *bool                                    `json:"enabled"`
	Exportable *bool                                    `json:"exportable"`
	ReuseKey   *bool                                    `json:"reuseKey"`
	Tags       []*CertificateProvisioningTagOptionInput `json:"tags,omitempty"`
}

// GetName returns CertificateProvisioningAzureOptionsInput.Name, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningAzureOptionsInput) GetName() *string { return v.Name }

// GetEnabled returns CertificateProvisioningAzureOptionsInput.Enabled, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningAzureOptionsInput) GetEnabled() *bool { return v.Enabled }

// GetExportable returns CertificateProvisioningAzureOptionsInput.Exportable, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningAzureOptionsInput) GetExportable() *bool { return v.Exportable }

// GetReuseKey returns CertificateProvisioningAzureOptionsInput.ReuseKey, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningAzureOptionsInput) GetReuseKey() *bool { return v.ReuseKey }

// GetTags returns CertificateProvisioningAzureOptionsInput.Tags, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningAzureOptionsInput) GetTags() []*CertificateProvisioningTagOptionInput {
	return v.Tags
}

type CertificateProvisioningGCPOptionsInput struct {
	// A GCM certificate id is a user provider name and must be unique within a region. The id must be a 1-63 character string, start with a letter and only contain lower case letters, digits and hyphens.
	Id          *string                                  `json:"id"`
	Description *string                                  `json:"description"`
	Scope       *GCMCertificateScope                     `json:"scope"`
	Labels      []*CertificateProvisioningTagOptionInput `json:"labels,omitempty"`
}

// GetId returns CertificateProvisioningGCPOptionsInput.Id, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningGCPOptionsInput) GetId() *string { return v.Id }

// GetDescription returns CertificateProvisioningGCPOptionsInput.Description, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningGCPOptionsInput) GetDescription() *string { return v.Description }

// GetScope returns CertificateProvisioningGCPOptionsInput.Scope, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningGCPOptionsInput) GetScope() *GCMCertificateScope { return v.Scope }

// GetLabels returns CertificateProvisioningGCPOptionsInput.Labels, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningGCPOptionsInput) GetLabels() []*CertificateProvisioningTagOptionInput {
	return v.Labels
}

type CertificateProvisioningOptionsInput struct {
	AwsOptions   *CertificateProvisioningAWSOptionsInput   `json:"awsOptions,omitempty"`
	AzureOptions *CertificateProvisioningAzureOptionsInput `json:"azureOptions,omitempty"`
	GcpOptions   *CertificateProvisioningGCPOptionsInput   `json:"gcpOptions,omitempty"`
}

// GetAwsOptions returns CertificateProvisioningOptionsInput.AwsOptions, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningOptionsInput) GetAwsOptions() *CertificateProvisioningAWSOptionsInput {
	return v.AwsOptions
}

// GetAzureOptions returns CertificateProvisioningOptionsInput.AzureOptions, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningOptionsInput) GetAzureOptions() *CertificateProvisioningAzureOptionsInput {
	return v.AzureOptions
}

// GetGcpOptions returns CertificateProvisioningOptionsInput.GcpOptions, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningOptionsInput) GetGcpOptions() *CertificateProvisioningGCPOptionsInput {
	return v.GcpOptions
}

type CertificateProvisioningTagOptionInput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// GetName returns CertificateProvisioningTagOptionInput.Name, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningTagOptionInput) GetName() string { return v.Name }

// GetValue returns CertificateProvisioningTagOptionInput.Value, and is useful for accessing the field via an interface.
func (v *CertificateProvisioningTagOptionInput) GetValue() string { return v.Value }

// Indicates the type of a Cloud Keystore
type CloudKeystoreType string

const (
	// AWS Certificate Manager
	CloudKeystoreTypeAcm CloudKeystoreType = "ACM"
	// Azure Key Vault
	CloudKeystoreTypeAkv CloudKeystoreType = "AKV"
	// Google Certificate Manager
	CloudKeystoreTypeGcm CloudKeystoreType = "GCM"
)

// Indicates the status of a cloud provider
type CloudProviderStatus string

const (
	// The cloud provider is successfully validated
	CloudProviderStatusValidated CloudProviderStatus = "VALIDATED"
	// The cloud provider is not validated or the validation fails for some reason
	CloudProviderStatusNotValidated CloudProviderStatus = "NOT_VALIDATED"
)

// Indicates the type of a Cloud Provider
type CloudProviderType string

const (
	// AWS cloud provider type
	CloudProviderTypeAws CloudProviderType = "AWS"
	// Azure cloud provider type
	CloudProviderTypeAzure CloudProviderType = "AZURE"
	// Google cloud provider type
	CloudProviderTypeGcp CloudProviderType = "GCP"
)

// Indicates the Scope for a certificate provisioned to GCP Certificate Manager
type GCMCertificateScope string

const (
	// Certificates with default scope are served from core Google data centers. If unsure, choose this option.
	GCMCertificateScopeDefault GCMCertificateScope = "DEFAULT"
	// Certificates with scope EDGE_CACHE are special-purposed certificates, served from Edge Points of Presence.
	// See https://cloud.google.com/vpc/docs/edge-locations.
	GCMCertificateScopeEdgeCache GCMCertificateScope = "EDGE_CACHE"
)

// GetCloudKeystoresCloudKeystoresCloudKeystoreConnection includes the requested fields of the GraphQL type CloudKeystoreConnection.
// The GraphQL type's documentation follows.
//
// A page of CloudKeystore results
type GetCloudKeystoresCloudKeystoresCloudKeystoreConnection struct {
	// CloudKeystores in the current page, without cursor
	Nodes []*GetCloudKeystoresCloudKeystoresCloudKeystoreConnectionNodesCloudKeystore `json:"nodes"`
}

// GetNodes returns GetCloudKeystoresCloudKeystoresCloudKeystoreConnection.Nodes, and is useful for accessing the field via an interface.
func (v *GetCloudKeystoresCloudKeystoresCloudKeystoreConnection) GetNodes() []*GetCloudKeystoresCloudKeystoresCloudKeystoreConnectionNodesCloudKeystore {
	return v.Nodes
}

// GetCloudKeystoresCloudKeystoresCloudKeystoreConnectionNodesCloudKeystore includes the requested fields of the GraphQL type CloudKeystore.
type GetCloudKeystoresCloudKeystoresCloudKeystoreConnectionNodesCloudKeystore struct {
	Id string `json:"id"`
	// Cloud Keystore name
	//
	// A string between 3 and 250 characters
	Name string            `json:"name"`
	Type CloudKeystoreType `json:"type"`
}

// GetId returns GetCloudKeystoresCloudKeystoresCloudKeystoreConnectionNodesCloudKeystore.Id, and is useful for accessing the field via an interface.
func (v *GetCloudKeystoresCloudKeystoresCloudKeystoreConnectionNodesCloudKeystore) GetId() string {
	return v.Id
}

// GetName returns GetCloudKeystoresCloudKeystoresCloudKeystoreConnectionNodesCloudKeystore.Name, and is useful for accessing the field via an interface.
func (v *GetCloudKeystoresCloudKeystoresCloudKeystoreConnectionNodesCloudKeystore) GetName() string {
	return v.Name
}

// GetType returns GetCloudKeystoresCloudKeystoresCloudKeystoreConnectionNodesCloudKeystore.Type, and is useful for accessing the field via an interface.
func (v *GetCloudKeystoresCloudKeystoresCloudKeystoreConnectionNodesCloudKeystore) GetType() CloudKeystoreType {
	return v.Type
}

// GetCloudKeystoresResponse is returned by GetCloudKeystores on success.
type GetCloudKeystoresResponse struct {
	// Retrieves Cloud Keystores.
	// The pagination can be either forward or backward. To enable forward pagination, two arguments
	// are used: `after` and `first`. To enable backward pagination, two arguments are used: `before` and `last`.
	// If arguments for both forward and backward pagination are supplied, forward pagination wil be used. If no arguments
	// are supplied, it returns the first page of 10 cloud keystores (i.e. defaults `first` to 10). The result is sorted by
	// the added on date in ascending order.
	// - after: returns the elements in the list that come after the specified cursor. Defaults to empty string, meaning
	// that we return the first page of cloud providers, if `first` value is supplied
	// - first: non-negative integer, denoting the first `n` number of records to return after the `after` cursor value.
	// Max value is 100
	// - before: returns the elements in the list that come before the specified cursor. By default is the empty string,
	// meaning that the results will be the last page, if `last` value is supplied
	// - last: non-negative integer, denoting the last `n` number of records to return before the `before` cursor value.
	// Max value is 100
	CloudKeystores *GetCloudKeystoresCloudKeystoresCloudKeystoreConnection `json:"cloudKeystores"`
}

// GetCloudKeystores returns GetCloudKeystoresResponse.CloudKeystores, and is useful for accessing the field via an interface.
func (v *GetCloudKeystoresResponse) GetCloudKeystores() *GetCloudKeystoresCloudKeystoresCloudKeystoreConnection {
	return v.CloudKeystores
}

// GetCloudProviderByNameCloudProvidersCloudProviderConnection includes the requested fields of the GraphQL type CloudProviderConnection.
// The GraphQL type's documentation follows.
//
// A page of CloudProvider results
type GetCloudProviderByNameCloudProvidersCloudProviderConnection struct {
	// CloudProviders in the current page, without cursor
	Nodes []*GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider `json:"nodes"`
}

// GetNodes returns GetCloudProviderByNameCloudProvidersCloudProviderConnection.Nodes, and is useful for accessing the field via an interface.
func (v *GetCloudProviderByNameCloudProvidersCloudProviderConnection) GetNodes() []*GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider {
	return v.Nodes
}

// GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider includes the requested fields of the GraphQL type CloudProvider.
type GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider struct {
	Id             string              `json:"id"`
	Name           string              `json:"name"`
	Type           CloudProviderType   `json:"type"`
	Status         CloudProviderStatus `json:"status"`
	StatusDetails  *string             `json:"statusDetails"`
	KeystoresCount int                 `json:"keystoresCount"`
}

// GetId returns GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider.Id, and is useful for accessing the field via an interface.
func (v *GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider) GetId() string {
	return v.Id
}

// GetName returns GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider.Name, and is useful for accessing the field via an interface.
func (v *GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider) GetName() string {
	return v.Name
}

// GetType returns GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider.Type, and is useful for accessing the field via an interface.
func (v *GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider) GetType() CloudProviderType {
	return v.Type
}

// GetStatus returns GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider.Status, and is useful for accessing the field via an interface.
func (v *GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider) GetStatus() CloudProviderStatus {
	return v.Status
}

// GetStatusDetails returns GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider.StatusDetails, and is useful for accessing the field via an interface.
func (v *GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider) GetStatusDetails() *string {
	return v.StatusDetails
}

// GetKeystoresCount returns GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider.KeystoresCount, and is useful for accessing the field via an interface.
func (v *GetCloudProviderByNameCloudProvidersCloudProviderConnectionNodesCloudProvider) GetKeystoresCount() int {
	return v.KeystoresCount
}

// GetCloudProviderByNameResponse is returned by GetCloudProviderByName on success.
type GetCloudProviderByNameResponse struct {
	// Retrieves Cloud Providers.
	// The pagination can be either forward or backward. To enable forward pagination, two arguments
	// are used: `after` and `first`. To enable backward pagination, two arguments are used: `before` and `last`.
	// If arguments for both forward and backward pagination are supplied, forward pagination wil be used. If no arguments
	// are supplied, it returns the first page of 10 cloud providers (i.e. defaults `first` to 10). The result is sorted by
	// the added on date in ascending order.
	// - after: returns the elements in the list that come after the specified cursor. Defaults to empty string, meaning
	// that we return the first page of cloud providers, if `first` value is supplied
	// - first: non-negative integer, denoting the first `n` number of records to return after the `after` cursor value.
	// Max value is 100
	// - before: returns the elements in the list that come before the specified cursor. By default is the empty string,
	// meaning that the results will be the last page, if `last` value is supplied
	// - last: non-negative integer, denoting the last `n` number of records to return before the `before` cursor value.
	// Max value is 100
	CloudProviders *GetCloudProviderByNameCloudProvidersCloudProviderConnection `json:"cloudProviders"`
}

// GetCloudProviders returns GetCloudProviderByNameResponse.CloudProviders, and is useful for accessing the field via an interface.
func (v *GetCloudProviderByNameResponse) GetCloudProviders() *GetCloudProviderByNameCloudProvidersCloudProviderConnection {
	return v.CloudProviders
}

// ProvisionCertificateProvisionToCloudKeystoreWorkflowResult includes the requested fields of the GraphQL type WorkflowResult.
type ProvisionCertificateProvisionToCloudKeystoreWorkflowResult struct {
	WorkflowId   string `json:"workflowId"`
	WorkflowName string `json:"workflowName"`
}

// GetWorkflowId returns ProvisionCertificateProvisionToCloudKeystoreWorkflowResult.WorkflowId, and is useful for accessing the field via an interface.
func (v *ProvisionCertificateProvisionToCloudKeystoreWorkflowResult) GetWorkflowId() string {
	return v.WorkflowId
}

// GetWorkflowName returns ProvisionCertificateProvisionToCloudKeystoreWorkflowResult.WorkflowName, and is useful for accessing the field via an interface.
func (v *ProvisionCertificateProvisionToCloudKeystoreWorkflowResult) GetWorkflowName() string {
	return v.WorkflowName
}

// ProvisionCertificateResponse is returned by ProvisionCertificate on success.
type ProvisionCertificateResponse struct {
	// Provision a certificate to a Cloud Keystore
	ProvisionToCloudKeystore *ProvisionCertificateProvisionToCloudKeystoreWorkflowResult `json:"provisionToCloudKeystore"`
}

// GetProvisionToCloudKeystore returns ProvisionCertificateResponse.ProvisionToCloudKeystore, and is useful for accessing the field via an interface.
func (v *ProvisionCertificateResponse) GetProvisionToCloudKeystore() *ProvisionCertificateProvisionToCloudKeystoreWorkflowResult {
	return v.ProvisionToCloudKeystore
}

// __GetCloudKeystoresInput is used internally by genqlient
type __GetCloudKeystoresInput struct {
	CloudKeystoreId   *string `json:"cloudKeystoreId"`
	CloudKeystoreName *string `json:"cloudKeystoreName"`
	CloudProviderId   *string `json:"cloudProviderId"`
	CloudProviderName *string `json:"cloudProviderName"`
}

// GetCloudKeystoreId returns __GetCloudKeystoresInput.CloudKeystoreId, and is useful for accessing the field via an interface.
func (v *__GetCloudKeystoresInput) GetCloudKeystoreId() *string { return v.CloudKeystoreId }

// GetCloudKeystoreName returns __GetCloudKeystoresInput.CloudKeystoreName, and is useful for accessing the field via an interface.
func (v *__GetCloudKeystoresInput) GetCloudKeystoreName() *string { return v.CloudKeystoreName }

// GetCloudProviderId returns __GetCloudKeystoresInput.CloudProviderId, and is useful for accessing the field via an interface.
func (v *__GetCloudKeystoresInput) GetCloudProviderId() *string { return v.CloudProviderId }

// GetCloudProviderName returns __GetCloudKeystoresInput.CloudProviderName, and is useful for accessing the field via an interface.
func (v *__GetCloudKeystoresInput) GetCloudProviderName() *string { return v.CloudProviderName }

// __GetCloudProviderByNameInput is used internally by genqlient
type __GetCloudProviderByNameInput struct {
	Name string `json:"name"`
}

// GetName returns __GetCloudProviderByNameInput.Name, and is useful for accessing the field via an interface.
func (v *__GetCloudProviderByNameInput) GetName() string { return v.Name }

// __ProvisionCertificateInput is used internally by genqlient
type __ProvisionCertificateInput struct {
	CertificateId   string                               `json:"certificateId"`
	CloudKeystoreId string                               `json:"cloudKeystoreId"`
	WsClientId      string                               `json:"wsClientId"`
	Options         *CertificateProvisioningOptionsInput `json:"options,omitempty"`
}

// GetCertificateId returns __ProvisionCertificateInput.CertificateId, and is useful for accessing the field via an interface.
func (v *__ProvisionCertificateInput) GetCertificateId() string { return v.CertificateId }

// GetCloudKeystoreId returns __ProvisionCertificateInput.CloudKeystoreId, and is useful for accessing the field via an interface.
func (v *__ProvisionCertificateInput) GetCloudKeystoreId() string { return v.CloudKeystoreId }

// GetWsClientId returns __ProvisionCertificateInput.WsClientId, and is useful for accessing the field via an interface.
func (v *__ProvisionCertificateInput) GetWsClientId() string { return v.WsClientId }

// GetOptions returns __ProvisionCertificateInput.Options, and is useful for accessing the field via an interface.
func (v *__ProvisionCertificateInput) GetOptions() *CertificateProvisioningOptionsInput {
	return v.Options
}

// The query or mutation executed by GetCloudKeystores.
const GetCloudKeystores_Operation = `
query GetCloudKeystores ($cloudKeystoreId: UUID, $cloudKeystoreName: String, $cloudProviderId: UUID, $cloudProviderName: String) {
	cloudKeystores(filter: {cloudKeystoreId:$cloudKeystoreId,cloudKeystoreName:$cloudKeystoreName,cloudProviderId:$cloudProviderId,cloudProviderName:$cloudProviderName}) {
		nodes {
			id
			name
			type
		}
	}
}
`

func GetCloudKeystores(
	ctx_ context.Context,
	client_ graphql.Client,
	cloudKeystoreId *string,
	cloudKeystoreName *string,
	cloudProviderId *string,
	cloudProviderName *string,
) (*GetCloudKeystoresResponse, error) {
	req_ := &graphql.Request{
		OpName: "GetCloudKeystores",
		Query:  GetCloudKeystores_Operation,
		Variables: &__GetCloudKeystoresInput{
			CloudKeystoreId:   cloudKeystoreId,
			CloudKeystoreName: cloudKeystoreName,
			CloudProviderId:   cloudProviderId,
			CloudProviderName: cloudProviderName,
		},
	}
	var err_ error

	var data_ GetCloudKeystoresResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by GetCloudProviderByName.
const GetCloudProviderByName_Operation = `
query GetCloudProviderByName ($name: String!) {
	cloudProviders(filter: {name:$name}) {
		nodes {
			id
			name
			type
			status
			statusDetails
			keystoresCount
		}
	}
}
`

func GetCloudProviderByName(
	ctx_ context.Context,
	client_ graphql.Client,
	name string,
) (*GetCloudProviderByNameResponse, error) {
	req_ := &graphql.Request{
		OpName: "GetCloudProviderByName",
		Query:  GetCloudProviderByName_Operation,
		Variables: &__GetCloudProviderByNameInput{
			Name: name,
		},
	}
	var err_ error

	var data_ GetCloudProviderByNameResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by ProvisionCertificate.
const ProvisionCertificate_Operation = `
mutation ProvisionCertificate ($certificateId: UUID!, $cloudKeystoreId: UUID!, $wsClientId: UUID!, $options: CertificateProvisioningOptionsInput) {
	provisionToCloudKeystore(certificateId: $certificateId, cloudKeystoreId: $cloudKeystoreId, wsClientId: $wsClientId, options: $options) {
		workflowId
		workflowName
	}
}
`

func ProvisionCertificate(
	ctx_ context.Context,
	client_ graphql.Client,
	certificateId string,
	cloudKeystoreId string,
	wsClientId string,
	options *CertificateProvisioningOptionsInput,
) (*ProvisionCertificateResponse, error) {
	req_ := &graphql.Request{
		OpName: "ProvisionCertificate",
		Query:  ProvisionCertificate_Operation,
		Variables: &__ProvisionCertificateInput{
			CertificateId:   certificateId,
			CloudKeystoreId: cloudKeystoreId,
			WsClientId:      wsClientId,
			Options:         options,
		},
	}
	var err_ error

	var data_ ProvisionCertificateResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}
