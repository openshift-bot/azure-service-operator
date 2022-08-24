// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type Servers_SpecARM struct {
	// Location: The location the resource resides in.
	Location *string `json:"location,omitempty"`

	// Name: The name of the server.
	Name string `json:"name,omitempty"`

	// Properties: The properties used to create a new server.
	Properties *ServerPropertiesForCreateARM `json:"properties,omitempty"`

	// Sku: Billing information related properties of a server.
	Sku *SkuARM `json:"sku,omitempty"`

	// Tags: Application-specific metadata in the form of key-value pairs.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Servers_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-06-01"
func (servers Servers_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (servers *Servers_SpecARM) GetName() string {
	return servers.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMariaDB/servers"
func (servers *Servers_SpecARM) GetType() string {
	return "Microsoft.DBforMariaDB/servers"
}

// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/definitions/ServerPropertiesForCreate
type ServerPropertiesForCreateARM struct {
	// ServerPropertiesForDefaultCreate: Mutually exclusive with all other properties
	ServerPropertiesForDefaultCreate *ServerPropertiesForDefaultCreateARM `json:"serverPropertiesForDefaultCreate,omitempty"`

	// ServerPropertiesForGeoRestore: Mutually exclusive with all other properties
	ServerPropertiesForGeoRestore *ServerPropertiesForGeoRestoreARM `json:"serverPropertiesForGeoRestore,omitempty"`

	// ServerPropertiesForReplica: Mutually exclusive with all other properties
	ServerPropertiesForReplica *ServerPropertiesForReplicaARM `json:"serverPropertiesForReplica,omitempty"`

	// ServerPropertiesForRestore: Mutually exclusive with all other properties
	ServerPropertiesForRestore *ServerPropertiesForRestoreARM `json:"serverPropertiesForRestore,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because ServerPropertiesForCreateARM represents a discriminated union (JSON OneOf)
func (create ServerPropertiesForCreateARM) MarshalJSON() ([]byte, error) {
	if create.ServerPropertiesForDefaultCreate != nil {
		return json.Marshal(create.ServerPropertiesForDefaultCreate)
	}
	if create.ServerPropertiesForGeoRestore != nil {
		return json.Marshal(create.ServerPropertiesForGeoRestore)
	}
	if create.ServerPropertiesForReplica != nil {
		return json.Marshal(create.ServerPropertiesForReplica)
	}
	if create.ServerPropertiesForRestore != nil {
		return json.Marshal(create.ServerPropertiesForRestore)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the ServerPropertiesForCreateARM
func (create *ServerPropertiesForCreateARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["createMode"]
	if discriminator == "Default" {
		create.ServerPropertiesForDefaultCreate = &ServerPropertiesForDefaultCreateARM{}
		return json.Unmarshal(data, create.ServerPropertiesForDefaultCreate)
	}
	if discriminator == "GeoRestore" {
		create.ServerPropertiesForGeoRestore = &ServerPropertiesForGeoRestoreARM{}
		return json.Unmarshal(data, create.ServerPropertiesForGeoRestore)
	}
	if discriminator == "PointInTimeRestore" {
		create.ServerPropertiesForRestore = &ServerPropertiesForRestoreARM{}
		return json.Unmarshal(data, create.ServerPropertiesForRestore)
	}
	if discriminator == "Replica" {
		create.ServerPropertiesForReplica = &ServerPropertiesForReplicaARM{}
		return json.Unmarshal(data, create.ServerPropertiesForReplica)
	}

	// No error
	return nil
}

// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/definitions/Sku
type SkuARM struct {
	// Capacity: The scale up/out capacity, representing server's compute units.
	Capacity *int `json:"capacity,omitempty"`

	// Family: The family of hardware.
	Family *string `json:"family,omitempty"`

	// Name: The name of the sku, typically, tier + family + cores, e.g. B_Gen4_1, GP_Gen5_8.
	Name *string `json:"name,omitempty"`

	// Size: The size code, to be interpreted by resource as appropriate.
	Size *string `json:"size,omitempty"`

	// Tier: The tier of the particular SKU, e.g. Basic.
	Tier *Sku_Tier `json:"tier,omitempty"`
}

type ServerPropertiesForDefaultCreateARM struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AdministratorLoginPassword: The password of the administrator login.
	AdministratorLoginPassword string                                                                `json:"administratorLoginPassword,omitempty"`
	CreateMode                 ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_CreateMode `json:"createMode,omitempty"`

	// MinimalTlsVersion: Enforce a minimal Tls version for the server.
	MinimalTlsVersion *ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_MinimalTlsVersion `json:"minimalTlsVersion,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// SslEnforcement: Enable ssl enforcement or not when connect to server.
	SslEnforcement *ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_SslEnforcement `json:"sslEnforcement,omitempty"`

	// StorageProfile: Storage Profile properties of a server
	StorageProfile *StorageProfileARM `json:"storageProfile,omitempty"`

	// Version: Server version.
	Version *ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_Version `json:"version,omitempty"`
}

type ServerPropertiesForGeoRestoreARM struct {
	CreateMode ServerPropertiesForCreate_ServerPropertiesForGeoRestore_CreateMode `json:"createMode,omitempty"`

	// MinimalTlsVersion: Enforce a minimal Tls version for the server.
	MinimalTlsVersion *ServerPropertiesForCreate_ServerPropertiesForGeoRestore_MinimalTlsVersion `json:"minimalTlsVersion,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *ServerPropertiesForCreate_ServerPropertiesForGeoRestore_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// SourceServerId: The source server id to restore from.
	SourceServerId *string `json:"sourceServerId,omitempty"`

	// SslEnforcement: Enable ssl enforcement or not when connect to server.
	SslEnforcement *ServerPropertiesForCreate_ServerPropertiesForGeoRestore_SslEnforcement `json:"sslEnforcement,omitempty"`

	// StorageProfile: Storage Profile properties of a server
	StorageProfile *StorageProfileARM `json:"storageProfile,omitempty"`

	// Version: Server version.
	Version *ServerPropertiesForCreate_ServerPropertiesForGeoRestore_Version `json:"version,omitempty"`
}

type ServerPropertiesForReplicaARM struct {
	CreateMode ServerPropertiesForCreate_ServerPropertiesForReplica_CreateMode `json:"createMode,omitempty"`

	// MinimalTlsVersion: Enforce a minimal Tls version for the server.
	MinimalTlsVersion *ServerPropertiesForCreate_ServerPropertiesForReplica_MinimalTlsVersion `json:"minimalTlsVersion,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *ServerPropertiesForCreate_ServerPropertiesForReplica_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// SourceServerId: The master server id to create replica from.
	SourceServerId *string `json:"sourceServerId,omitempty"`

	// SslEnforcement: Enable ssl enforcement or not when connect to server.
	SslEnforcement *ServerPropertiesForCreate_ServerPropertiesForReplica_SslEnforcement `json:"sslEnforcement,omitempty"`

	// StorageProfile: Storage Profile properties of a server
	StorageProfile *StorageProfileARM `json:"storageProfile,omitempty"`

	// Version: Server version.
	Version *ServerPropertiesForCreate_ServerPropertiesForReplica_Version `json:"version,omitempty"`
}

type ServerPropertiesForRestoreARM struct {
	CreateMode ServerPropertiesForCreate_ServerPropertiesForRestore_CreateMode `json:"createMode,omitempty"`

	// MinimalTlsVersion: Enforce a minimal Tls version for the server.
	MinimalTlsVersion *ServerPropertiesForCreate_ServerPropertiesForRestore_MinimalTlsVersion `json:"minimalTlsVersion,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess *ServerPropertiesForCreate_ServerPropertiesForRestore_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// RestorePointInTime: Restore point creation time (ISO8601 format), specifying the time to restore from.
	RestorePointInTime *string `json:"restorePointInTime,omitempty"`

	// SourceServerId: The source server id to restore from.
	SourceServerId *string `json:"sourceServerId,omitempty"`

	// SslEnforcement: Enable ssl enforcement or not when connect to server.
	SslEnforcement *ServerPropertiesForCreate_ServerPropertiesForRestore_SslEnforcement `json:"sslEnforcement,omitempty"`

	// StorageProfile: Storage Profile properties of a server
	StorageProfile *StorageProfileARM `json:"storageProfile,omitempty"`

	// Version: Server version.
	Version *ServerPropertiesForCreate_ServerPropertiesForRestore_Version `json:"version,omitempty"`
}

// +kubebuilder:validation:Enum={"Basic","GeneralPurpose","MemoryOptimized"}
type Sku_Tier string

const (
	Sku_Tier_Basic           = Sku_Tier("Basic")
	Sku_Tier_GeneralPurpose  = Sku_Tier("GeneralPurpose")
	Sku_Tier_MemoryOptimized = Sku_Tier("MemoryOptimized")
)

// +kubebuilder:validation:Enum={"Default"}
type ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_CreateMode string

const ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_CreateMode_Default = ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_CreateMode("Default")

// +kubebuilder:validation:Enum={"TLS1_0","TLS1_1","TLS1_2","TLSEnforcementDisabled"}
type ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_MinimalTlsVersion string

const (
	ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_MinimalTlsVersion_TLS1_0                 = ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_MinimalTlsVersion("TLS1_0")
	ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_MinimalTlsVersion_TLS1_1                 = ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_MinimalTlsVersion("TLS1_1")
	ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_MinimalTlsVersion_TLS1_2                 = ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_MinimalTlsVersion("TLS1_2")
	ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_MinimalTlsVersion_TLSEnforcementDisabled = ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_MinimalTlsVersion("TLSEnforcementDisabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_PublicNetworkAccess string

const (
	ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_PublicNetworkAccess_Disabled = ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_PublicNetworkAccess("Disabled")
	ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_PublicNetworkAccess_Enabled  = ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_PublicNetworkAccess("Enabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_SslEnforcement string

const (
	ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_SslEnforcement_Disabled = ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_SslEnforcement("Disabled")
	ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_SslEnforcement_Enabled  = ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_SslEnforcement("Enabled")
)

// +kubebuilder:validation:Enum={"10.2","10.3"}
type ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_Version string

const (
	ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_Version_102 = ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_Version("10.2")
	ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_Version_103 = ServerPropertiesForCreate_ServerPropertiesForDefaultCreate_Version("10.3")
)

// +kubebuilder:validation:Enum={"GeoRestore"}
type ServerPropertiesForCreate_ServerPropertiesForGeoRestore_CreateMode string

const ServerPropertiesForCreate_ServerPropertiesForGeoRestore_CreateMode_GeoRestore = ServerPropertiesForCreate_ServerPropertiesForGeoRestore_CreateMode("GeoRestore")

// +kubebuilder:validation:Enum={"TLS1_0","TLS1_1","TLS1_2","TLSEnforcementDisabled"}
type ServerPropertiesForCreate_ServerPropertiesForGeoRestore_MinimalTlsVersion string

const (
	ServerPropertiesForCreate_ServerPropertiesForGeoRestore_MinimalTlsVersion_TLS1_0                 = ServerPropertiesForCreate_ServerPropertiesForGeoRestore_MinimalTlsVersion("TLS1_0")
	ServerPropertiesForCreate_ServerPropertiesForGeoRestore_MinimalTlsVersion_TLS1_1                 = ServerPropertiesForCreate_ServerPropertiesForGeoRestore_MinimalTlsVersion("TLS1_1")
	ServerPropertiesForCreate_ServerPropertiesForGeoRestore_MinimalTlsVersion_TLS1_2                 = ServerPropertiesForCreate_ServerPropertiesForGeoRestore_MinimalTlsVersion("TLS1_2")
	ServerPropertiesForCreate_ServerPropertiesForGeoRestore_MinimalTlsVersion_TLSEnforcementDisabled = ServerPropertiesForCreate_ServerPropertiesForGeoRestore_MinimalTlsVersion("TLSEnforcementDisabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreate_ServerPropertiesForGeoRestore_PublicNetworkAccess string

const (
	ServerPropertiesForCreate_ServerPropertiesForGeoRestore_PublicNetworkAccess_Disabled = ServerPropertiesForCreate_ServerPropertiesForGeoRestore_PublicNetworkAccess("Disabled")
	ServerPropertiesForCreate_ServerPropertiesForGeoRestore_PublicNetworkAccess_Enabled  = ServerPropertiesForCreate_ServerPropertiesForGeoRestore_PublicNetworkAccess("Enabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreate_ServerPropertiesForGeoRestore_SslEnforcement string

const (
	ServerPropertiesForCreate_ServerPropertiesForGeoRestore_SslEnforcement_Disabled = ServerPropertiesForCreate_ServerPropertiesForGeoRestore_SslEnforcement("Disabled")
	ServerPropertiesForCreate_ServerPropertiesForGeoRestore_SslEnforcement_Enabled  = ServerPropertiesForCreate_ServerPropertiesForGeoRestore_SslEnforcement("Enabled")
)

// +kubebuilder:validation:Enum={"10.2","10.3"}
type ServerPropertiesForCreate_ServerPropertiesForGeoRestore_Version string

const (
	ServerPropertiesForCreate_ServerPropertiesForGeoRestore_Version_102 = ServerPropertiesForCreate_ServerPropertiesForGeoRestore_Version("10.2")
	ServerPropertiesForCreate_ServerPropertiesForGeoRestore_Version_103 = ServerPropertiesForCreate_ServerPropertiesForGeoRestore_Version("10.3")
)

// +kubebuilder:validation:Enum={"Replica"}
type ServerPropertiesForCreate_ServerPropertiesForReplica_CreateMode string

const ServerPropertiesForCreate_ServerPropertiesForReplica_CreateMode_Replica = ServerPropertiesForCreate_ServerPropertiesForReplica_CreateMode("Replica")

// +kubebuilder:validation:Enum={"TLS1_0","TLS1_1","TLS1_2","TLSEnforcementDisabled"}
type ServerPropertiesForCreate_ServerPropertiesForReplica_MinimalTlsVersion string

const (
	ServerPropertiesForCreate_ServerPropertiesForReplica_MinimalTlsVersion_TLS1_0                 = ServerPropertiesForCreate_ServerPropertiesForReplica_MinimalTlsVersion("TLS1_0")
	ServerPropertiesForCreate_ServerPropertiesForReplica_MinimalTlsVersion_TLS1_1                 = ServerPropertiesForCreate_ServerPropertiesForReplica_MinimalTlsVersion("TLS1_1")
	ServerPropertiesForCreate_ServerPropertiesForReplica_MinimalTlsVersion_TLS1_2                 = ServerPropertiesForCreate_ServerPropertiesForReplica_MinimalTlsVersion("TLS1_2")
	ServerPropertiesForCreate_ServerPropertiesForReplica_MinimalTlsVersion_TLSEnforcementDisabled = ServerPropertiesForCreate_ServerPropertiesForReplica_MinimalTlsVersion("TLSEnforcementDisabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreate_ServerPropertiesForReplica_PublicNetworkAccess string

const (
	ServerPropertiesForCreate_ServerPropertiesForReplica_PublicNetworkAccess_Disabled = ServerPropertiesForCreate_ServerPropertiesForReplica_PublicNetworkAccess("Disabled")
	ServerPropertiesForCreate_ServerPropertiesForReplica_PublicNetworkAccess_Enabled  = ServerPropertiesForCreate_ServerPropertiesForReplica_PublicNetworkAccess("Enabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreate_ServerPropertiesForReplica_SslEnforcement string

const (
	ServerPropertiesForCreate_ServerPropertiesForReplica_SslEnforcement_Disabled = ServerPropertiesForCreate_ServerPropertiesForReplica_SslEnforcement("Disabled")
	ServerPropertiesForCreate_ServerPropertiesForReplica_SslEnforcement_Enabled  = ServerPropertiesForCreate_ServerPropertiesForReplica_SslEnforcement("Enabled")
)

// +kubebuilder:validation:Enum={"10.2","10.3"}
type ServerPropertiesForCreate_ServerPropertiesForReplica_Version string

const (
	ServerPropertiesForCreate_ServerPropertiesForReplica_Version_102 = ServerPropertiesForCreate_ServerPropertiesForReplica_Version("10.2")
	ServerPropertiesForCreate_ServerPropertiesForReplica_Version_103 = ServerPropertiesForCreate_ServerPropertiesForReplica_Version("10.3")
)

// +kubebuilder:validation:Enum={"PointInTimeRestore"}
type ServerPropertiesForCreate_ServerPropertiesForRestore_CreateMode string

const ServerPropertiesForCreate_ServerPropertiesForRestore_CreateMode_PointInTimeRestore = ServerPropertiesForCreate_ServerPropertiesForRestore_CreateMode("PointInTimeRestore")

// +kubebuilder:validation:Enum={"TLS1_0","TLS1_1","TLS1_2","TLSEnforcementDisabled"}
type ServerPropertiesForCreate_ServerPropertiesForRestore_MinimalTlsVersion string

const (
	ServerPropertiesForCreate_ServerPropertiesForRestore_MinimalTlsVersion_TLS1_0                 = ServerPropertiesForCreate_ServerPropertiesForRestore_MinimalTlsVersion("TLS1_0")
	ServerPropertiesForCreate_ServerPropertiesForRestore_MinimalTlsVersion_TLS1_1                 = ServerPropertiesForCreate_ServerPropertiesForRestore_MinimalTlsVersion("TLS1_1")
	ServerPropertiesForCreate_ServerPropertiesForRestore_MinimalTlsVersion_TLS1_2                 = ServerPropertiesForCreate_ServerPropertiesForRestore_MinimalTlsVersion("TLS1_2")
	ServerPropertiesForCreate_ServerPropertiesForRestore_MinimalTlsVersion_TLSEnforcementDisabled = ServerPropertiesForCreate_ServerPropertiesForRestore_MinimalTlsVersion("TLSEnforcementDisabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreate_ServerPropertiesForRestore_PublicNetworkAccess string

const (
	ServerPropertiesForCreate_ServerPropertiesForRestore_PublicNetworkAccess_Disabled = ServerPropertiesForCreate_ServerPropertiesForRestore_PublicNetworkAccess("Disabled")
	ServerPropertiesForCreate_ServerPropertiesForRestore_PublicNetworkAccess_Enabled  = ServerPropertiesForCreate_ServerPropertiesForRestore_PublicNetworkAccess("Enabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesForCreate_ServerPropertiesForRestore_SslEnforcement string

const (
	ServerPropertiesForCreate_ServerPropertiesForRestore_SslEnforcement_Disabled = ServerPropertiesForCreate_ServerPropertiesForRestore_SslEnforcement("Disabled")
	ServerPropertiesForCreate_ServerPropertiesForRestore_SslEnforcement_Enabled  = ServerPropertiesForCreate_ServerPropertiesForRestore_SslEnforcement("Enabled")
)

// +kubebuilder:validation:Enum={"10.2","10.3"}
type ServerPropertiesForCreate_ServerPropertiesForRestore_Version string

const (
	ServerPropertiesForCreate_ServerPropertiesForRestore_Version_102 = ServerPropertiesForCreate_ServerPropertiesForRestore_Version("10.2")
	ServerPropertiesForCreate_ServerPropertiesForRestore_Version_103 = ServerPropertiesForCreate_ServerPropertiesForRestore_Version("10.3")
)

// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/definitions/StorageProfile
type StorageProfileARM struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// GeoRedundantBackup: Enable Geo-redundant or not for server backup.
	GeoRedundantBackup *StorageProfile_GeoRedundantBackup `json:"geoRedundantBackup,omitempty"`

	// StorageAutogrow: Enable Storage Auto Grow.
	StorageAutogrow *StorageProfile_StorageAutogrow `json:"storageAutogrow,omitempty"`

	// StorageMB: Max storage allowed for a server.
	StorageMB *int `json:"storageMB,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type StorageProfile_GeoRedundantBackup string

const (
	StorageProfile_GeoRedundantBackup_Disabled = StorageProfile_GeoRedundantBackup("Disabled")
	StorageProfile_GeoRedundantBackup_Enabled  = StorageProfile_GeoRedundantBackup("Enabled")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type StorageProfile_StorageAutogrow string

const (
	StorageProfile_StorageAutogrow_Disabled = StorageProfile_StorageAutogrow("Disabled")
	StorageProfile_StorageAutogrow_Enabled  = StorageProfile_StorageAutogrow("Enabled")
)
