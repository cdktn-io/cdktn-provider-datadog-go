// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package costbudget

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CostBudgetBudgetLineList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CostBudgetBudgetLineList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CostBudgetBudgetLineList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CostBudgetBudgetLineList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CostBudgetBudgetLineList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CostBudgetBudgetLineList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CostBudgetBudgetLineList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCostBudgetBudgetLineListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

