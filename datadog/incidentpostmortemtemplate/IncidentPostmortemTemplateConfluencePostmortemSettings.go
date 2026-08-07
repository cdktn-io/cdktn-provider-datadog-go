// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package incidentpostmortemtemplate


type IncidentPostmortemTemplateConfluencePostmortemSettings struct {
	// The ID of the Confluence account, a Datadog connected-account UUID (e.g. `3f9b1c2a-8d4e-4a11-9c2f-0b7e5d6a1f23`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/incident_postmortem_template#account_id IncidentPostmortemTemplate#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The ID of the parent Confluence page under which postmortems are created: a numeric page ID (e.g. `393217`), not a page path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/incident_postmortem_template#parent_id IncidentPostmortemTemplate#parent_id}
	ParentId *string `field:"optional" json:"parentId" yaml:"parentId"`
	// The Confluence space key (e.g. `ENG`), not a numeric space ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/incident_postmortem_template#space_id IncidentPostmortemTemplate#space_id}
	SpaceId *string `field:"optional" json:"spaceId" yaml:"spaceId"`
}

