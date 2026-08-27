// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package governancecontrol

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GovernanceControlConfig struct {
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
	// The detection type that uniquely identifies the control, for example `unused_api_keys`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/governance_control#detection_type GovernanceControl#detection_type}
	DetectionType *string `field:"required" json:"detectionType" yaml:"detectionType"`
	// Detection parameters for the control, as a JSON-encoded map of parameter names to their configured values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/governance_control#detection_parameters GovernanceControl#detection_parameters}
	DetectionParameters *string `field:"optional" json:"detectionParameters" yaml:"detectionParameters"`
	// Mitigation parameters for the control, as a JSON-encoded map of parameter names to their configured values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/governance_control#mitigation_parameters GovernanceControl#mitigation_parameters}
	MitigationParameters *string `field:"optional" json:"mitigationParameters" yaml:"mitigationParameters"`
	// The mitigation type configured for the control. Empty when not configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/governance_control#mitigation_type GovernanceControl#mitigation_type}
	MitigationType *string `field:"optional" json:"mitigationType" yaml:"mitigationType"`
	// The notification settings for the control, one entry per event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.19.0/docs/resources/governance_control#notification_settings GovernanceControl#notification_settings}
	NotificationSettings interface{} `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
}

