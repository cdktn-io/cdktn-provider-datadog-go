// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2Config struct {
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
	// The layout type of the dashboard. Valid values are `ordered`, `free`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#layout_type DashboardV2#layout_type}
	LayoutType *string `field:"required" json:"layoutType" yaml:"layoutType"`
	// The title of the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#title DashboardV2#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// A list of dashboard lists this dashboard belongs to.
	//
	// This attribute should not be set if managing the corresponding dashboard lists using Terraform as it causes inconsistent behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#dashboard_lists DashboardV2#dashboard_lists}
	DashboardLists *[]*float64 `field:"optional" json:"dashboardLists" yaml:"dashboardLists"`
	// A list of dashboard lists this dashboard should be removed from. Internal only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#dashboard_lists_removed DashboardV2#dashboard_lists_removed}
	DashboardListsRemoved *[]*float64 `field:"optional" json:"dashboardListsRemoved" yaml:"dashboardListsRemoved"`
	// The description of the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#description DashboardV2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#id DashboardV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether this dashboard is read-only.
	//
	// **Deprecated.** This field is deprecated and non-functional. Use `restricted_roles` instead to define which roles are required to edit the dashboard. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#is_read_only DashboardV2#is_read_only}
	IsReadOnly interface{} `field:"optional" json:"isReadOnly" yaml:"isReadOnly"`
	// The list of handles for the users to notify when changes are made to this dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#notify_list DashboardV2#notify_list}
	NotifyList *[]*string `field:"optional" json:"notifyList" yaml:"notifyList"`
	// The reflow type of a new dashboard layout.
	//
	// Set this only when layout type is `ordered`. If set to `fixed`, the dashboard expects all widgets to have a layout, and if it's set to `auto`, widgets should not have layouts. Valid values are `auto`, `fixed`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#reflow_type DashboardV2#reflow_type}
	ReflowType *string `field:"optional" json:"reflowType" yaml:"reflowType"`
	// A list of role identifiers.
	//
	// Only the author and users associated with at least one of these roles can edit this dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#restricted_roles DashboardV2#restricted_roles}
	RestrictedRoles *[]*string `field:"optional" json:"restrictedRoles" yaml:"restrictedRoles"`
	// tab block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#tab DashboardV2#tab}
	Tab interface{} `field:"optional" json:"tab" yaml:"tab"`
	// A list of tags assigned to the Dashboard. Only team names of the form `team:<name>` are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#tags DashboardV2#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// template_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#template_variable DashboardV2#template_variable}
	TemplateVariable interface{} `field:"optional" json:"templateVariable" yaml:"templateVariable"`
	// template_variable_preset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#template_variable_preset DashboardV2#template_variable_preset}
	TemplateVariablePreset interface{} `field:"optional" json:"templateVariablePreset" yaml:"templateVariablePreset"`
	// The URL of the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#url DashboardV2#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// widget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.8.0/docs/resources/dashboard_v2#widget DashboardV2#widget}
	Widget interface{} `field:"optional" json:"widget" yaml:"widget"`
}

