// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference interface {
	cdktn.ComplexObject
	ApmQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestApmQueryOutputReference
	ApmQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestApmQuery
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
	Formula() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaList
	FormulaInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequest
	SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequest)
	LogQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestLogQueryOutputReference
	LogQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestLogQuery
	ProcessQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestProcessQueryOutputReference
	ProcessQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	Query() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryList
	QueryInput() interface{}
	RumQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestRumQueryOutputReference
	RumQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestRumQuery
	SecurityQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSecurityQueryOutputReference
	SecurityQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSecurityQuery
	Sort() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSortOutputReference
	SortInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSort
	Style() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestStyleOutputReference
	StyleInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestStyle
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
	PutApmQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestApmQuery)
	PutFormula(value interface{})
	PutLogQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestLogQuery)
	PutProcessQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestProcessQuery)
	PutQuery(value interface{})
	PutRumQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestRumQuery)
	PutSecurityQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSecurityQuery)
	PutSort(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSort)
	PutStyle(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestStyle)
	ResetApmQuery()
	ResetDisplayType()
	ResetFormula()
	ResetLogQuery()
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

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ApmQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestApmQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ApmQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestApmQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) DisplayType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) DisplayTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) Formula() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaList {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaList
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) FormulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) InternalValue() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequest {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) LogQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestLogQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) LogQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestLogQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ProcessQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestProcessQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ProcessQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestProcessQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) Query() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryList {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestQueryList
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) RumQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestRumQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) RumQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestRumQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) SecurityQuery() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSecurityQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) SecurityQueryInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSecurityQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) Sort() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSortOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSortOutputReference
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) SortInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSort {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSort
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) Style() DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestStyleOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) StyleInput() *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestStyle {
	var returns *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference)SetDisplayType(val *string) {
	if err := j.validateSetDisplayTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayType",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference)SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) PutApmQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestApmQuery) {
	if err := d.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) PutFormula(value interface{}) {
	if err := d.validatePutFormulaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFormula",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) PutLogQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestLogQuery) {
	if err := d.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) PutProcessQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) PutQuery(value interface{}) {
	if err := d.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) PutRumQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestRumQuery) {
	if err := d.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) PutSecurityQuery(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSecurityQuery) {
	if err := d.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) PutSort(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestSort) {
	if err := d.validatePutSortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSort",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) PutStyle(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestStyle) {
	if err := d.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStyle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ResetDisplayType() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ResetFormula() {
	_jsii_.InvokeVoid(
		d,
		"resetFormula",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		d,
		"resetQ",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		d,
		"resetSort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		d,
		"resetStyle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

