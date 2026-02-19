// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package deploymentgate

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeploymentGateRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGateRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeploymentGateRuleList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeploymentGateRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentGateRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeploymentGateRuleList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeploymentGateRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeploymentGateRuleListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

