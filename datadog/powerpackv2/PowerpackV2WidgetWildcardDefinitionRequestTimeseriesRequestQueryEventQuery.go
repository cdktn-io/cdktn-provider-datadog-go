// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuery struct {
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#compute PowerpackV2#compute}
	Compute interface{} `field:"required" json:"compute" yaml:"compute"`
	// The data source for event platform-based queries.
	//
	// Valid values are `logs`, `spans`, `network`, `rum`, `security_signals`, `profiles`, `audit`, `events`, `ci_tests`, `ci_pipelines`, `incident_analytics`, `product_analytics`, `on_call_events`, `errors`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#data_source PowerpackV2#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#name PowerpackV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#cross_org_uuids PowerpackV2#cross_org_uuids}
	CrossOrgUuids *[]*string `field:"optional" json:"crossOrgUuids" yaml:"crossOrgUuids"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#group_by PowerpackV2#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// group_by_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#group_by_fields PowerpackV2#group_by_fields}
	GroupByFields *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQueryGroupByFields `field:"optional" json:"groupByFields" yaml:"groupByFields"`
	// An array of index names to query in the stream.
	//
	// Omit or use `[]` to query all indexes at once.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#indexes PowerpackV2#indexes}
	Indexes *[]*string `field:"optional" json:"indexes" yaml:"indexes"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#search PowerpackV2#search}
	Search *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryEventQuerySearch `field:"optional" json:"search" yaml:"search"`
	// Option for storage location. Feature in Private Beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/powerpack_v2#storage PowerpackV2#storage}
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
}

