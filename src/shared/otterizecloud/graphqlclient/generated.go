// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package graphqlclient

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type ComponentType string

const (
	ComponentTypeIntentsOperator     ComponentType = "INTENTS_OPERATOR"
	ComponentTypeCredentialsOperator ComponentType = "CREDENTIALS_OPERATOR"
	ComponentTypeNetworkMapper       ComponentType = "NETWORK_MAPPER"
)

type HTTPConfigInput struct {
	Path   string     `json:"path"`
	Method HTTPMethod `json:"method"`
}

// GetPath returns HTTPConfigInput.Path, and is useful for accessing the field via an interface.
func (v *HTTPConfigInput) GetPath() string { return v.Path }

// GetMethod returns HTTPConfigInput.Method, and is useful for accessing the field via an interface.
func (v *HTTPConfigInput) GetMethod() HTTPMethod { return v.Method }

type HTTPMethod string

const (
	HTTPMethodGet     HTTPMethod = "GET"
	HTTPMethodPost    HTTPMethod = "POST"
	HTTPMethodPut     HTTPMethod = "PUT"
	HTTPMethodDelete  HTTPMethod = "DELETE"
	HTTPMethodOptions HTTPMethod = "OPTIONS"
	HTTPMethodTrace   HTTPMethod = "TRACE"
	HTTPMethodPatch   HTTPMethod = "PATCH"
	HTTPMethodConnect HTTPMethod = "CONNECT"
)

type IntentInput struct {
	Namespace       string             `json:"namespace"`
	ClientName      string             `json:"clientName"`
	ServerName      string             `json:"serverName"`
	ServerNamespace string             `json:"serverNamespace"`
	Type            IntentType         `json:"type"`
	Topics          []KafkaConfigInput `json:"topics"`
	Resources       []HTTPConfigInput  `json:"resources"`
}

// GetNamespace returns IntentInput.Namespace, and is useful for accessing the field via an interface.
func (v *IntentInput) GetNamespace() string { return v.Namespace }

// GetClientName returns IntentInput.ClientName, and is useful for accessing the field via an interface.
func (v *IntentInput) GetClientName() string { return v.ClientName }

// GetServerName returns IntentInput.ServerName, and is useful for accessing the field via an interface.
func (v *IntentInput) GetServerName() string { return v.ServerName }

// GetServerNamespace returns IntentInput.ServerNamespace, and is useful for accessing the field via an interface.
func (v *IntentInput) GetServerNamespace() string { return v.ServerNamespace }

// GetType returns IntentInput.Type, and is useful for accessing the field via an interface.
func (v *IntentInput) GetType() IntentType { return v.Type }

// GetTopics returns IntentInput.Topics, and is useful for accessing the field via an interface.
func (v *IntentInput) GetTopics() []KafkaConfigInput { return v.Topics }

// GetResources returns IntentInput.Resources, and is useful for accessing the field via an interface.
func (v *IntentInput) GetResources() []HTTPConfigInput { return v.Resources }

type IntentType string

const (
	IntentTypeHttp  IntentType = "HTTP"
	IntentTypeKafka IntentType = "KAFKA"
)

type IntentsOperatorConfigurationInput struct {
	EnableEnforcement bool `json:"enableEnforcement"`
}

// GetEnableEnforcement returns IntentsOperatorConfigurationInput.EnableEnforcement, and is useful for accessing the field via an interface.
func (v *IntentsOperatorConfigurationInput) GetEnableEnforcement() bool { return v.EnableEnforcement }

type KafkaConfigInput struct {
	Name       string           `json:"name"`
	Operations []KafkaOperation `json:"operations"`
}

// GetName returns KafkaConfigInput.Name, and is useful for accessing the field via an interface.
func (v *KafkaConfigInput) GetName() string { return v.Name }

// GetOperations returns KafkaConfigInput.Operations, and is useful for accessing the field via an interface.
func (v *KafkaConfigInput) GetOperations() []KafkaOperation { return v.Operations }

type KafkaOperation string

const (
	KafkaOperationConsume         KafkaOperation = "CONSUME"
	KafkaOperationProduce         KafkaOperation = "PRODUCE"
	KafkaOperationCreate          KafkaOperation = "CREATE"
	KafkaOperationAlter           KafkaOperation = "ALTER"
	KafkaOperationDelete          KafkaOperation = "DELETE"
	KafkaOperationDescribe        KafkaOperation = "DESCRIBE"
	KafkaOperationClusterAction   KafkaOperation = "CLUSTER_ACTION"
	KafkaOperationDescribeConfigs KafkaOperation = "DESCRIBE_CONFIGS"
	KafkaOperationAlterConfigs    KafkaOperation = "ALTER_CONFIGS"
	KafkaOperationIdempotentWrite KafkaOperation = "IDEMPOTENT_WRITE"
)

type KafkaServerConfigInput struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Address   string            `json:"address"`
	Topics    []KafkaTopicInput `json:"topics"`
}

// GetName returns KafkaServerConfigInput.Name, and is useful for accessing the field via an interface.
func (v *KafkaServerConfigInput) GetName() string { return v.Name }

// GetNamespace returns KafkaServerConfigInput.Namespace, and is useful for accessing the field via an interface.
func (v *KafkaServerConfigInput) GetNamespace() string { return v.Namespace }

// GetAddress returns KafkaServerConfigInput.Address, and is useful for accessing the field via an interface.
func (v *KafkaServerConfigInput) GetAddress() string { return v.Address }

// GetTopics returns KafkaServerConfigInput.Topics, and is useful for accessing the field via an interface.
func (v *KafkaServerConfigInput) GetTopics() []KafkaTopicInput { return v.Topics }

type KafkaTopicInput struct {
	ClientIdentityRequired bool   `json:"clientIdentityRequired"`
	IntentsRequired        bool   `json:"intentsRequired"`
	Pattern                string `json:"pattern"`
	Topic                  string `json:"topic"`
}

// GetClientIdentityRequired returns KafkaTopicInput.ClientIdentityRequired, and is useful for accessing the field via an interface.
func (v *KafkaTopicInput) GetClientIdentityRequired() bool { return v.ClientIdentityRequired }

// GetIntentsRequired returns KafkaTopicInput.IntentsRequired, and is useful for accessing the field via an interface.
func (v *KafkaTopicInput) GetIntentsRequired() bool { return v.IntentsRequired }

// GetPattern returns KafkaTopicInput.Pattern, and is useful for accessing the field via an interface.
func (v *KafkaTopicInput) GetPattern() string { return v.Pattern }

// GetTopic returns KafkaTopicInput.Topic, and is useful for accessing the field via an interface.
func (v *KafkaTopicInput) GetTopic() string { return v.Topic }

// ReportAppliedKubernetesIntentsResponse is returned by ReportAppliedKubernetesIntents on success.
type ReportAppliedKubernetesIntentsResponse struct {
	ReportAppliedKubernetesIntents bool `json:"reportAppliedKubernetesIntents"`
}

// GetReportAppliedKubernetesIntents returns ReportAppliedKubernetesIntentsResponse.ReportAppliedKubernetesIntents, and is useful for accessing the field via an interface.
func (v *ReportAppliedKubernetesIntentsResponse) GetReportAppliedKubernetesIntents() bool {
	return v.ReportAppliedKubernetesIntents
}

// ReportComponentStatusResponse is returned by ReportComponentStatus on success.
type ReportComponentStatusResponse struct {
	ReportIntegrationComponentStatus bool `json:"reportIntegrationComponentStatus"`
}

// GetReportIntegrationComponentStatus returns ReportComponentStatusResponse.ReportIntegrationComponentStatus, and is useful for accessing the field via an interface.
func (v *ReportComponentStatusResponse) GetReportIntegrationComponentStatus() bool {
	return v.ReportIntegrationComponentStatus
}

// ReportIntentsOperatorConfigurationResponse is returned by ReportIntentsOperatorConfiguration on success.
type ReportIntentsOperatorConfigurationResponse struct {
	ReportIntentsOperatorConfiguration string `json:"reportIntentsOperatorConfiguration"`
}

// GetReportIntentsOperatorConfiguration returns ReportIntentsOperatorConfigurationResponse.ReportIntentsOperatorConfiguration, and is useful for accessing the field via an interface.
func (v *ReportIntentsOperatorConfigurationResponse) GetReportIntentsOperatorConfiguration() string {
	return v.ReportIntentsOperatorConfiguration
}

// ReportKafkaServerConfigResponse is returned by ReportKafkaServerConfig on success.
type ReportKafkaServerConfigResponse struct {
	ReportKafkaServerConfig bool `json:"reportKafkaServerConfig"`
}

// GetReportKafkaServerConfig returns ReportKafkaServerConfigResponse.ReportKafkaServerConfig, and is useful for accessing the field via an interface.
func (v *ReportKafkaServerConfigResponse) GetReportKafkaServerConfig() bool {
	return v.ReportKafkaServerConfig
}

// __ReportAppliedKubernetesIntentsInput is used internally by genqlient
type __ReportAppliedKubernetesIntentsInput struct {
	Namespace string        `json:"namespace"`
	Intents   []IntentInput `json:"intents"`
}

// GetNamespace returns __ReportAppliedKubernetesIntentsInput.Namespace, and is useful for accessing the field via an interface.
func (v *__ReportAppliedKubernetesIntentsInput) GetNamespace() string { return v.Namespace }

// GetIntents returns __ReportAppliedKubernetesIntentsInput.Intents, and is useful for accessing the field via an interface.
func (v *__ReportAppliedKubernetesIntentsInput) GetIntents() []IntentInput { return v.Intents }

// __ReportComponentStatusInput is used internally by genqlient
type __ReportComponentStatusInput struct {
	Component ComponentType `json:"component"`
}

// GetComponent returns __ReportComponentStatusInput.Component, and is useful for accessing the field via an interface.
func (v *__ReportComponentStatusInput) GetComponent() ComponentType { return v.Component }

// __ReportIntentsOperatorConfigurationInput is used internally by genqlient
type __ReportIntentsOperatorConfigurationInput struct {
	Configuration IntentsOperatorConfigurationInput `json:"configuration"`
}

// GetConfiguration returns __ReportIntentsOperatorConfigurationInput.Configuration, and is useful for accessing the field via an interface.
func (v *__ReportIntentsOperatorConfigurationInput) GetConfiguration() IntentsOperatorConfigurationInput {
	return v.Configuration
}

// __ReportKafkaServerConfigInput is used internally by genqlient
type __ReportKafkaServerConfigInput struct {
	Server KafkaServerConfigInput `json:"server"`
}

// GetServer returns __ReportKafkaServerConfigInput.Server, and is useful for accessing the field via an interface.
func (v *__ReportKafkaServerConfigInput) GetServer() KafkaServerConfigInput { return v.Server }

func ReportAppliedKubernetesIntents(
	ctx context.Context,
	client graphql.Client,
	namespace string,
	intents []IntentInput,
) (*ReportAppliedKubernetesIntentsResponse, error) {
	req := &graphql.Request{
		OpName: "ReportAppliedKubernetesIntents",
		Query: `
mutation ReportAppliedKubernetesIntents ($namespace: String!, $intents: [IntentInput!]!) {
	reportAppliedKubernetesIntents(namespace: $namespace, intents: $intents)
}
`,
		Variables: &__ReportAppliedKubernetesIntentsInput{
			Namespace: namespace,
			Intents:   intents,
		},
	}
	var err error

	var data ReportAppliedKubernetesIntentsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func ReportComponentStatus(
	ctx context.Context,
	client graphql.Client,
	component ComponentType,
) (*ReportComponentStatusResponse, error) {
	req := &graphql.Request{
		OpName: "ReportComponentStatus",
		Query: `
mutation ReportComponentStatus ($component: ComponentType!) {
	reportIntegrationComponentStatus(component: $component)
}
`,
		Variables: &__ReportComponentStatusInput{
			Component: component,
		},
	}
	var err error

	var data ReportComponentStatusResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func ReportIntentsOperatorConfiguration(
	ctx context.Context,
	client graphql.Client,
	configuration IntentsOperatorConfigurationInput,
) (*ReportIntentsOperatorConfigurationResponse, error) {
	req := &graphql.Request{
		OpName: "ReportIntentsOperatorConfiguration",
		Query: `
mutation ReportIntentsOperatorConfiguration ($configuration: IntentsOperatorConfigurationInput!) {
	reportIntentsOperatorConfiguration(configuration: $configuration)
}
`,
		Variables: &__ReportIntentsOperatorConfigurationInput{
			Configuration: configuration,
		},
	}
	var err error

	var data ReportIntentsOperatorConfigurationResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func ReportKafkaServerConfig(
	ctx context.Context,
	client graphql.Client,
	server KafkaServerConfigInput,
) (*ReportKafkaServerConfigResponse, error) {
	req := &graphql.Request{
		OpName: "ReportKafkaServerConfig",
		Query: `
mutation ReportKafkaServerConfig ($server: KafkaServerConfigInput!) {
	reportKafkaServerConfig(serverConfig: $server)
}
`,
		Variables: &__ReportKafkaServerConfigInput{
			Server: server,
		},
	}
	var err error

	var data ReportKafkaServerConfigResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
