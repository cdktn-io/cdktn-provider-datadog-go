// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionGcp struct {
	// service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#service_account ActionConnection#service_account}
	ServiceAccount *ActionConnectionGcpServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

