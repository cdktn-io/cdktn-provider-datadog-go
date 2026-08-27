// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount


type IntegrationAwsAccountMetricsConfigMetricNameFilters struct {
	// The AWS CloudWatch namespace to which this metric name filter applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/integration_aws_account#namespace IntegrationAwsAccount#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Exclude metric names matching one of these patterns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/integration_aws_account#exclude_only IntegrationAwsAccount#exclude_only}
	ExcludeOnly *[]*string `field:"optional" json:"excludeOnly" yaml:"excludeOnly"`
	// Include only metric names matching one of these patterns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/integration_aws_account#include_only IntegrationAwsAccount#include_only}
	IncludeOnly *[]*string `field:"optional" json:"includeOnly" yaml:"includeOnly"`
}

