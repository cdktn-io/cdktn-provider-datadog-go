// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionServiceNow struct {
	// basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/action_connection#basic_auth ActionConnection#basic_auth}
	BasicAuth *ActionConnectionServiceNowBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
}

