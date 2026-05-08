// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package teamsync


type TeamSyncSelectionStateExternalId struct {
	// The type of external identifier. Valid values are `team`, `organization`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/team_sync#type TeamSync#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The external identifier value from the source platform (e.g. a GitHub organization ID or team ID).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/team_sync#value TeamSync#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

