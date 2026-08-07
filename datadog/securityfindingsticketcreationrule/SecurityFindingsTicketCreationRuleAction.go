// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityfindingsticketcreationrule


type SecurityFindingsTicketCreationRuleAction struct {
	// The maximum number of tickets the rule may create per day.
	//
	// If exceeded, one final ticket will be created, explaining the limit was hit and linking back to the responsible rule. Value must be between 1 and 500.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_findings_ticket_creation_rule#max_tickets_per_day SecurityFindingsTicketCreationRule#max_tickets_per_day}
	MaxTicketsPerDay *float64 `field:"required" json:"maxTicketsPerDay" yaml:"maxTicketsPerDay"`
	// The UUID of the Case Management project. Must be a valid UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_findings_ticket_creation_rule#project_id SecurityFindingsTicketCreationRule#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// The ticketing system to create tickets in. Valid values are `jira`, `case_management`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_findings_ticket_creation_rule#target SecurityFindingsTicketCreationRule#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// The UUID of the default assignee for created tickets. Must be a valid UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_findings_ticket_creation_rule#assignee_id SecurityFindingsTicketCreationRule#assignee_id}
	AssigneeId *string `field:"optional" json:"assigneeId" yaml:"assigneeId"`
	// A JSON-encoded object of custom fields of the Jira issue to create.
	//
	// For the list of available fields, see the [Jira documentation](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-createmeta-projectidorkey-issuetypes-issuetypeid-get).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_findings_ticket_creation_rule#fields SecurityFindingsTicketCreationRule#fields}
	Fields *string `field:"optional" json:"fields" yaml:"fields"`
}

