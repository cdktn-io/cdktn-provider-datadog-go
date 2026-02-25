// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package orgconnection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OrgConnectionConfig struct {
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
	// Set of connection types to enable for this connection (., metrics, logs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/org_connection#connection_types OrgConnection#connection_types}
	ConnectionTypes *[]*string `field:"required" json:"connectionTypes" yaml:"connectionTypes"`
	// UUID of the sink (destination) organization. Must be a valid UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/org_connection#sink_org_id OrgConnection#sink_org_id}
	SinkOrgId *string `field:"required" json:"sinkOrgId" yaml:"sinkOrgId"`
}

