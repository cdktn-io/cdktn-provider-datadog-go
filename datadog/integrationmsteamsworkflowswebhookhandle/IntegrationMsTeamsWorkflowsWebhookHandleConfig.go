// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationmsteamsworkflowswebhookhandle

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IntegrationMsTeamsWorkflowsWebhookHandleConfig struct {
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
	// Your Microsoft Workflows webhook handle name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/integration_ms_teams_workflows_webhook_handle#name IntegrationMsTeamsWorkflowsWebhookHandle#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Your Microsoft Workflows webhook URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/integration_ms_teams_workflows_webhook_handle#url IntegrationMsTeamsWorkflowsWebhookHandle#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

