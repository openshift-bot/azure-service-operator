// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import (
	"fmt"
	alpha20210401s "github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Deprecated version of StorageAccountsQueueService. Use v1beta20210401.StorageAccountsQueueService instead
type StorageAccountsQueueService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccounts_QueueServices_Spec `json:"spec,omitempty"`
	Status            QueueServiceProperties_STATUS      `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsQueueService{}

// GetConditions returns the conditions of the resource
func (service *StorageAccountsQueueService) GetConditions() conditions.Conditions {
	return service.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (service *StorageAccountsQueueService) SetConditions(conditions conditions.Conditions) {
	service.Status.Conditions = conditions
}

var _ conversion.Convertible = &StorageAccountsQueueService{}

// ConvertFrom populates our StorageAccountsQueueService from the provided hub StorageAccountsQueueService
func (service *StorageAccountsQueueService) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20210401s.StorageAccountsQueueService

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = service.AssignProperties_From_StorageAccountsQueueService(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to service")
	}

	return nil
}

// ConvertTo populates the provided hub StorageAccountsQueueService from our StorageAccountsQueueService
func (service *StorageAccountsQueueService) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20210401s.StorageAccountsQueueService
	err := service.AssignProperties_To_StorageAccountsQueueService(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from service")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-storage-azure-com-v1alpha1api20210401-storageaccountsqueueservice,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsqueueservices,verbs=create;update,versions=v1alpha1api20210401,name=default.v1alpha1api20210401.storageaccountsqueueservices.storage.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &StorageAccountsQueueService{}

// Default applies defaults to the StorageAccountsQueueService resource
func (service *StorageAccountsQueueService) Default() {
	service.defaultImpl()
	var temp interface{} = service
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the StorageAccountsQueueService resource
func (service *StorageAccountsQueueService) defaultImpl() {}

var _ genruntime.KubernetesResource = &StorageAccountsQueueService{}

// AzureName returns the Azure name of the resource (always "default")
func (service *StorageAccountsQueueService) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (service StorageAccountsQueueService) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (service *StorageAccountsQueueService) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (service *StorageAccountsQueueService) GetSpec() genruntime.ConvertibleSpec {
	return &service.Spec
}

// GetStatus returns the status of this resource
func (service *StorageAccountsQueueService) GetStatus() genruntime.ConvertibleStatus {
	return &service.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/queueServices"
func (service *StorageAccountsQueueService) GetType() string {
	return "Microsoft.Storage/storageAccounts/queueServices"
}

// NewEmptyStatus returns a new empty (blank) status
func (service *StorageAccountsQueueService) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &QueueServiceProperties_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (service *StorageAccountsQueueService) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(service.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  service.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (service *StorageAccountsQueueService) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*QueueServiceProperties_STATUS); ok {
		service.Status = *st
		return nil
	}

	// Convert status to required version
	var st QueueServiceProperties_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	service.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-storage-azure-com-v1alpha1api20210401-storageaccountsqueueservice,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsqueueservices,verbs=create;update,versions=v1alpha1api20210401,name=validate.v1alpha1api20210401.storageaccountsqueueservices.storage.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &StorageAccountsQueueService{}

// ValidateCreate validates the creation of the resource
func (service *StorageAccountsQueueService) ValidateCreate() error {
	validations := service.createValidations()
	var temp interface{} = service
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (service *StorageAccountsQueueService) ValidateDelete() error {
	validations := service.deleteValidations()
	var temp interface{} = service
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (service *StorageAccountsQueueService) ValidateUpdate(old runtime.Object) error {
	validations := service.updateValidations()
	var temp interface{} = service
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (service *StorageAccountsQueueService) createValidations() []func() error {
	return []func() error{service.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (service *StorageAccountsQueueService) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (service *StorageAccountsQueueService) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return service.validateResourceReferences()
		},
		service.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (service *StorageAccountsQueueService) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&service.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (service *StorageAccountsQueueService) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*StorageAccountsQueueService)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, service)
}

// AssignProperties_From_StorageAccountsQueueService populates our StorageAccountsQueueService from the provided source StorageAccountsQueueService
func (service *StorageAccountsQueueService) AssignProperties_From_StorageAccountsQueueService(source *alpha20210401s.StorageAccountsQueueService) error {

	// ObjectMeta
	service.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec StorageAccounts_QueueServices_Spec
	err := spec.AssignProperties_From_StorageAccounts_QueueServices_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_StorageAccounts_QueueServices_Spec() to populate field Spec")
	}
	service.Spec = spec

	// Status
	var status QueueServiceProperties_STATUS
	err = status.AssignProperties_From_QueueServiceProperties_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_QueueServiceProperties_STATUS() to populate field Status")
	}
	service.Status = status

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsQueueService populates the provided destination StorageAccountsQueueService from our StorageAccountsQueueService
func (service *StorageAccountsQueueService) AssignProperties_To_StorageAccountsQueueService(destination *alpha20210401s.StorageAccountsQueueService) error {

	// ObjectMeta
	destination.ObjectMeta = *service.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20210401s.StorageAccounts_QueueServices_Spec
	err := service.Spec.AssignProperties_To_StorageAccounts_QueueServices_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_StorageAccounts_QueueServices_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20210401s.QueueServiceProperties_STATUS
	err = service.Status.AssignProperties_To_QueueServiceProperties_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_QueueServiceProperties_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (service *StorageAccountsQueueService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: service.Spec.OriginalVersion(),
		Kind:    "StorageAccountsQueueService",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of StorageAccountsQueueService. Use v1beta20210401.StorageAccountsQueueService instead
type StorageAccountsQueueServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsQueueService `json:"items"`
}

// Deprecated version of QueueServiceProperties_STATUS. Use v1beta20210401.QueueServiceProperties_STATUS instead
type QueueServiceProperties_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`
	Cors       *CorsRules_STATUS      `json:"cors,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &QueueServiceProperties_STATUS{}

// ConvertStatusFrom populates our QueueServiceProperties_STATUS from the provided source
func (properties *QueueServiceProperties_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20210401s.QueueServiceProperties_STATUS)
	if ok {
		// Populate our instance from source
		return properties.AssignProperties_From_QueueServiceProperties_STATUS(src)
	}

	// Convert to an intermediate form
	src = &alpha20210401s.QueueServiceProperties_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = properties.AssignProperties_From_QueueServiceProperties_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our QueueServiceProperties_STATUS
func (properties *QueueServiceProperties_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20210401s.QueueServiceProperties_STATUS)
	if ok {
		// Populate destination from our instance
		return properties.AssignProperties_To_QueueServiceProperties_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210401s.QueueServiceProperties_STATUS{}
	err := properties.AssignProperties_To_QueueServiceProperties_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &QueueServiceProperties_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (properties *QueueServiceProperties_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &QueueServiceProperties_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (properties *QueueServiceProperties_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(QueueServiceProperties_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected QueueServiceProperties_STATUSARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Cors’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Cors != nil {
			var cors1 CorsRules_STATUS
			err := cors1.PopulateFromARM(owner, *typedInput.Properties.Cors)
			if err != nil {
				return err
			}
			cors := cors1
			properties.Cors = &cors
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		properties.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		properties.Name = &name
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		properties.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_QueueServiceProperties_STATUS populates our QueueServiceProperties_STATUS from the provided source QueueServiceProperties_STATUS
func (properties *QueueServiceProperties_STATUS) AssignProperties_From_QueueServiceProperties_STATUS(source *alpha20210401s.QueueServiceProperties_STATUS) error {

	// Conditions
	properties.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Cors
	if source.Cors != nil {
		var cor CorsRules_STATUS
		err := cor.AssignProperties_From_CorsRules_STATUS(source.Cors)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_CorsRules_STATUS() to populate field Cors")
		}
		properties.Cors = &cor
	} else {
		properties.Cors = nil
	}

	// Id
	properties.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	properties.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	properties.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_QueueServiceProperties_STATUS populates the provided destination QueueServiceProperties_STATUS from our QueueServiceProperties_STATUS
func (properties *QueueServiceProperties_STATUS) AssignProperties_To_QueueServiceProperties_STATUS(destination *alpha20210401s.QueueServiceProperties_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(properties.Conditions)

	// Cors
	if properties.Cors != nil {
		var cor alpha20210401s.CorsRules_STATUS
		err := properties.Cors.AssignProperties_To_CorsRules_STATUS(&cor)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_CorsRules_STATUS() to populate field Cors")
		}
		destination.Cors = &cor
	} else {
		destination.Cors = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(properties.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(properties.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(properties.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type StorageAccounts_QueueServices_Spec struct {
	Cors     *CorsRules `json:"cors,omitempty"`
	Location *string    `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccount resource
	Owner *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccount"`
	Tags  map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &StorageAccounts_QueueServices_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (services *StorageAccounts_QueueServices_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if services == nil {
		return nil, nil
	}
	result := &StorageAccounts_QueueServices_SpecARM{}

	// Set property ‘Location’:
	if services.Location != nil {
		location := *services.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if services.Cors != nil {
		result.Properties = &QueueServicePropertiesPropertiesARM{}
	}
	if services.Cors != nil {
		corsARM, err := (*services.Cors).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		cors := *corsARM.(*CorsRulesARM)
		result.Properties.Cors = &cors
	}

	// Set property ‘Tags’:
	if services.Tags != nil {
		result.Tags = make(map[string]string, len(services.Tags))
		for key, value := range services.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (services *StorageAccounts_QueueServices_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &StorageAccounts_QueueServices_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (services *StorageAccounts_QueueServices_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(StorageAccounts_QueueServices_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected StorageAccounts_QueueServices_SpecARM, got %T", armInput)
	}

	// Set property ‘Cors’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Cors != nil {
			var cors1 CorsRules
			err := cors1.PopulateFromARM(owner, *typedInput.Properties.Cors)
			if err != nil {
				return err
			}
			cors := cors1
			services.Cors = &cors
		}
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		services.Location = &location
	}

	// Set property ‘Owner’:
	services.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		services.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			services.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &StorageAccounts_QueueServices_Spec{}

// ConvertSpecFrom populates our StorageAccounts_QueueServices_Spec from the provided source
func (services *StorageAccounts_QueueServices_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20210401s.StorageAccounts_QueueServices_Spec)
	if ok {
		// Populate our instance from source
		return services.AssignProperties_From_StorageAccounts_QueueServices_Spec(src)
	}

	// Convert to an intermediate form
	src = &alpha20210401s.StorageAccounts_QueueServices_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = services.AssignProperties_From_StorageAccounts_QueueServices_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our StorageAccounts_QueueServices_Spec
func (services *StorageAccounts_QueueServices_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20210401s.StorageAccounts_QueueServices_Spec)
	if ok {
		// Populate destination from our instance
		return services.AssignProperties_To_StorageAccounts_QueueServices_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210401s.StorageAccounts_QueueServices_Spec{}
	err := services.AssignProperties_To_StorageAccounts_QueueServices_Spec(dst)
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

// AssignProperties_From_StorageAccounts_QueueServices_Spec populates our StorageAccounts_QueueServices_Spec from the provided source StorageAccounts_QueueServices_Spec
func (services *StorageAccounts_QueueServices_Spec) AssignProperties_From_StorageAccounts_QueueServices_Spec(source *alpha20210401s.StorageAccounts_QueueServices_Spec) error {

	// Cors
	if source.Cors != nil {
		var cor CorsRules
		err := cor.AssignProperties_From_CorsRules(source.Cors)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_CorsRules() to populate field Cors")
		}
		services.Cors = &cor
	} else {
		services.Cors = nil
	}

	// Location
	services.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		services.Owner = &owner
	} else {
		services.Owner = nil
	}

	// Tags
	services.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_StorageAccounts_QueueServices_Spec populates the provided destination StorageAccounts_QueueServices_Spec from our StorageAccounts_QueueServices_Spec
func (services *StorageAccounts_QueueServices_Spec) AssignProperties_To_StorageAccounts_QueueServices_Spec(destination *alpha20210401s.StorageAccounts_QueueServices_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Cors
	if services.Cors != nil {
		var cor alpha20210401s.CorsRules
		err := services.Cors.AssignProperties_To_CorsRules(&cor)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_CorsRules() to populate field Cors")
		}
		destination.Cors = &cor
	} else {
		destination.Cors = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(services.Location)

	// OriginalVersion
	destination.OriginalVersion = services.OriginalVersion()

	// Owner
	if services.Owner != nil {
		owner := services.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(services.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (services *StorageAccounts_QueueServices_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

func init() {
	SchemeBuilder.Register(&StorageAccountsQueueService{}, &StorageAccountsQueueServiceList{})
}
