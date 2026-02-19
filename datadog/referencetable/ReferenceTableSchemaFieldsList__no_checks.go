// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package referencetable

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReferenceTableSchemaFieldsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_ReferenceTableSchemaFieldsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReferenceTableSchemaFieldsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReferenceTableSchemaFieldsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ReferenceTableSchemaFieldsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReferenceTableSchemaFieldsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReferenceTableSchemaFieldsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReferenceTableSchemaFieldsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

