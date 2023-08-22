//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20210701storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedLocation) DeepCopyInto(out *ExtendedLocation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedLocation.
func (in *ExtendedLocation) DeepCopy() *ExtendedLocation {
	if in == nil {
		return nil
	}
	out := new(ExtendedLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedLocation_STATUS) DeepCopyInto(out *ExtendedLocation_STATUS) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedLocation_STATUS.
func (in *ExtendedLocation_STATUS) DeepCopy() *ExtendedLocation_STATUS {
	if in == nil {
		return nil
	}
	out := new(ExtendedLocation_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Image) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageDataDisk) DeepCopyInto(out *ImageDataDisk) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(string)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.Lun != nil {
		in, out := &in.Lun, &out.Lun
		*out = new(int)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDataDisk.
func (in *ImageDataDisk) DeepCopy() *ImageDataDisk {
	if in == nil {
		return nil
	}
	out := new(ImageDataDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageDataDisk_STATUS) DeepCopyInto(out *ImageDataDisk_STATUS) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(string)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.Lun != nil {
		in, out := &in.Lun, &out.Lun
		*out = new(int)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDataDisk_STATUS.
func (in *ImageDataDisk_STATUS) DeepCopy() *ImageDataDisk_STATUS {
	if in == nil {
		return nil
	}
	out := new(ImageDataDisk_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageList) DeepCopyInto(out *ImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Image, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageList.
func (in *ImageList) DeepCopy() *ImageList {
	if in == nil {
		return nil
	}
	out := new(ImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageOSDisk) DeepCopyInto(out *ImageOSDisk) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(string)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.OsState != nil {
		in, out := &in.OsState, &out.OsState
		*out = new(string)
		**out = **in
	}
	if in.OsType != nil {
		in, out := &in.OsType, &out.OsType
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageOSDisk.
func (in *ImageOSDisk) DeepCopy() *ImageOSDisk {
	if in == nil {
		return nil
	}
	out := new(ImageOSDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageOSDisk_STATUS) DeepCopyInto(out *ImageOSDisk_STATUS) {
	*out = *in
	if in.BlobUri != nil {
		in, out := &in.BlobUri, &out.BlobUri
		*out = new(string)
		**out = **in
	}
	if in.Caching != nil {
		in, out := &in.Caching, &out.Caching
		*out = new(string)
		**out = **in
	}
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.OsState != nil {
		in, out := &in.OsState, &out.OsState
		*out = new(string)
		**out = **in
	}
	if in.OsType != nil {
		in, out := &in.OsType, &out.OsType
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Snapshot != nil {
		in, out := &in.Snapshot, &out.Snapshot
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageAccountType != nil {
		in, out := &in.StorageAccountType, &out.StorageAccountType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageOSDisk_STATUS.
func (in *ImageOSDisk_STATUS) DeepCopy() *ImageOSDisk_STATUS {
	if in == nil {
		return nil
	}
	out := new(ImageOSDisk_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStorageProfile) DeepCopyInto(out *ImageStorageProfile) {
	*out = *in
	if in.DataDisks != nil {
		in, out := &in.DataDisks, &out.DataDisks
		*out = make([]ImageDataDisk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OsDisk != nil {
		in, out := &in.OsDisk, &out.OsDisk
		*out = new(ImageOSDisk)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ZoneResilient != nil {
		in, out := &in.ZoneResilient, &out.ZoneResilient
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStorageProfile.
func (in *ImageStorageProfile) DeepCopy() *ImageStorageProfile {
	if in == nil {
		return nil
	}
	out := new(ImageStorageProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStorageProfile_STATUS) DeepCopyInto(out *ImageStorageProfile_STATUS) {
	*out = *in
	if in.DataDisks != nil {
		in, out := &in.DataDisks, &out.DataDisks
		*out = make([]ImageDataDisk_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OsDisk != nil {
		in, out := &in.OsDisk, &out.OsDisk
		*out = new(ImageOSDisk_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ZoneResilient != nil {
		in, out := &in.ZoneResilient, &out.ZoneResilient
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStorageProfile_STATUS.
func (in *ImageStorageProfile_STATUS) DeepCopy() *ImageStorageProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(ImageStorageProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image_STATUS) DeepCopyInto(out *Image_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtendedLocation != nil {
		in, out := &in.ExtendedLocation, &out.ExtendedLocation
		*out = new(ExtendedLocation_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.HyperVGeneration != nil {
		in, out := &in.HyperVGeneration, &out.HyperVGeneration
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.SourceVirtualMachine != nil {
		in, out := &in.SourceVirtualMachine, &out.SourceVirtualMachine
		*out = new(SubResource_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(ImageStorageProfile_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image_STATUS.
func (in *Image_STATUS) DeepCopy() *Image_STATUS {
	if in == nil {
		return nil
	}
	out := new(Image_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image_Spec) DeepCopyInto(out *Image_Spec) {
	*out = *in
	if in.ExtendedLocation != nil {
		in, out := &in.ExtendedLocation, &out.ExtendedLocation
		*out = new(ExtendedLocation)
		(*in).DeepCopyInto(*out)
	}
	if in.HyperVGeneration != nil {
		in, out := &in.HyperVGeneration, &out.HyperVGeneration
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SourceVirtualMachine != nil {
		in, out := &in.SourceVirtualMachine, &out.SourceVirtualMachine
		*out = new(SubResource)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(ImageStorageProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image_Spec.
func (in *Image_Spec) DeepCopy() *Image_Spec {
	if in == nil {
		return nil
	}
	out := new(Image_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubResource) DeepCopyInto(out *SubResource) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubResource.
func (in *SubResource) DeepCopy() *SubResource {
	if in == nil {
		return nil
	}
	out := new(SubResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubResource_STATUS) DeepCopyInto(out *SubResource_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubResource_STATUS.
func (in *SubResource_STATUS) DeepCopy() *SubResource_STATUS {
	if in == nil {
		return nil
	}
	out := new(SubResource_STATUS)
	in.DeepCopyInto(out)
	return out
}
