// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	Metric() *string
	SetMetric(val *string)
	MetricInput() *string
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
	ResetInterval()
	ResetMetric()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference
type jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) Aggregation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) AggregationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) Metric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) MetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference_Override(d DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference)SetAggregation(val *string) {
	if err := j.validateSetAggregationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregation",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference)SetInterval(val *float64) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference)SetMetric(val *string) {
	if err := j.validateSetMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metric",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		d,
		"resetMetric",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinitionRequestYQueryEventQueryComputeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

