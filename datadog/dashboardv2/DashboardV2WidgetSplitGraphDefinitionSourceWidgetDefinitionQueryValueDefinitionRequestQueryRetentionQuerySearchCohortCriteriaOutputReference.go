// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference interface {
	cdktn.ComplexObject
	BaseQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference
	BaseQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaBaseQuery
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
	InternalValue() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteria
	SetInternalValue(val *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteria)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeInterval() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaTimeIntervalOutputReference
	TimeIntervalInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaTimeInterval
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
	PutBaseQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaBaseQuery)
	PutTimeInterval(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaTimeInterval)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference
type jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) BaseQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference
	_jsii_.Get(
		j,
		"baseQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) BaseQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaBaseQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaBaseQuery
	_jsii_.Get(
		j,
		"baseQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) InternalValue() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteria {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) TimeInterval() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaTimeIntervalOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaTimeIntervalOutputReference
	_jsii_.Get(
		j,
		"timeInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) TimeIntervalInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaTimeInterval {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaTimeInterval
	_jsii_.Get(
		j,
		"timeIntervalInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference_Override(d DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference)SetInternalValue(val *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteria) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) PutBaseQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaBaseQuery) {
	if err := d.validatePutBaseQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBaseQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) PutTimeInterval(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaTimeInterval) {
	if err := d.validatePutTimeIntervalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeInterval",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinitionRequestQueryRetentionQuerySearchCohortCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

