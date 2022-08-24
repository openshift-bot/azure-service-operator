// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"fmt"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1alpha1api20210515.MongodbDatabaseThroughputSetting
// Deprecated version of MongodbDatabaseThroughputSetting. Use v1beta20210515.MongodbDatabaseThroughputSetting instead
type MongodbDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec `json:"spec,omitempty"`
	Status            ThroughputSettingsGetResults_STATUS                       `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabaseThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *MongodbDatabaseThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *MongodbDatabaseThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ conversion.Convertible = &MongodbDatabaseThroughputSetting{}

// ConvertFrom populates our MongodbDatabaseThroughputSetting from the provided hub MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210515s.MongodbDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/MongodbDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignProperties_From_MongodbDatabaseThroughputSetting(source)
}

// ConvertTo populates the provided hub MongodbDatabaseThroughputSetting from our MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.MongodbDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/MongodbDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignProperties_To_MongodbDatabaseThroughputSetting(destination)
}

var _ genruntime.KubernetesResource = &MongodbDatabaseThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *MongodbDatabaseThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting MongodbDatabaseThroughputSetting) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (setting *MongodbDatabaseThroughputSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *MongodbDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *MongodbDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
func (setting *MongodbDatabaseThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *MongodbDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ThroughputSettingsGetResults_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (setting *MongodbDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  setting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (setting *MongodbDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ThroughputSettingsGetResults_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ThroughputSettingsGetResults_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// AssignProperties_From_MongodbDatabaseThroughputSetting populates our MongodbDatabaseThroughputSetting from the provided source MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) AssignProperties_From_MongodbDatabaseThroughputSetting(source *v20210515s.MongodbDatabaseThroughputSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec
	err := spec.AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status ThroughputSettingsGetResults_STATUS
	err = status.AssignProperties_From_ThroughputSettingsGetResults_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_ThroughputSettingsGetResults_STATUS() to populate field Status")
	}
	setting.Status = status

	// No error
	return nil
}

// AssignProperties_To_MongodbDatabaseThroughputSetting populates the provided destination MongodbDatabaseThroughputSetting from our MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) AssignProperties_To_MongodbDatabaseThroughputSetting(destination *v20210515s.MongodbDatabaseThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec
	err := setting.Spec.AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.ThroughputSettingsGetResults_STATUS
	err = setting.Status.AssignProperties_To_ThroughputSettingsGetResults_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_ThroughputSettingsGetResults_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *MongodbDatabaseThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion,
		Kind:    "MongodbDatabaseThroughputSetting",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210515.MongodbDatabaseThroughputSetting
// Deprecated version of MongodbDatabaseThroughputSetting. Use v1beta20210515.MongodbDatabaseThroughputSetting instead
type MongodbDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseThroughputSetting `json:"items"`
}

// Storage version of v1alpha1api20210515.DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec
type DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec struct {
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/MongodbDatabase resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"MongodbDatabase"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource        `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec from the provided source
func (settings *DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec)
	if ok {
		// Populate our instance from source
		return settings.AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = settings.AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec
func (settings *DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec)
	if ok {
		// Populate destination from our instance
		return settings.AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec{}
	err := settings.AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec populates our DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec from the provided source DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec
func (settings *DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec) AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(source *v20210515s.DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Location
	settings.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	settings.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		settings.Owner = &owner
	} else {
		settings.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.AssignProperties_From_ThroughputSettingsResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_ThroughputSettingsResource() to populate field Resource")
		}
		settings.Resource = &resource
	} else {
		settings.Resource = nil
	}

	// Tags
	settings.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		settings.PropertyBag = propertyBag
	} else {
		settings.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec populates the provided destination DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec from our DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec
func (settings *DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec) AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec(destination *v20210515s.DatabaseAccounts_MongodbDatabases_ThroughputSettings_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(settings.PropertyBag)

	// Location
	destination.Location = genruntime.ClonePointerToString(settings.Location)

	// OriginalVersion
	destination.OriginalVersion = settings.OriginalVersion

	// Owner
	if settings.Owner != nil {
		owner := settings.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if settings.Resource != nil {
		var resource v20210515s.ThroughputSettingsResource
		err := settings.Resource.AssignProperties_To_ThroughputSettingsResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_ThroughputSettingsResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(settings.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseThroughputSetting{}, &MongodbDatabaseThroughputSettingList{})
}
