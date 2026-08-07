// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinitionView struct {
	// The 2-letter ISO code of a country to focus the map on, or `WORLD` for global view, or a region (`EMEA`, `APAC`, `LATAM`), or a continent (`NORTH_AMERICA`, `SOUTH_AMERICA`, `EUROPE`, `AFRICA`, `ASIA`, `OCEANIA`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard_v2#focus DashboardV2#focus}
	Focus *string `field:"required" json:"focus" yaml:"focus"`
}

