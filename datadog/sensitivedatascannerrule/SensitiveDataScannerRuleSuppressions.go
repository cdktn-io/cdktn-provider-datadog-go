// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sensitivedatascannerrule


type SensitiveDataScannerRuleSuppressions struct {
	// Any match that ends with a value in this list will be suppressed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/sensitive_data_scanner_rule#ends_with SensitiveDataScannerRule#ends_with}
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// Any match that appears in this list will be suppressed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/sensitive_data_scanner_rule#exact_match SensitiveDataScannerRule#exact_match}
	ExactMatch *[]*string `field:"optional" json:"exactMatch" yaml:"exactMatch"`
	// Any match that starts with a value in this list will be suppressed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.9.0/docs/resources/sensitive_data_scanner_rule#starts_with SensitiveDataScannerRule#starts_with}
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

