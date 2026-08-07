// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardTab struct {
	// The name of the tab.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of widget references for this tab.
	//
	// Use.
	WidgetIds *[]*string `field:"required" json:"widgetIds" yaml:"widgetIds"`
}

