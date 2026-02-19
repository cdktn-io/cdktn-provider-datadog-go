// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/observabilitypipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityPipelineConfigProcessorGroupProcessorOutputReference interface {
	cdktn.ComplexObject
	AddEnvVars() ObservabilityPipelineConfigProcessorGroupProcessorAddEnvVarsList
	AddEnvVarsInput() interface{}
	AddFields() ObservabilityPipelineConfigProcessorGroupProcessorAddFieldsList
	AddFieldsInput() interface{}
	AddHostname() ObservabilityPipelineConfigProcessorGroupProcessorAddHostnameList
	AddHostnameInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomProcessor() ObservabilityPipelineConfigProcessorGroupProcessorCustomProcessorList
	CustomProcessorInput() interface{}
	DatadogTags() ObservabilityPipelineConfigProcessorGroupProcessorDatadogTagsList
	DatadogTagsInput() interface{}
	Dedupe() ObservabilityPipelineConfigProcessorGroupProcessorDedupeList
	DedupeInput() interface{}
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EnrichmentTable() ObservabilityPipelineConfigProcessorGroupProcessorEnrichmentTableList
	EnrichmentTableInput() interface{}
	Filter() ObservabilityPipelineConfigProcessorGroupProcessorFilterList
	FilterInput() interface{}
	// Experimental.
	Fqn() *string
	GenerateDatadogMetrics() ObservabilityPipelineConfigProcessorGroupProcessorGenerateDatadogMetricsList
	GenerateDatadogMetricsInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Include() *string
	SetInclude(val *string)
	IncludeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MetricTags() ObservabilityPipelineConfigProcessorGroupProcessorMetricTagsList
	MetricTagsInput() interface{}
	OcsfMapper() ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperList
	OcsfMapperInput() interface{}
	ParseGrok() ObservabilityPipelineConfigProcessorGroupProcessorParseGrokList
	ParseGrokInput() interface{}
	ParseJson() ObservabilityPipelineConfigProcessorGroupProcessorParseJsonList
	ParseJsonInput() interface{}
	ParseXml() ObservabilityPipelineConfigProcessorGroupProcessorParseXmlList
	ParseXmlInput() interface{}
	Quota() ObservabilityPipelineConfigProcessorGroupProcessorQuotaList
	QuotaInput() interface{}
	Reduce() ObservabilityPipelineConfigProcessorGroupProcessorReduceList
	ReduceInput() interface{}
	RemoveFields() ObservabilityPipelineConfigProcessorGroupProcessorRemoveFieldsList
	RemoveFieldsInput() interface{}
	RenameFields() ObservabilityPipelineConfigProcessorGroupProcessorRenameFieldsList
	RenameFieldsInput() interface{}
	Sample() ObservabilityPipelineConfigProcessorGroupProcessorSampleList
	SampleInput() interface{}
	SensitiveDataScanner() ObservabilityPipelineConfigProcessorGroupProcessorSensitiveDataScannerList
	SensitiveDataScannerInput() interface{}
	SplitArray() ObservabilityPipelineConfigProcessorGroupProcessorSplitArrayList
	SplitArrayInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Throttle() ObservabilityPipelineConfigProcessorGroupProcessorThrottleList
	ThrottleInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutAddEnvVars(value interface{})
	PutAddFields(value interface{})
	PutAddHostname(value interface{})
	PutCustomProcessor(value interface{})
	PutDatadogTags(value interface{})
	PutDedupe(value interface{})
	PutEnrichmentTable(value interface{})
	PutFilter(value interface{})
	PutGenerateDatadogMetrics(value interface{})
	PutMetricTags(value interface{})
	PutOcsfMapper(value interface{})
	PutParseGrok(value interface{})
	PutParseJson(value interface{})
	PutParseXml(value interface{})
	PutQuota(value interface{})
	PutReduce(value interface{})
	PutRemoveFields(value interface{})
	PutRenameFields(value interface{})
	PutSample(value interface{})
	PutSensitiveDataScanner(value interface{})
	PutSplitArray(value interface{})
	PutThrottle(value interface{})
	ResetAddEnvVars()
	ResetAddFields()
	ResetAddHostname()
	ResetCustomProcessor()
	ResetDatadogTags()
	ResetDedupe()
	ResetDisplayName()
	ResetEnrichmentTable()
	ResetFilter()
	ResetGenerateDatadogMetrics()
	ResetMetricTags()
	ResetOcsfMapper()
	ResetParseGrok()
	ResetParseJson()
	ResetParseXml()
	ResetQuota()
	ResetReduce()
	ResetRemoveFields()
	ResetRenameFields()
	ResetSample()
	ResetSensitiveDataScanner()
	ResetSplitArray()
	ResetThrottle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityPipelineConfigProcessorGroupProcessorOutputReference
type jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) AddEnvVars() ObservabilityPipelineConfigProcessorGroupProcessorAddEnvVarsList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorAddEnvVarsList
	_jsii_.Get(
		j,
		"addEnvVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) AddEnvVarsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addEnvVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) AddFields() ObservabilityPipelineConfigProcessorGroupProcessorAddFieldsList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorAddFieldsList
	_jsii_.Get(
		j,
		"addFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) AddFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) AddHostname() ObservabilityPipelineConfigProcessorGroupProcessorAddHostnameList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorAddHostnameList
	_jsii_.Get(
		j,
		"addHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) AddHostnameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) CustomProcessor() ObservabilityPipelineConfigProcessorGroupProcessorCustomProcessorList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorCustomProcessorList
	_jsii_.Get(
		j,
		"customProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) CustomProcessorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) DatadogTags() ObservabilityPipelineConfigProcessorGroupProcessorDatadogTagsList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorDatadogTagsList
	_jsii_.Get(
		j,
		"datadogTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) DatadogTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datadogTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) Dedupe() ObservabilityPipelineConfigProcessorGroupProcessorDedupeList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorDedupeList
	_jsii_.Get(
		j,
		"dedupe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) DedupeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedupeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) EnrichmentTable() ObservabilityPipelineConfigProcessorGroupProcessorEnrichmentTableList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorEnrichmentTableList
	_jsii_.Get(
		j,
		"enrichmentTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) EnrichmentTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrichmentTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) Filter() ObservabilityPipelineConfigProcessorGroupProcessorFilterList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) GenerateDatadogMetrics() ObservabilityPipelineConfigProcessorGroupProcessorGenerateDatadogMetricsList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorGenerateDatadogMetricsList
	_jsii_.Get(
		j,
		"generateDatadogMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) GenerateDatadogMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateDatadogMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) Include() *string {
	var returns *string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) IncludeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) MetricTags() ObservabilityPipelineConfigProcessorGroupProcessorMetricTagsList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorMetricTagsList
	_jsii_.Get(
		j,
		"metricTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) MetricTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) OcsfMapper() ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperList
	_jsii_.Get(
		j,
		"ocsfMapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) OcsfMapperInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocsfMapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ParseGrok() ObservabilityPipelineConfigProcessorGroupProcessorParseGrokList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorParseGrokList
	_jsii_.Get(
		j,
		"parseGrok",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ParseGrokInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseGrokInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ParseJson() ObservabilityPipelineConfigProcessorGroupProcessorParseJsonList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorParseJsonList
	_jsii_.Get(
		j,
		"parseJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ParseJsonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ParseXml() ObservabilityPipelineConfigProcessorGroupProcessorParseXmlList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorParseXmlList
	_jsii_.Get(
		j,
		"parseXml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ParseXmlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseXmlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) Quota() ObservabilityPipelineConfigProcessorGroupProcessorQuotaList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorQuotaList
	_jsii_.Get(
		j,
		"quota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) QuotaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) Reduce() ObservabilityPipelineConfigProcessorGroupProcessorReduceList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorReduceList
	_jsii_.Get(
		j,
		"reduce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ReduceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reduceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) RemoveFields() ObservabilityPipelineConfigProcessorGroupProcessorRemoveFieldsList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorRemoveFieldsList
	_jsii_.Get(
		j,
		"removeFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) RemoveFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) RenameFields() ObservabilityPipelineConfigProcessorGroupProcessorRenameFieldsList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorRenameFieldsList
	_jsii_.Get(
		j,
		"renameFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) RenameFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renameFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) Sample() ObservabilityPipelineConfigProcessorGroupProcessorSampleList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorSampleList
	_jsii_.Get(
		j,
		"sample",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) SampleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sampleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) SensitiveDataScanner() ObservabilityPipelineConfigProcessorGroupProcessorSensitiveDataScannerList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorSensitiveDataScannerList
	_jsii_.Get(
		j,
		"sensitiveDataScanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) SensitiveDataScannerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sensitiveDataScannerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) SplitArray() ObservabilityPipelineConfigProcessorGroupProcessorSplitArrayList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorSplitArrayList
	_jsii_.Get(
		j,
		"splitArray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) SplitArrayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitArrayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) Throttle() ObservabilityPipelineConfigProcessorGroupProcessorThrottleList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorThrottleList
	_jsii_.Get(
		j,
		"throttle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ThrottleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttleInput",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigProcessorGroupProcessorOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityPipelineConfigProcessorGroupProcessorOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigProcessorGroupProcessorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorGroupProcessorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigProcessorGroupProcessorOutputReference_Override(o ObservabilityPipelineConfigProcessorGroupProcessorOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorGroupProcessorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference)SetInclude(val *string) {
	if err := j.validateSetIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutAddEnvVars(value interface{}) {
	if err := o.validatePutAddEnvVarsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAddEnvVars",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutAddFields(value interface{}) {
	if err := o.validatePutAddFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAddFields",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutAddHostname(value interface{}) {
	if err := o.validatePutAddHostnameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAddHostname",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutCustomProcessor(value interface{}) {
	if err := o.validatePutCustomProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCustomProcessor",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutDatadogTags(value interface{}) {
	if err := o.validatePutDatadogTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDatadogTags",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutDedupe(value interface{}) {
	if err := o.validatePutDedupeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDedupe",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutEnrichmentTable(value interface{}) {
	if err := o.validatePutEnrichmentTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEnrichmentTable",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutFilter(value interface{}) {
	if err := o.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putFilter",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutGenerateDatadogMetrics(value interface{}) {
	if err := o.validatePutGenerateDatadogMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGenerateDatadogMetrics",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutMetricTags(value interface{}) {
	if err := o.validatePutMetricTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMetricTags",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutOcsfMapper(value interface{}) {
	if err := o.validatePutOcsfMapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOcsfMapper",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutParseGrok(value interface{}) {
	if err := o.validatePutParseGrokParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putParseGrok",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutParseJson(value interface{}) {
	if err := o.validatePutParseJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putParseJson",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutParseXml(value interface{}) {
	if err := o.validatePutParseXmlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putParseXml",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutQuota(value interface{}) {
	if err := o.validatePutQuotaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putQuota",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutReduce(value interface{}) {
	if err := o.validatePutReduceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putReduce",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutRemoveFields(value interface{}) {
	if err := o.validatePutRemoveFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRemoveFields",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutRenameFields(value interface{}) {
	if err := o.validatePutRenameFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRenameFields",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutSample(value interface{}) {
	if err := o.validatePutSampleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSample",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutSensitiveDataScanner(value interface{}) {
	if err := o.validatePutSensitiveDataScannerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSensitiveDataScanner",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutSplitArray(value interface{}) {
	if err := o.validatePutSplitArrayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSplitArray",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) PutThrottle(value interface{}) {
	if err := o.validatePutThrottleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putThrottle",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetAddEnvVars() {
	_jsii_.InvokeVoid(
		o,
		"resetAddEnvVars",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetAddFields() {
	_jsii_.InvokeVoid(
		o,
		"resetAddFields",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetAddHostname() {
	_jsii_.InvokeVoid(
		o,
		"resetAddHostname",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetCustomProcessor() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomProcessor",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetDatadogTags() {
	_jsii_.InvokeVoid(
		o,
		"resetDatadogTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetDedupe() {
	_jsii_.InvokeVoid(
		o,
		"resetDedupe",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		o,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetEnrichmentTable() {
	_jsii_.InvokeVoid(
		o,
		"resetEnrichmentTable",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		o,
		"resetFilter",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetGenerateDatadogMetrics() {
	_jsii_.InvokeVoid(
		o,
		"resetGenerateDatadogMetrics",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetMetricTags() {
	_jsii_.InvokeVoid(
		o,
		"resetMetricTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetOcsfMapper() {
	_jsii_.InvokeVoid(
		o,
		"resetOcsfMapper",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetParseGrok() {
	_jsii_.InvokeVoid(
		o,
		"resetParseGrok",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetParseJson() {
	_jsii_.InvokeVoid(
		o,
		"resetParseJson",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetParseXml() {
	_jsii_.InvokeVoid(
		o,
		"resetParseXml",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetQuota() {
	_jsii_.InvokeVoid(
		o,
		"resetQuota",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetReduce() {
	_jsii_.InvokeVoid(
		o,
		"resetReduce",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetRemoveFields() {
	_jsii_.InvokeVoid(
		o,
		"resetRemoveFields",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetRenameFields() {
	_jsii_.InvokeVoid(
		o,
		"resetRenameFields",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetSample() {
	_jsii_.InvokeVoid(
		o,
		"resetSample",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetSensitiveDataScanner() {
	_jsii_.InvokeVoid(
		o,
		"resetSensitiveDataScanner",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetSplitArray() {
	_jsii_.InvokeVoid(
		o,
		"resetSplitArray",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ResetThrottle() {
	_jsii_.InvokeVoid(
		o,
		"resetThrottle",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

