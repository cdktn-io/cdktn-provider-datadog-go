// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetGroupDefinitionWidgetSankeyDefinitionTimeFixed struct {
	// Start time in seconds since epoch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#from PowerpackV2#from}
	From *float64 `field:"required" json:"from" yaml:"from"`
	// End time in seconds since epoch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/powerpack_v2#to PowerpackV2#to}
	To *float64 `field:"required" json:"to" yaml:"to"`
}

