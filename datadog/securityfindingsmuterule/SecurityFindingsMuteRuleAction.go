// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityfindingsmuterule


type SecurityFindingsMuteRuleAction struct {
	// The reason for muting a security finding. Valid values are `duplicate`, `false_positive`, `no_fix`, `other`, `pending_fix`, `risk_accepted`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/security_findings_mute_rule#reason SecurityFindingsMuteRule#reason}
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	// The Unix timestamp in milliseconds at which the mute expires. If omitted, the mute does not expire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/security_findings_mute_rule#expire_at SecurityFindingsMuteRule#expire_at}
	ExpireAt *float64 `field:"optional" json:"expireAt" yaml:"expireAt"`
	// An optional description providing more context for the mute reason.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/security_findings_mute_rule#reason_description SecurityFindingsMuteRule#reason_description}
	ReasonDescription *string `field:"optional" json:"reasonDescription" yaml:"reasonDescription"`
}

