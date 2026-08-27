// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityfindingsduedaterule


type SecurityFindingsDueDateRuleRule struct {
	// The list of security finding types that the automation rule applies to.
	//
	// Valid values are `api_security`, `attack_path`, `host_and_container_vulnerability`, `iac_misconfiguration`, `identity_risk`, `library_vulnerability`, `misconfiguration`, `runtime_code_vulnerability`, `secret`, `static_code_vulnerability`, `workload_activity`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/security_findings_due_date_rule#finding_types SecurityFindingsDueDateRule#finding_types}
	FindingTypes *[]*string `field:"required" json:"findingTypes" yaml:"findingTypes"`
	// A search query to further filter the findings matched by this rule.
	//
	// The `@workflow.*` namespace and `@status` fields are not permitted. For a reference of available fields, see the [Security Findings schema documentation](https://docs.datadoghq.com/security/guide/findings-schema/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/security_findings_due_date_rule#query SecurityFindingsDueDateRule#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

