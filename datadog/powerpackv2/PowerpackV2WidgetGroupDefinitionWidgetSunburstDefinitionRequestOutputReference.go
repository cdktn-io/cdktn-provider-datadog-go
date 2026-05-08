// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference interface {
	cdktn.ComplexObject
	ApmQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQueryOutputReference
	ApmQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQuery
	AuditQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQueryOutputReference
	AuditQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQuery
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
	Formula() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestFormulaList
	FormulaInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQueryOutputReference
	LogQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQuery
	NetworkQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQueryOutputReference
	NetworkQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQuery
	ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQueryOutputReference
	ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	Query() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestQueryList
	QueryInput() interface{}
	RumQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQueryOutputReference
	RumQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQuery
	SecurityQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQueryOutputReference
	SecurityQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQuery
	Sort() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSortOutputReference
	SortInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSort
	Style() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyleOutputReference
	StyleInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyle
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
	PutApmQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQuery)
	PutAuditQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQuery)
	PutFormula(value interface{})
	PutLogQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQuery)
	PutNetworkQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQuery)
	PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQuery)
	PutQuery(value interface{})
	PutRumQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQuery)
	PutSecurityQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQuery)
	PutSort(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSort)
	PutStyle(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyle)
	ResetApmQuery()
	ResetAuditQuery()
	ResetFormula()
	ResetLogQuery()
	ResetNetworkQuery()
	ResetProcessQuery()
	ResetQ()
	ResetQuery()
	ResetRumQuery()
	ResetSecurityQuery()
	ResetSort()
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

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ApmQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ApmQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) AuditQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQueryOutputReference
	_jsii_.Get(
		j,
		"auditQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) AuditQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQuery
	_jsii_.Get(
		j,
		"auditQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Formula() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestFormulaList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestFormulaList
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) FormulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) LogQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) LogQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) NetworkQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQueryOutputReference
	_jsii_.Get(
		j,
		"networkQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) NetworkQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQuery
	_jsii_.Get(
		j,
		"networkQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Query() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestQueryList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestQueryList
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) RumQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) RumQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) SecurityQuery() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) SecurityQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Sort() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSortOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSortOutputReference
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) SortInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSort {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSort
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Style() PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyleOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) StyleInput() *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyle {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutApmQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQuery) {
	if err := p.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutAuditQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQuery) {
	if err := p.validatePutAuditQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAuditQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutFormula(value interface{}) {
	if err := p.validatePutFormulaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFormula",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutLogQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQuery) {
	if err := p.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutNetworkQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQuery) {
	if err := p.validatePutNetworkQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNetworkQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutQuery(value interface{}) {
	if err := p.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutRumQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQuery) {
	if err := p.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutSecurityQuery(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQuery) {
	if err := p.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutSort(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSort) {
	if err := p.validatePutSortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSort",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutStyle(value *PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyle) {
	if err := p.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putStyle",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetAuditQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetAuditQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetFormula() {
	_jsii_.InvokeVoid(
		p,
		"resetFormula",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetNetworkQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetNetworkQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		p,
		"resetQ",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		p,
		"resetSort",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		p,
		"resetStyle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

