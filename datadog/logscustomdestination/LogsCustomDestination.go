// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logscustomdestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/logscustomdestination/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/logs_custom_destination datadog_logs_custom_destination}.
type LogsCustomDestination interface {
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
	ElasticsearchDestination() LogsCustomDestinationElasticsearchDestinationList
	ElasticsearchDestinationInput() interface{}
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	ForwardTags() interface{}
	SetForwardTags(val interface{})
	ForwardTagsInput() interface{}
	ForwardTagsRestrictionList() *[]*string
	SetForwardTagsRestrictionList(val *[]*string)
	ForwardTagsRestrictionListInput() *[]*string
	ForwardTagsRestrictionListType() *string
	SetForwardTagsRestrictionListType(val *string)
	ForwardTagsRestrictionListTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpDestination() LogsCustomDestinationHttpDestinationList
	HttpDestinationInput() interface{}
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MicrosoftSentinelDestination() LogsCustomDestinationMicrosoftSentinelDestinationList
	MicrosoftSentinelDestinationInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	// Experimental.
	RawOverrides() interface{}
	SplunkDestination() LogsCustomDestinationSplunkDestinationList
	SplunkDestinationInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutElasticsearchDestination(value interface{})
	PutHttpDestination(value interface{})
	PutMicrosoftSentinelDestination(value interface{})
	PutSplunkDestination(value interface{})
	ResetElasticsearchDestination()
	ResetEnabled()
	ResetForwardTags()
	ResetForwardTagsRestrictionList()
	ResetForwardTagsRestrictionListType()
	ResetHttpDestination()
	ResetMicrosoftSentinelDestination()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQuery()
	ResetSplunkDestination()
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

// The jsii proxy struct for LogsCustomDestination
type jsiiProxy_LogsCustomDestination struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_LogsCustomDestination) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) ElasticsearchDestination() LogsCustomDestinationElasticsearchDestinationList {
	var returns LogsCustomDestinationElasticsearchDestinationList
	_jsii_.Get(
		j,
		"elasticsearchDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) ElasticsearchDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) ForwardTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) ForwardTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) ForwardTagsRestrictionList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forwardTagsRestrictionList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) ForwardTagsRestrictionListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forwardTagsRestrictionListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) ForwardTagsRestrictionListType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardTagsRestrictionListType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) ForwardTagsRestrictionListTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardTagsRestrictionListTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) HttpDestination() LogsCustomDestinationHttpDestinationList {
	var returns LogsCustomDestinationHttpDestinationList
	_jsii_.Get(
		j,
		"httpDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) HttpDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) MicrosoftSentinelDestination() LogsCustomDestinationMicrosoftSentinelDestinationList {
	var returns LogsCustomDestinationMicrosoftSentinelDestinationList
	_jsii_.Get(
		j,
		"microsoftSentinelDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) MicrosoftSentinelDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftSentinelDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) SplunkDestination() LogsCustomDestinationSplunkDestinationList {
	var returns LogsCustomDestinationSplunkDestinationList
	_jsii_.Get(
		j,
		"splunkDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) SplunkDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogsCustomDestination) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/logs_custom_destination datadog_logs_custom_destination} Resource.
func NewLogsCustomDestination(scope constructs.Construct, id *string, config *LogsCustomDestinationConfig) LogsCustomDestination {
	_init_.Initialize()

	if err := validateNewLogsCustomDestinationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogsCustomDestination{}

	_jsii_.Create(
		"@cdktn/provider-datadog.logsCustomDestination.LogsCustomDestination",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/logs_custom_destination datadog_logs_custom_destination} Resource.
func NewLogsCustomDestination_Override(l LogsCustomDestination, scope constructs.Construct, id *string, config *LogsCustomDestinationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.logsCustomDestination.LogsCustomDestination",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LogsCustomDestination)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LogsCustomDestination)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LogsCustomDestination)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LogsCustomDestination)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_LogsCustomDestination)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LogsCustomDestination)SetForwardTags(val interface{}) {
	if err := j.validateSetForwardTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardTags",
		val,
	)
}

func (j *jsiiProxy_LogsCustomDestination)SetForwardTagsRestrictionList(val *[]*string) {
	if err := j.validateSetForwardTagsRestrictionListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardTagsRestrictionList",
		val,
	)
}

func (j *jsiiProxy_LogsCustomDestination)SetForwardTagsRestrictionListType(val *string) {
	if err := j.validateSetForwardTagsRestrictionListTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardTagsRestrictionListType",
		val,
	)
}

func (j *jsiiProxy_LogsCustomDestination)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LogsCustomDestination)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LogsCustomDestination)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LogsCustomDestination)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LogsCustomDestination)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

// Generates CDKTN code for importing a LogsCustomDestination resource upon running "cdktn plan <stack-name>".
func LogsCustomDestination_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateLogsCustomDestination_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.logsCustomDestination.LogsCustomDestination",
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
func LogsCustomDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogsCustomDestination_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.logsCustomDestination.LogsCustomDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogsCustomDestination_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogsCustomDestination_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.logsCustomDestination.LogsCustomDestination",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogsCustomDestination_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogsCustomDestination_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.logsCustomDestination.LogsCustomDestination",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LogsCustomDestination_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-datadog.logsCustomDestination.LogsCustomDestination",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LogsCustomDestination) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LogsCustomDestination) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LogsCustomDestination) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LogsCustomDestination) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogsCustomDestination) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LogsCustomDestination) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogsCustomDestination) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LogsCustomDestination) PutElasticsearchDestination(value interface{}) {
	if err := l.validatePutElasticsearchDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putElasticsearchDestination",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomDestination) PutHttpDestination(value interface{}) {
	if err := l.validatePutHttpDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHttpDestination",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomDestination) PutMicrosoftSentinelDestination(value interface{}) {
	if err := l.validatePutMicrosoftSentinelDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMicrosoftSentinelDestination",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomDestination) PutSplunkDestination(value interface{}) {
	if err := l.validatePutSplunkDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSplunkDestination",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogsCustomDestination) ResetElasticsearchDestination() {
	_jsii_.InvokeVoid(
		l,
		"resetElasticsearchDestination",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomDestination) ResetEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomDestination) ResetForwardTags() {
	_jsii_.InvokeVoid(
		l,
		"resetForwardTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomDestination) ResetForwardTagsRestrictionList() {
	_jsii_.InvokeVoid(
		l,
		"resetForwardTagsRestrictionList",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomDestination) ResetForwardTagsRestrictionListType() {
	_jsii_.InvokeVoid(
		l,
		"resetForwardTagsRestrictionListType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomDestination) ResetHttpDestination() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpDestination",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomDestination) ResetMicrosoftSentinelDestination() {
	_jsii_.InvokeVoid(
		l,
		"resetMicrosoftSentinelDestination",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomDestination) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomDestination) ResetQuery() {
	_jsii_.InvokeVoid(
		l,
		"resetQuery",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomDestination) ResetSplunkDestination() {
	_jsii_.InvokeVoid(
		l,
		"resetSplunkDestination",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogsCustomDestination) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogsCustomDestination) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		l,
		"with",
		args,
		&returns,
	)

	return returns
}

