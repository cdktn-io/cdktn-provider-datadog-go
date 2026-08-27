// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package incidentuserdefinedfield


type IncidentUserDefinedFieldValidValue struct {
	// The human-readable display name for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#display_name IncidentUserDefinedField#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The identifier that is stored when this option is selected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#value IncidentUserDefinedField#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// A detailed description of the valid value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#description IncidentUserDefinedField#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A short description of the valid value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#short_description IncidentUserDefinedField#short_description}
	ShortDescription *string `field:"optional" json:"shortDescription" yaml:"shortDescription"`
}

