// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package powerpackv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validatePutComputeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryCompute:
		value := value.(*[]*PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryCompute)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryCompute:
		value_ := value.([]*PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryCompute)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryCompute; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validatePutGroupByParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryGroupBy:
		value := value.(*[]*PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryGroupBy)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryGroupBy:
		value_ := value.([]*PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryGroupBy)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryGroupBy; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validatePutGroupByFieldsParameters(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryGroupByFields) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validatePutSearchParameters(value *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQuerySearch) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateSetCrossOrgUuidsParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateSetDataSourceParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateSetIndexesParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateSetInternalValueParameters(val *PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQuery) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateSetStorageParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPowerpackV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestYQueryEventQueryOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

