// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicelevelobjective


type ServiceLevelObjectiveSliSpecificationCountQueries struct {
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/service_level_objective#metric_query ServiceLevelObjective#metric_query}
	MetricQuery *ServiceLevelObjectiveSliSpecificationCountQueriesMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
}

