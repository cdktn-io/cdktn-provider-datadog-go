// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorAssets struct {
	// Type of asset the entity represents on a monitor. Valid values are `runbook`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/monitor#category Monitor#category}
	Category *string `field:"required" json:"category" yaml:"category"`
	// Name for the monitor asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/monitor#name Monitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// URL for the asset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/monitor#url Monitor#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Identifier of the internal Datadog resource that this asset represents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/monitor#resource_key Monitor#resource_key}
	ResourceKey *string `field:"optional" json:"resourceKey" yaml:"resourceKey"`
	// Type of internal Datadog resource associated with a monitor asset. Valid values are `notebook`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/monitor#resource_type Monitor#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

