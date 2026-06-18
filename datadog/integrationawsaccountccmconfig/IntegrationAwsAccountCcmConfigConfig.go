// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccountccmconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IntegrationAwsAccountCcmConfigConfig struct {
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
	// Unique Datadog ID of the AWS Account Integration Config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/integration_aws_account_ccm_config#aws_account_config_id IntegrationAwsAccountCcmConfig#aws_account_config_id}
	AwsAccountConfigId *string `field:"required" json:"awsAccountConfigId" yaml:"awsAccountConfigId"`
	// ccm_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.13.0/docs/resources/integration_aws_account_ccm_config#ccm_config IntegrationAwsAccountCcmConfig#ccm_config}
	CcmConfig *IntegrationAwsAccountCcmConfigCcmConfig `field:"optional" json:"ccmConfig" yaml:"ccmConfig"`
}

