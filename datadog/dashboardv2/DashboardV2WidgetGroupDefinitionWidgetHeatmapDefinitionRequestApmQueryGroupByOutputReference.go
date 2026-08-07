// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference interface {
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
	SortQuery() DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupBySortQueryOutputReference
	SortQueryInput() *DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupBySortQuery
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
	PutSortQuery(value *DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupBySortQuery)
	ResetFacet()
	ResetLimit()
	ResetSortQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) Facet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"facet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) FacetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"facetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) LimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) SortQuery() DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupBySortQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupBySortQueryOutputReference
	_jsii_.Get(
		j,
		"sortQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) SortQueryInput() *DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupBySortQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupBySortQuery
	_jsii_.Get(
		j,
		"sortQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference)SetFacet(val *string) {
	if err := j.validateSetFacetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference)SetLimit(val *float64) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) PutSortQuery(value *DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupBySortQuery) {
	if err := d.validatePutSortQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSortQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) ResetFacet() {
	_jsii_.InvokeVoid(
		d,
		"resetFacet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) ResetSortQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSortQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryGroupByOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

