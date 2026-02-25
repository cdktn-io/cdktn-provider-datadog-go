// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sensitivedatascannergroup


type SensitiveDataScannerGroupFilter struct {
	// Query to filter the events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/sensitive_data_scanner_group#query SensitiveDataScannerGroup#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

