// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logsarchive


type LogsArchiveS3Archive struct {
	// Name of your s3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_archive#bucket LogsArchive#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Your AWS access key id, used as an alternative to `account_id`/`role_name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_archive#access_key_id LogsArchive#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Your AWS account id. Required with `role_name`; mutually exclusive with `access_key_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_archive#account_id LogsArchive#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The AWS KMS encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_archive#encryption_key LogsArchive#encryption_key}
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The type of encryption on your archive. Valid values are `NO_OVERRIDE`, `SSE_S3`, `SSE_KMS`. Defaults to `"NO_OVERRIDE"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_archive#encryption_type LogsArchive#encryption_type}
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// Path where the archive is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_archive#path LogsArchive#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Your AWS role name. Required with `account_id`; mutually exclusive with `access_key_id`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_archive#role_name LogsArchive#role_name}
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// The AWS S3 storage class used to upload the logs.
	//
	// Valid values are `STANDARD`, `STANDARD_IA`, `ONEZONE_IA`, `INTELLIGENT_TIERING`, `GLACIER_IR`. Defaults to `"STANDARD"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.17.0/docs/resources/logs_archive#storage_class LogsArchive#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

