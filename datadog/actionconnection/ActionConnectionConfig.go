// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ActionConnectionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#name ActionConnection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// anthropic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#anthropic ActionConnection#anthropic}
	Anthropic *ActionConnectionAnthropic `field:"optional" json:"anthropic" yaml:"anthropic"`
	// asana block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#asana ActionConnection#asana}
	Asana *ActionConnectionAsana `field:"optional" json:"asana" yaml:"asana"`
	// aws block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#aws ActionConnection#aws}
	Aws *ActionConnectionAws `field:"optional" json:"aws" yaml:"aws"`
	// azure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#azure ActionConnection#azure}
	Azure *ActionConnectionAzure `field:"optional" json:"azure" yaml:"azure"`
	// circle_ci block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#circle_ci ActionConnection#circle_ci}
	CircleCi *ActionConnectionCircleCi `field:"optional" json:"circleCi" yaml:"circleCi"`
	// clickup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#clickup ActionConnection#clickup}
	Clickup *ActionConnectionClickup `field:"optional" json:"clickup" yaml:"clickup"`
	// cloudflare block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#cloudflare ActionConnection#cloudflare}
	Cloudflare *ActionConnectionCloudflare `field:"optional" json:"cloudflare" yaml:"cloudflare"`
	// config_cat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#config_cat ActionConnection#config_cat}
	ConfigCat *ActionConnectionConfigCat `field:"optional" json:"configCat" yaml:"configCat"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#datadog ActionConnection#datadog}
	Datadog *ActionConnectionDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// fastly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#fastly ActionConnection#fastly}
	Fastly *ActionConnectionFastly `field:"optional" json:"fastly" yaml:"fastly"`
	// freshservice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#freshservice ActionConnection#freshservice}
	Freshservice *ActionConnectionFreshservice `field:"optional" json:"freshservice" yaml:"freshservice"`
	// gcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#gcp ActionConnection#gcp}
	Gcp *ActionConnectionGcp `field:"optional" json:"gcp" yaml:"gcp"`
	// gemini block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#gemini ActionConnection#gemini}
	Gemini *ActionConnectionGemini `field:"optional" json:"gemini" yaml:"gemini"`
	// gitlab block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#gitlab ActionConnection#gitlab}
	Gitlab *ActionConnectionGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
	// grey_noise block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#grey_noise ActionConnection#grey_noise}
	GreyNoise *ActionConnectionGreyNoise `field:"optional" json:"greyNoise" yaml:"greyNoise"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#http ActionConnection#http}
	Http *ActionConnectionHttp `field:"optional" json:"http" yaml:"http"`
	// launch_darkly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#launch_darkly ActionConnection#launch_darkly}
	LaunchDarkly *ActionConnectionLaunchDarkly `field:"optional" json:"launchDarkly" yaml:"launchDarkly"`
	// notion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#notion ActionConnection#notion}
	Notion *ActionConnectionNotion `field:"optional" json:"notion" yaml:"notion"`
	// okta block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#okta ActionConnection#okta}
	Okta *ActionConnectionOkta `field:"optional" json:"okta" yaml:"okta"`
	// openai block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#openai ActionConnection#openai}
	Openai *ActionConnectionOpenai `field:"optional" json:"openai" yaml:"openai"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#service_now ActionConnection#service_now}
	ServiceNow *ActionConnectionServiceNow `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// split block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#split ActionConnection#split}
	Split *ActionConnectionSplit `field:"optional" json:"split" yaml:"split"`
	// statsig block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#statsig ActionConnection#statsig}
	Statsig *ActionConnectionStatsig `field:"optional" json:"statsig" yaml:"statsig"`
	// virus_total block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection#virus_total ActionConnection#virus_total}
	VirusTotal *ActionConnectionVirusTotal `field:"optional" json:"virusTotal" yaml:"virusTotal"`
}

