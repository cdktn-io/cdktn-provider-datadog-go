// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package incidentuserdefinedfield

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IncidentUserDefinedFieldConfig struct {
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
	// The ID of the incident type this field is associated with. Changing the incident type forces a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#incident_type IncidentUserDefinedField#incident_type}
	IncidentType *string `field:"required" json:"incidentType" yaml:"incidentType"`
	// The unique identifier of the field.
	//
	// Must start with a letter or digit and contain only letters, digits, underscores, or periods. Changing the name forces a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#name IncidentUserDefinedField#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of the field.
	//
	// Changing the type forces a new resource. Valid values are `dropdown`, `multiselect`, `textbox`, `textarray`, `metrictag`, `autocomplete`, `number`, `datetime`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#type IncidentUserDefinedField#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The section in which the field appears: `what_happened` or `why_it_happened`. When unset, the field appears in the Attributes section.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#category IncidentUserDefinedField#category}
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The default value for the field. Must be one of the valid values when `valid_values` is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#default_value IncidentUserDefinedField#default_value}
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The human-readable name shown in the UI. Defaults to a formatted version of the name if not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#display_name IncidentUserDefinedField#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// A decimal string representing the field's display order in the UI. Assigned by the server when not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#ordinal IncidentUserDefinedField#ordinal}
	Ordinal *string `field:"optional" json:"ordinal" yaml:"ordinal"`
	// When true, users must fill out this field on incidents. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#required IncidentUserDefinedField#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// For metric tag-type fields only, the metric tag key that powers the autocomplete options.
	//
	// Changing the tag key forces a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#tag_key IncidentUserDefinedField#tag_key}
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// valid_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_user_defined_field#valid_value IncidentUserDefinedField#valid_value}
	ValidValue interface{} `field:"optional" json:"validValue" yaml:"validValue"`
}

