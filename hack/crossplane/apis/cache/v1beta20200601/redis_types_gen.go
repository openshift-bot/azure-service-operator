// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:rbac:groups=cache.azure.com,resources=redis,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.azure.com,resources={redis/status,redis/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.Cache.json#/resourceDefinitions/redis
type Redis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Redis_Spec           `json:"spec,omitempty"`
	Status            RedisResource_STATUS `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.Cache.json#/resourceDefinitions/redis
type RedisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Redis `json:"items"`
}

// +kubebuilder:validation:Enum={"2020-06-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-06-01")

type Redis_Spec struct {
	v1alpha1.ResourceSpec `json:",inline,omitempty"`
	ForProvider           RedisParameters `json:"forProvider,omitempty"`
}

type RedisResource_STATUS struct {
	v1alpha1.ResourceStatus `json:",inline,omitempty"`
	AtProvider              RedisResourceObservation `json:"atProvider,omitempty"`
}

type RedisParameters struct {
	// EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1',
	// '1.2').
	MinimumTlsVersion *RedisCreateProperties_MinimumTlsVersion `json:"minimumTlsVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Name: The name of the Redis cache.
	Name string `json:"name,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method. Default value is
	// 'Enabled'.
	PublicNetworkAccess *RedisCreateProperties_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// RedisConfiguration: All Redis Settings. Few possible keys:
	// rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
	// etc.
	RedisConfiguration *RedisCommonPropertiesRedisConfiguration `json:"redisConfiguration,omitempty"`

	// ReplicasPerMaster: The number of replicas to be created per master.
	ReplicasPerMaster         *int                `json:"replicasPerMaster,omitempty"`
	ResourceGroupName         string              `json:"resourceGroupName,omitempty"`
	ResourceGroupNameRef      *v1alpha1.Reference `json:"resourceGroupNameRef,omitempty"`
	ResourceGroupNameSelector *v1alpha1.Selector  `json:"resourceGroupNameSelector,omitempty"`

	// ShardCount: The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int `json:"shardCount,omitempty"`

	// +kubebuilder:validation:Required
	// Sku: SKU parameters supplied to the create Redis operation.
	Sku *Sku `json:"sku,omitempty"`

	// +kubebuilder:validation:Pattern="^\\d+\\.\\d+\\.\\d+\\.\\d+$"
	// StaticIP: Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual
	// Network; auto assigned by default.
	StaticIP *string `json:"staticIP,omitempty"`

	// +kubebuilder:validation:Pattern="^/subscriptions/[^/]*/resourceGroups/[^/]*/providers/Microsoft.(ClassicNetwork|Network)/virtualNetworks/[^/]*/subnets/[^/]*$"
	// SubnetId: The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetId *string `json:"subnetId,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	// TenantSettings: A dictionary of tenant settings
	TenantSettings map[string]string `json:"tenantSettings,omitempty"`

	// Zones: A list of availability zones denoting where the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type RedisResourceObservation struct {
	// AccessKeys: The keys of the Redis cache - not set if this object is not the response to Create or Update redis cache
	AccessKeys *RedisAccessKeys_STATUS `json:"accessKeys,omitempty"`

	// EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`

	// HostName: Redis host name.
	HostName *string `json:"hostName,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Instances: List of the Redis instances associated with the cache
	Instances []RedisInstanceDetails_STATUS `json:"instances,omitempty"`

	// LinkedServers: List of the linked servers associated with the cache
	LinkedServers []RedisLinkedServer_STATUS `json:"linkedServers,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1',
	// '1.2')
	MinimumTlsVersion *RedisProperties_STATUS_MinimumTlsVersion `json:"minimumTlsVersion,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Port: Redis non-SSL port.
	Port *int `json:"port,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connection associated with the specified redis cache
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: Redis instance provisioning status.
	ProvisioningState *RedisProperties_STATUS_ProvisioningState `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method. Default value is
	// 'Enabled'
	PublicNetworkAccess *RedisProperties_STATUS_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// RedisConfiguration: All Redis Settings. Few possible keys:
	// rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
	// etc.
	RedisConfiguration *RedisProperties_STATUS_RedisConfiguration `json:"redisConfiguration,omitempty"`

	// RedisVersion: Redis version.
	RedisVersion *string `json:"redisVersion,omitempty"`

	// ReplicasPerMaster: The number of replicas to be created per master.
	ReplicasPerMaster *int `json:"replicasPerMaster,omitempty"`

	// ShardCount: The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int `json:"shardCount,omitempty"`

	// Sku: The SKU of the Redis cache to deploy.
	Sku *Sku_STATUS `json:"sku,omitempty"`

	// SslPort: Redis SSL port.
	SslPort *int `json:"sslPort,omitempty"`

	// StaticIP: Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual
	// Network; auto assigned by default.
	StaticIP *string `json:"staticIP,omitempty"`

	// SubnetId: The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/vnet1/subnets/subnet1
	SubnetId *string `json:"subnetId,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// TenantSettings: A dictionary of tenant settings
	TenantSettings map[string]string `json:"tenantSettings,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`

	// Zones: A list of availability zones denoting where the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type PrivateEndpointConnection_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// PrivateEndpoint: The resource of private end point.
	PrivateEndpoint *PrivateEndpoint_STATUS `json:"privateEndpoint,omitempty"`

	// PrivateLinkServiceConnectionState: A collection of information about the state of the connection between service
	// consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState_STATUS `json:"privateLinkServiceConnectionState,omitempty"`

	// ProvisioningState: The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type RedisAccessKeys_STATUS struct {
	// PrimaryKey: The current primary key that clients can use to authenticate with Redis cache.
	PrimaryKey *string `json:"primaryKey,omitempty"`

	// SecondaryKey: The current secondary key that clients can use to authenticate with Redis cache.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.Cache.json#/definitions/RedisCommonPropertiesRedisConfiguration
type RedisCommonPropertiesRedisConfiguration struct {
	// AdditionalProperties: Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty"`

	// AofStorageConnectionString0: First storage account connection string
	AofStorageConnectionString0 *string `json:"aof-storage-connection-string-0,omitempty"`

	// AofStorageConnectionString1: First storage account connection string
	AofStorageConnectionString1 *string `json:"aof-storage-connection-string-1,omitempty"`

	// MaxfragmentationmemoryReserved: Value in megabytes reserved for fragmentation per shard
	MaxfragmentationmemoryReserved *string `json:"maxfragmentationmemory-reserved,omitempty"`

	// MaxmemoryDelta: Value in megabytes reserved for non-cache usage per shard e.g. failover.
	MaxmemoryDelta *string `json:"maxmemory-delta,omitempty"`

	// MaxmemoryPolicy: The eviction strategy used when your data won't fit within its memory limit.
	MaxmemoryPolicy *string `json:"maxmemory-policy,omitempty"`

	// MaxmemoryReserved: Value in megabytes reserved for non-cache usage per shard e.g. failover.
	MaxmemoryReserved *string `json:"maxmemory-reserved,omitempty"`

	// RdbBackupEnabled: Specifies whether the rdb backup is enabled
	RdbBackupEnabled *string `json:"rdb-backup-enabled,omitempty"`

	// RdbBackupFrequency: Specifies the frequency for creating rdb backup
	RdbBackupFrequency *string `json:"rdb-backup-frequency,omitempty"`

	// RdbBackupMaxSnapshotCount: Specifies the maximum number of snapshots for rdb backup
	RdbBackupMaxSnapshotCount *string `json:"rdb-backup-max-snapshot-count,omitempty"`

	// RdbStorageConnectionString: The storage account connection string for storing rdb file
	RdbStorageConnectionString *string `json:"rdb-storage-connection-string,omitempty"`
}

// +kubebuilder:validation:Enum={"1.0","1.1","1.2"}
type RedisCreateProperties_MinimumTlsVersion string

const (
	RedisCreateProperties_MinimumTlsVersion_10 = RedisCreateProperties_MinimumTlsVersion("1.0")
	RedisCreateProperties_MinimumTlsVersion_11 = RedisCreateProperties_MinimumTlsVersion("1.1")
	RedisCreateProperties_MinimumTlsVersion_12 = RedisCreateProperties_MinimumTlsVersion("1.2")
)

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type RedisCreateProperties_PublicNetworkAccess string

const (
	RedisCreateProperties_PublicNetworkAccess_Disabled = RedisCreateProperties_PublicNetworkAccess("Disabled")
	RedisCreateProperties_PublicNetworkAccess_Enabled  = RedisCreateProperties_PublicNetworkAccess("Enabled")
)

type RedisInstanceDetails_STATUS struct {
	// IsMaster: Specifies whether the instance is a master node.
	IsMaster *bool `json:"isMaster,omitempty"`

	// NonSslPort: If enableNonSslPort is true, provides Redis instance Non-SSL port.
	NonSslPort *int `json:"nonSslPort,omitempty"`

	// ShardId: If clustering is enabled, the Shard ID of Redis Instance
	ShardId *int `json:"shardId,omitempty"`

	// SslPort: Redis instance SSL port.
	SslPort *int `json:"sslPort,omitempty"`

	// Zone: If the Cache uses availability zones, specifies availability zone where this instance is located.
	Zone *string `json:"zone,omitempty"`
}

type RedisLinkedServer_STATUS struct {
	// Id: Linked server Id.
	Id *string `json:"id,omitempty"`
}

type RedisProperties_STATUS_MinimumTlsVersion string

const (
	RedisProperties_STATUS_MinimumTlsVersion_10 = RedisProperties_STATUS_MinimumTlsVersion("1.0")
	RedisProperties_STATUS_MinimumTlsVersion_11 = RedisProperties_STATUS_MinimumTlsVersion("1.1")
	RedisProperties_STATUS_MinimumTlsVersion_12 = RedisProperties_STATUS_MinimumTlsVersion("1.2")
)

type RedisProperties_STATUS_ProvisioningState string

const (
	RedisProperties_STATUS_ProvisioningState_Creating               = RedisProperties_STATUS_ProvisioningState("Creating")
	RedisProperties_STATUS_ProvisioningState_Deleting               = RedisProperties_STATUS_ProvisioningState("Deleting")
	RedisProperties_STATUS_ProvisioningState_Disabled               = RedisProperties_STATUS_ProvisioningState("Disabled")
	RedisProperties_STATUS_ProvisioningState_Failed                 = RedisProperties_STATUS_ProvisioningState("Failed")
	RedisProperties_STATUS_ProvisioningState_Linking                = RedisProperties_STATUS_ProvisioningState("Linking")
	RedisProperties_STATUS_ProvisioningState_Provisioning           = RedisProperties_STATUS_ProvisioningState("Provisioning")
	RedisProperties_STATUS_ProvisioningState_RecoveringScaleFailure = RedisProperties_STATUS_ProvisioningState("RecoveringScaleFailure")
	RedisProperties_STATUS_ProvisioningState_Scaling                = RedisProperties_STATUS_ProvisioningState("Scaling")
	RedisProperties_STATUS_ProvisioningState_Succeeded              = RedisProperties_STATUS_ProvisioningState("Succeeded")
	RedisProperties_STATUS_ProvisioningState_Unlinking              = RedisProperties_STATUS_ProvisioningState("Unlinking")
	RedisProperties_STATUS_ProvisioningState_Unprovisioning         = RedisProperties_STATUS_ProvisioningState("Unprovisioning")
	RedisProperties_STATUS_ProvisioningState_Updating               = RedisProperties_STATUS_ProvisioningState("Updating")
)

type RedisProperties_STATUS_PublicNetworkAccess string

const (
	RedisProperties_STATUS_PublicNetworkAccess_Disabled = RedisProperties_STATUS_PublicNetworkAccess("Disabled")
	RedisProperties_STATUS_PublicNetworkAccess_Enabled  = RedisProperties_STATUS_PublicNetworkAccess("Enabled")
)

type RedisProperties_STATUS_RedisConfiguration struct {
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty"`

	// AofStorageConnectionString0: First storage account connection string
	AofStorageConnectionString0 *string `json:"aof-storage-connection-string-0,omitempty"`

	// AofStorageConnectionString1: Second storage account connection string
	AofStorageConnectionString1 *string `json:"aof-storage-connection-string-1,omitempty"`

	// Maxclients: The max clients config
	Maxclients *string `json:"maxclients,omitempty"`

	// MaxfragmentationmemoryReserved: Value in megabytes reserved for fragmentation per shard
	MaxfragmentationmemoryReserved *string `json:"maxfragmentationmemory-reserved,omitempty"`

	// MaxmemoryDelta: Value in megabytes reserved for non-cache usage per shard e.g. failover.
	MaxmemoryDelta *string `json:"maxmemory-delta,omitempty"`

	// MaxmemoryPolicy: The eviction strategy used when your data won't fit within its memory limit.
	MaxmemoryPolicy *string `json:"maxmemory-policy,omitempty"`

	// MaxmemoryReserved: Value in megabytes reserved for non-cache usage per shard e.g. failover.
	MaxmemoryReserved *string `json:"maxmemory-reserved,omitempty"`

	// RdbBackupEnabled: Specifies whether the rdb backup is enabled
	RdbBackupEnabled *string `json:"rdb-backup-enabled,omitempty"`

	// RdbBackupFrequency: Specifies the frequency for creating rdb backup
	RdbBackupFrequency *string `json:"rdb-backup-frequency,omitempty"`

	// RdbBackupMaxSnapshotCount: Specifies the maximum number of snapshots for rdb backup
	RdbBackupMaxSnapshotCount *string `json:"rdb-backup-max-snapshot-count,omitempty"`

	// RdbStorageConnectionString: The storage account connection string for storing rdb file
	RdbStorageConnectionString *string `json:"rdb-storage-connection-string,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.Cache.json#/definitions/Sku
type Sku struct {
	// +kubebuilder:validation:Required
	// Capacity: The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for
	// P (Premium) family (1, 2, 3, 4).
	Capacity *int `json:"capacity,omitempty"`

	// +kubebuilder:validation:Required
	// Family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
	Family *Sku_Family `json:"family,omitempty"`

	// +kubebuilder:validation:Required
	// Name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium).
	Name *Sku_Name `json:"name,omitempty"`
}

type Sku_STATUS struct {
	// Capacity: The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for
	// P (Premium) family (1, 2, 3, 4).
	Capacity *int `json:"capacity,omitempty"`

	// Family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
	Family *Sku_STATUS_Family `json:"family,omitempty"`

	// Name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
	Name *Sku_STATUS_Name `json:"name,omitempty"`
}

type PrivateEndpoint_STATUS struct {
	// Id: The ARM identifier for Private Endpoint
	Id *string `json:"id,omitempty"`
}

type PrivateEndpointConnectionProvisioningState_STATUS string

const (
	PrivateEndpointConnectionProvisioningState_STATUS_Creating  = PrivateEndpointConnectionProvisioningState_STATUS("Creating")
	PrivateEndpointConnectionProvisioningState_STATUS_Deleting  = PrivateEndpointConnectionProvisioningState_STATUS("Deleting")
	PrivateEndpointConnectionProvisioningState_STATUS_Failed    = PrivateEndpointConnectionProvisioningState_STATUS("Failed")
	PrivateEndpointConnectionProvisioningState_STATUS_Succeeded = PrivateEndpointConnectionProvisioningState_STATUS("Succeeded")
)

type PrivateLinkServiceConnectionState_STATUS struct {
	// ActionsRequired: A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string `json:"actionsRequired,omitempty"`

	// Description: The reason for approval/rejection of the connection.
	Description *string `json:"description,omitempty"`

	// Status: Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus_STATUS `json:"status,omitempty"`
}

// +kubebuilder:validation:Enum={"C","P"}
type Sku_Family string

const (
	Sku_Family_C = Sku_Family("C")
	Sku_Family_P = Sku_Family("P")
)

// +kubebuilder:validation:Enum={"Basic","Premium","Standard"}
type Sku_Name string

const (
	Sku_Name_Basic    = Sku_Name("Basic")
	Sku_Name_Premium  = Sku_Name("Premium")
	Sku_Name_Standard = Sku_Name("Standard")
)

type Sku_STATUS_Family string

const (
	Sku_STATUS_Family_C = Sku_STATUS_Family("C")
	Sku_STATUS_Family_P = Sku_STATUS_Family("P")
)

type Sku_STATUS_Name string

const (
	Sku_STATUS_Name_Basic    = Sku_STATUS_Name("Basic")
	Sku_STATUS_Name_Premium  = Sku_STATUS_Name("Premium")
	Sku_STATUS_Name_Standard = Sku_STATUS_Name("Standard")
)

type PrivateEndpointServiceConnectionStatus_STATUS string

const (
	PrivateEndpointServiceConnectionStatus_STATUS_Approved = PrivateEndpointServiceConnectionStatus_STATUS("Approved")
	PrivateEndpointServiceConnectionStatus_STATUS_Pending  = PrivateEndpointServiceConnectionStatus_STATUS("Pending")
	PrivateEndpointServiceConnectionStatus_STATUS_Rejected = PrivateEndpointServiceConnectionStatus_STATUS("Rejected")
)

func init() {
	SchemeBuilder.Register(&Redis{}, &RedisList{})
}
