// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package incidentpostmortemtemplate


type IncidentPostmortemTemplateGoogleDocsPostmortemSettings struct {
	// The ID of the Google Drive account, a Datadog connected-account UUID (e.g. `a1b2c3d4-e5f6-4789-8abc-1234567890ab`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_postmortem_template#account_id IncidentPostmortemTemplate#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The Google Drive folder ID where postmortems are created, taken from the folder URL (e.g. `1eCqLAKQqRHt49J2aqQLGUcnPMzGHkt2B`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_postmortem_template#parent_folder_id IncidentPostmortemTemplate#parent_folder_id}
	ParentFolderId *string `field:"optional" json:"parentFolderId" yaml:"parentFolderId"`
}

