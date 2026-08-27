// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference interface {
	cdktn.ComplexObject
	AudienceFilters() DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference
	AudienceFiltersInput() *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersAudienceFilters
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
	GraphFilter() DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersGraphFilterList
	GraphFilterInput() interface{}
	InternalValue() *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFilters
	SetInternalValue(val *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFilters)
	StringFilter() *string
	SetStringFilter(val *string)
	StringFilterInput() *string
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
	PutAudienceFilters(value *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersAudienceFilters)
	PutGraphFilter(value interface{})
	ResetAudienceFilters()
	ResetGraphFilter()
	ResetStringFilter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference
type jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) AudienceFilters() DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference {
	var returns DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference
	_jsii_.Get(
		j,
		"audienceFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) AudienceFiltersInput() *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersAudienceFilters {
	var returns *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersAudienceFilters
	_jsii_.Get(
		j,
		"audienceFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) GraphFilter() DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersGraphFilterList {
	var returns DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersGraphFilterList
	_jsii_.Get(
		j,
		"graphFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) GraphFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"graphFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) InternalValue() *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFilters {
	var returns *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) StringFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) StringFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference_Override(d DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference)SetInternalValue(val *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference)SetStringFilter(val *string) {
	if err := j.validateSetStringFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringFilter",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) PutAudienceFilters(value *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersAudienceFilters) {
	if err := d.validatePutAudienceFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAudienceFilters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) PutGraphFilter(value interface{}) {
	if err := d.validatePutGraphFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGraphFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) ResetAudienceFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetAudienceFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) ResetGraphFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetGraphFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) ResetStringFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetStringFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestQueryUserJourneyQuerySearchFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

