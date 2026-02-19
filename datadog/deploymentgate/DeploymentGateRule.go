// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deploymentgate


type DeploymentGateRule struct {
	// The rule name. Must be unique within the deployment gate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/deployment_gate#name DeploymentGate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The rule type (e.g., 'faulty_deployment_detection', 'monitor').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/deployment_gate#type DeploymentGate#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Whether the rule is in dry run mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/deployment_gate#dry_run DeploymentGate#dry_run}
	DryRun interface{} `field:"optional" json:"dryRun" yaml:"dryRun"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/deployment_gate#options DeploymentGate#options}
	Options *DeploymentGateRuleOptions `field:"optional" json:"options" yaml:"options"`
}

