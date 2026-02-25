// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rumapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RumApplicationConfig struct {
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
	// Name of the RUM application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/rum_application#name RumApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Controls the retention policy for Product Analytics data derived from RUM events. Valid values are `MAX`, `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/rum_application#product_analytics_retention_state RumApplication#product_analytics_retention_state}
	ProductAnalyticsRetentionState *string `field:"optional" json:"productAnalyticsRetentionState" yaml:"productAnalyticsRetentionState"`
	// Configures which RUM events are processed and stored for the application. Valid values are `ALL`, `ERROR_FOCUSED_MODE`, `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/rum_application#rum_event_processing_state RumApplication#rum_event_processing_state}
	RumEventProcessingState *string `field:"optional" json:"rumEventProcessingState" yaml:"rumEventProcessingState"`
	// Type of the RUM application. Supported values are `browser`, `ios`, `android`, `react-native`, `flutter`. Defaults to `"browser"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/rum_application#type RumApplication#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

