// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package childorganization

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChildOrganizationApiKeyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ChildOrganizationApiKeyList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ChildOrganizationApiKeyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ChildOrganizationApiKeyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ChildOrganizationApiKeyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ChildOrganizationApiKeyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewChildOrganizationApiKeyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

