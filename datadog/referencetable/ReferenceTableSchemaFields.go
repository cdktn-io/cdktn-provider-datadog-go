// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package referencetable


type ReferenceTableSchemaFields struct {
	// The name of the field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#name ReferenceTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of the field. Must be one of: STRING, INT32. Valid values are `STRING`, `INT32`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/reference_table#type ReferenceTable#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

