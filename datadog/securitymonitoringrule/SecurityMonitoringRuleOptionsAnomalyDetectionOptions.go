// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleOptionsAnomalyDetectionOptions struct {
	// Duration in seconds of the time buckets used to aggregate events matched by the rule.
	//
	// Valid values are 300, 600, 900, 1800, 3600, 10800. Valid values are `300`, `600`, `900`, `1800`, `3600`, `10800`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/security_monitoring_rule#bucket_duration SecurityMonitoringRule#bucket_duration}
	BucketDuration *float64 `field:"optional" json:"bucketDuration" yaml:"bucketDuration"`
	// An optional parameter that sets how permissive anomaly detection is.
	//
	// Higher values require higher deviations before triggering a signal. Valid values are 1, 2, 3, 4, 5. Valid values are `1`, `2`, `3`, `4`, `5`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/security_monitoring_rule#detection_tolerance SecurityMonitoringRule#detection_tolerance}
	DetectionTolerance *float64 `field:"optional" json:"detectionTolerance" yaml:"detectionTolerance"`
	// Learning duration in hours.
	//
	// Anomaly detection waits for at least this amount of historical data before it starts evaluating. Valid values are 1, 6, 12, 24, 48, 168, 336. Valid values are `1`, `6`, `12`, `24`, `48`, `168`, `336`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/security_monitoring_rule#learning_duration SecurityMonitoringRule#learning_duration}
	LearningDuration *float64 `field:"optional" json:"learningDuration" yaml:"learningDuration"`
	// An optional override baseline to apply while the rule is in the learning period.
	//
	// Must be greater than or equal to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/security_monitoring_rule#learning_period_baseline SecurityMonitoringRule#learning_period_baseline}
	LearningPeriodBaseline *float64 `field:"optional" json:"learningPeriodBaseline" yaml:"learningPeriodBaseline"`
}

