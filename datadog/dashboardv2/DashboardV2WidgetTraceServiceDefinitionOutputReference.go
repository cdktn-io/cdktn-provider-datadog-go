// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetTraceServiceDefinitionOutputReference interface {
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
	DisplayFormat() *string
	SetDisplayFormat(val *string)
	DisplayFormatInput() *string
	Env() *string
	SetEnv(val *string)
	EnvInput() *string
	// Experimental.
	Fqn() *string
	HideIncompleteCostData() interface{}
	SetHideIncompleteCostData(val interface{})
	HideIncompleteCostDataInput() interface{}
	InternalValue() *DashboardV2WidgetTraceServiceDefinition
	SetInternalValue(val *DashboardV2WidgetTraceServiceDefinition)
	LiveSpan() *string
	SetLiveSpan(val *string)
	LiveSpanInput() *string
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	ShowBreakdown() interface{}
	SetShowBreakdown(val interface{})
	ShowBreakdownInput() interface{}
	ShowDistribution() interface{}
	SetShowDistribution(val interface{})
	ShowDistributionInput() interface{}
	ShowErrors() interface{}
	SetShowErrors(val interface{})
	ShowErrorsInput() interface{}
	ShowHits() interface{}
	SetShowHits(val interface{})
	ShowHitsInput() interface{}
	ShowLatency() interface{}
	SetShowLatency(val interface{})
	ShowLatencyInput() interface{}
	ShowResourceList() interface{}
	SetShowResourceList(val interface{})
	ShowResourceListInput() interface{}
	SizeFormat() *string
	SetSizeFormat(val *string)
	SizeFormatInput() *string
	SpanName() *string
	SetSpanName(val *string)
	SpanNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Time() DashboardV2WidgetTraceServiceDefinitionTimeOutputReference
	TimeInput() *DashboardV2WidgetTraceServiceDefinitionTime
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
	PutTime(value *DashboardV2WidgetTraceServiceDefinitionTime)
	ResetDescription()
	ResetDisplayFormat()
	ResetHideIncompleteCostData()
	ResetLiveSpan()
	ResetShowBreakdown()
	ResetShowDistribution()
	ResetShowErrors()
	ResetShowHits()
	ResetShowLatency()
	ResetShowResourceList()
	ResetSizeFormat()
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

// The jsii proxy struct for DashboardV2WidgetTraceServiceDefinitionOutputReference
type jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) DisplayFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) DisplayFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) Env() *string {
	var returns *string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) EnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) HideIncompleteCostData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) HideIncompleteCostDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hideIncompleteCostDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) InternalValue() *DashboardV2WidgetTraceServiceDefinition {
	var returns *DashboardV2WidgetTraceServiceDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) LiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) LiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ShowBreakdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showBreakdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ShowBreakdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showBreakdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ShowDistribution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ShowDistributionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showDistributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ShowErrors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ShowErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ShowHits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showHits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ShowHitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showHitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ShowLatency() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ShowLatencyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ShowResourceList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showResourceList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ShowResourceListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showResourceListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) SizeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) SizeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) SpanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spanName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) SpanNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spanNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) Time() DashboardV2WidgetTraceServiceDefinitionTimeOutputReference {
	var returns DashboardV2WidgetTraceServiceDefinitionTimeOutputReference
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) TimeInput() *DashboardV2WidgetTraceServiceDefinitionTime {
	var returns *DashboardV2WidgetTraceServiceDefinitionTime
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) TitleAlign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) TitleAlignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleAlignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) TitleSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) TitleSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleSizeInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetTraceServiceDefinitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetTraceServiceDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetTraceServiceDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetTraceServiceDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetTraceServiceDefinitionOutputReference_Override(d DashboardV2WidgetTraceServiceDefinitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetTraceServiceDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetDisplayFormat(val *string) {
	if err := j.validateSetDisplayFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayFormat",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetEnv(val *string) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetHideIncompleteCostData(val interface{}) {
	if err := j.validateSetHideIncompleteCostDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hideIncompleteCostData",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetInternalValue(val *DashboardV2WidgetTraceServiceDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetLiveSpan(val *string) {
	if err := j.validateSetLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"liveSpan",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetShowBreakdown(val interface{}) {
	if err := j.validateSetShowBreakdownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showBreakdown",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetShowDistribution(val interface{}) {
	if err := j.validateSetShowDistributionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showDistribution",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetShowErrors(val interface{}) {
	if err := j.validateSetShowErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showErrors",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetShowHits(val interface{}) {
	if err := j.validateSetShowHitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showHits",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetShowLatency(val interface{}) {
	if err := j.validateSetShowLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showLatency",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetShowResourceList(val interface{}) {
	if err := j.validateSetShowResourceListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showResourceList",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetSizeFormat(val *string) {
	if err := j.validateSetSizeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeFormat",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetSpanName(val *string) {
	if err := j.validateSetSpanNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spanName",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetTitleAlign(val *string) {
	if err := j.validateSetTitleAlignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleAlign",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference)SetTitleSize(val *string) {
	if err := j.validateSetTitleSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"titleSize",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) PutTime(value *DashboardV2WidgetTraceServiceDefinitionTime) {
	if err := d.validatePutTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTime",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetDisplayFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetHideIncompleteCostData() {
	_jsii_.InvokeVoid(
		d,
		"resetHideIncompleteCostData",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetLiveSpan() {
	_jsii_.InvokeVoid(
		d,
		"resetLiveSpan",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetShowBreakdown() {
	_jsii_.InvokeVoid(
		d,
		"resetShowBreakdown",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetShowDistribution() {
	_jsii_.InvokeVoid(
		d,
		"resetShowDistribution",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetShowErrors() {
	_jsii_.InvokeVoid(
		d,
		"resetShowErrors",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetShowHits() {
	_jsii_.InvokeVoid(
		d,
		"resetShowHits",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetShowLatency() {
	_jsii_.InvokeVoid(
		d,
		"resetShowLatency",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetShowResourceList() {
	_jsii_.InvokeVoid(
		d,
		"resetShowResourceList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetSizeFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetSizeFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetTime() {
	_jsii_.InvokeVoid(
		d,
		"resetTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		d,
		"resetTitle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetTitleAlign() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleAlign",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ResetTitleSize() {
	_jsii_.InvokeVoid(
		d,
		"resetTitleSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetTraceServiceDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

