// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationdatabricksaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/integrationdatabricksaccount/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account datadog_integration_databricks_account}.
type IntegrationDatabricksAccount interface {
	cdktn.TerraformResource
	AuthConfig() IntegrationDatabricksAccountAuthConfigOutputReference
	AuthConfigInput() interface{}
	CcmEnabled() interface{}
	SetCcmEnabled(val interface{})
	CcmEnabledInput() interface{}
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
	DdApiKeyId() *string
	SetDdApiKeyId(val *string)
	DdApiKeyIdInput() *string
	DdApiKeySecret() *string
	SetDdApiKeySecret(val *string)
	DdApiKeySecretInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DjmClusterPolicyEnabled() interface{}
	SetDjmClusterPolicyEnabled(val interface{})
	DjmClusterPolicyEnabledInput() interface{}
	DjmEnabled() interface{}
	SetDjmEnabled(val interface{})
	DjmEnabledInput() interface{}
	DjmGlobalInitScriptEnabled() interface{}
	SetDjmGlobalInitScriptEnabled(val interface{})
	DjmGlobalInitScriptEnabledInput() interface{}
	DoCrawlersCron() *string
	SetDoCrawlersCron(val *string)
	DoCrawlersCronInput() *string
	DoEnabled() interface{}
	SetDoEnabled(val interface{})
	DoEnabledInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ModelServingEndpointName() *string
	SetModelServingEndpointName(val *string)
	ModelServingEndpointNameInput() *string
	ModelServingMetricsEnabled() interface{}
	SetModelServingMetricsEnabled(val interface{})
	ModelServingMetricsEnabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PrivateActionRunnerConfiguration() IntegrationDatabricksAccountPrivateActionRunnerConfigurationOutputReference
	PrivateActionRunnerConfigurationInput() interface{}
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
	ScriptGpumEnabled() interface{}
	SetScriptGpumEnabled(val interface{})
	ScriptGpumEnabledInput() interface{}
	ScriptLogsEnabled() interface{}
	SetScriptLogsEnabled(val interface{})
	ScriptLogsEnabledInput() interface{}
	ServerlessJobsEnabled() interface{}
	SetServerlessJobsEnabled(val interface{})
	ServerlessJobsEnabledInput() interface{}
	SystemTablesSqlWarehouseId() *string
	SetSystemTablesSqlWarehouseId(val *string)
	SystemTablesSqlWarehouseIdInput() *string
	TableLineageEnabled() interface{}
	SetTableLineageEnabled(val interface{})
	TableLineageEnabledInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UcVolumePath() *string
	SetUcVolumePath(val *string)
	UcVolumePathInput() *string
	WorkspaceUrl() *string
	SetWorkspaceUrl(val *string)
	WorkspaceUrlInput() *string
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
	PutAuthConfig(value *IntegrationDatabricksAccountAuthConfig)
	PutPrivateActionRunnerConfiguration(value *IntegrationDatabricksAccountPrivateActionRunnerConfiguration)
	ResetAuthConfig()
	ResetCcmEnabled()
	ResetDdApiKeyId()
	ResetDdApiKeySecret()
	ResetDjmClusterPolicyEnabled()
	ResetDjmEnabled()
	ResetDjmGlobalInitScriptEnabled()
	ResetDoCrawlersCron()
	ResetDoEnabled()
	ResetModelServingEndpointName()
	ResetModelServingMetricsEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateActionRunnerConfiguration()
	ResetScriptGpumEnabled()
	ResetScriptLogsEnabled()
	ResetServerlessJobsEnabled()
	ResetSystemTablesSqlWarehouseId()
	ResetTableLineageEnabled()
	ResetUcVolumePath()
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

// The jsii proxy struct for IntegrationDatabricksAccount
type jsiiProxy_IntegrationDatabricksAccount struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_IntegrationDatabricksAccount) AuthConfig() IntegrationDatabricksAccountAuthConfigOutputReference {
	var returns IntegrationDatabricksAccountAuthConfigOutputReference
	_jsii_.Get(
		j,
		"authConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) AuthConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) CcmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ccmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) CcmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ccmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DdApiKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddApiKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DdApiKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddApiKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DdApiKeySecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddApiKeySecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DdApiKeySecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ddApiKeySecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DjmClusterPolicyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"djmClusterPolicyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DjmClusterPolicyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"djmClusterPolicyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DjmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"djmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DjmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"djmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DjmGlobalInitScriptEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"djmGlobalInitScriptEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DjmGlobalInitScriptEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"djmGlobalInitScriptEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DoCrawlersCron() *string {
	var returns *string
	_jsii_.Get(
		j,
		"doCrawlersCron",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DoCrawlersCronInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"doCrawlersCronInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DoEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"doEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) DoEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"doEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) ModelServingEndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelServingEndpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) ModelServingEndpointNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelServingEndpointNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) ModelServingMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelServingMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) ModelServingMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelServingMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) PrivateActionRunnerConfiguration() IntegrationDatabricksAccountPrivateActionRunnerConfigurationOutputReference {
	var returns IntegrationDatabricksAccountPrivateActionRunnerConfigurationOutputReference
	_jsii_.Get(
		j,
		"privateActionRunnerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) PrivateActionRunnerConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateActionRunnerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) ScriptGpumEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scriptGpumEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) ScriptGpumEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scriptGpumEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) ScriptLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scriptLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) ScriptLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scriptLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) ServerlessJobsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessJobsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) ServerlessJobsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessJobsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) SystemTablesSqlWarehouseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemTablesSqlWarehouseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) SystemTablesSqlWarehouseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemTablesSqlWarehouseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) TableLineageEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableLineageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) TableLineageEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableLineageEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) UcVolumePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ucVolumePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) UcVolumePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ucVolumePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) WorkspaceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationDatabricksAccount) WorkspaceUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceUrlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account datadog_integration_databricks_account} Resource.
func NewIntegrationDatabricksAccount(scope constructs.Construct, id *string, config *IntegrationDatabricksAccountConfig) IntegrationDatabricksAccount {
	_init_.Initialize()

	if err := validateNewIntegrationDatabricksAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationDatabricksAccount{}

	_jsii_.Create(
		"@cdktn/provider-datadog.integrationDatabricksAccount.IntegrationDatabricksAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account datadog_integration_databricks_account} Resource.
func NewIntegrationDatabricksAccount_Override(i IntegrationDatabricksAccount, scope constructs.Construct, id *string, config *IntegrationDatabricksAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.integrationDatabricksAccount.IntegrationDatabricksAccount",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetCcmEnabled(val interface{}) {
	if err := j.validateSetCcmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ccmEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetDdApiKeyId(val *string) {
	if err := j.validateSetDdApiKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ddApiKeyId",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetDdApiKeySecret(val *string) {
	if err := j.validateSetDdApiKeySecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ddApiKeySecret",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetDjmClusterPolicyEnabled(val interface{}) {
	if err := j.validateSetDjmClusterPolicyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"djmClusterPolicyEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetDjmEnabled(val interface{}) {
	if err := j.validateSetDjmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"djmEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetDjmGlobalInitScriptEnabled(val interface{}) {
	if err := j.validateSetDjmGlobalInitScriptEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"djmGlobalInitScriptEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetDoCrawlersCron(val *string) {
	if err := j.validateSetDoCrawlersCronParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"doCrawlersCron",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetDoEnabled(val interface{}) {
	if err := j.validateSetDoEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"doEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetModelServingEndpointName(val *string) {
	if err := j.validateSetModelServingEndpointNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelServingEndpointName",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetModelServingMetricsEnabled(val interface{}) {
	if err := j.validateSetModelServingMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelServingMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetScriptGpumEnabled(val interface{}) {
	if err := j.validateSetScriptGpumEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptGpumEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetScriptLogsEnabled(val interface{}) {
	if err := j.validateSetScriptLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetServerlessJobsEnabled(val interface{}) {
	if err := j.validateSetServerlessJobsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverlessJobsEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetSystemTablesSqlWarehouseId(val *string) {
	if err := j.validateSetSystemTablesSqlWarehouseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemTablesSqlWarehouseId",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetTableLineageEnabled(val interface{}) {
	if err := j.validateSetTableLineageEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableLineageEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetUcVolumePath(val *string) {
	if err := j.validateSetUcVolumePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ucVolumePath",
		val,
	)
}

func (j *jsiiProxy_IntegrationDatabricksAccount)SetWorkspaceUrl(val *string) {
	if err := j.validateSetWorkspaceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceUrl",
		val,
	)
}

// Generates CDKTN code for importing a IntegrationDatabricksAccount resource upon running "cdktn plan <stack-name>".
func IntegrationDatabricksAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateIntegrationDatabricksAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationDatabricksAccount.IntegrationDatabricksAccount",
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
func IntegrationDatabricksAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationDatabricksAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationDatabricksAccount.IntegrationDatabricksAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationDatabricksAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationDatabricksAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationDatabricksAccount.IntegrationDatabricksAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationDatabricksAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationDatabricksAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationDatabricksAccount.IntegrationDatabricksAccount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IntegrationDatabricksAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-datadog.integrationDatabricksAccount.IntegrationDatabricksAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IntegrationDatabricksAccount) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IntegrationDatabricksAccount) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IntegrationDatabricksAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IntegrationDatabricksAccount) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IntegrationDatabricksAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IntegrationDatabricksAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IntegrationDatabricksAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IntegrationDatabricksAccount) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IntegrationDatabricksAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IntegrationDatabricksAccount) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IntegrationDatabricksAccount) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) PutAuthConfig(value *IntegrationDatabricksAccountAuthConfig) {
	if err := i.validatePutAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAuthConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) PutPrivateActionRunnerConfiguration(value *IntegrationDatabricksAccountPrivateActionRunnerConfiguration) {
	if err := i.validatePutPrivateActionRunnerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPrivateActionRunnerConfiguration",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetAuthConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetCcmEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetCcmEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetDdApiKeyId() {
	_jsii_.InvokeVoid(
		i,
		"resetDdApiKeyId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetDdApiKeySecret() {
	_jsii_.InvokeVoid(
		i,
		"resetDdApiKeySecret",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetDjmClusterPolicyEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetDjmClusterPolicyEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetDjmEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetDjmEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetDjmGlobalInitScriptEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetDjmGlobalInitScriptEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetDoCrawlersCron() {
	_jsii_.InvokeVoid(
		i,
		"resetDoCrawlersCron",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetDoEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetDoEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetModelServingEndpointName() {
	_jsii_.InvokeVoid(
		i,
		"resetModelServingEndpointName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetModelServingMetricsEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetModelServingMetricsEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetPrivateActionRunnerConfiguration() {
	_jsii_.InvokeVoid(
		i,
		"resetPrivateActionRunnerConfiguration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetScriptGpumEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetScriptGpumEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetScriptLogsEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetScriptLogsEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetServerlessJobsEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetServerlessJobsEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetSystemTablesSqlWarehouseId() {
	_jsii_.InvokeVoid(
		i,
		"resetSystemTablesSqlWarehouseId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetTableLineageEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetTableLineageEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ResetUcVolumePath() {
	_jsii_.InvokeVoid(
		i,
		"resetUcVolumePath",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationDatabricksAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationDatabricksAccount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationDatabricksAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationDatabricksAccount) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

