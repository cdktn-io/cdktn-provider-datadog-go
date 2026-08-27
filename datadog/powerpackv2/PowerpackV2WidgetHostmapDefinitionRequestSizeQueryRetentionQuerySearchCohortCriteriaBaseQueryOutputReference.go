// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference interface {
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
	InternalValue() *PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQuery
	SetInternalValue(val *PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQuery)
	Search() PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQuerySearchOutputReference
	SearchInput() *PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQuerySearch
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
	PutSearch(value *PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQuerySearch)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference
type jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) InternalValue() *PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQuery {
	var returns *PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) Search() PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQuerySearchOutputReference {
	var returns PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQuerySearchOutputReference
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) SearchInput() *PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQuerySearch {
	var returns *PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQuerySearch
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference_Override(p PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference)SetInternalValue(val *PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) PutSearch(value *PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQuerySearch) {
	if err := p.validatePutSearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSearch",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetHostmapDefinitionRequestSizeQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

