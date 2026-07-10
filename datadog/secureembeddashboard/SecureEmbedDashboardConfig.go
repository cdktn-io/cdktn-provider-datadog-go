// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secureembeddashboard

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecureEmbedDashboardConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the dashboard to create a secure embed for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/secure_embed_dashboard#dashboard_id SecureEmbedDashboard#dashboard_id}
	DashboardId *string `field:"required" json:"dashboardId" yaml:"dashboardId"`
	// Title of the secure embed share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/secure_embed_dashboard#title SecureEmbedDashboard#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// The live span for the global time, e.g. `1h`, `4h`, `1d`, `2d`, `1w`. Defaults to `"1h"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/secure_embed_dashboard#global_time_live_span SecureEmbedDashboard#global_time_live_span}
	GlobalTimeLiveSpan *string `field:"optional" json:"globalTimeLiveSpan" yaml:"globalTimeLiveSpan"`
	// Whether viewers can change the global time range. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/secure_embed_dashboard#global_time_selectable SecureEmbedDashboard#global_time_selectable}
	GlobalTimeSelectable interface{} `field:"optional" json:"globalTimeSelectable" yaml:"globalTimeSelectable"`
	// selectable_template_vars block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/secure_embed_dashboard#selectable_template_vars SecureEmbedDashboard#selectable_template_vars}
	SelectableTemplateVars interface{} `field:"optional" json:"selectableTemplateVars" yaml:"selectableTemplateVars"`
	// Status of the secure embed. Valid values are `active` and `paused`. Defaults to `"active"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/secure_embed_dashboard#status SecureEmbedDashboard#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Whether to display the dashboard in high density mode. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/secure_embed_dashboard#viewing_preferences_high_density SecureEmbedDashboard#viewing_preferences_high_density}
	ViewingPreferencesHighDensity interface{} `field:"optional" json:"viewingPreferencesHighDensity" yaml:"viewingPreferencesHighDensity"`
	// Display theme for the embedded dashboard. Valid values are `system`, `dark`, `light`. Defaults to `"system"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/secure_embed_dashboard#viewing_preferences_theme SecureEmbedDashboard#viewing_preferences_theme}
	ViewingPreferencesTheme *string `field:"optional" json:"viewingPreferencesTheme" yaml:"viewingPreferencesTheme"`
}

