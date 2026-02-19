// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package deploymentgate

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeploymentGate) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateMoveToIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (d *jsiiProxy_DeploymentGate) validatePutRuleParameters(value interface{}) error {
	return nil
}

func validateDeploymentGate_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDeploymentGate_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDeploymentGate_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDeploymentGate_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentGate) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentGate) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentGate) validateSetDryRunParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentGate) validateSetEnvParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeploymentGate) validateSetIdentifierParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeploymentGate) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_DeploymentGate) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentGate) validateSetServiceParameters(val *string) error {
	return nil
}

func validateNewDeploymentGateParameters(scope constructs.Construct, id *string, config *DeploymentGateConfig) error {
	return nil
}

