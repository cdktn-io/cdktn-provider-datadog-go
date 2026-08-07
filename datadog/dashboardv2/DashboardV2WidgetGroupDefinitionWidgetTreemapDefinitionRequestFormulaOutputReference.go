// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference interface {
	cdktn.ComplexObject
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	CellDisplayMode() *string
	SetCellDisplayMode(val *string)
	CellDisplayModeInput() *string
	CellDisplayModeOptions() DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaCellDisplayModeOptionsOutputReference
	CellDisplayModeOptionsInput() *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaCellDisplayModeOptions
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
	ConditionalFormats() DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaConditionalFormatsList
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
	Limit() DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaLimitOutputReference
	LimitInput() *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaLimit
	NumberFormat() DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference
	NumberFormatInput() *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaNumberFormat
	Style() DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaStyleOutputReference
	StyleInput() *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaStyle
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
	PutCellDisplayModeOptions(value *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaCellDisplayModeOptions)
	PutConditionalFormats(value interface{})
	PutLimit(value *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaLimit)
	PutNumberFormat(value *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaNumberFormat)
	PutStyle(value *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaStyle)
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

// The jsii proxy struct for DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference
type jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) CellDisplayMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cellDisplayMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) CellDisplayModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cellDisplayModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) CellDisplayModeOptions() DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaCellDisplayModeOptionsOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaCellDisplayModeOptionsOutputReference
	_jsii_.Get(
		j,
		"cellDisplayModeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) CellDisplayModeOptionsInput() *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaCellDisplayModeOptions {
	var returns *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaCellDisplayModeOptions
	_jsii_.Get(
		j,
		"cellDisplayModeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) ConditionalFormats() DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaConditionalFormatsList {
	var returns DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaConditionalFormatsList
	_jsii_.Get(
		j,
		"conditionalFormats",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) ConditionalFormatsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionalFormatsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) FormulaExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formulaExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) FormulaExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formulaExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) Limit() DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaLimitOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaLimitOutputReference
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) LimitInput() *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaLimit {
	var returns *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaLimit
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) NumberFormat() DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaNumberFormatOutputReference
	_jsii_.Get(
		j,
		"numberFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) NumberFormatInput() *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaNumberFormat {
	var returns *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaNumberFormat
	_jsii_.Get(
		j,
		"numberFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) Style() DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaStyleOutputReference {
	var returns DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaStyleOutputReference
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) StyleInput() *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaStyle {
	var returns *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaStyle
	_jsii_.Get(
		j,
		"styleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference_Override(d DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference)SetAlias(val *string) {
	if err := j.validateSetAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference)SetCellDisplayMode(val *string) {
	if err := j.validateSetCellDisplayModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cellDisplayMode",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference)SetFormulaExpression(val *string) {
	if err := j.validateSetFormulaExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formulaExpression",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) PutCellDisplayModeOptions(value *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaCellDisplayModeOptions) {
	if err := d.validatePutCellDisplayModeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCellDisplayModeOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) PutConditionalFormats(value interface{}) {
	if err := d.validatePutConditionalFormatsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConditionalFormats",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) PutLimit(value *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaLimit) {
	if err := d.validatePutLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLimit",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) PutNumberFormat(value *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaNumberFormat) {
	if err := d.validatePutNumberFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNumberFormat",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) PutStyle(value *DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaStyle) {
	if err := d.validatePutStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStyle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) ResetAlias() {
	_jsii_.InvokeVoid(
		d,
		"resetAlias",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) ResetCellDisplayMode() {
	_jsii_.InvokeVoid(
		d,
		"resetCellDisplayMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) ResetCellDisplayModeOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetCellDisplayModeOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) ResetConditionalFormats() {
	_jsii_.InvokeVoid(
		d,
		"resetConditionalFormats",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) ResetLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) ResetNumberFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) ResetStyle() {
	_jsii_.InvokeVoid(
		d,
		"resetStyle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetTreemapDefinitionRequestFormulaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

