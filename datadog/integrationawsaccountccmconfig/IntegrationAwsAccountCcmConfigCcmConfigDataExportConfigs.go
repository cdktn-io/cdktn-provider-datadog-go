// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccountccmconfig


type IntegrationAwsAccountCcmConfigCcmConfigDataExportConfigs struct {
	// Name of the S3 bucket where the Cost and Usage Report is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/integration_aws_account_ccm_config#bucket_name IntegrationAwsAccountCcmConfig#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// AWS region of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/integration_aws_account_ccm_config#bucket_region IntegrationAwsAccountCcmConfig#bucket_region}
	BucketRegion *string `field:"optional" json:"bucketRegion" yaml:"bucketRegion"`
	// Name of the Cost and Usage Report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/integration_aws_account_ccm_config#report_name IntegrationAwsAccountCcmConfig#report_name}
	ReportName *string `field:"optional" json:"reportName" yaml:"reportName"`
	// S3 prefix where the Cost and Usage Report is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/integration_aws_account_ccm_config#report_prefix IntegrationAwsAccountCcmConfig#report_prefix}
	ReportPrefix *string `field:"optional" json:"reportPrefix" yaml:"reportPrefix"`
	// Type of the Cost and Usage Report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/integration_aws_account_ccm_config#report_type IntegrationAwsAccountCcmConfig#report_type}
	ReportType *string `field:"optional" json:"reportType" yaml:"reportType"`
}

