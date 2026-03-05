// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourceHttpClient struct {
	// The decoding format used to interpret incoming logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#decoding ObservabilityPipeline#decoding}
	Decoding *string `field:"required" json:"decoding" yaml:"decoding"`
	// Optional authentication strategy for HTTP requests. Valid values are `none`, `basic`, `bearer`, `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#auth_strategy ObservabilityPipeline#auth_strategy}
	AuthStrategy *string `field:"optional" json:"authStrategy" yaml:"authStrategy"`
	// Name of the environment variable or secret that holds a custom header value (used with custom auth strategies).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#custom_key ObservabilityPipeline#custom_key}
	CustomKey *string `field:"optional" json:"customKey" yaml:"customKey"`
	// Name of the environment variable or secret that holds the HTTP endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#endpoint_url_key ObservabilityPipeline#endpoint_url_key}
	EndpointUrlKey *string `field:"optional" json:"endpointUrlKey" yaml:"endpointUrlKey"`
	// Name of the environment variable or secret that holds the password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#password_key ObservabilityPipeline#password_key}
	PasswordKey *string `field:"optional" json:"passwordKey" yaml:"passwordKey"`
	// The interval (in seconds) between HTTP scrape requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#scrape_interval_secs ObservabilityPipeline#scrape_interval_secs}
	ScrapeIntervalSecs *float64 `field:"optional" json:"scrapeIntervalSecs" yaml:"scrapeIntervalSecs"`
	// The timeout (in seconds) for each scrape request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#scrape_timeout_secs ObservabilityPipeline#scrape_timeout_secs}
	ScrapeTimeoutSecs *float64 `field:"optional" json:"scrapeTimeoutSecs" yaml:"scrapeTimeoutSecs"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// Name of the environment variable or secret that holds the authentication token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#token_key ObservabilityPipeline#token_key}
	TokenKey *string `field:"optional" json:"tokenKey" yaml:"tokenKey"`
	// Name of the environment variable or secret that holds the username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#username_key ObservabilityPipeline#username_key}
	UsernameKey *string `field:"optional" json:"usernameKey" yaml:"usernameKey"`
}

