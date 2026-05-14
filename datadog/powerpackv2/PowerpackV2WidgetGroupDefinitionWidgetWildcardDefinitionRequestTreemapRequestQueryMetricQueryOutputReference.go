// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference interface {
	cdktn.ComplexObject
	Aggregator() *string
	SetAggregator(val *string)
	AggregatorInput() *string
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
	CrossOrgUuids() *[]*string
	SetCrossOrgUuids(val *[]*string)
	CrossOrgUuidsInput() *[]*string
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQuery
	SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQuery)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	SemanticMode() *string
	SetSemanticMode(val *string)
	SemanticModeInput() *string
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
	ResetAggregator()
	ResetCrossOrgUuids()
	ResetDataSource()
	ResetSemanticMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) Aggregator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) AggregatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) CrossOrgUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) CrossOrgUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) InternalValue() *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) SemanticMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"semanticMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) SemanticModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"semanticModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference)SetAggregator(val *string) {
	if err := j.validateSetAggregatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregator",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference)SetCrossOrgUuids(val *[]*string) {
	if err := j.validateSetCrossOrgUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossOrgUuids",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference)SetInternalValue(val *PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference)SetSemanticMode(val *string) {
	if err := j.validateSetSemanticModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"semanticMode",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) ResetAggregator() {
	_jsii_.InvokeVoid(
		p,
		"resetAggregator",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) ResetCrossOrgUuids() {
	_jsii_.InvokeVoid(
		p,
		"resetCrossOrgUuids",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) ResetDataSource() {
	_jsii_.InvokeVoid(
		p,
		"resetDataSource",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) ResetSemanticMode() {
	_jsii_.InvokeVoid(
		p,
		"resetSemanticMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTreemapRequestQueryMetricQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

