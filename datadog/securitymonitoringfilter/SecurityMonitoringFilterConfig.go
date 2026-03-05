// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringfilter

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityMonitoringFilterConfig struct {
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
	// Whether the security filter is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/security_monitoring_filter#is_enabled SecurityMonitoringFilter#is_enabled}
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// The name of the security filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/security_monitoring_filter#name SecurityMonitoringFilter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The query of the security filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/security_monitoring_filter#query SecurityMonitoringFilter#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// exclusion_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/security_monitoring_filter#exclusion_filter SecurityMonitoringFilter#exclusion_filter}
	ExclusionFilter interface{} `field:"optional" json:"exclusionFilter" yaml:"exclusionFilter"`
	// The filtered data type. Valid values are `logs`. Defaults to `"logs"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/security_monitoring_filter#filtered_data_type SecurityMonitoringFilter#filtered_data_type}
	FilteredDataType *string `field:"optional" json:"filteredDataType" yaml:"filteredDataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/security_monitoring_filter#id SecurityMonitoringFilter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

