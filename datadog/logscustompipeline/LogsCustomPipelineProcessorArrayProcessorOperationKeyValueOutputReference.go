// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/logscustompipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference interface {
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
	InternalValue() *LogsCustomPipelineProcessorArrayProcessorOperationKeyValue
	SetInternalValue(val *LogsCustomPipelineProcessorArrayProcessorOperationKeyValue)
	KeyToExtract() *string
	SetKeyToExtract(val *string)
	KeyToExtractInput() *string
	OverrideOnConflict() interface{}
	SetOverrideOnConflict(val interface{})
	OverrideOnConflictInput() interface{}
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	Target() *string
	SetTarget(val *string)
	TargetInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ValueToExtract() *string
	SetValueToExtract(val *string)
	ValueToExtractInput() *string
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
	ResetOverrideOnConflict()
	ResetTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference
type jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) InternalValue() *LogsCustomPipelineProcessorArrayProcessorOperationKeyValue {
	var returns *LogsCustomPipelineProcessorArrayProcessorOperationKeyValue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) KeyToExtract() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyToExtract",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) KeyToExtractInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyToExtractInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) OverrideOnConflict() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideOnConflict",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) OverrideOnConflictInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideOnConflictInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) ValueToExtract() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueToExtract",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) ValueToExtractInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueToExtractInput",
		&returns,
	)
	return returns
}


func NewLogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference {
	_init_.Initialize()

	if err := validateNewLogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference_Override(l LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.logsCustomPipeline.LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference)SetInternalValue(val *LogsCustomPipelineProcessorArrayProcessorOperationKeyValue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference)SetKeyToExtract(val *string) {
	if err := j.validateSetKeyToExtractParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyToExtract",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference)SetOverrideOnConflict(val interface{}) {
	if err := j.validateSetOverrideOnConflictParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideOnConflict",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference)SetValueToExtract(val *string) {
	if err := j.validateSetValueToExtractParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueToExtract",
		val,
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) ResetOverrideOnConflict() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideOnConflict",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		l,
		"resetTarget",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (l *jsiiProxy_LogsCustomPipelineProcessorArrayProcessorOperationKeyValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

