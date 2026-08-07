// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboard/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference interface {
	cdktn.ComplexObject
	Aggregation() *string
	SetAggregation(val *string)
	AggregationInput() *string
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
	InternalValue() *DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSort
	SetInternalValue(val *DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSort)
	Metric() *string
	SetMetric(val *string)
	MetricInput() *string
	Order() *string
	SetOrder(val *string)
	OrderInput() *string
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
	ResetMetric()
	ResetOrder()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference
type jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) Aggregation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) AggregationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) InternalValue() *DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSort {
	var returns *DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSort
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) Metric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) MetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) Order() *string {
	var returns *string
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) OrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboard.DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference_Override(d DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboard.DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference)SetAggregation(val *string) {
	if err := j.validateSetAggregationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregation",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference)SetInternalValue(val *DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSort) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference)SetMetric(val *string) {
	if err := j.validateSetMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metric",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference)SetOrder(val *string) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		d,
		"resetMetric",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		d,
		"resetOrder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardWidgetChangeDefinitionRequestQueryEventQueryGroupByFieldsSortOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

