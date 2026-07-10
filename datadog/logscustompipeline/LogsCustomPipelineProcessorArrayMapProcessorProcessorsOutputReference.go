// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/logscustompipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference interface {
	cdktn.ComplexObject
	ArithmeticProcessor() LogsCustomPipelineProcessorArrayMapProcessorProcessorsArithmeticProcessorOutputReference
	ArithmeticProcessorInput() *LogsCustomPipelineProcessorArrayMapProcessorProcessorsArithmeticProcessor
	AttributeRemapper() LogsCustomPipelineProcessorArrayMapProcessorProcessorsAttributeRemapperOutputReference
	AttributeRemapperInput() *LogsCustomPipelineProcessorArrayMapProcessorProcessorsAttributeRemapper
	CategoryProcessor() LogsCustomPipelineProcessorArrayMapProcessorProcessorsCategoryProcessorOutputReference
	CategoryProcessorInput() *LogsCustomPipelineProcessorArrayMapProcessorProcessorsCategoryProcessor
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StringBuilderProcessor() LogsCustomPipelineProcessorArrayMapProcessorProcessorsStringBuilderProcessorOutputReference
	StringBuilderProcessorInput() *LogsCustomPipelineProcessorArrayMapProcessorProcessorsStringBuilderProcessor
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
	PutArithmeticProcessor(value *LogsCustomPipelineProcessorArrayMapProcessorProcessorsArithmeticProcessor)
	PutAttributeRemapper(value *LogsCustomPipelineProcessorArrayMapProcessorProcessorsAttributeRemapper)
	PutCategoryProcessor(value *LogsCustomPipelineProcessorArrayMapProcessorProcessorsCategoryProcessor)
	PutStringBuilderProcessor(value *LogsCustomPipelineProcessorArrayMapProcessorProcessorsStringBuilderProcessor)
	ResetArithmeticProcessor()
	ResetAttributeRemapper()
	ResetCategoryProcessor()
	ResetStringBuilderProcessor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference
type jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) ArithmeticProcessor() LogsCustomPipelineProcessorArrayMapProcessorProcessorsArithmeticProcessorOutputReference {
	var returns LogsCustomPipelineProcessorArrayMapProcessorProcessorsArithmeticProcessorOutputReference
	_jsii_.Get(
		j,
		"arithmeticProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) ArithmeticProcessorInput() *LogsCustomPipelineProcessorArrayMapProcessorProcessorsArithmeticProcessor {
	var returns *LogsCustomPipelineProcessorArrayMapProcessorProcessorsArithmeticProcessor
	_jsii_.Get(
		j,
		"arithmeticProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) AttributeRemapper() LogsCustomPipelineProcessorArrayMapProcessorProcessorsAttributeRemapperOutputReference {
	var returns LogsCustomPipelineProcessorArrayMapProcessorProcessorsAttributeRemapperOutputReference
	_jsii_.Get(
		j,
		"attributeRemapper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) AttributeRemapperInput() *LogsCustomPipelineProcessorArrayMapProcessorProcessorsAttributeRemapper {
	var returns *LogsCustomPipelineProcessorArrayMapProcessorProcessorsAttributeRemapper
	_jsii_.Get(
		j,
		"attributeRemapperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) CategoryProcessor() LogsCustomPipelineProcessorArrayMapProcessorProcessorsCategoryProcessorOutputReference {
	var returns LogsCustomPipelineProcessorArrayMapProcessorProcessorsCategoryProcessorOutputReference
	_jsii_.Get(
		j,
		"categoryProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) CategoryProcessorInput() *LogsCustomPipelineProcessorArrayMapProcessorProcessorsCategoryProcessor {
	var returns *LogsCustomPipelineProcessorArrayMapProcessorProcessorsCategoryProcessor
	_jsii_.Get(
		j,
		"categoryProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) StringBuilderProcessor() LogsCustomPipelineProcessorArrayMapProcessorProcessorsStringBuilderProcessorOutputReference {
	var returns LogsCustomPipelineProcessorArrayMapProcessorProcessorsStringBuilderProcessorOutputReference
	_jsii_.Get(
		j,
		"stringBuilderProcessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) StringBuilderProcessorInput() *LogsCustomPipelineProcessorArrayMapProcessorProcessorsStringBuilderProcessor {
	var returns *LogsCustomPipelineProcessorArrayMapProcessorProcessorsStringBuilderProcessor
	_jsii_.Get(
		j,
		"stringBuilderProcessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference {
	_init_.Initialize()

	if err := validateNewLogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference_Override(l LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) PutArithmeticProcessor(value *LogsCustomPipelineProcessorArrayMapProcessorProcessorsArithmeticProcessor) {
	if err := l.validatePutArithmeticProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putArithmeticProcessor",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) PutAttributeRemapper(value *LogsCustomPipelineProcessorArrayMapProcessorProcessorsAttributeRemapper) {
	if err := l.validatePutAttributeRemapperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAttributeRemapper",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) PutCategoryProcessor(value *LogsCustomPipelineProcessorArrayMapProcessorProcessorsCategoryProcessor) {
	if err := l.validatePutCategoryProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCategoryProcessor",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) PutStringBuilderProcessor(value *LogsCustomPipelineProcessorArrayMapProcessorProcessorsStringBuilderProcessor) {
	if err := l.validatePutStringBuilderProcessorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStringBuilderProcessor",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) ResetArithmeticProcessor() {
	_jsii_.InvokeVoid(
		l,
		"resetArithmeticProcessor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) ResetAttributeRemapper() {
	_jsii_.InvokeVoid(
		l,
		"resetAttributeRemapper",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) ResetCategoryProcessor() {
	_jsii_.InvokeVoid(
		l,
		"resetCategoryProcessor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) ResetStringBuilderProcessor() {
	_jsii_.InvokeVoid(
		l,
		"resetStringBuilderProcessor",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayMapProcessorProcessorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

