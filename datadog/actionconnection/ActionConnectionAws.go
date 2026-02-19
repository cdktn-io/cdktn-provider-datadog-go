// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionAws struct {
	// assume_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/action_connection#assume_role ActionConnection#assume_role}
	AssumeRole *ActionConnectionAwsAssumeRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
}

