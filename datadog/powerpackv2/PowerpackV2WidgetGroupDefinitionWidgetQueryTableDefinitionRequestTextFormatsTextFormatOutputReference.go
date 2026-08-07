// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference interface {
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
	CustomBgColor() *string
	SetCustomBgColor(val *string)
	CustomBgColorInput() *string
	CustomFgColor() *string
	SetCustomFgColor(val *string)
	CustomFgColorInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Match() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatMatchOutputReference
	MatchInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatMatch
	Palette() *string
	SetPalette(val *string)
	PaletteInput() *string
	Replace() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference
	ReplaceInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace
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
	PutMatch(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatMatch)
	PutReplace(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace)
	ResetCustomBgColor()
	ResetCustomFgColor()
	ResetPalette()
	ResetReplace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) CustomBgColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customBgColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) CustomBgColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customBgColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) CustomFgColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFgColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) CustomFgColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customFgColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) Match() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatMatchOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatMatchOutputReference
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) MatchInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatMatch {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatMatch
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) Palette() *string {
	var returns *string
	_jsii_.Get(
		j,
		"palette",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) PaletteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paletteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) Replace() PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplaceOutputReference
	_jsii_.Get(
		j,
		"replace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) ReplaceInput() *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace
	_jsii_.Get(
		j,
		"replaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference)SetCustomBgColor(val *string) {
	if err := j.validateSetCustomBgColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customBgColor",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference)SetCustomFgColor(val *string) {
	if err := j.validateSetCustomFgColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customFgColor",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference)SetPalette(val *string) {
	if err := j.validateSetPaletteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"palette",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) PutMatch(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatMatch) {
	if err := p.validatePutMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putMatch",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) PutReplace(value *PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace) {
	if err := p.validatePutReplaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putReplace",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) ResetCustomBgColor() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomBgColor",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) ResetCustomFgColor() {
	_jsii_.InvokeVoid(
		p,
		"resetCustomFgColor",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) ResetPalette() {
	_jsii_.InvokeVoid(
		p,
		"resetPalette",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) ResetReplace() {
	_jsii_.InvokeVoid(
		p,
		"resetReplace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetQueryTableDefinitionRequestTextFormatsTextFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

