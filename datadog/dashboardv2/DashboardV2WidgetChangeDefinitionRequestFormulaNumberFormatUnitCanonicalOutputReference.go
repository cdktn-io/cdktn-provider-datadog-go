// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference interface {
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
	InternalValue() *DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonical
	SetInternalValue(val *DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonical)
	PerUnitName() *string
	SetPerUnitName(val *string)
	PerUnitNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UnitName() *string
	SetUnitName(val *string)
	UnitNameInput() *string
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
	ResetPerUnitName()
	ResetUnitName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference
type jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) InternalValue() *DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonical {
	var returns *DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonical
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) PerUnitName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perUnitName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) PerUnitNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perUnitNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) UnitName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) UnitNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitNameInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference_Override(d DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetInternalValue(val *DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonical) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetPerUnitName(val *string) {
	if err := j.validateSetPerUnitNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perUnitName",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference)SetUnitName(val *string) {
	if err := j.validateSetUnitNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitName",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) ResetPerUnitName() {
	_jsii_.InvokeVoid(
		d,
		"resetPerUnitName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) ResetUnitName() {
	_jsii_.InvokeVoid(
		d,
		"resetUnitName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetChangeDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

