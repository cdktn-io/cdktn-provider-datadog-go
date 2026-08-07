// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package teamsync


type TeamSyncSelectionState struct {
	// external_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/team_sync#external_id TeamSync#external_id}
	ExternalId *TeamSyncSelectionStateExternalId `field:"required" json:"externalId" yaml:"externalId"`
	// The operation to perform on the selected hierarchy. Valid values are `include`. Defaults to `"include"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/team_sync#operation TeamSync#operation}
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// The scope of the selection. Valid values are `subtree`. Defaults to `"subtree"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/team_sync#scope TeamSync#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

