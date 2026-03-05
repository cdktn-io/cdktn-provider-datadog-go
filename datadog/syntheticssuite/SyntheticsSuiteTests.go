// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticssuite


type SyntheticsSuiteTests struct {
	// Public ID of the test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/synthetics_suite#public_id SyntheticsSuite#public_id}
	PublicId *string `field:"required" json:"publicId" yaml:"publicId"`
	// Alerting criticality for the test. Valid values are `ignore`, `critical`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/synthetics_suite#alerting_criticality SyntheticsSuite#alerting_criticality}
	AlertingCriticality *string `field:"optional" json:"alertingCriticality" yaml:"alertingCriticality"`
}

