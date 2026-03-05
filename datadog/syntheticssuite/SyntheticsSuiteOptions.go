// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticssuite


type SyntheticsSuiteOptions struct {
	// Alerting threshold for the suite. Value must be between 0.000000 and 1.000000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/synthetics_suite#alerting_threshold SyntheticsSuite#alerting_threshold}
	AlertingThreshold *float64 `field:"required" json:"alertingThreshold" yaml:"alertingThreshold"`
}

