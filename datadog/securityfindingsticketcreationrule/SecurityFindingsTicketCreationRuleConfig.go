// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityfindingsticketcreationrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityFindingsTicketCreationRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The action to take when the ticket creation rule matches a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_findings_ticket_creation_rule#action SecurityFindingsTicketCreationRule#action}
	Action *SecurityFindingsTicketCreationRuleAction `field:"required" json:"action" yaml:"action"`
	// The name of the ticket creation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_findings_ticket_creation_rule#name SecurityFindingsTicketCreationRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Defines the scope of findings to which the automation rule applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_findings_ticket_creation_rule#rule SecurityFindingsTicketCreationRule#rule}
	Rule *SecurityFindingsTicketCreationRuleRule `field:"required" json:"rule" yaml:"rule"`
	// Whether the ticket creation rule is enabled. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/security_findings_ticket_creation_rule#enabled SecurityFindingsTicketCreationRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

