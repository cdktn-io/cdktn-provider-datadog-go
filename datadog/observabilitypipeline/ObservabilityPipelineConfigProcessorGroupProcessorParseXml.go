// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorGroupProcessorParseXml struct {
	// The path to the log field on which you want to parse XML.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#field ObservabilityPipeline#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// Whether to always store text inside an object using the text key even when no attributes exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#always_use_text_key ObservabilityPipeline#always_use_text_key}
	AlwaysUseTextKey interface{} `field:"optional" json:"alwaysUseTextKey" yaml:"alwaysUseTextKey"`
	// The prefix to use for XML attributes in the parsed output.
	//
	// If the field is left empty, the original attribute key is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#attr_prefix ObservabilityPipeline#attr_prefix}
	AttrPrefix *string `field:"optional" json:"attrPrefix" yaml:"attrPrefix"`
	// Whether to include XML attributes in the parsed output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#include_attr ObservabilityPipeline#include_attr}
	IncludeAttr interface{} `field:"optional" json:"includeAttr" yaml:"includeAttr"`
	// Whether to parse boolean values from strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#parse_bool ObservabilityPipeline#parse_bool}
	ParseBool interface{} `field:"optional" json:"parseBool" yaml:"parseBool"`
	// Whether to parse null values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#parse_null ObservabilityPipeline#parse_null}
	ParseNull interface{} `field:"optional" json:"parseNull" yaml:"parseNull"`
	// Whether to parse numeric values from strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#parse_number ObservabilityPipeline#parse_number}
	ParseNumber interface{} `field:"optional" json:"parseNumber" yaml:"parseNumber"`
	// The key name to use for the text node when XML attributes are appended.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/observability_pipeline#text_key ObservabilityPipeline#text_key}
	TextKey *string `field:"optional" json:"textKey" yaml:"textKey"`
}

