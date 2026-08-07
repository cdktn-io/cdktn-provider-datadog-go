// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2


type PowerpackV2WidgetQueryValueDefinitionRequestRumQueryGroupBySortQuery struct {
	// The aggregation method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#aggregation PowerpackV2#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// Widget sorting methods. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#order PowerpackV2#order}
	Order *string `field:"required" json:"order" yaml:"order"`
	// The facet name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/powerpack_v2#facet PowerpackV2#facet}
	Facet *string `field:"optional" json:"facet" yaml:"facet"`
}

