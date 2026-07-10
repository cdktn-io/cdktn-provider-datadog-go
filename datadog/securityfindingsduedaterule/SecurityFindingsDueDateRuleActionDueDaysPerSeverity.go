// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityfindingsduedaterule


type SecurityFindingsDueDateRuleActionDueDaysPerSeverity struct {
	// The number of days from the reference point until the finding is due.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/security_findings_due_date_rule#due_in_days SecurityFindingsDueDateRule#due_in_days}
	DueInDays *float64 `field:"required" json:"dueInDays" yaml:"dueInDays"`
	// A severity level used to configure due date thresholds. Valid values are `critical`, `high`, `medium`, `low`, `info`, `none`, `unknown`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.15.0/docs/resources/security_findings_due_date_rule#severity SecurityFindingsDueDateRule#severity}
	Severity *string `field:"required" json:"severity" yaml:"severity"`
}

