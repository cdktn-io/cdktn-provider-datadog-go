// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference interface {
	cdktn.ComplexObject
	BaseQuery() PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaBaseQueryOutputReference
	BaseQueryInput() *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaBaseQuery
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
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteria
	SetInternalValue(val *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteria)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeInterval() PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaTimeIntervalOutputReference
	TimeIntervalInput() *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaTimeInterval
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
	PutBaseQuery(value *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaBaseQuery)
	PutTimeInterval(value *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaTimeInterval)
	ResetTimeInterval()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference
type jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) BaseQuery() PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaBaseQueryOutputReference {
	var returns PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaBaseQueryOutputReference
	_jsii_.Get(
		j,
		"baseQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) BaseQueryInput() *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaBaseQuery {
	var returns *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaBaseQuery
	_jsii_.Get(
		j,
		"baseQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) InternalValue() *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteria {
	var returns *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) TimeInterval() PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaTimeIntervalOutputReference {
	var returns PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaTimeIntervalOutputReference
	_jsii_.Get(
		j,
		"timeInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) TimeIntervalInput() *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaTimeInterval {
	var returns *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaTimeInterval
	_jsii_.Get(
		j,
		"timeIntervalInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference_Override(p PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference)SetInternalValue(val *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteria) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) PutBaseQuery(value *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaBaseQuery) {
	if err := p.validatePutBaseQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putBaseQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) PutTimeInterval(value *PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaTimeInterval) {
	if err := p.validatePutTimeIntervalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeInterval",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) ResetTimeInterval() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeInterval",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetRetentionCurveDefinitionRequestQuerySearchReturnCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

