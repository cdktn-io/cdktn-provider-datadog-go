// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/securitymonitoringrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference interface {
	cdktn.ComplexObject
	BucketDuration() *float64
	SetBucketDuration(val *float64)
	BucketDurationInput() *float64
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
	DetectionTolerance() *float64
	SetDetectionTolerance(val *float64)
	DetectionToleranceInput() *float64
	// Experimental.
	Fqn() *string
	InstantaneousBaseline() interface{}
	SetInstantaneousBaseline(val interface{})
	InstantaneousBaselineInput() interface{}
	InternalValue() *SecurityMonitoringRuleOptionsAnomalyDetectionOptions
	SetInternalValue(val *SecurityMonitoringRuleOptionsAnomalyDetectionOptions)
	LearningDuration() *float64
	SetLearningDuration(val *float64)
	LearningDurationInput() *float64
	LearningPeriodBaseline() *float64
	SetLearningPeriodBaseline(val *float64)
	LearningPeriodBaselineInput() *float64
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
	ResetBucketDuration()
	ResetDetectionTolerance()
	ResetInstantaneousBaseline()
	ResetLearningDuration()
	ResetLearningPeriodBaseline()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference
type jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) BucketDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bucketDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) BucketDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bucketDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) DetectionTolerance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"detectionTolerance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) DetectionToleranceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"detectionToleranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) InstantaneousBaseline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instantaneousBaseline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) InstantaneousBaselineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instantaneousBaselineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) InternalValue() *SecurityMonitoringRuleOptionsAnomalyDetectionOptions {
	var returns *SecurityMonitoringRuleOptionsAnomalyDetectionOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) LearningDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"learningDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) LearningDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"learningDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) LearningPeriodBaseline() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"learningPeriodBaseline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) LearningPeriodBaselineInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"learningPeriodBaselineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.securityMonitoringRule.SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference_Override(s SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.securityMonitoringRule.SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference)SetBucketDuration(val *float64) {
	if err := j.validateSetBucketDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketDuration",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference)SetDetectionTolerance(val *float64) {
	if err := j.validateSetDetectionToleranceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectionTolerance",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference)SetInstantaneousBaseline(val interface{}) {
	if err := j.validateSetInstantaneousBaselineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instantaneousBaseline",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference)SetInternalValue(val *SecurityMonitoringRuleOptionsAnomalyDetectionOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference)SetLearningDuration(val *float64) {
	if err := j.validateSetLearningDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"learningDuration",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference)SetLearningPeriodBaseline(val *float64) {
	if err := j.validateSetLearningPeriodBaselineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"learningPeriodBaseline",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) ResetBucketDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetBucketDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) ResetDetectionTolerance() {
	_jsii_.InvokeVoid(
		s,
		"resetDetectionTolerance",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) ResetInstantaneousBaseline() {
	_jsii_.InvokeVoid(
		s,
		"resetInstantaneousBaseline",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) ResetLearningDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetLearningDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) ResetLearningPeriodBaseline() {
	_jsii_.InvokeVoid(
		s,
		"resetLearningPeriodBaseline",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityMonitoringRuleOptionsAnomalyDetectionOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

