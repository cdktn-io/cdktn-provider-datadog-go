// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deploymentgate


type DeploymentGateRuleOptions struct {
	// The duration for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/deployment_gate#duration DeploymentGate#duration}
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
	// Resources to exclude from faulty deployment detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/deployment_gate#excluded_resources DeploymentGate#excluded_resources}
	ExcludedResources *[]*string `field:"optional" json:"excludedResources" yaml:"excludedResources"`
	// The query for monitor rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/deployment_gate#query DeploymentGate#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

