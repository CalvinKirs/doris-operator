//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseSpec) DeepCopyInto(out *BaseSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ExportService)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]ConfigMap, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HostAliases != nil {
		in, out := &in.HostAliases, &out.HostAliases
		*out = make([]corev1.HostAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PersistentVolume != nil {
		in, out := &in.PersistentVolume, &out.PersistentVolume
		*out = new(PersistentVolume)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseSpec.
func (in *BaseSpec) DeepCopy() *BaseSpec {
	if in == nil {
		return nil
	}
	out := new(BaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseStatus) DeepCopyInto(out *BaseStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseStatus.
func (in *BaseStatus) DeepCopy() *BaseStatus {
	if in == nil {
		return nil
	}
	out := new(BaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMap) DeepCopyInto(out *ConfigMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMap.
func (in *ConfigMap) DeepCopy() *ConfigMap {
	if in == nil {
		return nil
	}
	out := new(ConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DorisDisaggregatedMetaService) DeepCopyInto(out *DorisDisaggregatedMetaService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DorisDisaggregatedMetaService.
func (in *DorisDisaggregatedMetaService) DeepCopy() *DorisDisaggregatedMetaService {
	if in == nil {
		return nil
	}
	out := new(DorisDisaggregatedMetaService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DorisDisaggregatedMetaService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DorisDisaggregatedMetaServiceList) DeepCopyInto(out *DorisDisaggregatedMetaServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DorisDisaggregatedMetaService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DorisDisaggregatedMetaServiceList.
func (in *DorisDisaggregatedMetaServiceList) DeepCopy() *DorisDisaggregatedMetaServiceList {
	if in == nil {
		return nil
	}
	out := new(DorisDisaggregatedMetaServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DorisDisaggregatedMetaServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DorisDisaggregatedMetaServiceSpec) DeepCopyInto(out *DorisDisaggregatedMetaServiceSpec) {
	*out = *in
	if in.FDB != nil {
		in, out := &in.FDB, &out.FDB
		*out = new(FoundationDB)
		(*in).DeepCopyInto(*out)
	}
	if in.MS != nil {
		in, out := &in.MS, &out.MS
		*out = new(MetaService)
		(*in).DeepCopyInto(*out)
	}
	if in.Recycler != nil {
		in, out := &in.Recycler, &out.Recycler
		*out = new(Recycler)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DorisDisaggregatedMetaServiceSpec.
func (in *DorisDisaggregatedMetaServiceSpec) DeepCopy() *DorisDisaggregatedMetaServiceSpec {
	if in == nil {
		return nil
	}
	out := new(DorisDisaggregatedMetaServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DorisDisaggregatedMetaServiceStatus) DeepCopyInto(out *DorisDisaggregatedMetaServiceStatus) {
	*out = *in
	out.FDBStatus = in.FDBStatus
	out.MSStatus = in.MSStatus
	out.RecyclerStatus = in.RecyclerStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DorisDisaggregatedMetaServiceStatus.
func (in *DorisDisaggregatedMetaServiceStatus) DeepCopy() *DorisDisaggregatedMetaServiceStatus {
	if in == nil {
		return nil
	}
	out := new(DorisDisaggregatedMetaServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportService) DeepCopyInto(out *ExportService) {
	*out = *in
	if in.PortMaps != nil {
		in, out := &in.PortMaps, &out.PortMaps
		*out = make([]PortMap, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportService.
func (in *ExportService) DeepCopy() *ExportService {
	if in == nil {
		return nil
	}
	out := new(ExportService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDBStatus) DeepCopyInto(out *FDBStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FDBStatus.
func (in *FDBStatus) DeepCopy() *FDBStatus {
	if in == nil {
		return nil
	}
	out := new(FDBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoundationDB) DeepCopyInto(out *FoundationDB) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.VolumeClaimTemplate != nil {
		in, out := &in.VolumeClaimTemplate, &out.VolumeClaimTemplate
		*out = new(corev1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoundationDB.
func (in *FoundationDB) DeepCopy() *FoundationDB {
	if in == nil {
		return nil
	}
	out := new(FoundationDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaService) DeepCopyInto(out *MetaService) {
	*out = *in
	in.BaseSpec.DeepCopyInto(&out.BaseSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaService.
func (in *MetaService) DeepCopy() *MetaService {
	if in == nil {
		return nil
	}
	out := new(MetaService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolume) DeepCopyInto(out *PersistentVolume) {
	*out = *in
	in.PersistentVolumeClaimSpec.DeepCopyInto(&out.PersistentVolumeClaimSpec)
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolume.
func (in *PersistentVolume) DeepCopy() *PersistentVolume {
	if in == nil {
		return nil
	}
	out := new(PersistentVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortMap) DeepCopyInto(out *PortMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortMap.
func (in *PortMap) DeepCopy() *PortMap {
	if in == nil {
		return nil
	}
	out := new(PortMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Recycler) DeepCopyInto(out *Recycler) {
	*out = *in
	in.BaseSpec.DeepCopyInto(&out.BaseSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Recycler.
func (in *Recycler) DeepCopy() *Recycler {
	if in == nil {
		return nil
	}
	out := new(Recycler)
	in.DeepCopyInto(out)
	return out
}
