// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package powerpackv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegment:
		val := val.(*PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegment)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegment:
		val_ := val.(PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegment)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegment; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateSetSegmentIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPowerpackV2WidgetScatterplotDefinitionRequestScatterplotTableQueryUserJourneyQuerySearchFiltersAudienceFiltersSegmentOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

