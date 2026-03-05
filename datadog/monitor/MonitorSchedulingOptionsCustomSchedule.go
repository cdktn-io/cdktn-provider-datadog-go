// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorSchedulingOptionsCustomSchedule struct {
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/monitor#recurrence Monitor#recurrence}
	Recurrence *MonitorSchedulingOptionsCustomScheduleRecurrence `field:"required" json:"recurrence" yaml:"recurrence"`
}

