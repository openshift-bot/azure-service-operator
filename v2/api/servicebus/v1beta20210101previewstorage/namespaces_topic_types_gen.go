// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101previewstorage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=servicebus.azure.com,resources=namespacestopics,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=servicebus.azure.com,resources={namespacestopics/status,namespacestopics/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210101preview.NamespacesTopic
// Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces_topics
type NamespacesTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_Topics_Spec `json:"spec,omitempty"`
	Status            SBTopic_STATUS         `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesTopic{}

// GetConditions returns the conditions of the resource
func (topic *NamespacesTopic) GetConditions() conditions.Conditions {
	return topic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (topic *NamespacesTopic) SetConditions(conditions conditions.Conditions) {
	topic.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &NamespacesTopic{}

// AzureName returns the Azure name of the resource
func (topic *NamespacesTopic) AzureName() string {
	return topic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (topic NamespacesTopic) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (topic *NamespacesTopic) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (topic *NamespacesTopic) GetSpec() genruntime.ConvertibleSpec {
	return &topic.Spec
}

// GetStatus returns the status of this resource
func (topic *NamespacesTopic) GetStatus() genruntime.ConvertibleStatus {
	return &topic.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/topics"
func (topic *NamespacesTopic) GetType() string {
	return "Microsoft.ServiceBus/namespaces/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (topic *NamespacesTopic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SBTopic_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (topic *NamespacesTopic) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(topic.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  topic.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (topic *NamespacesTopic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SBTopic_STATUS); ok {
		topic.Status = *st
		return nil
	}

	// Convert status to required version
	var st SBTopic_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	topic.Status = st
	return nil
}

// Hub marks that this NamespacesTopic is the hub type for conversion
func (topic *NamespacesTopic) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (topic *NamespacesTopic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: topic.Spec.OriginalVersion,
		Kind:    "NamespacesTopic",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210101preview.NamespacesTopic
// Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces_topics
type NamespacesTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesTopic `json:"items"`
}

// Storage version of v1beta20210101preview.Namespaces_Topics_Spec
type Namespaces_Topics_Spec struct {
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                           string  `json:"azureName,omitempty"`
	DefaultMessageTimeToLive            *string `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool   `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool   `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool   `json:"enablePartitioning,omitempty"`
	Location                            *string `json:"location,omitempty"`
	MaxSizeInMegabytes                  *int    `json:"maxSizeInMegabytes,omitempty"`
	OriginalVersion                     string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a servicebus.azure.com/Namespace resource
	Owner                      *genruntime.KnownResourceReference `group:"servicebus.azure.com" json:"owner,omitempty" kind:"Namespace"`
	PropertyBag                genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection *bool                              `json:"requiresDuplicateDetection,omitempty"`
	SupportOrdering            *bool                              `json:"supportOrdering,omitempty"`
	Tags                       map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Namespaces_Topics_Spec{}

// ConvertSpecFrom populates our Namespaces_Topics_Spec from the provided source
func (topics *Namespaces_Topics_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == topics {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(topics)
}

// ConvertSpecTo populates the provided destination from our Namespaces_Topics_Spec
func (topics *Namespaces_Topics_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == topics {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(topics)
}

// Storage version of v1beta20210101preview.SBTopic_STATUS
type SBTopic_STATUS struct {
	AccessedAt                          *string                     `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string                     `json:"autoDeleteOnIdle,omitempty"`
	Conditions                          []conditions.Condition      `json:"conditions,omitempty"`
	CountDetails                        *MessageCountDetails_STATUS `json:"countDetails,omitempty"`
	CreatedAt                           *string                     `json:"createdAt,omitempty"`
	DefaultMessageTimeToLive            *string                     `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                     `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                       `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                       `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                       `json:"enablePartitioning,omitempty"`
	Id                                  *string                     `json:"id,omitempty"`
	MaxSizeInMegabytes                  *int                        `json:"maxSizeInMegabytes,omitempty"`
	Name                                *string                     `json:"name,omitempty"`
	PropertyBag                         genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection          *bool                       `json:"requiresDuplicateDetection,omitempty"`
	SizeInBytes                         *int                        `json:"sizeInBytes,omitempty"`
	Status                              *string                     `json:"status,omitempty"`
	SubscriptionCount                   *int                        `json:"subscriptionCount,omitempty"`
	SupportOrdering                     *bool                       `json:"supportOrdering,omitempty"`
	SystemData                          *SystemData_STATUS          `json:"systemData,omitempty"`
	Type                                *string                     `json:"type,omitempty"`
	UpdatedAt                           *string                     `json:"updatedAt,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SBTopic_STATUS{}

// ConvertStatusFrom populates our SBTopic_STATUS from the provided source
func (topic *SBTopic_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == topic {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(topic)
}

// ConvertStatusTo populates the provided destination from our SBTopic_STATUS
func (topic *SBTopic_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == topic {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(topic)
}

func init() {
	SchemeBuilder.Register(&NamespacesTopic{}, &NamespacesTopicList{})
}
