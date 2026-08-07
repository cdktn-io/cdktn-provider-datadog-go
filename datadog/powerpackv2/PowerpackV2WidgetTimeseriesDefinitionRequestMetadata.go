// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetTimeseriesDefinitionRequestMetadata struct {
	// The expression name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#expression PowerpackV2#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The expression alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#alias_name PowerpackV2#alias_name}
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
}

