// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorQuota struct {
	// The name of the quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#name ObservabilityPipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether to drop events exceeding the limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#drop_events ObservabilityPipeline#drop_events}
	DropEvents interface{} `field:"optional" json:"dropEvents" yaml:"dropEvents"`
	// Whether to ignore when partition fields are missing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#ignore_when_missing_partitions ObservabilityPipeline#ignore_when_missing_partitions}
	IgnoreWhenMissingPartitions interface{} `field:"optional" json:"ignoreWhenMissingPartitions" yaml:"ignoreWhenMissingPartitions"`
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#limit ObservabilityPipeline#limit}
	Limit interface{} `field:"optional" json:"limit" yaml:"limit"`
	// The action to take when the quota is exceeded: `drop`, `no_action`, or `overflow_routing`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#overflow_action ObservabilityPipeline#overflow_action}
	OverflowAction *string `field:"optional" json:"overflowAction" yaml:"overflowAction"`
	// override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#override ObservabilityPipeline#override}
	Override interface{} `field:"optional" json:"override" yaml:"override"`
	// List of partition fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#partition_fields ObservabilityPipeline#partition_fields}
	PartitionFields *[]*string `field:"optional" json:"partitionFields" yaml:"partitionFields"`
	// The action to take when the max number of buckets is exceeded: `drop`, `no_action`, or `overflow_routing`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#too_many_buckets_action ObservabilityPipeline#too_many_buckets_action}
	TooManyBucketsAction *string `field:"optional" json:"tooManyBucketsAction" yaml:"tooManyBucketsAction"`
}

