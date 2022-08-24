// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Namespaces_Eventhubs_Spec. Use v1beta20211101.Namespaces_Eventhubs_Spec instead
type Namespaces_Eventhubs_SpecARM struct {
	Location   *string                                  `json:"location,omitempty"`
	Name       string                                   `json:"name,omitempty"`
	Properties *Namespaces_Eventhubs_Spec_PropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                        `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Namespaces_Eventhubs_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (eventhubs Namespaces_Eventhubs_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (eventhubs *Namespaces_Eventhubs_SpecARM) GetName() string {
	return eventhubs.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs"
func (eventhubs *Namespaces_Eventhubs_SpecARM) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs"
}

// Deprecated version of Namespaces_Eventhubs_Spec_Properties. Use v1beta20211101.Namespaces_Eventhubs_Spec_Properties instead
type Namespaces_Eventhubs_Spec_PropertiesARM struct {
	CaptureDescription     *Namespaces_Eventhubs_Spec_Properties_CaptureDescriptionARM `json:"captureDescription,omitempty"`
	MessageRetentionInDays *int                                                        `json:"messageRetentionInDays,omitempty"`
	PartitionCount         *int                                                        `json:"partitionCount,omitempty"`
}

// Deprecated version of Namespaces_Eventhubs_Spec_Properties_CaptureDescription. Use v1beta20211101.Namespaces_Eventhubs_Spec_Properties_CaptureDescription instead
type Namespaces_Eventhubs_Spec_Properties_CaptureDescriptionARM struct {
	Destination       *Namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationARM `json:"destination,omitempty"`
	Enabled           *bool                                                                   `json:"enabled,omitempty"`
	Encoding          *Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Encoding       `json:"encoding,omitempty"`
	IntervalInSeconds *int                                                                    `json:"intervalInSeconds,omitempty"`
	SizeLimitInBytes  *int                                                                    `json:"sizeLimitInBytes,omitempty"`
	SkipEmptyArchives *bool                                                                   `json:"skipEmptyArchives,omitempty"`
}

// Deprecated version of Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination. Use v1beta20211101.Namespaces_Eventhubs_Spec_Properties_CaptureDescription_Destination instead
type Namespaces_Eventhubs_Spec_Properties_CaptureDescription_DestinationARM struct {
	Name       *string                   `json:"name,omitempty"`
	Properties *DestinationPropertiesARM `json:"properties,omitempty"`
}

// Deprecated version of DestinationProperties. Use v1beta20211101.DestinationProperties instead
type DestinationPropertiesARM struct {
	ArchiveNameFormat        *string `json:"archiveNameFormat,omitempty"`
	BlobContainer            *string `json:"blobContainer,omitempty"`
	DataLakeAccountName      *string `json:"dataLakeAccountName,omitempty"`
	DataLakeFolderPath       *string `json:"dataLakeFolderPath,omitempty"`
	DataLakeSubscriptionId   *string `json:"dataLakeSubscriptionId,omitempty"`
	StorageAccountResourceId *string `json:"storageAccountResourceId,omitempty"`
}
