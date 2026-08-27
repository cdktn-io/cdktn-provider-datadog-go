// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference interface {
	cdktn.ComplexObject
	ComparisonSegments() *[]*string
	SetComparisonSegments(val *[]*string)
	ComparisonSegmentsInput() *[]*string
	ComparisonTime() PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestComparisonTimeOutputReference
	ComparisonTimeInput() *PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestComparisonTime
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Query() PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestQueryOutputReference
	QueryInput() *PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestQuery
	RequestType() *string
	SetRequestType(val *string)
	RequestTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutComparisonTime(value *PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestComparisonTime)
	PutQuery(value *PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestQuery)
	ResetComparisonSegments()
	ResetComparisonTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference
type jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) ComparisonSegments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"comparisonSegments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) ComparisonSegmentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"comparisonSegmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) ComparisonTime() PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestComparisonTimeOutputReference {
	var returns PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestComparisonTimeOutputReference
	_jsii_.Get(
		j,
		"comparisonTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) ComparisonTimeInput() *PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestComparisonTime {
	var returns *PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestComparisonTime
	_jsii_.Get(
		j,
		"comparisonTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) Query() PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestQueryOutputReference {
	var returns PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestQueryOutputReference
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) QueryInput() *PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestQuery {
	var returns *PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestQuery
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) RequestType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) RequestTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference_Override(p PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference)SetComparisonSegments(val *[]*string) {
	if err := j.validateSetComparisonSegmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparisonSegments",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference)SetRequestType(val *string) {
	if err := j.validateSetRequestTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestType",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) PutComparisonTime(value *PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestComparisonTime) {
	if err := p.validatePutComparisonTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putComparisonTime",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) PutQuery(value *PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestQuery) {
	if err := p.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) ResetComparisonSegments() {
	_jsii_.InvokeVoid(
		p,
		"resetComparisonSegments",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) ResetComparisonTime() {
	_jsii_.InvokeVoid(
		p,
		"resetComparisonTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetProductAnalyticsFunnelDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

