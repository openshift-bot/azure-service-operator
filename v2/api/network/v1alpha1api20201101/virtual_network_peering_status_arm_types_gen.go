// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

// Deprecated version of VirtualNetworkPeering_STATUS. Use v1beta20201101.VirtualNetworkPeering_STATUS instead
type VirtualNetworkPeering_STATUSARM struct {
	Etag       *string                                          `json:"etag,omitempty"`
	Id         *string                                          `json:"id,omitempty"`
	Name       *string                                          `json:"name,omitempty"`
	Properties *VirtualNetworkPeeringPropertiesFormat_STATUSARM `json:"properties,omitempty"`
	Type       *string                                          `json:"type,omitempty"`
}

// Deprecated version of VirtualNetworkPeeringPropertiesFormat_STATUS. Use v1beta20201101.VirtualNetworkPeeringPropertiesFormat_STATUS instead
type VirtualNetworkPeeringPropertiesFormat_STATUSARM struct {
	AllowForwardedTraffic     *bool                                                      `json:"allowForwardedTraffic,omitempty"`
	AllowGatewayTransit       *bool                                                      `json:"allowGatewayTransit,omitempty"`
	AllowVirtualNetworkAccess *bool                                                      `json:"allowVirtualNetworkAccess,omitempty"`
	DoNotVerifyRemoteGateways *bool                                                      `json:"doNotVerifyRemoteGateways,omitempty"`
	PeeringState              *VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState `json:"peeringState,omitempty"`
	ProvisioningState         *ProvisioningState_STATUS                                  `json:"provisioningState,omitempty"`
	RemoteAddressSpace        *AddressSpace_STATUSARM                                    `json:"remoteAddressSpace,omitempty"`
	RemoteBgpCommunities      *VirtualNetworkBgpCommunities_STATUSARM                    `json:"remoteBgpCommunities,omitempty"`
	RemoteVirtualNetwork      *SubResource_STATUSARM                                     `json:"remoteVirtualNetwork,omitempty"`
	ResourceGuid              *string                                                    `json:"resourceGuid,omitempty"`
	UseRemoteGateways         *bool                                                      `json:"useRemoteGateways,omitempty"`
}

// Deprecated version of VirtualNetworkBgpCommunities_STATUS. Use v1beta20201101.VirtualNetworkBgpCommunities_STATUS instead
type VirtualNetworkBgpCommunities_STATUSARM struct {
	RegionalCommunity       *string `json:"regionalCommunity,omitempty"`
	VirtualNetworkCommunity *string `json:"virtualNetworkCommunity,omitempty"`
}

// Deprecated version of VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState. Use
// v1beta20201101.VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState instead
type VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState string

const (
	VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState_Connected    = VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState("Connected")
	VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState_Disconnected = VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState("Disconnected")
	VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState_Initiated    = VirtualNetworkPeeringPropertiesFormat_STATUS_PeeringState("Initiated")
)
