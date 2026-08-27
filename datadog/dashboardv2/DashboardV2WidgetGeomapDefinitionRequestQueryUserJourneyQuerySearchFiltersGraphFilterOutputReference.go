// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	Target() DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterTargetOutputReference
	TargetInput() *DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterTarget
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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
	PutTarget(value *DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterTarget)
	ResetName()
	ResetOperator()
	ResetTarget()
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

// The jsii proxy struct for DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference
type jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) Target() DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterTargetOutputReference {
	var returns DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterTargetOutputReference
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) TargetInput() *DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterTarget {
	var returns *DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterTarget
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference_Override(d DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference)SetValue(val *float64) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) PutTarget(value *DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterTarget) {
	if err := d.validatePutTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTarget",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) ResetOperator() {
	_jsii_.InvokeVoid(
		d,
		"resetOperator",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		d,
		"resetValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestQueryUserJourneyQuerySearchFiltersGraphFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

