// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSplitConfigSort struct {
	// Widget sorting methods. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#order DashboardV2#order}
	Order *string `field:"required" json:"order" yaml:"order"`
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#compute DashboardV2#compute}
	Compute *DashboardV2WidgetSplitGraphDefinitionSplitConfigSortCompute `field:"optional" json:"compute" yaml:"compute"`
}

