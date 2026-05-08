// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2


type DashboardV2WidgetPowerpackDefinitionTemplateVariables struct {
	// controlled_by_powerpack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#controlled_by_powerpack DashboardV2#controlled_by_powerpack}
	ControlledByPowerpack interface{} `field:"optional" json:"controlledByPowerpack" yaml:"controlledByPowerpack"`
	// controlled_externally block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#controlled_externally DashboardV2#controlled_externally}
	ControlledExternally interface{} `field:"optional" json:"controlledExternally" yaml:"controlledExternally"`
}

