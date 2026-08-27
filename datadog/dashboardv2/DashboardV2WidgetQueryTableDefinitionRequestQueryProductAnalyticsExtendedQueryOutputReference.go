// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference interface {
	cdktn.ComplexObject
	AudienceFilters() DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference
	AudienceFiltersInput() *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFilters
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
	Compute() DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference
	ComputeInput() *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryCompute
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
	GroupBy() DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryGroupByList
	GroupByInput() interface{}
	Indexes() *[]*string
	SetIndexes(val *[]*string)
	IndexesInput() *[]*string
	InternalValue() *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQuery
	SetInternalValue(val *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQuery)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Query() DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryQueryOutputReference
	QueryInput() *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryQuery
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
	PutAudienceFilters(value *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFilters)
	PutCompute(value *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryCompute)
	PutGroupBy(value interface{})
	PutQuery(value *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryQuery)
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

// The jsii proxy struct for DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference
type jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) AudienceFilters() DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference {
	var returns DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference
	_jsii_.Get(
		j,
		"audienceFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) AudienceFiltersInput() *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFilters {
	var returns *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFilters
	_jsii_.Get(
		j,
		"audienceFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) Compute() DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference {
	var returns DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryComputeOutputReference
	_jsii_.Get(
		j,
		"compute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) ComputeInput() *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryCompute {
	var returns *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryCompute
	_jsii_.Get(
		j,
		"computeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) GroupBy() DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryGroupByList {
	var returns DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryGroupByList
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) GroupByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) Indexes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) IndexesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) InternalValue() *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQuery {
	var returns *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) Query() DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryQueryOutputReference {
	var returns DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryQueryOutputReference
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) QueryInput() *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryQuery {
	var returns *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryQuery
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference_Override(d DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference)SetIndexes(val *[]*string) {
	if err := j.validateSetIndexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexes",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference)SetInternalValue(val *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) PutAudienceFilters(value *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFilters) {
	if err := d.validatePutAudienceFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAudienceFilters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) PutCompute(value *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryCompute) {
	if err := d.validatePutComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCompute",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) PutGroupBy(value interface{}) {
	if err := d.validatePutGroupByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGroupBy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) PutQuery(value *DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryQuery) {
	if err := d.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) ResetAudienceFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetAudienceFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) ResetIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryProductAnalyticsExtendedQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

