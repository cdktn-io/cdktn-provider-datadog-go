// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package dashboardv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validatePutCellDisplayModeOptionsParameters(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaCellDisplayModeOptions) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validatePutConditionalFormatsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaConditionalFormats:
		value := value.(*[]*DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaConditionalFormats)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaConditionalFormats:
		value_ := value.([]*DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaConditionalFormats)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaConditionalFormats; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validatePutLimitParameters(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaLimit) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validatePutNumberFormatParameters(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaNumberFormat) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validatePutStyleParameters(value *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaStyle) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateSetAliasParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateSetCellDisplayModeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateSetFormulaExpressionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormula:
		val := val.(*DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormula)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormula:
		val_ := val.(DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormula)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormula; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDashboardV2WidgetGroupDefinitionWidgetWildcardDefinitionRequestTimeseriesRequestFormulaOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

