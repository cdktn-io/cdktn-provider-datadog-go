// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datastore

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Datastore) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateMoveToIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Datastore) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateDatastore_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDatastore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatastore_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDatastore_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Datastore) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Datastore) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Datastore) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Datastore) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Datastore) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Datastore) validateSetOrgAccessParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Datastore) validateSetPrimaryColumnNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Datastore) validateSetPrimaryKeyGenerationStrategyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Datastore) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewDatastoreParameters(scope constructs.Construct, id *string, config *DatastoreConfig) error {
	return nil
}

