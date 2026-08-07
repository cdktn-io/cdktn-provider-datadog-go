// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sensitivedatascannerrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/sensitivedatascannerrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SensitiveDataScannerRuleSuppressionsOutputReference interface {
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
	EndsWith() *[]*string
	SetEndsWith(val *[]*string)
	EndsWithInput() *[]*string
	ExactMatch() *[]*string
	SetExactMatch(val *[]*string)
	ExactMatchInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SensitiveDataScannerRuleSuppressions
	SetInternalValue(val *SensitiveDataScannerRuleSuppressions)
	StartsWith() *[]*string
	SetStartsWith(val *[]*string)
	StartsWithInput() *[]*string
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
	ResetEndsWith()
	ResetExactMatch()
	ResetStartsWith()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SensitiveDataScannerRuleSuppressionsOutputReference
type jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) EndsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) EndsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"endsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) ExactMatch() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exactMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) ExactMatchInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exactMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) InternalValue() *SensitiveDataScannerRuleSuppressions {
	var returns *SensitiveDataScannerRuleSuppressions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) StartsWith() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"startsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) StartsWithInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"startsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSensitiveDataScannerRuleSuppressionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SensitiveDataScannerRuleSuppressionsOutputReference {
	_init_.Initialize()

	if err := validateNewSensitiveDataScannerRuleSuppressionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.sensitiveDataScannerRule.SensitiveDataScannerRuleSuppressionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSensitiveDataScannerRuleSuppressionsOutputReference_Override(s SensitiveDataScannerRuleSuppressionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.sensitiveDataScannerRule.SensitiveDataScannerRuleSuppressionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference)SetEndsWith(val *[]*string) {
	if err := j.validateSetEndsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endsWith",
		val,
	)
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference)SetExactMatch(val *[]*string) {
	if err := j.validateSetExactMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exactMatch",
		val,
	)
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference)SetInternalValue(val *SensitiveDataScannerRuleSuppressions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference)SetStartsWith(val *[]*string) {
	if err := j.validateSetStartsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startsWith",
		val,
	)
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) ResetEndsWith() {
	_jsii_.InvokeVoid(
		s,
		"resetEndsWith",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) ResetExactMatch() {
	_jsii_.InvokeVoid(
		s,
		"resetExactMatch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) ResetStartsWith() {
	_jsii_.InvokeVoid(
		s,
		"resetStartsWith",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SensitiveDataScannerRuleSuppressionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

