// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionCircleCi struct {
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/action_connection#api_key ActionConnection#api_key}
	ApiKey *ActionConnectionCircleCiApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
}

