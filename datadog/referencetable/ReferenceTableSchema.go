// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package referencetable


type ReferenceTableSchema struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/reference_table#fields ReferenceTable#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// List of field names that serve as primary keys for the table. Currently only one primary key is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.82.0/docs/resources/reference_table#primary_keys ReferenceTable#primary_keys}
	PrimaryKeys *[]*string `field:"optional" json:"primaryKeys" yaml:"primaryKeys"`
}

