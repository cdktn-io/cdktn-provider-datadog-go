// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionOkta struct {
	// api_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/action_connection#api_token ActionConnection#api_token}
	ApiToken *ActionConnectionOktaApiToken `field:"optional" json:"apiToken" yaml:"apiToken"`
}

