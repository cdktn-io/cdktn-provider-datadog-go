// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/observabilitypipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference interface {
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
	Contains() *string
	SetContains(val *string)
	ContainsInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EqualsSource() *string
	SetEqualsSource(val *string)
	EqualsSourceInput() *string
	EqualTo() *string
	SetEqualTo(val *string)
	EqualToInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Matches() *string
	SetMatches(val *string)
	MatchesInput() *string
	NotMatches() *string
	SetNotMatches(val *string)
	NotMatchesInput() *string
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
	ResetContains()
	ResetEqualsSource()
	ResetEqualTo()
	ResetMatches()
	ResetNotMatches()
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

// The jsii proxy struct for ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference
type jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) Contains() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) ContainsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) EqualsSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"equalsSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) EqualsSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"equalsSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) EqualTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"equalTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) EqualToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"equalToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) Matches() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) MatchesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) NotMatches() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notMatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) NotMatchesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notMatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference_Override(o ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference)SetContains(val *string) {
	if err := j.validateSetContainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contains",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference)SetEqualsSource(val *string) {
	if err := j.validateSetEqualsSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"equalsSource",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference)SetEqualTo(val *string) {
	if err := j.validateSetEqualToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"equalTo",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference)SetMatches(val *string) {
	if err := j.validateSetMatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matches",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference)SetNotMatches(val *string) {
	if err := j.validateSetNotMatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notMatches",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) ResetContains() {
	_jsii_.InvokeVoid(
		o,
		"resetContains",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) ResetEqualsSource() {
	_jsii_.InvokeVoid(
		o,
		"resetEqualsSource",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) ResetEqualTo() {
	_jsii_.InvokeVoid(
		o,
		"resetEqualTo",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) ResetMatches() {
	_jsii_.InvokeVoid(
		o,
		"resetMatches",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) ResetNotMatches() {
	_jsii_.InvokeVoid(
		o,
		"resetNotMatches",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		o,
		"resetValue",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigProcessorGroupProcessorOcsfMapperMappingCustomMappingMappingLookupTableOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

