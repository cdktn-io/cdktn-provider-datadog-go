// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionConfigCatSdkKey struct {
	// ConfigCat Public Management API password. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/action_connection#api_password ActionConnection#api_password}
	ApiPassword *string `field:"optional" json:"apiPassword" yaml:"apiPassword"`
	// ConfigCat Public Management API username. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/action_connection#api_username ActionConnection#api_username}
	ApiUsername *string `field:"optional" json:"apiUsername" yaml:"apiUsername"`
	// ConfigCat SDK key. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/action_connection#sdk_key ActionConnection#sdk_key}
	SdkKey *string `field:"optional" json:"sdkKey" yaml:"sdkKey"`
}

