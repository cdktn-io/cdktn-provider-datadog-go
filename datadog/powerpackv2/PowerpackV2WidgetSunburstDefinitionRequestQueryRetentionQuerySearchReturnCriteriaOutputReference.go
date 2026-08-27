// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference interface {
	cdktn.ComplexObject
	BaseQuery() PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaBaseQueryOutputReference
	BaseQueryInput() *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaBaseQuery
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
	InternalValue() *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteria
	SetInternalValue(val *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteria)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeInterval() PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaTimeIntervalOutputReference
	TimeIntervalInput() *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaTimeInterval
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
	PutBaseQuery(value *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaBaseQuery)
	PutTimeInterval(value *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaTimeInterval)
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

// The jsii proxy struct for PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference
type jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) BaseQuery() PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaBaseQueryOutputReference {
	var returns PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaBaseQueryOutputReference
	_jsii_.Get(
		j,
		"baseQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) BaseQueryInput() *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaBaseQuery {
	var returns *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaBaseQuery
	_jsii_.Get(
		j,
		"baseQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) InternalValue() *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteria {
	var returns *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) TimeInterval() PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaTimeIntervalOutputReference {
	var returns PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaTimeIntervalOutputReference
	_jsii_.Get(
		j,
		"timeInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) TimeIntervalInput() *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaTimeInterval {
	var returns *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaTimeInterval
	_jsii_.Get(
		j,
		"timeIntervalInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference_Override(p PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference)SetInternalValue(val *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteria) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) PutBaseQuery(value *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaBaseQuery) {
	if err := p.validatePutBaseQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putBaseQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) PutTimeInterval(value *PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaTimeInterval) {
	if err := p.validatePutTimeIntervalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeInterval",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) ResetTimeInterval() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeInterval",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetSunburstDefinitionRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

