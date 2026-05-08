// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference interface {
	cdktn.ComplexObject
	AdditionalQueryFilters() *string
	SetAdditionalQueryFilters(val *string)
	AdditionalQueryFiltersInput() *string
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
	GroupMode() *string
	SetGroupMode(val *string)
	GroupModeInput() *string
	InternalValue() *PowerpackV2WidgetGeomapDefinitionRequestQuerySloQuery
	SetInternalValue(val *PowerpackV2WidgetGeomapDefinitionRequestQuerySloQuery)
	Measure() *string
	SetMeasure(val *string)
	MeasureInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	SloId() *string
	SetSloId(val *string)
	SloIdInput() *string
	SloQueryType() *string
	SetSloQueryType(val *string)
	SloQueryTypeInput() *string
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
	ResetAdditionalQueryFilters()
	ResetCrossOrgUuids()
	ResetGroupMode()
	ResetName()
	ResetSloQueryType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference
type jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) AdditionalQueryFilters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalQueryFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) AdditionalQueryFiltersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalQueryFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) CrossOrgUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) CrossOrgUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GroupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GroupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) InternalValue() *PowerpackV2WidgetGeomapDefinitionRequestQuerySloQuery {
	var returns *PowerpackV2WidgetGeomapDefinitionRequestQuerySloQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) Measure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) MeasureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) SloId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) SloIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) SloQueryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloQueryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) SloQueryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sloQueryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference_Override(p PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetAdditionalQueryFilters(val *string) {
	if err := j.validateSetAdditionalQueryFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalQueryFilters",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetCrossOrgUuids(val *[]*string) {
	if err := j.validateSetCrossOrgUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossOrgUuids",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetGroupMode(val *string) {
	if err := j.validateSetGroupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupMode",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetInternalValue(val *PowerpackV2WidgetGeomapDefinitionRequestQuerySloQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetMeasure(val *string) {
	if err := j.validateSetMeasureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measure",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetSloId(val *string) {
	if err := j.validateSetSloIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sloId",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetSloQueryType(val *string) {
	if err := j.validateSetSloQueryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sloQueryType",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ResetAdditionalQueryFilters() {
	_jsii_.InvokeVoid(
		p,
		"resetAdditionalQueryFilters",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ResetCrossOrgUuids() {
	_jsii_.InvokeVoid(
		p,
		"resetCrossOrgUuids",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ResetGroupMode() {
	_jsii_.InvokeVoid(
		p,
		"resetGroupMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ResetSloQueryType() {
	_jsii_.InvokeVoid(
		p,
		"resetSloQueryType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGeomapDefinitionRequestQuerySloQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

