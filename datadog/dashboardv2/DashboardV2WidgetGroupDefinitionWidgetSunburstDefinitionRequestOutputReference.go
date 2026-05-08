// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference interface {
	cdktn.ComplexObject
	ApmQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQueryOutputReference
	ApmQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQuery
	AuditQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQueryOutputReference
	AuditQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQuery
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
	Formula() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestFormulaList
	FormulaInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQueryOutputReference
	LogQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQuery
	NetworkQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQueryOutputReference
	NetworkQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQuery
	ProcessQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQueryOutputReference
	ProcessQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	Query() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestQueryList
	QueryInput() interface{}
	RumQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQueryOutputReference
	RumQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQuery
	SecurityQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQueryOutputReference
	SecurityQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQuery
	Sort() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSortOutputReference
	SortInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSort
	Style() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyleOutputReference
	StyleInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyle
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
	PutApmQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQuery)
	PutAuditQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQuery)
	PutFormula(value interface{})
	PutLogQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQuery)
	PutNetworkQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQuery)
	PutProcessQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQuery)
	PutQuery(value interface{})
	PutRumQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQuery)
	PutSecurityQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQuery)
	PutSort(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSort)
	PutStyle(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyle)
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

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ApmQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ApmQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) AuditQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQueryOutputReference
	_jsii_.Get(
		j,
		"auditQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) AuditQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQuery
	_jsii_.Get(
		j,
		"auditQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Formula() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestFormulaList {
	var returns DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestFormulaList
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) FormulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) LogQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) LogQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) NetworkQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQueryOutputReference
	_jsii_.Get(
		j,
		"networkQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) NetworkQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQuery
	_jsii_.Get(
		j,
		"networkQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ProcessQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ProcessQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Query() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestQueryList {
	var returns DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestQueryList
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) RumQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) RumQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) SecurityQuery() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) SecurityQueryInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Sort() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSortOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSortOutputReference
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) SortInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSort {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSort
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Style() DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyleOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) StyleInput() *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyle {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutApmQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestApmQuery) {
	if err := d.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutAuditQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestAuditQuery) {
	if err := d.validatePutAuditQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAuditQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutFormula(value interface{}) {
	if err := d.validatePutFormulaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFormula",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutLogQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestLogQuery) {
	if err := d.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutNetworkQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestNetworkQuery) {
	if err := d.validatePutNetworkQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNetworkQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutProcessQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutQuery(value interface{}) {
	if err := d.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutRumQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestRumQuery) {
	if err := d.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutSecurityQuery(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSecurityQuery) {
	if err := d.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutSort(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestSort) {
	if err := d.validatePutSortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSort",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) PutStyle(value *DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestStyle) {
	if err := d.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStyle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetAuditQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetAuditQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetFormula() {
	_jsii_.InvokeVoid(
		d,
		"resetFormula",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetNetworkQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		d,
		"resetQ",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		d,
		"resetSort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		d,
		"resetStyle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSunburstDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

