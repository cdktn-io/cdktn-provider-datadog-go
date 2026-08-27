// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionConfigCat struct {
	// sdk_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/action_connection#sdk_key ActionConnection#sdk_key}
	SdkKey *ActionConnectionConfigCatSdkKey `field:"optional" json:"sdkKey" yaml:"sdkKey"`
}

