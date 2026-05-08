// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormat
	SetInternalValue(val *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormat)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Unit() DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnitOutputReference
	UnitInput() *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnit
	UnitScale() DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnitScaleOutputReference
	UnitScaleInput() *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnitScale
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
	PutUnit(value *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnit)
	PutUnitScale(value *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnitScale)
	ResetUnitScale()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference
type jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) InternalValue() *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormat {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormat
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) Unit() DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnitOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnitOutputReference
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) UnitInput() *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnit {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnit
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) UnitScale() DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnitScaleOutputReference {
	var returns DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnitScaleOutputReference
	_jsii_.Get(
		j,
		"unitScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) UnitScaleInput() *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnitScale {
	var returns *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnitScale
	_jsii_.Get(
		j,
		"unitScaleInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference_Override(d DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference)SetInternalValue(val *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormat) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) PutUnit(value *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnit) {
	if err := d.validatePutUnitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUnit",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) PutUnitScale(value *DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatUnitScale) {
	if err := d.validatePutUnitScaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUnitScale",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) ResetUnitScale() {
	_jsii_.InvokeVoid(
		d,
		"resetUnitScale",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetScatterplotDefinitionRequestYFormulaNumberFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

