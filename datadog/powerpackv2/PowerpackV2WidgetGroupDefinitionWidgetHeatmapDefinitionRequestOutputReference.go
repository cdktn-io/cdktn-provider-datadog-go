// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference interface {
	cdktn.ComplexObject
	ApmQuery() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryOutputReference
	ApmQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQuery
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
	Formula() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaList
	FormulaInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQueryOutputReference
	LogQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQuery
	ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQueryOutputReference
	ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	Query() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestQueryList
	QueryInput() interface{}
	RumQuery() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQueryOutputReference
	RumQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQuery
	SecurityQuery() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQueryOutputReference
	SecurityQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQuery
	Style() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyleOutputReference
	StyleInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyle
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
	PutApmQuery(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQuery)
	PutFormula(value interface{})
	PutLogQuery(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQuery)
	PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQuery)
	PutQuery(value interface{})
	PutRumQuery(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQuery)
	PutSecurityQuery(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQuery)
	PutStyle(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyle)
	ResetApmQuery()
	ResetFormula()
	ResetLogQuery()
	ResetProcessQuery()
	ResetQ()
	ResetQuery()
	ResetRumQuery()
	ResetSecurityQuery()
	ResetStyle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ApmQuery() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ApmQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) Formula() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaList
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) FormulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) LogQuery() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) LogQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) Query() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestQueryList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestQueryList
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) RumQuery() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) RumQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) SecurityQuery() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) SecurityQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) Style() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyleOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) StyleInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyle {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutApmQuery(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestApmQuery) {
	if err := p.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutFormula(value interface{}) {
	if err := p.validatePutFormulaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFormula",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutLogQuery(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestLogQuery) {
	if err := p.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutQuery(value interface{}) {
	if err := p.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutRumQuery(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestRumQuery) {
	if err := p.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutSecurityQuery(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestSecurityQuery) {
	if err := p.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) PutStyle(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestStyle) {
	if err := p.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putStyle",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetFormula() {
	_jsii_.InvokeVoid(
		p,
		"resetFormula",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		p,
		"resetQ",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		p,
		"resetStyle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

