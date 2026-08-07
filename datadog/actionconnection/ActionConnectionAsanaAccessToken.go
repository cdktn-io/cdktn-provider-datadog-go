// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionAsanaAccessToken struct {
	// Asana access token. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/action_connection#access_token ActionConnection#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
}

