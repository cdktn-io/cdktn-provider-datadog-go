// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicelevelobjective


type ServiceLevelObjectiveSliSpecificationCount struct {
	// The formula that specifies how to compute the good events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/service_level_objective#good_events_formula ServiceLevelObjective#good_events_formula}
	GoodEventsFormula *string `field:"required" json:"goodEventsFormula" yaml:"goodEventsFormula"`
	// queries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/service_level_objective#queries ServiceLevelObjective#queries}
	Queries interface{} `field:"required" json:"queries" yaml:"queries"`
	// The formula that specifies how to compute the total events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/service_level_objective#total_events_formula ServiceLevelObjective#total_events_formula}
	TotalEventsFormula *string `field:"required" json:"totalEventsFormula" yaml:"totalEventsFormula"`
}

