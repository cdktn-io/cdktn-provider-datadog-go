// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logscustomdestination


type LogsCustomDestinationSplunkDestinationSourcetype struct {
	// The source type string. Set to `null` to omit the sourcetype from forwarded events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/logs_custom_destination#value LogsCustomDestination#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

