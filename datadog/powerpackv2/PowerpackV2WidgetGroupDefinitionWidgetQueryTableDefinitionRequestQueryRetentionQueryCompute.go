// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryRetentionQueryCompute struct {
	// Aggregation for the retention query, including standard event aggregations and `pcNN` percentiles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#aggregation PowerpackV2#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// Metric for the retention computation. Valid values are `__dd.retention`, `__dd.retention_rate`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/powerpack_v2#metric PowerpackV2#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
}

