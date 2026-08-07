// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceHttpServerValidTokenPathToToken struct {
	// The name of the HTTP header that carries the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/observability_pipeline#header ObservabilityPipeline#header}
	Header *string `field:"optional" json:"header" yaml:"header"`
	// Built-in token location on the incoming HTTP request. One of `path`, `address`. Valid values are `path`, `address`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/observability_pipeline#location ObservabilityPipeline#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
}

