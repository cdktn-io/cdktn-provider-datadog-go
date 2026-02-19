// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deploymentgate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeploymentGateConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The target environment (example: dev).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/deployment_gate#env DeploymentGate#env}
	Env *string `field:"required" json:"env" yaml:"env"`
	// The service name (example: transaction-backend).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/deployment_gate#service DeploymentGate#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Enable Dry Run to test gate behavior without impacting deployments.
	//
	// The evaluation of a dry run gate always responds with a pass status, but the in-app result is the real status based on rules evaluation. This is particularly useful when performing an initial evaluation of the gate behavior without impacting the deployment pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/deployment_gate#dry_run DeploymentGate#dry_run}
	DryRun interface{} `field:"optional" json:"dryRun" yaml:"dryRun"`
	// Unique name for multiple gates on the same service/environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/deployment_gate#identifier DeploymentGate#identifier}
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/deployment_gate#rule DeploymentGate#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

