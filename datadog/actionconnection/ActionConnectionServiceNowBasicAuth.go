// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionServiceNowBasicAuth struct {
	// ServiceNow instance. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/action_connection#instance ActionConnection#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// ServiceNow password. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/action_connection#password ActionConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// ServiceNow username. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/action_connection#username ActionConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

