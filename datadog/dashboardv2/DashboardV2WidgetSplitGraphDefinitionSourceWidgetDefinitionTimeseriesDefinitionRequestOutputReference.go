// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference interface {
	cdktn.ComplexObject
	ApmQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestApmQueryOutputReference
	ApmQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestApmQuery
	AuditQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestAuditQueryOutputReference
	AuditQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestAuditQuery
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
	EventQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestEventQueryOutputReference
	EventQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestEventQuery
	Formula() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaList
	FormulaInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestLogQueryOutputReference
	LogQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestLogQuery
	Metadata() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestMetadataList
	MetadataInput() interface{}
	NetworkQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestNetworkQueryOutputReference
	NetworkQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestNetworkQuery
	OnRightYaxis() interface{}
	SetOnRightYaxis(val interface{})
	OnRightYaxisInput() interface{}
	ProcessQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProcessQueryOutputReference
	ProcessQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProcessQuery
	ProfileMetricsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProfileMetricsQueryOutputReference
	ProfileMetricsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProfileMetricsQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	Query() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestQueryList
	QueryInput() interface{}
	RumQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestRumQueryOutputReference
	RumQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestRumQuery
	SecurityQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestSecurityQueryOutputReference
	SecurityQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestSecurityQuery
	Style() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestStyleOutputReference
	StyleInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestStyle
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
	PutApmQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestApmQuery)
	PutAuditQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestAuditQuery)
	PutEventQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestEventQuery)
	PutFormula(value interface{})
	PutLogQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestLogQuery)
	PutMetadata(value interface{})
	PutNetworkQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestNetworkQuery)
	PutProcessQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProcessQuery)
	PutProfileMetricsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProfileMetricsQuery)
	PutQuery(value interface{})
	PutRumQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestRumQuery)
	PutSecurityQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestSecurityQuery)
	PutStyle(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestStyle)
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

// The jsii proxy struct for DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference
type jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ApmQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestApmQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ApmQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestApmQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) AuditQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestAuditQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestAuditQueryOutputReference
	_jsii_.Get(
		j,
		"auditQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) AuditQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestAuditQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestAuditQuery
	_jsii_.Get(
		j,
		"auditQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) DisplayType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) DisplayTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) EventQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestEventQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestEventQueryOutputReference
	_jsii_.Get(
		j,
		"eventQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) EventQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestEventQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestEventQuery
	_jsii_.Get(
		j,
		"eventQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) Formula() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaList {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestFormulaList
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) FormulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) LogQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestLogQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) LogQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestLogQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) Metadata() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestMetadataList {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestMetadataList
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) NetworkQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestNetworkQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestNetworkQueryOutputReference
	_jsii_.Get(
		j,
		"networkQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) NetworkQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestNetworkQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestNetworkQuery
	_jsii_.Get(
		j,
		"networkQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) OnRightYaxis() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onRightYaxis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) OnRightYaxisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onRightYaxisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ProcessQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProcessQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ProcessQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProcessQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ProfileMetricsQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProfileMetricsQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProfileMetricsQueryOutputReference
	_jsii_.Get(
		j,
		"profileMetricsQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ProfileMetricsQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProfileMetricsQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProfileMetricsQuery
	_jsii_.Get(
		j,
		"profileMetricsQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) Query() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestQueryList {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestQueryList
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) RumQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestRumQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) RumQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestRumQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) SecurityQuery() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestSecurityQueryOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) SecurityQueryInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestSecurityQuery {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) Style() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestStyleOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) StyleInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestStyle {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference_Override(d DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference)SetDisplayType(val *string) {
	if err := j.validateSetDisplayTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayType",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference)SetOnRightYaxis(val interface{}) {
	if err := j.validateSetOnRightYaxisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onRightYaxis",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) PutApmQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestApmQuery) {
	if err := d.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) PutAuditQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestAuditQuery) {
	if err := d.validatePutAuditQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAuditQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) PutEventQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestEventQuery) {
	if err := d.validatePutEventQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEventQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) PutFormula(value interface{}) {
	if err := d.validatePutFormulaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFormula",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) PutLogQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestLogQuery) {
	if err := d.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) PutMetadata(value interface{}) {
	if err := d.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetadata",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) PutNetworkQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestNetworkQuery) {
	if err := d.validatePutNetworkQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNetworkQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) PutProcessQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) PutProfileMetricsQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProfileMetricsQuery) {
	if err := d.validatePutProfileMetricsQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProfileMetricsQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) PutQuery(value interface{}) {
	if err := d.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) PutRumQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestRumQuery) {
	if err := d.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) PutSecurityQuery(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestSecurityQuery) {
	if err := d.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) PutStyle(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestStyle) {
	if err := d.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStyle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetAuditQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetAuditQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetDisplayType() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetEventQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetEventQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetFormula() {
	_jsii_.InvokeVoid(
		d,
		"resetFormula",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		d,
		"resetMetadata",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetNetworkQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetOnRightYaxis() {
	_jsii_.InvokeVoid(
		d,
		"resetOnRightYaxis",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetProfileMetricsQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProfileMetricsQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		d,
		"resetQ",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		d,
		"resetStyle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

