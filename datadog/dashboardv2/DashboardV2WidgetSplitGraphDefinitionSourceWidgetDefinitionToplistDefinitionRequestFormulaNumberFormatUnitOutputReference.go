// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference interface {
	cdktn.ComplexObject
	Canonical() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference
	CanonicalInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCanonical
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
	Custom() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCustomOutputReference
	CustomInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCustom
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnit
	SetInternalValue(val *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnit)
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
	PutCanonical(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCanonical)
	PutCustom(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCustom)
	ResetCanonical()
	ResetCustom()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference
type jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) Canonical() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference
	_jsii_.Get(
		j,
		"canonical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) CanonicalInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCanonical {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCanonical
	_jsii_.Get(
		j,
		"canonicalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) Custom() DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCustomOutputReference {
	var returns DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCustomOutputReference
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) CustomInput() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCustom {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCustom
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) InternalValue() *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnit {
	var returns *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnit
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference_Override(d DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference)SetInternalValue(val *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnit) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) PutCanonical(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCanonical) {
	if err := d.validatePutCanonicalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCanonical",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) PutCustom(value *DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitCustom) {
	if err := d.validatePutCustomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustom",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) ResetCanonical() {
	_jsii_.InvokeVoid(
		d,
		"resetCanonical",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		d,
		"resetCustom",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionRequestFormulaNumberFormatUnitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

