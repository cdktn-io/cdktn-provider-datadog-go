// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference interface {
	cdktn.ComplexObject
	AudienceFilters() DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference
	AudienceFiltersInput() *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryAudienceFilters
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
	Compute() DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference
	ComputeInput() *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryCompute
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	// Experimental.
	Fqn() *string
	GroupBy() DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryGroupByList
	GroupByInput() interface{}
	Indexes() *[]*string
	SetIndexes(val *[]*string)
	IndexesInput() *[]*string
	InternalValue() *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQuery
	SetInternalValue(val *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQuery)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Query() DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryQueryOutputReference
	QueryInput() *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryQuery
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
	PutAudienceFilters(value *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryAudienceFilters)
	PutCompute(value *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryCompute)
	PutGroupBy(value interface{})
	PutQuery(value *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryQuery)
	ResetAudienceFilters()
	ResetGroupBy()
	ResetIndexes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference
type jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) AudienceFilters() DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference
	_jsii_.Get(
		j,
		"audienceFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) AudienceFiltersInput() *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryAudienceFilters {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryAudienceFilters
	_jsii_.Get(
		j,
		"audienceFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) Compute() DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryComputeOutputReference
	_jsii_.Get(
		j,
		"compute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) ComputeInput() *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryCompute {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryCompute
	_jsii_.Get(
		j,
		"computeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) GroupBy() DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryGroupByList {
	var returns DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryGroupByList
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) GroupByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) Indexes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) IndexesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) InternalValue() *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQuery {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) Query() DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryQueryOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryQueryOutputReference
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) QueryInput() *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryQuery {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryQuery
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference_Override(d DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference)SetIndexes(val *[]*string) {
	if err := j.validateSetIndexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexes",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference)SetInternalValue(val *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) PutAudienceFilters(value *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryAudienceFilters) {
	if err := d.validatePutAudienceFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAudienceFilters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) PutCompute(value *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryCompute) {
	if err := d.validatePutComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCompute",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) PutGroupBy(value interface{}) {
	if err := d.validatePutGroupByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGroupBy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) PutQuery(value *DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryQuery) {
	if err := d.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) ResetAudienceFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetAudienceFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) ResetIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestXQueryProductAnalyticsExtendedQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

