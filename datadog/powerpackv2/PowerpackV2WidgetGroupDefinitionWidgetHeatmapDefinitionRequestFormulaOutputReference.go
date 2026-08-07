// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference interface {
	cdktn.ComplexObject
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	CellDisplayMode() *string
	SetCellDisplayMode(val *string)
	CellDisplayModeInput() *string
	CellDisplayModeOptions() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaCellDisplayModeOptionsOutputReference
	CellDisplayModeOptionsInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaCellDisplayModeOptions
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
	ConditionalFormats() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaConditionalFormatsList
	ConditionalFormatsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FormulaExpression() *string
	SetFormulaExpression(val *string)
	FormulaExpressionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Limit() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaLimitOutputReference
	LimitInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaLimit
	NumberFormat() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaNumberFormatOutputReference
	NumberFormatInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaNumberFormat
	Style() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaStyleOutputReference
	StyleInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaStyle
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
	PutCellDisplayModeOptions(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaCellDisplayModeOptions)
	PutConditionalFormats(value interface{})
	PutLimit(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaLimit)
	PutNumberFormat(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaNumberFormat)
	PutStyle(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaStyle)
	ResetAlias()
	ResetCellDisplayMode()
	ResetCellDisplayModeOptions()
	ResetConditionalFormats()
	ResetLimit()
	ResetNumberFormat()
	ResetStyle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference
type jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) CellDisplayMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cellDisplayMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) CellDisplayModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cellDisplayModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) CellDisplayModeOptions() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaCellDisplayModeOptionsOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaCellDisplayModeOptionsOutputReference
	_jsii_.Get(
		j,
		"cellDisplayModeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) CellDisplayModeOptionsInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaCellDisplayModeOptions {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaCellDisplayModeOptions
	_jsii_.Get(
		j,
		"cellDisplayModeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) ConditionalFormats() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaConditionalFormatsList {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaConditionalFormatsList
	_jsii_.Get(
		j,
		"conditionalFormats",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) ConditionalFormatsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionalFormatsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) FormulaExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formulaExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) FormulaExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formulaExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) Limit() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaLimitOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaLimitOutputReference
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) LimitInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaLimit {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaLimit
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) NumberFormat() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaNumberFormatOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaNumberFormatOutputReference
	_jsii_.Get(
		j,
		"numberFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) NumberFormatInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaNumberFormat {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaNumberFormat
	_jsii_.Get(
		j,
		"numberFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) Style() PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaStyleOutputReference {
	var returns PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) StyleInput() *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaStyle {
	var returns *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference_Override(p PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference)SetAlias(val *string) {
	if err := j.validateSetAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference)SetCellDisplayMode(val *string) {
	if err := j.validateSetCellDisplayModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cellDisplayMode",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference)SetFormulaExpression(val *string) {
	if err := j.validateSetFormulaExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formulaExpression",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) PutCellDisplayModeOptions(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaCellDisplayModeOptions) {
	if err := p.validatePutCellDisplayModeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCellDisplayModeOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) PutConditionalFormats(value interface{}) {
	if err := p.validatePutConditionalFormatsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putConditionalFormats",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) PutLimit(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaLimit) {
	if err := p.validatePutLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLimit",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) PutNumberFormat(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaNumberFormat) {
	if err := p.validatePutNumberFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNumberFormat",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) PutStyle(value *PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaStyle) {
	if err := p.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putStyle",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) ResetAlias() {
	_jsii_.InvokeVoid(
		p,
		"resetAlias",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) ResetCellDisplayMode() {
	_jsii_.InvokeVoid(
		p,
		"resetCellDisplayMode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) ResetCellDisplayModeOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetCellDisplayModeOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) ResetConditionalFormats() {
	_jsii_.InvokeVoid(
		p,
		"resetConditionalFormats",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		p,
		"resetLimit",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) ResetNumberFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetNumberFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		p,
		"resetStyle",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetHeatmapDefinitionRequestFormulaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

