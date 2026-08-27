// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package governancecontrol


type GovernanceControlNotificationSettings struct {
	// Whether notifications are enabled for this event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/governance_control#enabled GovernanceControl#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The event type the notification settings apply to, such as `new_detection`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/governance_control#event_type GovernanceControl#event_type}
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// The destinations that receive notifications for this event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/governance_control#targets GovernanceControl#targets}
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
}

