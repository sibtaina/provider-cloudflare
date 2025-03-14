// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type JobInitParameters struct {

	// (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) Name of the dataset. A list of supported datasets can be found on the Developer Docs.
	// Name of the dataset. A list of supported datasets can be found on the [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).
	Dataset *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// (String) Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included.
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included.
	DestinationConf *string `json:"destinationConf,omitempty" tf:"destination_conf,omitempty"`

	// (Boolean) Flag that indicates if the job is enabled.
	// Flag that indicates if the job is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) This field is deprecated. Please use max_upload_* parameters instead. The frequency at which Cloudflare sends batches of logs to your destination. Setting frequency to high sends your logs in larger quantities of smaller files. Setting frequency to low sends logs in smaller quantities of larger files.
	// This field is deprecated. Please use `max_upload_*` parameters instead. The frequency at which Cloudflare sends batches of logs to your destination. Setting frequency to high sends your logs in larger quantities of smaller files. Setting frequency to low sends logs in smaller quantities of larger files.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (String) The kind parameter  is used to differentiate between Logpush and Edge Log Delivery jobs. Currently, Edge Log Delivery is only supported for the http_requests dataset.
	// The kind parameter (optional) is used to differentiate between Logpush and Edge Log Delivery jobs. Currently, Edge Log Delivery is only supported for the `http_requests` dataset.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// (String) This field is deprecated. Use output_options instead. Configuration string. It specifies things like requested fields and timestamp formats. If migrating from the logpull api, copy the url (full url or just the query string) of your call here, and logpush will keep on making this call for you, setting start and end times appropriately.
	// This field is deprecated. Use `output_options` instead. Configuration string. It specifies things like requested fields and timestamp formats. If migrating from the logpull api, copy the url (full url or just the query string) of your call here, and logpush will keep on making this call for you, setting start and end times appropriately.
	LogpullOptions *string `json:"logpullOptions,omitempty" tf:"logpull_options,omitempty"`

	// (Number) The maximum uncompressed file size of a batch of logs. This setting value must be between 5 MB and 1 GB, or 0 to disable it. Note that you cannot set a minimum file size; this means that log files may be much smaller than this batch size. This parameter is not available for jobs with edge as its kind.
	// The maximum uncompressed file size of a batch of logs. This setting value must be between `5 MB` and `1 GB`, or `0` to disable it. Note that you cannot set a minimum file size; this means that log files may be much smaller than this batch size. This parameter is not available for jobs with `edge` as its kind.
	MaxUploadBytes *float64 `json:"maxUploadBytes,omitempty" tf:"max_upload_bytes,omitempty"`

	// (Number) The maximum interval in seconds for log batches. This setting must be between 30 and 300 seconds (5 minutes), or 0 to disable it. Note that you cannot specify a minimum interval for log batches; this means that log files may be sent in shorter intervals than this. This parameter is only used for jobs with edge as its kind.
	// The maximum interval in seconds for log batches. This setting must be between 30 and 300 seconds (5 minutes), or `0` to disable it. Note that you cannot specify a minimum interval for log batches; this means that log files may be sent in shorter intervals than this. This parameter is only used for jobs with `edge` as its kind.
	MaxUploadIntervalSeconds *float64 `json:"maxUploadIntervalSeconds,omitempty" tf:"max_upload_interval_seconds,omitempty"`

	// (Number) The maximum number of log lines per batch. This setting must be between 1000 and 1,000,000 lines, or 0 to disable it. Note that you cannot specify a minimum number of log lines per batch; this means that log files may contain many fewer lines than this. This parameter is not available for jobs with edge as its kind.
	// The maximum number of log lines per batch. This setting must be between 1000 and 1,000,000 lines, or `0` to disable it. Note that you cannot specify a minimum number of log lines per batch; this means that log files may contain many fewer lines than this. This parameter is not available for jobs with `edge` as its kind.
	MaxUploadRecords *float64 `json:"maxUploadRecords,omitempty" tf:"max_upload_records,omitempty"`

	// (Attributes) The structured replacement for logpull_options. When including this field, the logpull_option field will be ignored. (see below for nested schema)
	OutputOptions *OutputOptionsInitParameters `json:"outputOptions,omitempty" tf:"output_options,omitempty"`

	// (String) Ownership challenge token to prove destination ownership.
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge *string `json:"ownershipChallenge,omitempty" tf:"ownership_challenge,omitempty"`

	// (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type JobObservation struct {

	// (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) Name of the dataset. A list of supported datasets can be found on the Developer Docs.
	// Name of the dataset. A list of supported datasets can be found on the [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).
	Dataset *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// (String) Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included.
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included.
	DestinationConf *string `json:"destinationConf,omitempty" tf:"destination_conf,omitempty"`

	// (Boolean) Flag that indicates if the job is enabled.
	// Flag that indicates if the job is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) If not null, the job is currently failing. Failures are usually repetitive (example: no permissions to write to destination bucket). Only the last failure is recorded. On successful execution of a job the error_message and last_error are set to null.
	// If not null, the job is currently failing. Failures are usually repetitive (example: no permissions to write to destination bucket). Only the last failure is recorded. On successful execution of a job the error_message and last_error are set to null.
	ErrorMessage *string `json:"errorMessage,omitempty" tf:"error_message,omitempty"`

	// (String) This field is deprecated. Please use max_upload_* parameters instead. The frequency at which Cloudflare sends batches of logs to your destination. Setting frequency to high sends your logs in larger quantities of smaller files. Setting frequency to low sends logs in smaller quantities of larger files.
	// This field is deprecated. Please use `max_upload_*` parameters instead. The frequency at which Cloudflare sends batches of logs to your destination. Setting frequency to high sends your logs in larger quantities of smaller files. Setting frequency to low sends logs in smaller quantities of larger files.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (Number) Unique id of the job.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The kind parameter  is used to differentiate between Logpush and Edge Log Delivery jobs. Currently, Edge Log Delivery is only supported for the http_requests dataset.
	// The kind parameter (optional) is used to differentiate between Logpush and Edge Log Delivery jobs. Currently, Edge Log Delivery is only supported for the `http_requests` dataset.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// 07-23T10:00:00Z to 2018-07-23T10:01:00Z then the value of this field will be 2018-07-23T10:01:00Z. If the job has never run or has just been enabled and hasn't run yet then the field will be empty.
	// Records the last time for which logs have been successfully pushed. If the last successful push was for logs range 2018-07-23T10:00:00Z to 2018-07-23T10:01:00Z then the value of this field will be 2018-07-23T10:01:00Z. If the job has never run or has just been enabled and hasn't run yet then the field will be empty.
	LastComplete *string `json:"lastComplete,omitempty" tf:"last_complete,omitempty"`

	// (String) Records the last time the job failed. If not null, the job is currently failing. If null, the job has either never failed or has run successfully at least once since last failure. See also the error_message field.
	// Records the last time the job failed. If not null, the job is currently failing. If null, the job has either never failed or has run successfully at least once since last failure. See also the error_message field.
	LastError *string `json:"lastError,omitempty" tf:"last_error,omitempty"`

	// (String) This field is deprecated. Use output_options instead. Configuration string. It specifies things like requested fields and timestamp formats. If migrating from the logpull api, copy the url (full url or just the query string) of your call here, and logpush will keep on making this call for you, setting start and end times appropriately.
	// This field is deprecated. Use `output_options` instead. Configuration string. It specifies things like requested fields and timestamp formats. If migrating from the logpull api, copy the url (full url or just the query string) of your call here, and logpush will keep on making this call for you, setting start and end times appropriately.
	LogpullOptions *string `json:"logpullOptions,omitempty" tf:"logpull_options,omitempty"`

	// (Number) The maximum uncompressed file size of a batch of logs. This setting value must be between 5 MB and 1 GB, or 0 to disable it. Note that you cannot set a minimum file size; this means that log files may be much smaller than this batch size. This parameter is not available for jobs with edge as its kind.
	// The maximum uncompressed file size of a batch of logs. This setting value must be between `5 MB` and `1 GB`, or `0` to disable it. Note that you cannot set a minimum file size; this means that log files may be much smaller than this batch size. This parameter is not available for jobs with `edge` as its kind.
	MaxUploadBytes *float64 `json:"maxUploadBytes,omitempty" tf:"max_upload_bytes,omitempty"`

	// (Number) The maximum interval in seconds for log batches. This setting must be between 30 and 300 seconds (5 minutes), or 0 to disable it. Note that you cannot specify a minimum interval for log batches; this means that log files may be sent in shorter intervals than this. This parameter is only used for jobs with edge as its kind.
	// The maximum interval in seconds for log batches. This setting must be between 30 and 300 seconds (5 minutes), or `0` to disable it. Note that you cannot specify a minimum interval for log batches; this means that log files may be sent in shorter intervals than this. This parameter is only used for jobs with `edge` as its kind.
	MaxUploadIntervalSeconds *float64 `json:"maxUploadIntervalSeconds,omitempty" tf:"max_upload_interval_seconds,omitempty"`

	// (Number) The maximum number of log lines per batch. This setting must be between 1000 and 1,000,000 lines, or 0 to disable it. Note that you cannot specify a minimum number of log lines per batch; this means that log files may contain many fewer lines than this. This parameter is not available for jobs with edge as its kind.
	// The maximum number of log lines per batch. This setting must be between 1000 and 1,000,000 lines, or `0` to disable it. Note that you cannot specify a minimum number of log lines per batch; this means that log files may contain many fewer lines than this. This parameter is not available for jobs with `edge` as its kind.
	MaxUploadRecords *float64 `json:"maxUploadRecords,omitempty" tf:"max_upload_records,omitempty"`

	// (Attributes) The structured replacement for logpull_options. When including this field, the logpull_option field will be ignored. (see below for nested schema)
	OutputOptions *OutputOptionsObservation `json:"outputOptions,omitempty" tf:"output_options,omitempty"`

	// (String) Ownership challenge token to prove destination ownership.
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge *string `json:"ownershipChallenge,omitempty" tf:"ownership_challenge,omitempty"`

	// (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type JobParameters struct {

	// (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) Name of the dataset. A list of supported datasets can be found on the Developer Docs.
	// Name of the dataset. A list of supported datasets can be found on the [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).
	// +kubebuilder:validation:Optional
	Dataset *string `json:"dataset,omitempty" tf:"dataset,omitempty"`

	// (String) Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included.
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included.
	// +kubebuilder:validation:Optional
	DestinationConf *string `json:"destinationConf,omitempty" tf:"destination_conf,omitempty"`

	// (Boolean) Flag that indicates if the job is enabled.
	// Flag that indicates if the job is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) This field is deprecated. Please use max_upload_* parameters instead. The frequency at which Cloudflare sends batches of logs to your destination. Setting frequency to high sends your logs in larger quantities of smaller files. Setting frequency to low sends logs in smaller quantities of larger files.
	// This field is deprecated. Please use `max_upload_*` parameters instead. The frequency at which Cloudflare sends batches of logs to your destination. Setting frequency to high sends your logs in larger quantities of smaller files. Setting frequency to low sends logs in smaller quantities of larger files.
	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (String) The kind parameter  is used to differentiate between Logpush and Edge Log Delivery jobs. Currently, Edge Log Delivery is only supported for the http_requests dataset.
	// The kind parameter (optional) is used to differentiate between Logpush and Edge Log Delivery jobs. Currently, Edge Log Delivery is only supported for the `http_requests` dataset.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// (String) This field is deprecated. Use output_options instead. Configuration string. It specifies things like requested fields and timestamp formats. If migrating from the logpull api, copy the url (full url or just the query string) of your call here, and logpush will keep on making this call for you, setting start and end times appropriately.
	// This field is deprecated. Use `output_options` instead. Configuration string. It specifies things like requested fields and timestamp formats. If migrating from the logpull api, copy the url (full url or just the query string) of your call here, and logpush will keep on making this call for you, setting start and end times appropriately.
	// +kubebuilder:validation:Optional
	LogpullOptions *string `json:"logpullOptions,omitempty" tf:"logpull_options,omitempty"`

	// (Number) The maximum uncompressed file size of a batch of logs. This setting value must be between 5 MB and 1 GB, or 0 to disable it. Note that you cannot set a minimum file size; this means that log files may be much smaller than this batch size. This parameter is not available for jobs with edge as its kind.
	// The maximum uncompressed file size of a batch of logs. This setting value must be between `5 MB` and `1 GB`, or `0` to disable it. Note that you cannot set a minimum file size; this means that log files may be much smaller than this batch size. This parameter is not available for jobs with `edge` as its kind.
	// +kubebuilder:validation:Optional
	MaxUploadBytes *float64 `json:"maxUploadBytes,omitempty" tf:"max_upload_bytes,omitempty"`

	// (Number) The maximum interval in seconds for log batches. This setting must be between 30 and 300 seconds (5 minutes), or 0 to disable it. Note that you cannot specify a minimum interval for log batches; this means that log files may be sent in shorter intervals than this. This parameter is only used for jobs with edge as its kind.
	// The maximum interval in seconds for log batches. This setting must be between 30 and 300 seconds (5 minutes), or `0` to disable it. Note that you cannot specify a minimum interval for log batches; this means that log files may be sent in shorter intervals than this. This parameter is only used for jobs with `edge` as its kind.
	// +kubebuilder:validation:Optional
	MaxUploadIntervalSeconds *float64 `json:"maxUploadIntervalSeconds,omitempty" tf:"max_upload_interval_seconds,omitempty"`

	// (Number) The maximum number of log lines per batch. This setting must be between 1000 and 1,000,000 lines, or 0 to disable it. Note that you cannot specify a minimum number of log lines per batch; this means that log files may contain many fewer lines than this. This parameter is not available for jobs with edge as its kind.
	// The maximum number of log lines per batch. This setting must be between 1000 and 1,000,000 lines, or `0` to disable it. Note that you cannot specify a minimum number of log lines per batch; this means that log files may contain many fewer lines than this. This parameter is not available for jobs with `edge` as its kind.
	// +kubebuilder:validation:Optional
	MaxUploadRecords *float64 `json:"maxUploadRecords,omitempty" tf:"max_upload_records,omitempty"`

	// (Attributes) The structured replacement for logpull_options. When including this field, the logpull_option field will be ignored. (see below for nested schema)
	// +kubebuilder:validation:Optional
	OutputOptions *OutputOptionsParameters `json:"outputOptions,omitempty" tf:"output_options,omitempty"`

	// (String) Ownership challenge token to prove destination ownership.
	// Ownership challenge token to prove destination ownership.
	// +kubebuilder:validation:Optional
	OwnershipChallenge *string `json:"ownershipChallenge,omitempty" tf:"ownership_challenge,omitempty"`

	// (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type OutputOptionsInitParameters struct {

	// (String) String to be prepended before each batch.
	// String to be prepended before each batch.
	BatchPrefix *string `json:"batchPrefix,omitempty" tf:"batch_prefix,omitempty"`

	// (String) String to be appended after each batch.
	// String to be appended after each batch.
	BatchSuffix *string `json:"batchSuffix,omitempty" tf:"batch_suffix,omitempty"`

	// (Boolean) If set to true, will cause all occurrences of ${ in the generated files to be replaced with x{.
	// If set to true, will cause all occurrences of `${` in the generated files to be replaced with `x{`.
	Cve20214428 *bool `json:"cve20214428,omitempty" tf:"cve_2021_4428,omitempty"`

	// (String) String to join fields. This field be ignored when record_template is set.
	// String to join fields. This field be ignored when `record_template` is set.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// (List of String) List of field names to be included in the Logpush output. For the moment, there is no option to add all fields at once, so you must specify all the fields names you are interested in.
	// List of field names to be included in the Logpush output. For the moment, there is no option to add all fields at once, so you must specify all the fields names you are interested in.
	FieldNames []*string `json:"fieldNames,omitempty" tf:"field_names,omitempty"`

	// (String) Specifies the output type, such as ndjson or csv. This sets default values for the rest of the settings, depending on the chosen output type. Some formatting rules, like string quoting, are different between output types.
	// Specifies the output type, such as `ndjson` or `csv`. This sets default values for the rest of the settings, depending on the chosen output type. Some formatting rules, like string quoting, are different between output types.
	OutputType *string `json:"outputType,omitempty" tf:"output_type,omitempty"`

	// between the records as separator.
	// String to be inserted in-between the records as separator.
	RecordDelimiter *string `json:"recordDelimiter,omitempty" tf:"record_delimiter,omitempty"`

	// (String) String to be prepended before each record.
	// String to be prepended before each record.
	RecordPrefix *string `json:"recordPrefix,omitempty" tf:"record_prefix,omitempty"`

	// (String) String to be appended after each record.
	// String to be appended after each record.
	RecordSuffix *string `json:"recordSuffix,omitempty" tf:"record_suffix,omitempty"`

	// separated list. All fields used in the template must be present in field_names as well, otherwise they will end up as null. Format as a Go text/template without any standard functions, like conditionals, loops, sub-templates, etc.
	// String to use as template for each record instead of the default comma-separated list. All fields used in the template must be present in `field_names` as well, otherwise they will end up as null. Format as a Go `text/template` without any standard functions, like conditionals, loops, sub-templates, etc.
	RecordTemplate *string `json:"recordTemplate,omitempty" tf:"record_template,omitempty"`

	// (Number) Floating number to specify sampling rate. Sampling is applied on top of filtering, and regardless of the current sample_interval of the data.
	// Floating number to specify sampling rate. Sampling is applied on top of filtering, and regardless of the current `sample_interval` of the data.
	SampleRate *float64 `json:"sampleRate,omitempty" tf:"sample_rate,omitempty"`

	// (String) String to specify the format for timestamps, such as unixnano, unix, or rfc3339.
	// String to specify the format for timestamps, such as `unixnano`, `unix`, or `rfc3339`.
	TimestampFormat *string `json:"timestampFormat,omitempty" tf:"timestamp_format,omitempty"`
}

type OutputOptionsObservation struct {

	// (String) String to be prepended before each batch.
	// String to be prepended before each batch.
	BatchPrefix *string `json:"batchPrefix,omitempty" tf:"batch_prefix,omitempty"`

	// (String) String to be appended after each batch.
	// String to be appended after each batch.
	BatchSuffix *string `json:"batchSuffix,omitempty" tf:"batch_suffix,omitempty"`

	// (Boolean) If set to true, will cause all occurrences of ${ in the generated files to be replaced with x{.
	// If set to true, will cause all occurrences of `${` in the generated files to be replaced with `x{`.
	Cve20214428 *bool `json:"cve20214428,omitempty" tf:"cve_2021_4428,omitempty"`

	// (String) String to join fields. This field be ignored when record_template is set.
	// String to join fields. This field be ignored when `record_template` is set.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// (List of String) List of field names to be included in the Logpush output. For the moment, there is no option to add all fields at once, so you must specify all the fields names you are interested in.
	// List of field names to be included in the Logpush output. For the moment, there is no option to add all fields at once, so you must specify all the fields names you are interested in.
	FieldNames []*string `json:"fieldNames,omitempty" tf:"field_names,omitempty"`

	// (String) Specifies the output type, such as ndjson or csv. This sets default values for the rest of the settings, depending on the chosen output type. Some formatting rules, like string quoting, are different between output types.
	// Specifies the output type, such as `ndjson` or `csv`. This sets default values for the rest of the settings, depending on the chosen output type. Some formatting rules, like string quoting, are different between output types.
	OutputType *string `json:"outputType,omitempty" tf:"output_type,omitempty"`

	// between the records as separator.
	// String to be inserted in-between the records as separator.
	RecordDelimiter *string `json:"recordDelimiter,omitempty" tf:"record_delimiter,omitempty"`

	// (String) String to be prepended before each record.
	// String to be prepended before each record.
	RecordPrefix *string `json:"recordPrefix,omitempty" tf:"record_prefix,omitempty"`

	// (String) String to be appended after each record.
	// String to be appended after each record.
	RecordSuffix *string `json:"recordSuffix,omitempty" tf:"record_suffix,omitempty"`

	// separated list. All fields used in the template must be present in field_names as well, otherwise they will end up as null. Format as a Go text/template without any standard functions, like conditionals, loops, sub-templates, etc.
	// String to use as template for each record instead of the default comma-separated list. All fields used in the template must be present in `field_names` as well, otherwise they will end up as null. Format as a Go `text/template` without any standard functions, like conditionals, loops, sub-templates, etc.
	RecordTemplate *string `json:"recordTemplate,omitempty" tf:"record_template,omitempty"`

	// (Number) Floating number to specify sampling rate. Sampling is applied on top of filtering, and regardless of the current sample_interval of the data.
	// Floating number to specify sampling rate. Sampling is applied on top of filtering, and regardless of the current `sample_interval` of the data.
	SampleRate *float64 `json:"sampleRate,omitempty" tf:"sample_rate,omitempty"`

	// (String) String to specify the format for timestamps, such as unixnano, unix, or rfc3339.
	// String to specify the format for timestamps, such as `unixnano`, `unix`, or `rfc3339`.
	TimestampFormat *string `json:"timestampFormat,omitempty" tf:"timestamp_format,omitempty"`
}

type OutputOptionsParameters struct {

	// (String) String to be prepended before each batch.
	// String to be prepended before each batch.
	// +kubebuilder:validation:Optional
	BatchPrefix *string `json:"batchPrefix,omitempty" tf:"batch_prefix,omitempty"`

	// (String) String to be appended after each batch.
	// String to be appended after each batch.
	// +kubebuilder:validation:Optional
	BatchSuffix *string `json:"batchSuffix,omitempty" tf:"batch_suffix,omitempty"`

	// (Boolean) If set to true, will cause all occurrences of ${ in the generated files to be replaced with x{.
	// If set to true, will cause all occurrences of `${` in the generated files to be replaced with `x{`.
	// +kubebuilder:validation:Optional
	Cve20214428 *bool `json:"cve20214428,omitempty" tf:"cve_2021_4428,omitempty"`

	// (String) String to join fields. This field be ignored when record_template is set.
	// String to join fields. This field be ignored when `record_template` is set.
	// +kubebuilder:validation:Optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// (List of String) List of field names to be included in the Logpush output. For the moment, there is no option to add all fields at once, so you must specify all the fields names you are interested in.
	// List of field names to be included in the Logpush output. For the moment, there is no option to add all fields at once, so you must specify all the fields names you are interested in.
	// +kubebuilder:validation:Optional
	FieldNames []*string `json:"fieldNames,omitempty" tf:"field_names,omitempty"`

	// (String) Specifies the output type, such as ndjson or csv. This sets default values for the rest of the settings, depending on the chosen output type. Some formatting rules, like string quoting, are different between output types.
	// Specifies the output type, such as `ndjson` or `csv`. This sets default values for the rest of the settings, depending on the chosen output type. Some formatting rules, like string quoting, are different between output types.
	// +kubebuilder:validation:Optional
	OutputType *string `json:"outputType,omitempty" tf:"output_type,omitempty"`

	// between the records as separator.
	// String to be inserted in-between the records as separator.
	// +kubebuilder:validation:Optional
	RecordDelimiter *string `json:"recordDelimiter,omitempty" tf:"record_delimiter,omitempty"`

	// (String) String to be prepended before each record.
	// String to be prepended before each record.
	// +kubebuilder:validation:Optional
	RecordPrefix *string `json:"recordPrefix,omitempty" tf:"record_prefix,omitempty"`

	// (String) String to be appended after each record.
	// String to be appended after each record.
	// +kubebuilder:validation:Optional
	RecordSuffix *string `json:"recordSuffix,omitempty" tf:"record_suffix,omitempty"`

	// separated list. All fields used in the template must be present in field_names as well, otherwise they will end up as null. Format as a Go text/template without any standard functions, like conditionals, loops, sub-templates, etc.
	// String to use as template for each record instead of the default comma-separated list. All fields used in the template must be present in `field_names` as well, otherwise they will end up as null. Format as a Go `text/template` without any standard functions, like conditionals, loops, sub-templates, etc.
	// +kubebuilder:validation:Optional
	RecordTemplate *string `json:"recordTemplate,omitempty" tf:"record_template,omitempty"`

	// (Number) Floating number to specify sampling rate. Sampling is applied on top of filtering, and regardless of the current sample_interval of the data.
	// Floating number to specify sampling rate. Sampling is applied on top of filtering, and regardless of the current `sample_interval` of the data.
	// +kubebuilder:validation:Optional
	SampleRate *float64 `json:"sampleRate,omitempty" tf:"sample_rate,omitempty"`

	// (String) String to specify the format for timestamps, such as unixnano, unix, or rfc3339.
	// String to specify the format for timestamps, such as `unixnano`, `unix`, or `rfc3339`.
	// +kubebuilder:validation:Optional
	TimestampFormat *string `json:"timestampFormat,omitempty" tf:"timestamp_format,omitempty"`
}

// JobSpec defines the desired state of Job
type JobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider JobInitParameters `json:"initProvider,omitempty"`
}

// JobStatus defines the observed state of Job.
type JobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Job is the Schema for the Jobs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare}
type Job struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destinationConf) || (has(self.initProvider) && has(self.initProvider.destinationConf))",message="spec.forProvider.destinationConf is a required parameter"
	Spec   JobSpec   `json:"spec"`
	Status JobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobList contains a list of Jobs
type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Job `json:"items"`
}

// Repository type metadata.
var (
	Job_Kind             = "Job"
	Job_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Job_Kind}.String()
	Job_KindAPIVersion   = Job_Kind + "." + CRDGroupVersion.String()
	Job_GroupVersionKind = CRDGroupVersion.WithKind(Job_Kind)
)

func init() {
	SchemeBuilder.Register(&Job{}, &JobList{})
}
