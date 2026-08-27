// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference interface {
	cdktn.ComplexObject
	BaseQuery() DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaBaseQueryOutputReference
	BaseQueryInput() *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaBaseQuery
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
	InternalValue() *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteria
	SetInternalValue(val *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteria)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeInterval() DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaTimeIntervalOutputReference
	TimeIntervalInput() *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaTimeInterval
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
	PutBaseQuery(value *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaBaseQuery)
	PutTimeInterval(value *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaTimeInterval)
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

// The jsii proxy struct for DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference
type jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) BaseQuery() DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaBaseQueryOutputReference {
	var returns DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaBaseQueryOutputReference
	_jsii_.Get(
		j,
		"baseQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) BaseQueryInput() *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaBaseQuery {
	var returns *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaBaseQuery
	_jsii_.Get(
		j,
		"baseQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) InternalValue() *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteria {
	var returns *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) TimeInterval() DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaTimeIntervalOutputReference {
	var returns DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaTimeIntervalOutputReference
	_jsii_.Get(
		j,
		"timeInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) TimeIntervalInput() *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaTimeInterval {
	var returns *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaTimeInterval
	_jsii_.Get(
		j,
		"timeIntervalInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference_Override(d DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference)SetInternalValue(val *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteria) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) PutBaseQuery(value *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaBaseQuery) {
	if err := d.validatePutBaseQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBaseQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) PutTimeInterval(value *DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaTimeInterval) {
	if err := d.validatePutTimeIntervalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeInterval",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) ResetTimeInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTreemapRequestQueryRetentionQuerySearchReturnCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

