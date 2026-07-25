// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package incidentuserdefinedrole


type IncidentUserDefinedRolePolicy struct {
	// Whether this role can only be assigned to one responder at a time. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/incident_user_defined_role#is_single IncidentUserDefinedRole#is_single}
	IsSingle interface{} `field:"optional" json:"isSingle" yaml:"isSingle"`
}

