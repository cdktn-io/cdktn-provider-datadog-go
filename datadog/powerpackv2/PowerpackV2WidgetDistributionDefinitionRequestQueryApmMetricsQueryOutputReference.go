// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference interface {
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
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	// Experimental.
	Fqn() *string
	GroupBy() *[]*string
	SetGroupBy(val *[]*string)
	GroupByInput() *[]*string
	InternalValue() *PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQuery
	SetInternalValue(val *PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQuery)
	Name() *string
	SetName(val *string)
	NameInput() *string
	OperationMode() *string
	SetOperationMode(val *string)
	OperationModeInput() *string
	OperationName() *string
	SetOperationName(val *string)
	OperationNameInput() *string
	PeerTags() *[]*string
	SetPeerTags(val *[]*string)
	PeerTagsInput() *[]*string
	QueryFilter() *string
	SetQueryFilter(val *string)
	QueryFilterInput() *string
	ResourceHash() *string
	SetResourceHash(val *string)
	ResourceHashInput() *string
	ResourceName() *string
	SetResourceName(val *string)
	ResourceNameInput() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	SpanKind() *string
	SetSpanKind(val *string)
	SpanKindInput() *string
	Stat() *string
	SetStat(val *string)
	StatInput() *string
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
	ResetGroupBy()
	ResetOperationMode()
	ResetOperationName()
	ResetPeerTags()
	ResetQueryFilter()
	ResetResourceHash()
	ResetResourceName()
	ResetService()
	ResetSpanKind()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference
type jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) GroupBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) GroupByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) InternalValue() *PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQuery {
	var returns *PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) OperationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) OperationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) OperationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) PeerTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) PeerTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) QueryFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) QueryFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ResourceHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ResourceHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) SpanKind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spanKind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) SpanKindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spanKindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) Stat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) StatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference_Override(p PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetGroupBy(val *[]*string) {
	if err := j.validateSetGroupByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupBy",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetInternalValue(val *PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetOperationMode(val *string) {
	if err := j.validateSetOperationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationMode",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetOperationName(val *string) {
	if err := j.validateSetOperationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetPeerTags(val *[]*string) {
	if err := j.validateSetPeerTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerTags",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetQueryFilter(val *string) {
	if err := j.validateSetQueryFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryFilter",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetResourceHash(val *string) {
	if err := j.validateSetResourceHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceHash",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetResourceName(val *string) {
	if err := j.validateSetResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceName",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetSpanKind(val *string) {
	if err := j.validateSetSpanKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spanKind",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetStat(val *string) {
	if err := j.validateSetStatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stat",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ResetOperationMode() {
	_jsii_.InvokeVoid(
		p,
		"resetOperationMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ResetOperationName() {
	_jsii_.InvokeVoid(
		p,
		"resetOperationName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ResetPeerTags() {
	_jsii_.InvokeVoid(
		p,
		"resetPeerTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ResetQueryFilter() {
	_jsii_.InvokeVoid(
		p,
		"resetQueryFilter",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ResetResourceHash() {
	_jsii_.InvokeVoid(
		p,
		"resetResourceHash",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ResetResourceName() {
	_jsii_.InvokeVoid(
		p,
		"resetResourceName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		p,
		"resetService",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ResetSpanKind() {
	_jsii_.InvokeVoid(
		p,
		"resetSpanKind",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetDistributionDefinitionRequestQueryApmMetricsQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

