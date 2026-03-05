// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicelevelobjective


type ServiceLevelObjectiveSliSpecification struct {
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/service_level_objective#count ServiceLevelObjective#count}
	Count *ServiceLevelObjectiveSliSpecificationCount `field:"optional" json:"count" yaml:"count"`
	// time_slice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/service_level_objective#time_slice ServiceLevelObjective#time_slice}
	TimeSlice *ServiceLevelObjectiveSliSpecificationTimeSlice `field:"optional" json:"timeSlice" yaml:"timeSlice"`
}

