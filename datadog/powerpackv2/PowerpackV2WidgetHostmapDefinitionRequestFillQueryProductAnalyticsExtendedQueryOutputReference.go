// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference interface {
	cdktn.ComplexObject
	AudienceFilters() PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference
	AudienceFiltersInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryAudienceFilters
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
	Compute() PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryComputeOutputReference
	ComputeInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryCompute
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
	GroupBy() PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryGroupByList
	GroupByInput() interface{}
	Indexes() *[]*string
	SetIndexes(val *[]*string)
	IndexesInput() *[]*string
	InternalValue() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQuery
	SetInternalValue(val *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQuery)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Query() PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryQueryOutputReference
	QueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryQuery
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
	PutAudienceFilters(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryAudienceFilters)
	PutCompute(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryCompute)
	PutGroupBy(value interface{})
	PutQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryQuery)
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

// The jsii proxy struct for PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference
type jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) AudienceFilters() PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference {
	var returns PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference
	_jsii_.Get(
		j,
		"audienceFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) AudienceFiltersInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryAudienceFilters {
	var returns *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryAudienceFilters
	_jsii_.Get(
		j,
		"audienceFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) Compute() PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryComputeOutputReference {
	var returns PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryComputeOutputReference
	_jsii_.Get(
		j,
		"compute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) ComputeInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryCompute {
	var returns *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryCompute
	_jsii_.Get(
		j,
		"computeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) GroupBy() PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryGroupByList {
	var returns PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryGroupByList
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) GroupByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) Indexes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) IndexesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) InternalValue() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQuery {
	var returns *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) Query() PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryQueryOutputReference {
	var returns PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryQueryOutputReference
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) QueryInput() *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryQuery {
	var returns *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryQuery
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference_Override(p PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference)SetIndexes(val *[]*string) {
	if err := j.validateSetIndexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexes",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference)SetInternalValue(val *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) PutAudienceFilters(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryAudienceFilters) {
	if err := p.validatePutAudienceFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAudienceFilters",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) PutCompute(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryCompute) {
	if err := p.validatePutComputeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCompute",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) PutGroupBy(value interface{}) {
	if err := p.validatePutGroupByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGroupBy",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) PutQuery(value *PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryQuery) {
	if err := p.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) ResetAudienceFilters() {
	_jsii_.InvokeVoid(
		p,
		"resetAudienceFilters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) ResetIndexes() {
	_jsii_.InvokeVoid(
		p,
		"resetIndexes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestFillQueryProductAnalyticsExtendedQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

