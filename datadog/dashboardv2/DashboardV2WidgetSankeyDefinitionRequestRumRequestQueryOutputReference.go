// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference interface {
	cdktn.ComplexObject
	AudienceFilters() DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryAudienceFiltersOutputReference
	AudienceFiltersInput() *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryAudienceFilters
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
	InternalValue() *DashboardV2WidgetSankeyDefinitionRequestRumRequestQuery
	SetInternalValue(val *DashboardV2WidgetSankeyDefinitionRequestRumRequestQuery)
	JoinKeys() DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryJoinKeysOutputReference
	JoinKeysInput() *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryJoinKeys
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	NumberOfSteps() *float64
	SetNumberOfSteps(val *float64)
	NumberOfStepsInput() *float64
	Occurrence() DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOccurrenceOutputReference
	OccurrenceInput() *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOccurrence
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
	PutAudienceFilters(value *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryAudienceFilters)
	PutJoinKeys(value *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryJoinKeys)
	PutOccurrence(value *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOccurrence)
	ResetAudienceFilters()
	ResetEntriesPerStep()
	ResetJoinKeys()
	ResetNumberOfSteps()
	ResetOccurrence()
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

// The jsii proxy struct for DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference
type jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) AudienceFilters() DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryAudienceFiltersOutputReference {
	var returns DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryAudienceFiltersOutputReference
	_jsii_.Get(
		j,
		"audienceFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) AudienceFiltersInput() *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryAudienceFilters {
	var returns *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryAudienceFilters
	_jsii_.Get(
		j,
		"audienceFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) EntriesPerStep() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"entriesPerStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) EntriesPerStepInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"entriesPerStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) InternalValue() *DashboardV2WidgetSankeyDefinitionRequestRumRequestQuery {
	var returns *DashboardV2WidgetSankeyDefinitionRequestRumRequestQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) JoinKeys() DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryJoinKeysOutputReference {
	var returns DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryJoinKeysOutputReference
	_jsii_.Get(
		j,
		"joinKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) JoinKeysInput() *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryJoinKeys {
	var returns *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryJoinKeys
	_jsii_.Get(
		j,
		"joinKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) NumberOfSteps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) NumberOfStepsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfStepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) Occurrence() DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOccurrenceOutputReference {
	var returns DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOccurrenceOutputReference
	_jsii_.Get(
		j,
		"occurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) OccurrenceInput() *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOccurrence {
	var returns *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOccurrence
	_jsii_.Get(
		j,
		"occurrenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) QueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) SubqueryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subqueryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) SubqueryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subqueryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference_Override(d DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetEntriesPerStep(val *float64) {
	if err := j.validateSetEntriesPerStepParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entriesPerStep",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetInternalValue(val *DashboardV2WidgetSankeyDefinitionRequestRumRequestQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetNumberOfSteps(val *float64) {
	if err := j.validateSetNumberOfStepsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfSteps",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetQueryString(val *string) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetSubqueryId(val *string) {
	if err := j.validateSetSubqueryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subqueryId",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) PutAudienceFilters(value *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryAudienceFilters) {
	if err := d.validatePutAudienceFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAudienceFilters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) PutJoinKeys(value *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryJoinKeys) {
	if err := d.validatePutJoinKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJoinKeys",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) PutOccurrence(value *DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOccurrence) {
	if err := d.validatePutOccurrenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOccurrence",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ResetAudienceFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetAudienceFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ResetEntriesPerStep() {
	_jsii_.InvokeVoid(
		d,
		"resetEntriesPerStep",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ResetJoinKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetJoinKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ResetNumberOfSteps() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberOfSteps",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ResetOccurrence() {
	_jsii_.InvokeVoid(
		d,
		"resetOccurrence",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		d,
		"resetSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ResetSubqueryId() {
	_jsii_.InvokeVoid(
		d,
		"resetSubqueryId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSankeyDefinitionRequestRumRequestQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

