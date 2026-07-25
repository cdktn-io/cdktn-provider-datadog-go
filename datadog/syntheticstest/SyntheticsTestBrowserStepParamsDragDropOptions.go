// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestBrowserStepParamsDragDropOptions struct {
	// Delay in milliseconds before performing the action (0–9999).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/synthetics_test#delay SyntheticsTest#delay}
	Delay *float64 `field:"optional" json:"delay" yaml:"delay"`
	// offset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/synthetics_test#offset SyntheticsTest#offset}
	Offset *SyntheticsTestBrowserStepParamsDragDropOptionsOffset `field:"optional" json:"offset" yaml:"offset"`
}

