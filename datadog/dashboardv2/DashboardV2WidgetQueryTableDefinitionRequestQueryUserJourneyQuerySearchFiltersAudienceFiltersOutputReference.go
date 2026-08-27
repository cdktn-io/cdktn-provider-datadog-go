// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference interface {
	cdktn.ComplexObject
	Account() DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccountList
	AccountInput() interface{}
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
	FilterCondition() *string
	SetFilterCondition(val *string)
	FilterConditionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters
	SetInternalValue(val *DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters)
	Segment() DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentList
	SegmentInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	User() DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUserList
	UserInput() interface{}
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
	PutAccount(value interface{})
	PutSegment(value interface{})
	PutUser(value interface{})
	ResetAccount()
	ResetFilterCondition()
	ResetSegment()
	ResetUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference
type jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) Account() DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccountList {
	var returns DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccountList
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) AccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) FilterCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) FilterConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) InternalValue() *DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters {
	var returns *DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) Segment() DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentList {
	var returns DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentList
	_jsii_.Get(
		j,
		"segment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) SegmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"segmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) User() DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUserList {
	var returns DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUserList
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) UserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference_Override(d DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetFilterCondition(val *string) {
	if err := j.validateSetFilterConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterCondition",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetInternalValue(val *DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) PutAccount(value interface{}) {
	if err := d.validatePutAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAccount",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) PutSegment(value interface{}) {
	if err := d.validatePutSegmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSegment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) PutUser(value interface{}) {
	if err := d.validatePutUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUser",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ResetAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ResetFilterCondition() {
	_jsii_.InvokeVoid(
		d,
		"resetFilterCondition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ResetSegment() {
	_jsii_.InvokeVoid(
		d,
		"resetSegment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		d,
		"resetUser",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

