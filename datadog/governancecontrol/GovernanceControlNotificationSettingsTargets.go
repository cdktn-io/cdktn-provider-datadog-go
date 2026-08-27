// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package governancecontrol


type GovernanceControlNotificationSettingsTargets struct {
	// The handle of the notification target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/governance_control#handle GovernanceControl#handle}
	Handle *string `field:"required" json:"handle" yaml:"handle"`
	// The type of notification target: `email`, `slack`, `at_mention`, or `case`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/governance_control#type GovernanceControl#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

