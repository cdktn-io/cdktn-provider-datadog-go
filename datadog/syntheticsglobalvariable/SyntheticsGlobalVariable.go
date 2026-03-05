// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticsglobalvariable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/syntheticsglobalvariable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/synthetics_global_variable datadog_synthetics_global_variable}.
type SyntheticsGlobalVariable interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IsFido() interface{}
	SetIsFido(val interface{})
	IsFidoInput() interface{}
	IsTotp() interface{}
	SetIsTotp(val interface{})
	IsTotpInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Options() SyntheticsGlobalVariableOptionsList
	OptionsInput() interface{}
	ParseTestId() *string
	SetParseTestId(val *string)
	ParseTestIdInput() *string
	ParseTestOptions() SyntheticsGlobalVariableParseTestOptionsList
	ParseTestOptionsInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RestrictedRoles() *[]*string
	SetRestrictedRoles(val *[]*string)
	RestrictedRolesInput() *[]*string
	Secure() interface{}
	SetSecure(val interface{})
	SecureInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
	ValueWo() *string
	SetValueWo(val *string)
	ValueWoInput() *string
	ValueWoVersion() *string
	SetValueWoVersion(val *string)
	ValueWoVersionInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutOptions(value interface{})
	PutParseTestOptions(value interface{})
	ResetDescription()
	ResetIsFido()
	ResetIsTotp()
	ResetOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParseTestId()
	ResetParseTestOptions()
	ResetRestrictedRoles()
	ResetSecure()
	ResetTags()
	ResetValue()
	ResetValueWo()
	ResetValueWoVersion()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for SyntheticsGlobalVariable
type jsiiProxy_SyntheticsGlobalVariable struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SyntheticsGlobalVariable) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) IsFido() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFido",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) IsFidoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFidoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) IsTotp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTotp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) IsTotpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTotpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Options() SyntheticsGlobalVariableOptionsList {
	var returns SyntheticsGlobalVariableOptionsList
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) OptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) ParseTestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parseTestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) ParseTestIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parseTestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) ParseTestOptions() SyntheticsGlobalVariableParseTestOptionsList {
	var returns SyntheticsGlobalVariableParseTestOptionsList
	_jsii_.Get(
		j,
		"parseTestOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) ParseTestOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseTestOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) RestrictedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) RestrictedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Secure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) SecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) ValueWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) ValueWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) ValueWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsGlobalVariable) ValueWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueWoVersionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/synthetics_global_variable datadog_synthetics_global_variable} Resource.
func NewSyntheticsGlobalVariable(scope constructs.Construct, id *string, config *SyntheticsGlobalVariableConfig) SyntheticsGlobalVariable {
	_init_.Initialize()

	if err := validateNewSyntheticsGlobalVariableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsGlobalVariable{}

	_jsii_.Create(
		"@cdktn/provider-datadog.syntheticsGlobalVariable.SyntheticsGlobalVariable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/synthetics_global_variable datadog_synthetics_global_variable} Resource.
func NewSyntheticsGlobalVariable_Override(s SyntheticsGlobalVariable, scope constructs.Construct, id *string, config *SyntheticsGlobalVariableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.syntheticsGlobalVariable.SyntheticsGlobalVariable",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetIsFido(val interface{}) {
	if err := j.validateSetIsFidoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isFido",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetIsTotp(val interface{}) {
	if err := j.validateSetIsTotpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isTotp",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetParseTestId(val *string) {
	if err := j.validateSetParseTestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parseTestId",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetRestrictedRoles(val *[]*string) {
	if err := j.validateSetRestrictedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictedRoles",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetSecure(val interface{}) {
	if err := j.validateSetSecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secure",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetValueWo(val *string) {
	if err := j.validateSetValueWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueWo",
		val,
	)
}

func (j *jsiiProxy_SyntheticsGlobalVariable)SetValueWoVersion(val *string) {
	if err := j.validateSetValueWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueWoVersion",
		val,
	)
}

// Generates CDKTN code for importing a SyntheticsGlobalVariable resource upon running "cdktn plan <stack-name>".
func SyntheticsGlobalVariable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSyntheticsGlobalVariable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.syntheticsGlobalVariable.SyntheticsGlobalVariable",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SyntheticsGlobalVariable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsGlobalVariable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.syntheticsGlobalVariable.SyntheticsGlobalVariable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsGlobalVariable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsGlobalVariable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.syntheticsGlobalVariable.SyntheticsGlobalVariable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsGlobalVariable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsGlobalVariable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.syntheticsGlobalVariable.SyntheticsGlobalVariable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SyntheticsGlobalVariable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-datadog.syntheticsGlobalVariable.SyntheticsGlobalVariable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) PutOptions(value interface{}) {
	if err := s.validatePutOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) PutParseTestOptions(value interface{}) {
	if err := s.validatePutParseTestOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putParseTestOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ResetIsFido() {
	_jsii_.InvokeVoid(
		s,
		"resetIsFido",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ResetIsTotp() {
	_jsii_.InvokeVoid(
		s,
		"resetIsTotp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ResetOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ResetParseTestId() {
	_jsii_.InvokeVoid(
		s,
		"resetParseTestId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ResetParseTestOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetParseTestOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ResetRestrictedRoles() {
	_jsii_.InvokeVoid(
		s,
		"resetRestrictedRoles",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ResetSecure() {
	_jsii_.InvokeVoid(
		s,
		"resetSecure",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ResetValue() {
	_jsii_.InvokeVoid(
		s,
		"resetValue",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ResetValueWo() {
	_jsii_.InvokeVoid(
		s,
		"resetValueWo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ResetValueWoVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetValueWoVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsGlobalVariable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsGlobalVariable) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

