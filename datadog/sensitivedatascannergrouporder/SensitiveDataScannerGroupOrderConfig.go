// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sensitivedatascannergrouporder

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SensitiveDataScannerGroupOrderConfig struct {
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
	// The list of Sensitive Data Scanner group IDs, in order.
	//
	// Logs are tested against the query filter of each index one by one following the order of the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/sensitive_data_scanner_group_order#group_ids SensitiveDataScannerGroupOrder#group_ids}
	GroupIds *[]*string `field:"required" json:"groupIds" yaml:"groupIds"`
}

