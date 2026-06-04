// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetTimeseriesDefinitionRequestQueryApmDependencyStatsQuery struct {
	// The data source for APM Dependency Stats queries. Valid values are `apm_dependency_stats`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#data_source DashboardV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// APM environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#env DashboardV2#env}
	Env *string `field:"required" json:"env" yaml:"env"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#name DashboardV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of operation on service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#operation_name DashboardV2#operation_name}
	OperationName *string `field:"required" json:"operationName" yaml:"operationName"`
	// APM resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#resource_name DashboardV2#resource_name}
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// APM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#service DashboardV2#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// APM statistic. Valid values are `avg_duration`, `avg_root_duration`, `avg_spans_per_trace`, `error_rate`, `pct_exec_time`, `pct_of_traces`, `total_traces_count`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#stat DashboardV2#stat}
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#cross_org_uuids DashboardV2#cross_org_uuids}
	CrossOrgUuids *[]*string `field:"optional" json:"crossOrgUuids" yaml:"crossOrgUuids"`
	// Determines whether stats for upstream or downstream dependencies should be queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#is_upstream DashboardV2#is_upstream}
	IsUpstream interface{} `field:"optional" json:"isUpstream" yaml:"isUpstream"`
	// The name of the second primary tag used within APM; required when `primary_tag_value` is specified. See https://docs.datadoghq.com/tracing/guide/setting_primary_tags_to_scope/#add-a-second-primary-tag-in-datadog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#primary_tag_name DashboardV2#primary_tag_name}
	PrimaryTagName *string `field:"optional" json:"primaryTagName" yaml:"primaryTagName"`
	// Filter APM data by the second primary tag. `primary_tag_name` must also be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/dashboard_v2#primary_tag_value DashboardV2#primary_tag_value}
	PrimaryTagValue *string `field:"optional" json:"primaryTagValue" yaml:"primaryTagValue"`
}

