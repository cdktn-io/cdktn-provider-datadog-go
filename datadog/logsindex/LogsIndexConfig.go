// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsindex

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogsIndexConfig struct {
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
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/logs_index#filter LogsIndex#filter}
	Filter *LogsIndexFilter `field:"required" json:"filter" yaml:"filter"`
	// The name of the index.
	//
	// Index names cannot be modified after creation. If this value is changed, a new index will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/logs_index#name LogsIndex#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The number of log events you can send in this index per day before you are rate-limited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/logs_index#daily_limit LogsIndex#daily_limit}
	DailyLimit *float64 `field:"optional" json:"dailyLimit" yaml:"dailyLimit"`
	// daily_limit_reset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/logs_index#daily_limit_reset LogsIndex#daily_limit_reset}
	DailyLimitReset *LogsIndexDailyLimitReset `field:"optional" json:"dailyLimitReset" yaml:"dailyLimitReset"`
	// A percentage threshold of the daily quota at which a Datadog warning event is generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/logs_index#daily_limit_warning_threshold_percentage LogsIndex#daily_limit_warning_threshold_percentage}
	DailyLimitWarningThresholdPercentage *float64 `field:"optional" json:"dailyLimitWarningThresholdPercentage" yaml:"dailyLimitWarningThresholdPercentage"`
	// If true, sets the daily_limit value to null and the index is not limited on a daily basis (any specified daily_limit value in the request is ignored).
	//
	// If false or omitted, the index's current daily_limit is maintained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/logs_index#disable_daily_limit LogsIndex#disable_daily_limit}
	DisableDailyLimit interface{} `field:"optional" json:"disableDailyLimit" yaml:"disableDailyLimit"`
	// exclusion_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/logs_index#exclusion_filter LogsIndex#exclusion_filter}
	ExclusionFilter interface{} `field:"optional" json:"exclusionFilter" yaml:"exclusionFilter"`
	// The total number of days logs are stored in Standard and Flex Tier before being deleted from the index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/logs_index#flex_retention_days LogsIndex#flex_retention_days}
	FlexRetentionDays *float64 `field:"optional" json:"flexRetentionDays" yaml:"flexRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/logs_index#id LogsIndex#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The number of days logs are stored in Standard Tier before aging into the Flex Tier or being deleted from the index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/logs_index#retention_days LogsIndex#retention_days}
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
	// A list of tags for this index.
	//
	// Tags must be in `key:value` format. If default tags are present at the provider level, they will be added to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.90.0/docs/resources/logs_index#tags LogsIndex#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

