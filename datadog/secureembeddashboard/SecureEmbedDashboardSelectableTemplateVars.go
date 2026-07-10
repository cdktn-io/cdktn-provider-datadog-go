// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secureembeddashboard


type SecureEmbedDashboardSelectableTemplateVars struct {
	// The name of the template variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/secure_embed_dashboard#name SecureEmbedDashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The default values for this template variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/secure_embed_dashboard#default_values SecureEmbedDashboard#default_values}
	DefaultValues *[]*string `field:"optional" json:"defaultValues" yaml:"defaultValues"`
	// The tag prefix for this template variable. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/secure_embed_dashboard#prefix SecureEmbedDashboard#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The type of the template variable. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/secure_embed_dashboard#type SecureEmbedDashboard#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The visible tag values for this template variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/secure_embed_dashboard#visible_tags SecureEmbedDashboard#visible_tags}
	VisibleTags *[]*string `field:"optional" json:"visibleTags" yaml:"visibleTags"`
}

