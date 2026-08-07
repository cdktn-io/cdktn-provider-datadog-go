// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package teamsync

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TeamSyncConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The external source platform for team synchronization. Valid values are `github`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/team_sync#source TeamSync#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// The type of synchronization operation.
	//
	// `link` connects teams by matching names. `provision` creates new teams when no match is found. Valid values are `link`, `provision`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/team_sync#type TeamSync#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// How often the sync process should run. Valid values are `once`, `continuously`, `paused`. Defaults to `"once"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/team_sync#frequency TeamSync#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// selection_state block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/team_sync#selection_state TeamSync#selection_state}
	SelectionState interface{} `field:"optional" json:"selectionState" yaml:"selectionState"`
	// Whether to sync members from the external team to the Datadog team. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/team_sync#sync_membership TeamSync#sync_membership}
	SyncMembership interface{} `field:"optional" json:"syncMembership" yaml:"syncMembership"`
}

