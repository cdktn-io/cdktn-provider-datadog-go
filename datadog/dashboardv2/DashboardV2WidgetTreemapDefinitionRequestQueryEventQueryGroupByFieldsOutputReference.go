// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference interface {
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
	Fields() *[]*string
	SetFields(val *[]*string)
	FieldsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFields
	SetInternalValue(val *DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFields)
	Limit() *float64
	SetLimit(val *float64)
	LimitInput() *float64
	Sort() DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference
	SortInput() *DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsSort
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
	PutSort(value *DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsSort)
	ResetLimit()
	ResetSort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference
type jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) Fields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) FieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) InternalValue() *DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFields {
	var returns *DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFields
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) LimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) Sort() DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference {
	var returns DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) SortInput() *DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsSort {
	var returns *DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsSort
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference_Override(d DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetFields(val *[]*string) {
	if err := j.validateSetFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fields",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetInternalValue(val *DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFields) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetLimit(val *float64) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) PutSort(value *DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsSort) {
	if err := d.validatePutSortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSort",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		d,
		"resetSort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetTreemapDefinitionRequestQueryEventQueryGroupByFieldsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

