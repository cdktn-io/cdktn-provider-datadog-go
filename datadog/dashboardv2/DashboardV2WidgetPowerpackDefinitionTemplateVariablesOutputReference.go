// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference interface {
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
	ControlledByPowerpack() DashboardV2WidgetPowerpackDefinitionTemplateVariablesControlledByPowerpackList
	ControlledByPowerpackInput() interface{}
	ControlledExternally() DashboardV2WidgetPowerpackDefinitionTemplateVariablesControlledExternallyList
	ControlledExternallyInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardV2WidgetPowerpackDefinitionTemplateVariables
	SetInternalValue(val *DashboardV2WidgetPowerpackDefinitionTemplateVariables)
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
	PutControlledByPowerpack(value interface{})
	PutControlledExternally(value interface{})
	ResetControlledByPowerpack()
	ResetControlledExternally()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference
type jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) ControlledByPowerpack() DashboardV2WidgetPowerpackDefinitionTemplateVariablesControlledByPowerpackList {
	var returns DashboardV2WidgetPowerpackDefinitionTemplateVariablesControlledByPowerpackList
	_jsii_.Get(
		j,
		"controlledByPowerpack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) ControlledByPowerpackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"controlledByPowerpackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) ControlledExternally() DashboardV2WidgetPowerpackDefinitionTemplateVariablesControlledExternallyList {
	var returns DashboardV2WidgetPowerpackDefinitionTemplateVariablesControlledExternallyList
	_jsii_.Get(
		j,
		"controlledExternally",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) ControlledExternallyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"controlledExternallyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) InternalValue() *DashboardV2WidgetPowerpackDefinitionTemplateVariables {
	var returns *DashboardV2WidgetPowerpackDefinitionTemplateVariables
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference_Override(d DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference)SetInternalValue(val *DashboardV2WidgetPowerpackDefinitionTemplateVariables) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) PutControlledByPowerpack(value interface{}) {
	if err := d.validatePutControlledByPowerpackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putControlledByPowerpack",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) PutControlledExternally(value interface{}) {
	if err := d.validatePutControlledExternallyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putControlledExternally",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) ResetControlledByPowerpack() {
	_jsii_.InvokeVoid(
		d,
		"resetControlledByPowerpack",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) ResetControlledExternally() {
	_jsii_.InvokeVoid(
		d,
		"resetControlledExternally",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetPowerpackDefinitionTemplateVariablesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

