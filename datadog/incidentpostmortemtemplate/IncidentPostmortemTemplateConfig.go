// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package incidentpostmortemtemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IncidentPostmortemTemplateConfig struct {
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
	// The ID of the incident type this template is associated with. Immutable after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_postmortem_template#incident_type IncidentPostmortemTemplate#incident_type}
	IncidentType *string `field:"required" json:"incidentType" yaml:"incidentType"`
	// The name of the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_postmortem_template#name IncidentPostmortemTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// confluence_postmortem_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_postmortem_template#confluence_postmortem_settings IncidentPostmortemTemplate#confluence_postmortem_settings}
	ConfluencePostmortemSettings *IncidentPostmortemTemplateConfluencePostmortemSettings `field:"optional" json:"confluencePostmortemSettings" yaml:"confluencePostmortemSettings"`
	// The templated content of the postmortem, supporting Markdown and incident template variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_postmortem_template#content IncidentPostmortemTemplate#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// google_docs_postmortem_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_postmortem_template#google_docs_postmortem_settings IncidentPostmortemTemplate#google_docs_postmortem_settings}
	GoogleDocsPostmortemSettings *IncidentPostmortemTemplateGoogleDocsPostmortemSettings `field:"optional" json:"googleDocsPostmortemSettings" yaml:"googleDocsPostmortemSettings"`
	// Whether this template is a default for its incident type.
	//
	// The API stores a timestamp; the effective default for an incident type is the template with the most recent default timestamp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_postmortem_template#is_default IncidentPostmortemTemplate#is_default}
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
	// The location where the postmortem is created and stored. Valid values are: datadog_notebooks, confluence, google_docs. Defaults to datadog_notebooks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/incident_postmortem_template#location IncidentPostmortemTemplate#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
}

