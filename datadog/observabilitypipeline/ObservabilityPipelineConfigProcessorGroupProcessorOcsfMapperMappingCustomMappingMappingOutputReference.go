// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/observabilitypipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference interface {
	cdktn.ComplexObject
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
	Default() *string
	SetDefault(val *string)
	DefaultInput() *string
	Dest() *string
	SetDest(val *string)
	DestInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lookup() ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupList
	LookupInput() interface{}
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	Sources() *[]*string
	SetSources(val *[]*string)
	SourcesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	PutLookup(value interface{})
	ResetDefault()
	ResetLookup()
	ResetSource()
	ResetSources()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference
type jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) Default() *string {
	var returns *string
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) DefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) Dest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) DestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) Lookup() ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupList {
	var returns ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupList
	_jsii_.Get(
		j,
		"lookup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) LookupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lookupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) Sources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) SourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference_Override(o ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference)SetDefault(val *string) {
	if err := j.validateSetDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"default",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference)SetDest(val *string) {
	if err := j.validateSetDestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dest",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference)SetSources(val *[]*string) {
	if err := j.validateSetSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sources",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) PutLookup(value interface{}) {
	if err := o.validatePutLookupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLookup",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) ResetDefault() {
	_jsii_.InvokeVoid(
		o,
		"resetDefault",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) ResetLookup() {
	_jsii_.InvokeVoid(
		o,
		"resetLookup",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		o,
		"resetSource",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) ResetSources() {
	_jsii_.InvokeVoid(
		o,
		"resetSources",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		o,
		"resetValue",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

