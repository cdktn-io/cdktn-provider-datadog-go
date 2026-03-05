// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationazure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/integrationazure/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/integration_azure datadog_integration_azure}.
type IntegrationAzure interface {
	cdktn.TerraformResource
	AppServicePlanFilters() *string
	SetAppServicePlanFilters(val *string)
	AppServicePlanFiltersInput() *string
	Automute() interface{}
	SetAutomute(val interface{})
	AutomuteInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerAppFilters() *string
	SetContainerAppFilters(val *string)
	ContainerAppFiltersInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CspmEnabled() interface{}
	SetCspmEnabled(val interface{})
	CspmEnabledInput() interface{}
	CustomMetricsEnabled() interface{}
	SetCustomMetricsEnabled(val interface{})
	CustomMetricsEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostFilters() *string
	SetHostFilters(val *string)
	HostFiltersInput() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MetricsEnabled() interface{}
	SetMetricsEnabled(val interface{})
	MetricsEnabledDefault() interface{}
	SetMetricsEnabledDefault(val interface{})
	MetricsEnabledDefaultInput() interface{}
	MetricsEnabledInput() interface{}
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
	// Experimental.
	RawOverrides() interface{}
	ResourceCollectionEnabled() interface{}
	SetResourceCollectionEnabled(val interface{})
	ResourceCollectionEnabledInput() interface{}
	ResourceProviderConfigs() IntegrationAzureResourceProviderConfigsList
	ResourceProviderConfigsInput() interface{}
	TenantName() *string
	SetTenantName(val *string)
	TenantNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UsageMetricsEnabled() interface{}
	SetUsageMetricsEnabled(val interface{})
	UsageMetricsEnabledInput() interface{}
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
	PutResourceProviderConfigs(value interface{})
	ResetAppServicePlanFilters()
	ResetAutomute()
	ResetContainerAppFilters()
	ResetCspmEnabled()
	ResetCustomMetricsEnabled()
	ResetHostFilters()
	ResetMetricsEnabled()
	ResetMetricsEnabledDefault()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceCollectionEnabled()
	ResetResourceProviderConfigs()
	ResetUsageMetricsEnabled()
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

// The jsii proxy struct for IntegrationAzure
type jsiiProxy_IntegrationAzure struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_IntegrationAzure) AppServicePlanFilters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServicePlanFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) AppServicePlanFiltersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServicePlanFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) Automute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) AutomuteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) ContainerAppFilters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAppFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) ContainerAppFiltersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAppFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) CspmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cspmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) CspmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cspmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) CustomMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) CustomMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) HostFilters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) HostFiltersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) MetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) MetricsEnabledDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabledDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) MetricsEnabledDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabledDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) MetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) ResourceCollectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceCollectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) ResourceCollectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceCollectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) ResourceProviderConfigs() IntegrationAzureResourceProviderConfigsList {
	var returns IntegrationAzureResourceProviderConfigsList
	_jsii_.Get(
		j,
		"resourceProviderConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) ResourceProviderConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceProviderConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) TenantName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) TenantNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) UsageMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usageMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAzure) UsageMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usageMetricsEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/integration_azure datadog_integration_azure} Resource.
func NewIntegrationAzure(scope constructs.Construct, id *string, config *IntegrationAzureConfig) IntegrationAzure {
	_init_.Initialize()

	if err := validateNewIntegrationAzureParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationAzure{}

	_jsii_.Create(
		"@cdktn/provider-datadog.integrationAzure.IntegrationAzure",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/integration_azure datadog_integration_azure} Resource.
func NewIntegrationAzure_Override(i IntegrationAzure, scope constructs.Construct, id *string, config *IntegrationAzureConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.integrationAzure.IntegrationAzure",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetAppServicePlanFilters(val *string) {
	if err := j.validateSetAppServicePlanFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appServicePlanFilters",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetAutomute(val interface{}) {
	if err := j.validateSetAutomuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automute",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetContainerAppFilters(val *string) {
	if err := j.validateSetContainerAppFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerAppFilters",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetCspmEnabled(val interface{}) {
	if err := j.validateSetCspmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cspmEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetCustomMetricsEnabled(val interface{}) {
	if err := j.validateSetCustomMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetHostFilters(val *string) {
	if err := j.validateSetHostFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostFilters",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetMetricsEnabled(val interface{}) {
	if err := j.validateSetMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetMetricsEnabledDefault(val interface{}) {
	if err := j.validateSetMetricsEnabledDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsEnabledDefault",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetResourceCollectionEnabled(val interface{}) {
	if err := j.validateSetResourceCollectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceCollectionEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetTenantName(val *string) {
	if err := j.validateSetTenantNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantName",
		val,
	)
}

func (j *jsiiProxy_IntegrationAzure)SetUsageMetricsEnabled(val interface{}) {
	if err := j.validateSetUsageMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usageMetricsEnabled",
		val,
	)
}

// Generates CDKTN code for importing a IntegrationAzure resource upon running "cdktn plan <stack-name>".
func IntegrationAzure_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateIntegrationAzure_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationAzure.IntegrationAzure",
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
func IntegrationAzure_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationAzure_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationAzure.IntegrationAzure",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationAzure_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationAzure_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationAzure.IntegrationAzure",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationAzure_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationAzure_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationAzure.IntegrationAzure",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IntegrationAzure_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-datadog.integrationAzure.IntegrationAzure",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IntegrationAzure) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IntegrationAzure) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IntegrationAzure) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IntegrationAzure) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationAzure) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IntegrationAzure) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationAzure) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IntegrationAzure) PutResourceProviderConfigs(value interface{}) {
	if err := i.validatePutResourceProviderConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putResourceProviderConfigs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationAzure) ResetAppServicePlanFilters() {
	_jsii_.InvokeVoid(
		i,
		"resetAppServicePlanFilters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAzure) ResetAutomute() {
	_jsii_.InvokeVoid(
		i,
		"resetAutomute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAzure) ResetContainerAppFilters() {
	_jsii_.InvokeVoid(
		i,
		"resetContainerAppFilters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAzure) ResetCspmEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetCspmEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAzure) ResetCustomMetricsEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetCustomMetricsEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAzure) ResetHostFilters() {
	_jsii_.InvokeVoid(
		i,
		"resetHostFilters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAzure) ResetMetricsEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetMetricsEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAzure) ResetMetricsEnabledDefault() {
	_jsii_.InvokeVoid(
		i,
		"resetMetricsEnabledDefault",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAzure) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAzure) ResetResourceCollectionEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetResourceCollectionEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAzure) ResetResourceProviderConfigs() {
	_jsii_.InvokeVoid(
		i,
		"resetResourceProviderConfigs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAzure) ResetUsageMetricsEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetUsageMetricsEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAzure) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAzure) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

