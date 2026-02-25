// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationconfluentresource

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IntegrationConfluentResourceConfig struct {
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
	// Confluent Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/integration_confluent_resource#account_id IntegrationConfluentResource#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The ID associated with a Confluent resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/integration_confluent_resource#resource_id IntegrationConfluentResource#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Enable the `custom.consumer_lag_offset` metric, which contains extra metric tags. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/integration_confluent_resource#enable_custom_metrics IntegrationConfluentResource#enable_custom_metrics}
	EnableCustomMetrics interface{} `field:"optional" json:"enableCustomMetrics" yaml:"enableCustomMetrics"`
	// The resource type of the Resource. Can be `kafka`, `connector`, `ksql`, or `schema_registry`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/integration_confluent_resource#resource_type IntegrationConfluentResource#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// A list of strings representing tags. Can be a single key, or key-value pairs separated by a colon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/integration_confluent_resource#tags IntegrationConfluentResource#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

