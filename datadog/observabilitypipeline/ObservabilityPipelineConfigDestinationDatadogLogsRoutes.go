// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationDatadogLogsRoutes struct {
	// Name of the environment variable or secret that stores the Datadog API key used by this route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#api_key_key ObservabilityPipeline#api_key_key}
	ApiKeyKey *string `field:"required" json:"apiKeyKey" yaml:"apiKeyKey"`
	// A Datadog search query that determines which logs are forwarded using this route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#include ObservabilityPipeline#include}
	Include *string `field:"required" json:"include" yaml:"include"`
	// Unique identifier for this route within the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#route_id ObservabilityPipeline#route_id}
	RouteId *string `field:"required" json:"routeId" yaml:"routeId"`
	// Datadog site where matching logs are sent (for example, `us1`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#site ObservabilityPipeline#site}
	Site *string `field:"required" json:"site" yaml:"site"`
}

