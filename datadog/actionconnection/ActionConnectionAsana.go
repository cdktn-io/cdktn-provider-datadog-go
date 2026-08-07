// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionAsana struct {
	// access_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/action_connection#access_token ActionConnection#access_token}
	AccessToken *ActionConnectionAsanaAccessToken `field:"optional" json:"accessToken" yaml:"accessToken"`
}

