// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetTimeseriesDefinitionRequestMetadata struct {
	// The expression name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#expression DashboardV2#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The expression alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/dashboard_v2#alias_name DashboardV2#alias_name}
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
}

