// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceHttpClient struct {
	// The decoding format used to interpret incoming logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#decoding ObservabilityPipeline#decoding}
	Decoding *string `field:"required" json:"decoding" yaml:"decoding"`
	// Optional authentication strategy for HTTP requests. Valid values are `none`, `basic`, `bearer`, `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#auth_strategy ObservabilityPipeline#auth_strategy}
	AuthStrategy *string `field:"optional" json:"authStrategy" yaml:"authStrategy"`
	// The interval (in seconds) between HTTP scrape requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#scrape_interval_secs ObservabilityPipeline#scrape_interval_secs}
	ScrapeIntervalSecs *float64 `field:"optional" json:"scrapeIntervalSecs" yaml:"scrapeIntervalSecs"`
	// The timeout (in seconds) for each scrape request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#scrape_timeout_secs ObservabilityPipeline#scrape_timeout_secs}
	ScrapeTimeoutSecs *float64 `field:"optional" json:"scrapeTimeoutSecs" yaml:"scrapeTimeoutSecs"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

