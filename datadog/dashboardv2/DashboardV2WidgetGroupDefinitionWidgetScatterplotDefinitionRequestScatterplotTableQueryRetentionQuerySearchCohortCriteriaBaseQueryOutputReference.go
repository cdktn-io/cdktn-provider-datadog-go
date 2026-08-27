// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQuery
	SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQuery)
	Search() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQuerySearchOutputReference
	SearchInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQuerySearch
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
	PutSearch(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQuerySearch)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) InternalValue() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQuery {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) Search() DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQuerySearchOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQuerySearchOutputReference
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) SearchInput() *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQuerySearch {
	var returns *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQuerySearch
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference)SetDataSource(val *string) {
	if err := j.validateSetDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference)SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) PutSearch(value *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQuerySearch) {
	if err := d.validatePutSearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSearch",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryRetentionQuerySearchCohortCriteriaBaseQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

