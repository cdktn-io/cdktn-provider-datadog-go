// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/monitor/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MonitorSchedulingOptionsEvaluationWindowOutputReference interface {
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
	DayStarts() *string
	SetDayStarts(val *string)
	DayStartsInput() *string
	// Experimental.
	Fqn() *string
	HourStarts() *float64
	SetHourStarts(val *float64)
	HourStartsInput() *float64
	InternalValue() *MonitorSchedulingOptionsEvaluationWindow
	SetInternalValue(val *MonitorSchedulingOptionsEvaluationWindow)
	MonthStarts() *float64
	SetMonthStarts(val *float64)
	MonthStartsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
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
	ResetDayStarts()
	ResetHourStarts()
	ResetMonthStarts()
	ResetTimezone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorSchedulingOptionsEvaluationWindowOutputReference
type jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) DayStarts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayStarts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) DayStartsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayStartsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) HourStarts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourStarts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) HourStartsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourStartsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) InternalValue() *MonitorSchedulingOptionsEvaluationWindow {
	var returns *MonitorSchedulingOptionsEvaluationWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) MonthStarts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthStarts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) MonthStartsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthStartsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}


func NewMonitorSchedulingOptionsEvaluationWindowOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MonitorSchedulingOptionsEvaluationWindowOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorSchedulingOptionsEvaluationWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.monitor.MonitorSchedulingOptionsEvaluationWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorSchedulingOptionsEvaluationWindowOutputReference_Override(m MonitorSchedulingOptionsEvaluationWindowOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.monitor.MonitorSchedulingOptionsEvaluationWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference)SetDayStarts(val *string) {
	if err := j.validateSetDayStartsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayStarts",
		val,
	)
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference)SetHourStarts(val *float64) {
	if err := j.validateSetHourStartsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hourStarts",
		val,
	)
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference)SetInternalValue(val *MonitorSchedulingOptionsEvaluationWindow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference)SetMonthStarts(val *float64) {
	if err := j.validateSetMonthStartsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthStarts",
		val,
	)
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) ResetDayStarts() {
	_jsii_.InvokeVoid(
		m,
		"resetDayStarts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) ResetHourStarts() {
	_jsii_.InvokeVoid(
		m,
		"resetHourStarts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) ResetMonthStarts() {
	_jsii_.InvokeVoid(
		m,
		"resetMonthStarts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		m,
		"resetTimezone",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorSchedulingOptionsEvaluationWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

