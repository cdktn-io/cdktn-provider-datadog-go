// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatadogreferencetablerows

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatadogReferenceTableRowsConfig struct {
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
	// List of primary key values (row IDs) to retrieve.
	//
	// These are the values of the table's primary key field(s). Maximum 250 IDs per request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/reference_table_rows#row_ids DataDatadogReferenceTableRows#row_ids}
	RowIds *[]*string `field:"required" json:"rowIds" yaml:"rowIds"`
	// The UUID of the reference table to query rows from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/reference_table_rows#table_id DataDatadogReferenceTableRows#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
	// rows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/data-sources/reference_table_rows#rows DataDatadogReferenceTableRows#rows}
	Rows interface{} `field:"optional" json:"rows" yaml:"rows"`
}

