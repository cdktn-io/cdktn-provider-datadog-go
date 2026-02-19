// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package syntheticstest

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsTestConfigVariableList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SyntheticsTestConfigVariableList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsTestConfigVariableList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestConfigVariableList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestConfigVariableList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestConfigVariableList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SyntheticsTestConfigVariableList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSyntheticsTestConfigVariableListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

