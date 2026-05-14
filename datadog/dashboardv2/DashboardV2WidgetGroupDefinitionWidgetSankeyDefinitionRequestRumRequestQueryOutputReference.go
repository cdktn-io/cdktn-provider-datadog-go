// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference interface {
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
	EntriesPerStep() *float64
	SetEntriesPerStep(val *float64)
	EntriesPerStepInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQuery
	SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQuery)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	NumberOfSteps() *float64
	SetNumberOfSteps(val *float64)
	NumberOfStepsInput() *float64
	QueryString() *string
	SetQueryString(val *string)
	QueryStringInput() *string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	SubqueryId() *string
	SetSubqueryId(val *string)
	SubqueryIdInput() *string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
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
	ResetEntriesPerStep()
	ResetNumberOfSteps()
	ResetSource()
	ResetSubqueryId()
	ResetTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) EntriesPerStep() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"entriesPerStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) EntriesPerStepInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"entriesPerStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) InternalValue() *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) NumberOfSteps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) NumberOfStepsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfStepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) QueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) SubqueryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subqueryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) SubqueryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subqueryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetEntriesPerStep(val *float64) {
	if err := j.validateSetEntriesPerStepParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entriesPerStep",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetNumberOfSteps(val *float64) {
	if err := j.validateSetNumberOfStepsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfSteps",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetQueryString(val *string) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetSubqueryId(val *string) {
	if err := j.validateSetSubqueryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subqueryId",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ResetEntriesPerStep() {
	_jsii_.InvokeVoid(
		d,
		"resetEntriesPerStep",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ResetNumberOfSteps() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberOfSteps",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		d,
		"resetSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ResetSubqueryId() {
	_jsii_.InvokeVoid(
		d,
		"resetSubqueryId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

