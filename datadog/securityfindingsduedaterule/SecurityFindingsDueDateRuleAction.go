// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityfindingsduedaterule


type SecurityFindingsDueDateRuleAction struct {
	// A list of severity-to-due-date mappings. Each severity may appear at most once.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/security_findings_due_date_rule#due_days_per_severity SecurityFindingsDueDateRule#due_days_per_severity}
	DueDaysPerSeverity interface{} `field:"required" json:"dueDaysPerSeverity" yaml:"dueDaysPerSeverity"`
	// The reference point from which the due date is calculated.
	//
	// When `fix_available` is selected but not applicable to the finding type, `first_seen` is used instead. Valid values are `first_seen`, `fix_available`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/security_findings_due_date_rule#due_from SecurityFindingsDueDateRule#due_from}
	DueFrom *string `field:"required" json:"dueFrom" yaml:"dueFrom"`
	// An optional description providing more context for the due date assignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/security_findings_due_date_rule#reason_description SecurityFindingsDueDateRule#reason_description}
	ReasonDescription *string `field:"optional" json:"reasonDescription" yaml:"reasonDescription"`
}

