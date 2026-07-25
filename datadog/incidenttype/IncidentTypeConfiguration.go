// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package incidenttype


type IncidentTypeConfiguration struct {
	// Whether incidents of this type can be deleted. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/incident_type#allow_incident_deletion IncidentType#allow_incident_deletion}
	AllowIncidentDeletion interface{} `field:"optional" json:"allowIncidentDeletion" yaml:"allowIncidentDeletion"`
	// Whether users can manually run a workflow from an incident of this type. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/incident_type#allow_workflows IncidentType#allow_workflows}
	AllowWorkflows interface{} `field:"optional" json:"allowWorkflows" yaml:"allowWorkflows"`
	// An optional message shown to users when they declare an incident of this type. Defaults to an empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/incident_type#create_message IncidentType#create_message}
	CreateMessage *string `field:"optional" json:"createMessage" yaml:"createMessage"`
	// Whether responders can edit incident timestamps for incidents of this type. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/incident_type#editable_timestamps IncidentType#editable_timestamps}
	EditableTimestamps interface{} `field:"optional" json:"editableTimestamps" yaml:"editableTimestamps"`
	// Whether responders can create private incidents of this type. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/incident_type#private_incidents IncidentType#private_incidents}
	PrivateIncidents interface{} `field:"optional" json:"privateIncidents" yaml:"privateIncidents"`
	// Whether the private toggle is enabled by default in the incident creation modal for this type. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/incident_type#private_incidents_by_default IncidentType#private_incidents_by_default}
	PrivateIncidentsByDefault interface{} `field:"optional" json:"privateIncidentsByDefault" yaml:"privateIncidentsByDefault"`
	// The source used to derive the incident slug.
	//
	// When set to `servicenow`, incidents display the ServiceNow record ID instead of the public ID. If no ServiceNow integration exists, the public ID is displayed. Defaults to `default`. Valid values are `default`, `servicenow`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/incident_type#slug_source IncidentType#slug_source}
	SlugSource *string `field:"optional" json:"slugSource" yaml:"slugSource"`
	// Whether test incidents of this type can be created. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/incident_type#test_incidents IncidentType#test_incidents}
	TestIncidents interface{} `field:"optional" json:"testIncidents" yaml:"testIncidents"`
}

