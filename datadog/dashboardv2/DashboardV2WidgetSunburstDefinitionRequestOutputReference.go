// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetSunburstDefinitionRequestOutputReference interface {
	cdktn.ComplexObject
	ApmQuery() DashboardV2WidgetSunburstDefinitionRequestApmQueryOutputReference
	ApmQueryInput() *DashboardV2WidgetSunburstDefinitionRequestApmQuery
	AuditQuery() DashboardV2WidgetSunburstDefinitionRequestAuditQueryOutputReference
	AuditQueryInput() *DashboardV2WidgetSunburstDefinitionRequestAuditQuery
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
	Formula() DashboardV2WidgetSunburstDefinitionRequestFormulaList
	FormulaInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() DashboardV2WidgetSunburstDefinitionRequestLogQueryOutputReference
	LogQueryInput() *DashboardV2WidgetSunburstDefinitionRequestLogQuery
	NetworkQuery() DashboardV2WidgetSunburstDefinitionRequestNetworkQueryOutputReference
	NetworkQueryInput() *DashboardV2WidgetSunburstDefinitionRequestNetworkQuery
	ProcessQuery() DashboardV2WidgetSunburstDefinitionRequestProcessQueryOutputReference
	ProcessQueryInput() *DashboardV2WidgetSunburstDefinitionRequestProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	Query() DashboardV2WidgetSunburstDefinitionRequestQueryList
	QueryInput() interface{}
	RumQuery() DashboardV2WidgetSunburstDefinitionRequestRumQueryOutputReference
	RumQueryInput() *DashboardV2WidgetSunburstDefinitionRequestRumQuery
	SecurityQuery() DashboardV2WidgetSunburstDefinitionRequestSecurityQueryOutputReference
	SecurityQueryInput() *DashboardV2WidgetSunburstDefinitionRequestSecurityQuery
	Sort() DashboardV2WidgetSunburstDefinitionRequestSortOutputReference
	SortInput() *DashboardV2WidgetSunburstDefinitionRequestSort
	Style() DashboardV2WidgetSunburstDefinitionRequestStyleOutputReference
	StyleInput() *DashboardV2WidgetSunburstDefinitionRequestStyle
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
	PutApmQuery(value *DashboardV2WidgetSunburstDefinitionRequestApmQuery)
	PutAuditQuery(value *DashboardV2WidgetSunburstDefinitionRequestAuditQuery)
	PutFormula(value interface{})
	PutLogQuery(value *DashboardV2WidgetSunburstDefinitionRequestLogQuery)
	PutNetworkQuery(value *DashboardV2WidgetSunburstDefinitionRequestNetworkQuery)
	PutProcessQuery(value *DashboardV2WidgetSunburstDefinitionRequestProcessQuery)
	PutQuery(value interface{})
	PutRumQuery(value *DashboardV2WidgetSunburstDefinitionRequestRumQuery)
	PutSecurityQuery(value *DashboardV2WidgetSunburstDefinitionRequestSecurityQuery)
	PutSort(value *DashboardV2WidgetSunburstDefinitionRequestSort)
	PutStyle(value *DashboardV2WidgetSunburstDefinitionRequestStyle)
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

// The jsii proxy struct for DashboardV2WidgetSunburstDefinitionRequestOutputReference
type jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ApmQuery() DashboardV2WidgetSunburstDefinitionRequestApmQueryOutputReference {
	var returns DashboardV2WidgetSunburstDefinitionRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ApmQueryInput() *DashboardV2WidgetSunburstDefinitionRequestApmQuery {
	var returns *DashboardV2WidgetSunburstDefinitionRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) AuditQuery() DashboardV2WidgetSunburstDefinitionRequestAuditQueryOutputReference {
	var returns DashboardV2WidgetSunburstDefinitionRequestAuditQueryOutputReference
	_jsii_.Get(
		j,
		"auditQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) AuditQueryInput() *DashboardV2WidgetSunburstDefinitionRequestAuditQuery {
	var returns *DashboardV2WidgetSunburstDefinitionRequestAuditQuery
	_jsii_.Get(
		j,
		"auditQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) Formula() DashboardV2WidgetSunburstDefinitionRequestFormulaList {
	var returns DashboardV2WidgetSunburstDefinitionRequestFormulaList
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) FormulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) LogQuery() DashboardV2WidgetSunburstDefinitionRequestLogQueryOutputReference {
	var returns DashboardV2WidgetSunburstDefinitionRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) LogQueryInput() *DashboardV2WidgetSunburstDefinitionRequestLogQuery {
	var returns *DashboardV2WidgetSunburstDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) NetworkQuery() DashboardV2WidgetSunburstDefinitionRequestNetworkQueryOutputReference {
	var returns DashboardV2WidgetSunburstDefinitionRequestNetworkQueryOutputReference
	_jsii_.Get(
		j,
		"networkQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) NetworkQueryInput() *DashboardV2WidgetSunburstDefinitionRequestNetworkQuery {
	var returns *DashboardV2WidgetSunburstDefinitionRequestNetworkQuery
	_jsii_.Get(
		j,
		"networkQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ProcessQuery() DashboardV2WidgetSunburstDefinitionRequestProcessQueryOutputReference {
	var returns DashboardV2WidgetSunburstDefinitionRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ProcessQueryInput() *DashboardV2WidgetSunburstDefinitionRequestProcessQuery {
	var returns *DashboardV2WidgetSunburstDefinitionRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) Query() DashboardV2WidgetSunburstDefinitionRequestQueryList {
	var returns DashboardV2WidgetSunburstDefinitionRequestQueryList
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) RumQuery() DashboardV2WidgetSunburstDefinitionRequestRumQueryOutputReference {
	var returns DashboardV2WidgetSunburstDefinitionRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) RumQueryInput() *DashboardV2WidgetSunburstDefinitionRequestRumQuery {
	var returns *DashboardV2WidgetSunburstDefinitionRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) SecurityQuery() DashboardV2WidgetSunburstDefinitionRequestSecurityQueryOutputReference {
	var returns DashboardV2WidgetSunburstDefinitionRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) SecurityQueryInput() *DashboardV2WidgetSunburstDefinitionRequestSecurityQuery {
	var returns *DashboardV2WidgetSunburstDefinitionRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) Sort() DashboardV2WidgetSunburstDefinitionRequestSortOutputReference {
	var returns DashboardV2WidgetSunburstDefinitionRequestSortOutputReference
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) SortInput() *DashboardV2WidgetSunburstDefinitionRequestSort {
	var returns *DashboardV2WidgetSunburstDefinitionRequestSort
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) Style() DashboardV2WidgetSunburstDefinitionRequestStyleOutputReference {
	var returns DashboardV2WidgetSunburstDefinitionRequestStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) StyleInput() *DashboardV2WidgetSunburstDefinitionRequestStyle {
	var returns *DashboardV2WidgetSunburstDefinitionRequestStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetSunburstDefinitionRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetSunburstDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetSunburstDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSunburstDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetSunburstDefinitionRequestOutputReference_Override(d DashboardV2WidgetSunburstDefinitionRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSunburstDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) PutApmQuery(value *DashboardV2WidgetSunburstDefinitionRequestApmQuery) {
	if err := d.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) PutAuditQuery(value *DashboardV2WidgetSunburstDefinitionRequestAuditQuery) {
	if err := d.validatePutAuditQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAuditQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) PutFormula(value interface{}) {
	if err := d.validatePutFormulaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFormula",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) PutLogQuery(value *DashboardV2WidgetSunburstDefinitionRequestLogQuery) {
	if err := d.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) PutNetworkQuery(value *DashboardV2WidgetSunburstDefinitionRequestNetworkQuery) {
	if err := d.validatePutNetworkQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNetworkQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) PutProcessQuery(value *DashboardV2WidgetSunburstDefinitionRequestProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) PutQuery(value interface{}) {
	if err := d.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) PutRumQuery(value *DashboardV2WidgetSunburstDefinitionRequestRumQuery) {
	if err := d.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) PutSecurityQuery(value *DashboardV2WidgetSunburstDefinitionRequestSecurityQuery) {
	if err := d.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) PutSort(value *DashboardV2WidgetSunburstDefinitionRequestSort) {
	if err := d.validatePutSortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSort",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) PutStyle(value *DashboardV2WidgetSunburstDefinitionRequestStyle) {
	if err := d.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStyle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ResetAuditQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetAuditQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ResetFormula() {
	_jsii_.InvokeVoid(
		d,
		"resetFormula",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ResetNetworkQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		d,
		"resetQ",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		d,
		"resetSort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		d,
		"resetStyle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSunburstDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

