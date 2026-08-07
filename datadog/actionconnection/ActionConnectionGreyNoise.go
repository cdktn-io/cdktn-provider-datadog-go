// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionGreyNoise struct {
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/action_connection#api_key ActionConnection#api_key}
	ApiKey *ActionConnectionGreyNoiseApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
}

