// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference interface {
	cdktn.ComplexObject
	ApmResourceStatsQuery() DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryApmResourceStatsQueryOutputReference
	ApmResourceStatsQueryInput() *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryApmResourceStatsQuery
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
	EventQuery() DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryEventQueryOutputReference
	EventQueryInput() *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryEventQuery
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQuery
	SetInternalValue(val *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQuery)
	MetricQuery() DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryMetricQueryOutputReference
	MetricQueryInput() *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryMetricQuery
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
	PutApmResourceStatsQuery(value *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryApmResourceStatsQuery)
	PutEventQuery(value *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryEventQuery)
	PutMetricQuery(value *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryMetricQuery)
	ResetApmResourceStatsQuery()
	ResetEventQuery()
	ResetMetricQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference
type jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) ApmResourceStatsQuery() DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryApmResourceStatsQueryOutputReference {
	var returns DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryApmResourceStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmResourceStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) ApmResourceStatsQueryInput() *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryApmResourceStatsQuery {
	var returns *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"apmResourceStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) EventQuery() DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryEventQueryOutputReference {
	var returns DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) EventQueryInput() *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryEventQuery {
	var returns *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) InternalValue() *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQuery {
	var returns *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) MetricQuery() DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryMetricQueryOutputReference {
	var returns DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryMetricQueryOutputReference
	_jsii_.Get(
		j,
		"metricQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) MetricQueryInput() *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryMetricQuery {
	var returns *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryMetricQuery
	_jsii_.Get(
		j,
		"metricQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference_Override(d DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference)SetInternalValue(val *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) PutApmResourceStatsQuery(value *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryApmResourceStatsQuery) {
	if err := d.validatePutApmResourceStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmResourceStatsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) PutEventQuery(value *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryEventQuery) {
	if err := d.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) PutMetricQuery(value *DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryMetricQuery) {
	if err := d.validatePutMetricQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetricQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) ResetApmResourceStatsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmResourceStatsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) ResetMetricQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetMetricQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestHistogramRequestHistogramQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

