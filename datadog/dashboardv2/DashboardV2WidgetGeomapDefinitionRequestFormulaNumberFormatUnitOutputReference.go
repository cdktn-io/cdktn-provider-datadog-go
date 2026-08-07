// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference interface {
	cdktn.ComplexObject
	Canonical() DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference
	CanonicalInput() *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonical
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
	Custom() DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCustomOutputReference
	CustomInput() *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCustom
	// Experimental.
	Fqn() *string
	InternalValue() *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnit
	SetInternalValue(val *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnit)
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
	PutCanonical(value *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonical)
	PutCustom(value *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCustom)
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

// The jsii proxy struct for DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference
type jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) Canonical() DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference {
	var returns DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonicalOutputReference
	_jsii_.Get(
		j,
		"canonical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) CanonicalInput() *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonical {
	var returns *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonical
	_jsii_.Get(
		j,
		"canonicalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) Custom() DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCustomOutputReference {
	var returns DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCustomOutputReference
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) CustomInput() *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCustom {
	var returns *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCustom
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) InternalValue() *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnit {
	var returns *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnit
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference_Override(d DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference)SetInternalValue(val *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnit) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) PutCanonical(value *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonical) {
	if err := d.validatePutCanonicalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCanonical",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) PutCustom(value *DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitCustom) {
	if err := d.validatePutCustomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustom",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) ResetCanonical() {
	_jsii_.InvokeVoid(
		d,
		"resetCanonical",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		d,
		"resetCustom",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGeomapDefinitionRequestFormulaNumberFormatUnitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

