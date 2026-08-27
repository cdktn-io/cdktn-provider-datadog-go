// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetBarChartDefinitionRequestQueryApmMetricsQuery struct {
	// The data source for APM metrics queries. Valid values are `apm_metrics`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#data_source DashboardV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Name of this query to use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#name DashboardV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// APM metric stat name.
	//
	// Valid values are `errors`, `error_rate`, `errors_per_second`, `latency_avg`, `latency_max`, `latency_p50`, `latency_p75`, `latency_p90`, `latency_p95`, `latency_p99`, `latency_p999`, `latency_distribution`, `hits`, `hits_per_second`, `total_time`, `apdex`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#stat DashboardV2#stat}
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// Optional fields to group the query results by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#group_by DashboardV2#group_by}
	GroupBy *[]*string `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Optional operation mode used to aggregate across operation names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#operation_mode DashboardV2#operation_mode}
	OperationMode *string `field:"optional" json:"operationMode" yaml:"operationMode"`
	// Name of the operation on the service. If omitted, the primary operation name is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#operation_name DashboardV2#operation_name}
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// Tags to query for a specific downstream entity, such as `peer.service` or `peer.db_instance`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#peer_tags DashboardV2#peer_tags}
	PeerTags *[]*string `field:"optional" json:"peerTags" yaml:"peerTags"`
	// Additional filters for the query using metrics query syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#query_filter DashboardV2#query_filter}
	QueryFilter *string `field:"optional" json:"queryFilter" yaml:"queryFilter"`
	// The hash of a specific resource to filter by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#resource_hash DashboardV2#resource_hash}
	ResourceHash *string `field:"optional" json:"resourceHash" yaml:"resourceHash"`
	// The full name of a specific resource to filter by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#resource_name DashboardV2#resource_name}
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
	// APM service name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#service DashboardV2#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// The relationship between the span, its parents, and its children in a trace.
	//
	// Valid values are `consumer`, `server`, `client`, `producer`, `internal`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/dashboard_v2#span_kind DashboardV2#span_kind}
	SpanKind *string `field:"optional" json:"spanKind" yaml:"spanKind"`
}

