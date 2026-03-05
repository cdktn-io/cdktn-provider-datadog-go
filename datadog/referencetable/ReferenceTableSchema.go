// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package referencetable


type ReferenceTableSchema struct {
	// List of field names that serve as primary keys for the table. Currently only one primary key is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/reference_table#primary_keys ReferenceTable#primary_keys}
	PrimaryKeys *[]*string `field:"required" json:"primaryKeys" yaml:"primaryKeys"`
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/reference_table#fields ReferenceTable#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
}

