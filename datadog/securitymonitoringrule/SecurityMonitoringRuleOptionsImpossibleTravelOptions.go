// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleOptionsImpossibleTravelOptions struct {
	// If true, signals are suppressed for the first 24 hours.
	//
	// During that time, Datadog learns the user's regular access locations. This can be helpful to reduce noise and infer VPN usage or credentialed API access. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/security_monitoring_rule#baseline_user_locations SecurityMonitoringRule#baseline_user_locations}
	BaselineUserLocations interface{} `field:"optional" json:"baselineUserLocations" yaml:"baselineUserLocations"`
	// The duration in days during which Datadog learns a user's access locations before generating signals.
	//
	// Only applicable when `baseline_user_locations` is `true`. Defaults to `1` if unset. . Value must be between 1 and 30.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.0/docs/resources/security_monitoring_rule#baseline_user_locations_duration SecurityMonitoringRule#baseline_user_locations_duration}
	BaselineUserLocationsDuration *float64 `field:"optional" json:"baselineUserLocationsDuration" yaml:"baselineUserLocationsDuration"`
}

