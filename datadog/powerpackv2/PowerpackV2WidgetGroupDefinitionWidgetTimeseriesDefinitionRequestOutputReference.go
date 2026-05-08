// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference interface {
	cdktn.ComplexObject
	ApmQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestApmQueryOutputReference
	ApmQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestApmQuery
	AuditQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestAuditQueryOutputReference
	AuditQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestAuditQuery
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
	DisplayType() *string
	SetDisplayType(val *string)
	DisplayTypeInput() *string
	EventQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestEventQueryOutputReference
	EventQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestEventQuery
	Formula() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestFormulaList
	FormulaInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestLogQueryOutputReference
	LogQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestLogQuery
	Metadata() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestMetadataList
	MetadataInput() interface{}
	NetworkQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestNetworkQueryOutputReference
	NetworkQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestNetworkQuery
	OnRightYaxis() interface{}
	SetOnRightYaxis(val interface{})
	OnRightYaxisInput() interface{}
	ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProcessQueryOutputReference
	ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProcessQuery
	ProfileMetricsQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProfileMetricsQueryOutputReference
	ProfileMetricsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProfileMetricsQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	Query() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryList
	QueryInput() interface{}
	RumQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestRumQueryOutputReference
	RumQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestRumQuery
	SecurityQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestSecurityQueryOutputReference
	SecurityQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestSecurityQuery
	Style() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestStyleOutputReference
	StyleInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestStyle
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
	PutApmQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestApmQuery)
	PutAuditQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestAuditQuery)
	PutEventQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestEventQuery)
	PutFormula(value interface{})
	PutLogQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestLogQuery)
	PutMetadata(value interface{})
	PutNetworkQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestNetworkQuery)
	PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProcessQuery)
	PutProfileMetricsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProfileMetricsQuery)
	PutQuery(value interface{})
	PutRumQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestRumQuery)
	PutSecurityQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestSecurityQuery)
	PutStyle(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestStyle)
	ResetApmQuery()
	ResetAuditQuery()
	ResetDisplayType()
	ResetEventQuery()
	ResetFormula()
	ResetLogQuery()
	ResetMetadata()
	ResetNetworkQuery()
	ResetOnRightYaxis()
	ResetProcessQuery()
	ResetProfileMetricsQuery()
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

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ApmQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestApmQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ApmQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestApmQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) AuditQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestAuditQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestAuditQueryOutputReference
	_jsii_.Get(
		j,
		"auditQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) AuditQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestAuditQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestAuditQuery
	_jsii_.Get(
		j,
		"auditQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) DisplayType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) DisplayTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) EventQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestEventQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) EventQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestEventQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) Formula() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestFormulaList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestFormulaList
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) FormulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) LogQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestLogQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) LogQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestLogQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) Metadata() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestMetadataList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestMetadataList
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) NetworkQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestNetworkQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestNetworkQueryOutputReference
	_jsii_.Get(
		j,
		"networkQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) NetworkQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestNetworkQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestNetworkQuery
	_jsii_.Get(
		j,
		"networkQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) OnRightYaxis() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onRightYaxis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) OnRightYaxisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onRightYaxisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ProcessQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProcessQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ProcessQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProcessQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ProfileMetricsQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProfileMetricsQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProfileMetricsQueryOutputReference
	_jsii_.Get(
		j,
		"profileMetricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ProfileMetricsQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProfileMetricsQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProfileMetricsQuery
	_jsii_.Get(
		j,
		"profileMetricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) Query() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestQueryList
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) RumQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestRumQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) RumQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestRumQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) SecurityQuery() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestSecurityQueryOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) SecurityQueryInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestSecurityQuery {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) Style() PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestStyleOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) StyleInput() *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestStyle {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference)SetDisplayType(val *string) {
	if err := j.validateSetDisplayTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayType",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference)SetOnRightYaxis(val interface{}) {
	if err := j.validateSetOnRightYaxisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onRightYaxis",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) PutApmQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestApmQuery) {
	if err := p.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) PutAuditQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestAuditQuery) {
	if err := p.validatePutAuditQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAuditQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) PutEventQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestEventQuery) {
	if err := p.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) PutFormula(value interface{}) {
	if err := p.validatePutFormulaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFormula",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) PutLogQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestLogQuery) {
	if err := p.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) PutMetadata(value interface{}) {
	if err := p.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMetadata",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) PutNetworkQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestNetworkQuery) {
	if err := p.validatePutNetworkQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNetworkQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) PutProcessQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProcessQuery) {
	if err := p.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) PutProfileMetricsQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestProfileMetricsQuery) {
	if err := p.validatePutProfileMetricsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProfileMetricsQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) PutQuery(value interface{}) {
	if err := p.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) PutRumQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestRumQuery) {
	if err := p.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) PutSecurityQuery(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestSecurityQuery) {
	if err := p.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) PutStyle(value *PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestStyle) {
	if err := p.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putStyle",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetAuditQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetAuditQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetDisplayType() {
	_jsii_.InvokeVoid(
		p,
		"resetDisplayType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetFormula() {
	_jsii_.InvokeVoid(
		p,
		"resetFormula",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		p,
		"resetMetadata",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetNetworkQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetNetworkQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetOnRightYaxis() {
	_jsii_.InvokeVoid(
		p,
		"resetOnRightYaxis",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetProfileMetricsQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetProfileMetricsQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		p,
		"resetQ",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		p,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		p,
		"resetStyle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetTimeseriesDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

