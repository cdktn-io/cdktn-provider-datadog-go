// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference interface {
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
	CrossOrgUuids() *[]*string
	SetCrossOrgUuids(val *[]*string)
	CrossOrgUuidsInput() *[]*string
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	Env() *string
	SetEnv(val *string)
	EnvInput() *string
	// Experimental.
	Fqn() *string
	GroupBy() *[]*string
	SetGroupBy(val *[]*string)
	GroupByInput() *[]*string
	InternalValue() *DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQuery
	SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQuery)
	Name() *string
	SetName(val *string)
	NameInput() *string
	OperationName() *string
	SetOperationName(val *string)
	OperationNameInput() *string
	PrimaryTagName() *string
	SetPrimaryTagName(val *string)
	PrimaryTagNameInput() *string
	PrimaryTagValue() *string
	SetPrimaryTagValue(val *string)
	PrimaryTagValueInput() *string
	ResourceName() *string
	SetResourceName(val *string)
	ResourceNameInput() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	Stat() *string
	SetStat(val *string)
	StatInput() *string
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
	ResetCrossOrgUuids()
	ResetGroupBy()
	ResetOperationName()
	ResetPrimaryTagName()
	ResetPrimaryTagValue()
	ResetResourceName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) CrossOrgUuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) CrossOrgUuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossOrgUuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) Env() *string {
	var returns *string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) EnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) GroupBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) GroupByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) InternalValue() *DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) OperationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) PrimaryTagName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) PrimaryTagNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) PrimaryTagValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) PrimaryTagValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryTagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) Stat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) StatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetCrossOrgUuids(val *[]*string) {
	if err := j.validateSetCrossOrgUuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossOrgUuids",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetEnv(val *string) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetGroupBy(val *[]*string) {
	if err := j.validateSetGroupByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupBy",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetOperationName(val *string) {
	if err := j.validateSetOperationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetPrimaryTagName(val *string) {
	if err := j.validateSetPrimaryTagNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryTagName",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetPrimaryTagValue(val *string) {
	if err := j.validateSetPrimaryTagValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryTagValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetResourceName(val *string) {
	if err := j.validateSetResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceName",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetStat(val *string) {
	if err := j.validateSetStatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stat",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetCrossOrgUuids() {
	_jsii_.InvokeVoid(
		d,
		"resetCrossOrgUuids",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetGroupBy() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetOperationName() {
	_jsii_.InvokeVoid(
		d,
		"resetOperationName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetPrimaryTagName() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimaryTagName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetPrimaryTagValue() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimaryTagValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) ResetResourceName() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestQueryApmResourceStatsQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

