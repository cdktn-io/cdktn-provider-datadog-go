// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetManageStatusDefinitionOutputReference interface {
	cdktn.ComplexObject
	ColorPreference() *string
	SetColorPreference(val *string)
	ColorPreferenceInput() *string
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
	DisplayFormat() *string
	SetDisplayFormat(val *string)
	DisplayFormatInput() *string
	// Experimental.
	Fqn() *string
	HideIncompleteCostData() interface{}
	SetHideIncompleteCostData(val interface{})
	HideIncompleteCostDataInput() interface{}
	HideZeroCounts() interface{}
	SetHideZeroCounts(val interface{})
	HideZeroCountsInput() interface{}
	InternalValue() *DashboardV2WidgetManageStatusDefinition
	SetInternalValue(val *DashboardV2WidgetManageStatusDefinition)
	LiveSpan() *string
	SetLiveSpan(val *string)
	LiveSpanInput() *string
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	ShowLastTriggered() interface{}
	SetShowLastTriggered(val interface{})
	ShowLastTriggeredInput() interface{}
	ShowPriority() interface{}
	SetShowPriority(val interface{})
	ShowPriorityInput() interface{}
	Sort() *string
	SetSort(val *string)
	SortInput() *string
	SummaryType() *string
	SetSummaryType(val *string)
	SummaryTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Time() DashboardV2WidgetManageStatusDefinitionTimeOutputReference
	TimeInput() *DashboardV2WidgetManageStatusDefinitionTime
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
	PutTime(value *DashboardV2WidgetManageStatusDefinitionTime)
	ResetColorPreference()
	ResetDescription()
	ResetDisplayFormat()
	ResetHideIncompleteCostData()
	ResetHideZeroCounts()
	ResetLiveSpan()
	ResetShowLastTriggered()
	ResetShowPriority()
	ResetSort()
	ResetSummaryType()
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

// The jsii proxy struct for DashboardV2WidgetManageStatusDefinitionOutputReference
type jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ColorPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ColorPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) DisplayFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) DisplayFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) HideIncompleteCostData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) HideIncompleteCostDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) HideZeroCounts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideZeroCounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) HideZeroCountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideZeroCountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) InternalValue() *DashboardV2WidgetManageStatusDefinition {
	var returns *DashboardV2WidgetManageStatusDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) LiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) LiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ShowLastTriggered() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLastTriggered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ShowLastTriggeredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLastTriggeredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ShowPriority() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ShowPriorityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) Sort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) SortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) SummaryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) SummaryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"summaryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) Time() DashboardV2WidgetManageStatusDefinitionTimeOutputReference {
	var returns DashboardV2WidgetManageStatusDefinitionTimeOutputReference
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) TimeInput() *DashboardV2WidgetManageStatusDefinitionTime {
	var returns *DashboardV2WidgetManageStatusDefinitionTime
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetManageStatusDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetManageStatusDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetManageStatusDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetManageStatusDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetManageStatusDefinitionOutputReference_Override(d DashboardV2WidgetManageStatusDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetManageStatusDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetColorPreference(val *string) {
	if err := j.validateSetColorPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colorPreference",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetDisplayFormat(val *string) {
	if err := j.validateSetDisplayFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayFormat",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetHideIncompleteCostData(val interface{}) {
	if err := j.validateSetHideIncompleteCostDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIncompleteCostData",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetHideZeroCounts(val interface{}) {
	if err := j.validateSetHideZeroCountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideZeroCounts",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetInternalValue(val *DashboardV2WidgetManageStatusDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetLiveSpan(val *string) {
	if err := j.validateSetLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSpan",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetShowLastTriggered(val interface{}) {
	if err := j.validateSetShowLastTriggeredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showLastTriggered",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetShowPriority(val interface{}) {
	if err := j.validateSetShowPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showPriority",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetSort(val *string) {
	if err := j.validateSetSortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sort",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetSummaryType(val *string) {
	if err := j.validateSetSummaryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"summaryType",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) PutTime(value *DashboardV2WidgetManageStatusDefinitionTime) {
	if err := d.validatePutTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTime",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetColorPreference() {
	_jsii_.InvokeVoid(
		d,
		"resetColorPreference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetDisplayFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetHideIncompleteCostData() {
	_jsii_.InvokeVoid(
		d,
		"resetHideIncompleteCostData",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetHideZeroCounts() {
	_jsii_.InvokeVoid(
		d,
		"resetHideZeroCounts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetLiveSpan() {
	_jsii_.InvokeVoid(
		d,
		"resetLiveSpan",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetShowLastTriggered() {
	_jsii_.InvokeVoid(
		d,
		"resetShowLastTriggered",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetShowPriority() {
	_jsii_.InvokeVoid(
		d,
		"resetShowPriority",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		d,
		"resetSort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetSummaryType() {
	_jsii_.InvokeVoid(
		d,
		"resetSummaryType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetTime() {
	_jsii_.InvokeVoid(
		d,
		"resetTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetManageStatusDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

