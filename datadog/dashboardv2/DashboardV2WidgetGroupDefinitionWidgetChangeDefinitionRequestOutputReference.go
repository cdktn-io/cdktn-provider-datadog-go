// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference interface {
	cdktn.ComplexObject
	ApmQuery() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestApmQueryOutputReference
	ApmQueryInput() *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestApmQuery
	ChangeType() *string
	SetChangeType(val *string)
	ChangeTypeInput() *string
	CompareTo() *string
	SetCompareTo(val *string)
	CompareToInput() *string
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
	Formula() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestFormulaList
	FormulaInput() interface{}
	// Experimental.
	Fqn() *string
	IncreaseGood() interface{}
	SetIncreaseGood(val interface{})
	IncreaseGoodInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogQuery() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestLogQueryOutputReference
	LogQueryInput() *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestLogQuery
	OrderBy() *string
	SetOrderBy(val *string)
	OrderByInput() *string
	OrderDir() *string
	SetOrderDir(val *string)
	OrderDirInput() *string
	ProcessQuery() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestProcessQueryOutputReference
	ProcessQueryInput() *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestProcessQuery
	Q() *string
	SetQ(val *string)
	QInput() *string
	Query() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryList
	QueryInput() interface{}
	RumQuery() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestRumQueryOutputReference
	RumQueryInput() *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestRumQuery
	SecurityQuery() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestSecurityQueryOutputReference
	SecurityQueryInput() *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestSecurityQuery
	ShowPresent() interface{}
	SetShowPresent(val interface{})
	ShowPresentInput() interface{}
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
	PutApmQuery(value *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestApmQuery)
	PutFormula(value interface{})
	PutLogQuery(value *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestLogQuery)
	PutProcessQuery(value *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestProcessQuery)
	PutQuery(value interface{})
	PutRumQuery(value *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestRumQuery)
	PutSecurityQuery(value *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestSecurityQuery)
	ResetApmQuery()
	ResetChangeType()
	ResetCompareTo()
	ResetFormula()
	ResetIncreaseGood()
	ResetLogQuery()
	ResetOrderBy()
	ResetOrderDir()
	ResetProcessQuery()
	ResetQ()
	ResetQuery()
	ResetRumQuery()
	ResetSecurityQuery()
	ResetShowPresent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ApmQuery() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestApmQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestApmQueryOutputReference
	_jsii_.Get(
		j,
		"apmQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ApmQueryInput() *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestApmQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestApmQuery
	_jsii_.Get(
		j,
		"apmQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ChangeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ChangeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) CompareTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) CompareToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compareToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) Formula() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestFormulaList {
	var returns DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestFormulaList
	_jsii_.Get(
		j,
		"formula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) FormulaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"formulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) IncreaseGood() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"increaseGood",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) IncreaseGoodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"increaseGoodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) LogQuery() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestLogQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestLogQueryOutputReference
	_jsii_.Get(
		j,
		"logQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) LogQueryInput() *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestLogQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestLogQuery
	_jsii_.Get(
		j,
		"logQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) OrderBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) OrderByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) OrderDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) OrderDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ProcessQuery() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestProcessQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestProcessQueryOutputReference
	_jsii_.Get(
		j,
		"processQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ProcessQueryInput() *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestProcessQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestProcessQuery
	_jsii_.Get(
		j,
		"processQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) Q() *string {
	var returns *string
	_jsii_.Get(
		j,
		"q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) QInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) Query() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryList {
	var returns DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestQueryList
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) RumQuery() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestRumQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestRumQueryOutputReference
	_jsii_.Get(
		j,
		"rumQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) RumQueryInput() *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestRumQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestRumQuery
	_jsii_.Get(
		j,
		"rumQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) SecurityQuery() DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestSecurityQueryOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestSecurityQueryOutputReference
	_jsii_.Get(
		j,
		"securityQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) SecurityQueryInput() *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestSecurityQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestSecurityQuery
	_jsii_.Get(
		j,
		"securityQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ShowPresent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPresent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ShowPresentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPresentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference)SetChangeType(val *string) {
	if err := j.validateSetChangeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeType",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference)SetCompareTo(val *string) {
	if err := j.validateSetCompareToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compareTo",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference)SetIncreaseGood(val interface{}) {
	if err := j.validateSetIncreaseGoodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"increaseGood",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference)SetOrderBy(val *string) {
	if err := j.validateSetOrderByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orderBy",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference)SetOrderDir(val *string) {
	if err := j.validateSetOrderDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orderDir",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference)SetQ(val *string) {
	if err := j.validateSetQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"q",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference)SetShowPresent(val interface{}) {
	if err := j.validateSetShowPresentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showPresent",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) PutApmQuery(value *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestApmQuery) {
	if err := d.validatePutApmQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putApmQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) PutFormula(value interface{}) {
	if err := d.validatePutFormulaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFormula",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) PutLogQuery(value *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestLogQuery) {
	if err := d.validatePutLogQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLogQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) PutProcessQuery(value *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestProcessQuery) {
	if err := d.validatePutProcessQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProcessQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) PutQuery(value interface{}) {
	if err := d.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) PutRumQuery(value *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestRumQuery) {
	if err := d.validatePutRumQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRumQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) PutSecurityQuery(value *DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestSecurityQuery) {
	if err := d.validatePutSecurityQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecurityQuery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetApmQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetApmQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetChangeType() {
	_jsii_.InvokeVoid(
		d,
		"resetChangeType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetCompareTo() {
	_jsii_.InvokeVoid(
		d,
		"resetCompareTo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetFormula() {
	_jsii_.InvokeVoid(
		d,
		"resetFormula",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetIncreaseGood() {
	_jsii_.InvokeVoid(
		d,
		"resetIncreaseGood",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetLogQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetLogQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetOrderBy() {
	_jsii_.InvokeVoid(
		d,
		"resetOrderBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetOrderDir() {
	_jsii_.InvokeVoid(
		d,
		"resetOrderDir",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetProcessQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetProcessQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetQ() {
	_jsii_.InvokeVoid(
		d,
		"resetQ",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetRumQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetRumQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetSecurityQuery() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityQuery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ResetShowPresent() {
	_jsii_.InvokeVoid(
		d,
		"resetShowPresent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetChangeDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

