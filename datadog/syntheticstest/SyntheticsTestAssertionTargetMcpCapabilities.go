// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestAssertionTargetMcpCapabilities struct {
	// List of MCP server capabilities to assert against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/synthetics_test#capabilities SyntheticsTest#capabilities}
	Capabilities *[]*string `field:"required" json:"capabilities" yaml:"capabilities"`
}

