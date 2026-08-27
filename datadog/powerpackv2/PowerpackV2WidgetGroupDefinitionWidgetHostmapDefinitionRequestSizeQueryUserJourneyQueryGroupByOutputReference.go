// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference interface {
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
	Facet() *string
	SetFacet(val *string)
	FacetInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Limit() *float64
	SetLimit(val *float64)
	LimitInput() *float64
	ShouldExcludeMissing() interface{}
	SetShouldExcludeMissing(val interface{})
	ShouldExcludeMissingInput() interface{}
	Sort() PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupBySortOutputReference
	SortInput() *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupBySort
	Target() PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByTargetOutputReference
	TargetInput() *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByTarget
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
	PutSort(value *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupBySort)
	PutTarget(value *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByTarget)
	ResetLimit()
	ResetShouldExcludeMissing()
	ResetSort()
	ResetTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) Facet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"facet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) FacetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"facetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) LimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) ShouldExcludeMissing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldExcludeMissing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) ShouldExcludeMissingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldExcludeMissingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) Sort() PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupBySortOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupBySortOutputReference
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) SortInput() *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupBySort {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupBySort
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) Target() PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByTargetOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByTargetOutputReference
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) TargetInput() *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByTarget {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByTarget
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference)SetFacet(val *string) {
	if err := j.validateSetFacetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference)SetLimit(val *float64) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference)SetShouldExcludeMissing(val interface{}) {
	if err := j.validateSetShouldExcludeMissingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldExcludeMissing",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) PutSort(value *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupBySort) {
	if err := p.validatePutSortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSort",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) PutTarget(value *PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByTarget) {
	if err := p.validatePutTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTarget",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		p,
		"resetLimit",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) ResetShouldExcludeMissing() {
	_jsii_.InvokeVoid(
		p,
		"resetShouldExcludeMissing",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		p,
		"resetSort",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		p,
		"resetTarget",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeQueryUserJourneyQueryGroupByOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

