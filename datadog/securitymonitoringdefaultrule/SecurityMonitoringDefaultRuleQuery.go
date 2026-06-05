// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringdefaultrule


type SecurityMonitoringDefaultRuleQuery struct {
	// agent_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/security_monitoring_default_rule#agent_rule SecurityMonitoringDefaultRule#agent_rule}
	AgentRule interface{} `field:"optional" json:"agentRule" yaml:"agentRule"`
	// Query extension to append to the logs query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.12.1/docs/resources/security_monitoring_default_rule#custom_query_extension SecurityMonitoringDefaultRule#custom_query_extension}
	CustomQueryExtension *string `field:"optional" json:"customQueryExtension" yaml:"customQueryExtension"`
}

