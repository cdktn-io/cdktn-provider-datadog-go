// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package compliancecustomframework

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComplianceCustomFrameworkRequirementsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ComplianceCustomFrameworkRequirementsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ComplianceCustomFrameworkRequirementsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ComplianceCustomFrameworkRequirementsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ComplianceCustomFrameworkRequirementsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ComplianceCustomFrameworkRequirementsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ComplianceCustomFrameworkRequirementsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewComplianceCustomFrameworkRequirementsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

