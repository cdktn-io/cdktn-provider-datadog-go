// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package incidentuserdefinedrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IncidentUserDefinedRoleConfig struct {
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
	// The ID of the incident type this user-defined role is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/incident_user_defined_role#incident_type IncidentUserDefinedRole#incident_type}
	IncidentType *string `field:"required" json:"incidentType" yaml:"incidentType"`
	// The name of the user-defined role.
	//
	// Cannot be a reserved name ("Incident Commander" or "Responder") and must be at most 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/incident_user_defined_role#name IncidentUserDefinedRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the user-defined role. At most 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/incident_user_defined_role#description IncidentUserDefinedRole#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Policy configuration for the user-defined role. Defaults to a multi-assignee policy when omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/incident_user_defined_role#policy IncidentUserDefinedRole#policy}
	Policy *IncidentUserDefinedRolePolicy `field:"optional" json:"policy" yaml:"policy"`
}

