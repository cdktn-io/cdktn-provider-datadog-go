// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestQuerySloQuery struct {
	// The data source for SLO queries. Valid values are `slo`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#data_source DashboardV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// SLO measures queries. Valid values are `good_events`, `bad_events`, `good_minutes`, `bad_minutes`, `slo_status`, `error_budget_remaining`, `burn_rate`, `error_budget_burndown`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#measure DashboardV2#measure}
	Measure *string `field:"required" json:"measure" yaml:"measure"`
	// ID of an SLO to query measures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#slo_id DashboardV2#slo_id}
	SloId *string `field:"required" json:"sloId" yaml:"sloId"`
	// Additional filters applied to the SLO query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#additional_query_filters DashboardV2#additional_query_filters}
	AdditionalQueryFilters *string `field:"optional" json:"additionalQueryFilters" yaml:"additionalQueryFilters"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#cross_org_uuids DashboardV2#cross_org_uuids}
	CrossOrgUuids *[]*string `field:"optional" json:"crossOrgUuids" yaml:"crossOrgUuids"`
	// Group mode to query measures. Valid values are `overall`, `components`. Defaults to `"overall"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#group_mode DashboardV2#group_mode}
	GroupMode *string `field:"optional" json:"groupMode" yaml:"groupMode"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#name DashboardV2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// type of the SLO to query. Valid values are `metric`, `monitor`, `time_slice`. Defaults to `"metric"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/dashboard_v2#slo_query_type DashboardV2#slo_query_type}
	SloQueryType *string `field:"optional" json:"sloQueryType" yaml:"sloQueryType"`
}

