// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference interface {
	cdktn.ComplexObject
	ApmQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmQueryOutputReference
	ApmQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmQuery
	ApmStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmStatsQueryOutputReference
	ApmStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmStatsQuery
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
	Formula() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestFormulaList
	FormulaInput() interface{}
	// Experimental.
	Fqn() *string
	HistogramQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestHistogramQueryOutputReference
	HistogramQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestHistogramQuery
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestLogQueryOutputReference
	LogQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestLogQuery
	ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestProcessQueryOutputReference
	ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	Query() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryList
	QueryInput() interface{}
	RequestType() *string
	SetRequestType(val *string)
	RequestTypeInput() *string
	RumQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestRumQueryOutputReference
	RumQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestRumQuery
	SecurityQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestSecurityQueryOutputReference
	SecurityQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestSecurityQuery
	Style() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestStyleOutputReference
	StyleInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestStyle
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
	PutApmQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmQuery)
	PutApmStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmStatsQuery)
	PutFormula(value interface{})
	PutHistogramQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestHistogramQuery)
	PutLogQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestLogQuery)
	PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestProcessQuery)
	PutQuery(value interface{})
	PutRumQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestRumQuery)
	PutSecurityQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestSecurityQuery)
	PutStyle(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestStyle)
	ResetApmQuery()
	ResetApmStatsQuery()
	ResetFormula()
	ResetHistogramQuery()
	ResetLogQuery()
	ResetProcessQuery()
	ResetQ()
	ResetQuery()
	ResetRequestType()
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

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ApmQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ApmQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ApmStatsQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmStatsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmStatsQueryOutputReference
	_jsii_.Get(
		j,
		"apmStatsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ApmStatsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmStatsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmStatsQuery
	_jsii_.Get(
		j,
		"apmStatsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) Formula() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestFormulaList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestFormulaList
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) FormulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) HistogramQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestHistogramQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestHistogramQueryOutputReference
	_jsii_.Get(
		j,
		"histogramQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) HistogramQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestHistogramQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestHistogramQuery
	_jsii_.Get(
		j,
		"histogramQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) LogQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestLogQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) LogQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestLogQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestProcessQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestProcessQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) Query() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestQueryList
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) RequestType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) RequestTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) RumQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestRumQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) RumQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestRumQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) SecurityQuery() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestSecurityQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) SecurityQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestSecurityQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) Style() PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestStyleOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) StyleInput() *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestStyle {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference)SetRequestType(val *string) {
	if err := j.validateSetRequestTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestType",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) PutApmQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmQuery) {
	if err := p.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) PutApmStatsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestApmStatsQuery) {
	if err := p.validatePutApmStatsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmStatsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) PutFormula(value interface{}) {
	if err := p.validatePutFormulaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFormula",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) PutHistogramQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestHistogramQuery) {
	if err := p.validatePutHistogramQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHistogramQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) PutLogQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestLogQuery) {
	if err := p.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) PutQuery(value interface{}) {
	if err := p.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) PutRumQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestRumQuery) {
	if err := p.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) PutSecurityQuery(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestSecurityQuery) {
	if err := p.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) PutStyle(value *PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestStyle) {
	if err := p.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putStyle",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ResetApmStatsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmStatsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ResetFormula() {
	_jsii_.InvokeVoid(
		p,
		"resetFormula",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ResetHistogramQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetHistogramQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		p,
		"resetQ",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ResetRequestType() {
	_jsii_.InvokeVoid(
		p,
		"resetRequestType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		p,
		"resetStyle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetDistributionDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

