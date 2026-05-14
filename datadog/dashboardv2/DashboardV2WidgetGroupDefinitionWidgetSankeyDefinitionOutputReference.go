// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	HideIncompleteCostData() interface{}
	SetHideIncompleteCostData(val interface{})
	HideIncompleteCostDataInput() interface{}
	InternalValue() *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinition
	SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinition)
	LiveSpan() *string
	SetLiveSpan(val *string)
	LiveSpanInput() *string
	Request() DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestList
	RequestInput() interface{}
	ShowOtherLinks() interface{}
	SetShowOtherLinks(val interface{})
	ShowOtherLinksInput() interface{}
	SortNodes() interface{}
	SetSortNodes(val interface{})
	SortNodesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Time() DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionTimeOutputReference
	TimeInput() *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionTime
	Title() *string
	SetTitle(val *string)
	TitleAlign() *string
	SetTitleAlign(val *string)
	TitleAlignInput() *string
	TitleInput() *string
	TitleSize() *string
	SetTitleSize(val *string)
	TitleSizeInput() *string
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
	PutRequest(value interface{})
	PutTime(value *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionTime)
	ResetDescription()
	ResetHideIncompleteCostData()
	ResetLiveSpan()
	ResetRequest()
	ResetShowOtherLinks()
	ResetSortNodes()
	ResetTime()
	ResetTitle()
	ResetTitleAlign()
	ResetTitleSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) HideIncompleteCostData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) HideIncompleteCostDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) InternalValue() *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinition {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) LiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) LiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) Request() DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestList {
	var returns DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionRequestList
	_jsii_.Get(
		j,
		"request",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) RequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ShowOtherLinks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showOtherLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ShowOtherLinksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showOtherLinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) SortNodes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) SortNodesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) Time() DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionTimeOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionTimeOutputReference
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) TimeInput() *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionTime {
	var returns *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionTime
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference)SetHideIncompleteCostData(val interface{}) {
	if err := j.validateSetHideIncompleteCostDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIncompleteCostData",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference)SetInternalValue(val *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference)SetLiveSpan(val *string) {
	if err := j.validateSetLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSpan",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference)SetShowOtherLinks(val interface{}) {
	if err := j.validateSetShowOtherLinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showOtherLinks",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference)SetSortNodes(val interface{}) {
	if err := j.validateSetSortNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortNodes",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) PutRequest(value interface{}) {
	if err := d.validatePutRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRequest",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) PutTime(value *DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionTime) {
	if err := d.validatePutTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTime",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ResetHideIncompleteCostData() {
	_jsii_.InvokeVoid(
		d,
		"resetHideIncompleteCostData",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ResetLiveSpan() {
	_jsii_.InvokeVoid(
		d,
		"resetLiveSpan",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ResetRequest() {
	_jsii_.InvokeVoid(
		d,
		"resetRequest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ResetShowOtherLinks() {
	_jsii_.InvokeVoid(
		d,
		"resetShowOtherLinks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ResetSortNodes() {
	_jsii_.InvokeVoid(
		d,
		"resetSortNodes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ResetTime() {
	_jsii_.InvokeVoid(
		d,
		"resetTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetSankeyDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

