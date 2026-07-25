// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionAzure struct {
	// tenant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#tenant ActionConnection#tenant}
	Tenant *ActionConnectionAzureTenant `field:"optional" json:"tenant" yaml:"tenant"`
}

