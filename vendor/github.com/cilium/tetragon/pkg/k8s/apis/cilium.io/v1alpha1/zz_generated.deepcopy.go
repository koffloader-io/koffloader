//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionSelector) DeepCopyInto(out *ActionSelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionSelector.
func (in *ActionSelector) DeepCopy() *ActionSelector {
	if in == nil {
		return nil
	}
	out := new(ActionSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgSelector) DeepCopyInto(out *ArgSelector) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgSelector.
func (in *ArgSelector) DeepCopy() *ArgSelector {
	if in == nil {
		return nil
	}
	out := new(ArgSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BinarySelector) DeepCopyInto(out *BinarySelector) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BinarySelector.
func (in *BinarySelector) DeepCopy() *BinarySelector {
	if in == nil {
		return nil
	}
	out := new(BinarySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapabilitiesSelector) DeepCopyInto(out *CapabilitiesSelector) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapabilitiesSelector.
func (in *CapabilitiesSelector) DeepCopy() *CapabilitiesSelector {
	if in == nil {
		return nil
	}
	out := new(CapabilitiesSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnforcerSpec) DeepCopyInto(out *EnforcerSpec) {
	*out = *in
	if in.Calls != nil {
		in, out := &in.Calls, &out.Calls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnforcerSpec.
func (in *EnforcerSpec) DeepCopy() *EnforcerSpec {
	if in == nil {
		return nil
	}
	out := new(EnforcerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KProbeArg) DeepCopyInto(out *KProbeArg) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KProbeArg.
func (in *KProbeArg) DeepCopy() *KProbeArg {
	if in == nil {
		return nil
	}
	out := new(KProbeArg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KProbeSelector) DeepCopyInto(out *KProbeSelector) {
	*out = *in
	if in.MatchPIDs != nil {
		in, out := &in.MatchPIDs, &out.MatchPIDs
		*out = make([]PIDSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchArgs != nil {
		in, out := &in.MatchArgs, &out.MatchArgs
		*out = make([]ArgSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchActions != nil {
		in, out := &in.MatchActions, &out.MatchActions
		*out = make([]ActionSelector, len(*in))
		copy(*out, *in)
	}
	if in.MatchReturnArgs != nil {
		in, out := &in.MatchReturnArgs, &out.MatchReturnArgs
		*out = make([]ArgSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchReturnActions != nil {
		in, out := &in.MatchReturnActions, &out.MatchReturnActions
		*out = make([]ActionSelector, len(*in))
		copy(*out, *in)
	}
	if in.MatchBinaries != nil {
		in, out := &in.MatchBinaries, &out.MatchBinaries
		*out = make([]BinarySelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchNamespaces != nil {
		in, out := &in.MatchNamespaces, &out.MatchNamespaces
		*out = make([]NamespaceSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchNamespaceChanges != nil {
		in, out := &in.MatchNamespaceChanges, &out.MatchNamespaceChanges
		*out = make([]NamespaceChangesSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchCapabilities != nil {
		in, out := &in.MatchCapabilities, &out.MatchCapabilities
		*out = make([]CapabilitiesSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchCapabilityChanges != nil {
		in, out := &in.MatchCapabilityChanges, &out.MatchCapabilityChanges
		*out = make([]CapabilitiesSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KProbeSelector.
func (in *KProbeSelector) DeepCopy() *KProbeSelector {
	if in == nil {
		return nil
	}
	out := new(KProbeSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KProbeSpec) DeepCopyInto(out *KProbeSpec) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]KProbeArg, len(*in))
		copy(*out, *in)
	}
	if in.ReturnArg != nil {
		in, out := &in.ReturnArg, &out.ReturnArg
		*out = new(KProbeArg)
		**out = **in
	}
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]KProbeSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KProbeSpec.
func (in *KProbeSpec) DeepCopy() *KProbeSpec {
	if in == nil {
		return nil
	}
	out := new(KProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListSpec) DeepCopyInto(out *ListSpec) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Pattern != nil {
		in, out := &in.Pattern, &out.Pattern
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListSpec.
func (in *ListSpec) DeepCopy() *ListSpec {
	if in == nil {
		return nil
	}
	out := new(ListSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceChangesSelector) DeepCopyInto(out *NamespaceChangesSelector) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceChangesSelector.
func (in *NamespaceChangesSelector) DeepCopy() *NamespaceChangesSelector {
	if in == nil {
		return nil
	}
	out := new(NamespaceChangesSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceSelector) DeepCopyInto(out *NamespaceSelector) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceSelector.
func (in *NamespaceSelector) DeepCopy() *NamespaceSelector {
	if in == nil {
		return nil
	}
	out := new(NamespaceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionSpec) DeepCopyInto(out *OptionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionSpec.
func (in *OptionSpec) DeepCopy() *OptionSpec {
	if in == nil {
		return nil
	}
	out := new(OptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PIDSelector) DeepCopyInto(out *PIDSelector) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]uint32, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PIDSelector.
func (in *PIDSelector) DeepCopy() *PIDSelector {
	if in == nil {
		return nil
	}
	out := new(PIDSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodIP) DeepCopyInto(out *PodIP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodIP.
func (in *PodIP) DeepCopy() *PodIP {
	if in == nil {
		return nil
	}
	out := new(PodIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodInfo) DeepCopyInto(out *PodInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	out.WorkloadType = in.WorkloadType
	out.WorkloadObject = in.WorkloadObject
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodInfo.
func (in *PodInfo) DeepCopy() *PodInfo {
	if in == nil {
		return nil
	}
	out := new(PodInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodInfoList) DeepCopyInto(out *PodInfoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodInfoList.
func (in *PodInfoList) DeepCopy() *PodInfoList {
	if in == nil {
		return nil
	}
	out := new(PodInfoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodInfoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodInfoSpec) DeepCopyInto(out *PodInfoSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodInfoSpec.
func (in *PodInfoSpec) DeepCopy() *PodInfoSpec {
	if in == nil {
		return nil
	}
	out := new(PodInfoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodInfoStatus) DeepCopyInto(out *PodInfoStatus) {
	*out = *in
	if in.PodIPs != nil {
		in, out := &in.PodIPs, &out.PodIPs
		*out = make([]PodIP, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodInfoStatus.
func (in *PodInfoStatus) DeepCopy() *PodInfoStatus {
	if in == nil {
		return nil
	}
	out := new(PodInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracepointSpec) DeepCopyInto(out *TracepointSpec) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]KProbeArg, len(*in))
		copy(*out, *in)
	}
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]KProbeSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracepointSpec.
func (in *TracepointSpec) DeepCopy() *TracepointSpec {
	if in == nil {
		return nil
	}
	out := new(TracepointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingPolicy) DeepCopyInto(out *TracingPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingPolicy.
func (in *TracingPolicy) DeepCopy() *TracingPolicy {
	if in == nil {
		return nil
	}
	out := new(TracingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TracingPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingPolicyList) DeepCopyInto(out *TracingPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TracingPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingPolicyList.
func (in *TracingPolicyList) DeepCopy() *TracingPolicyList {
	if in == nil {
		return nil
	}
	out := new(TracingPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TracingPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingPolicyNamespaced) DeepCopyInto(out *TracingPolicyNamespaced) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingPolicyNamespaced.
func (in *TracingPolicyNamespaced) DeepCopy() *TracingPolicyNamespaced {
	if in == nil {
		return nil
	}
	out := new(TracingPolicyNamespaced)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TracingPolicyNamespaced) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingPolicyNamespacedList) DeepCopyInto(out *TracingPolicyNamespacedList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TracingPolicyNamespaced, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingPolicyNamespacedList.
func (in *TracingPolicyNamespacedList) DeepCopy() *TracingPolicyNamespacedList {
	if in == nil {
		return nil
	}
	out := new(TracingPolicyNamespacedList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TracingPolicyNamespacedList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracingPolicySpec) DeepCopyInto(out *TracingPolicySpec) {
	*out = *in
	if in.KProbes != nil {
		in, out := &in.KProbes, &out.KProbes
		*out = make([]KProbeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tracepoints != nil {
		in, out := &in.Tracepoints, &out.Tracepoints
		*out = make([]TracepointSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UProbes != nil {
		in, out := &in.UProbes, &out.UProbes
		*out = make([]UProbeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Lists != nil {
		in, out := &in.Lists, &out.Lists
		*out = make([]ListSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Enforcers != nil {
		in, out := &in.Enforcers, &out.Enforcers
		*out = make([]EnforcerSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]OptionSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracingPolicySpec.
func (in *TracingPolicySpec) DeepCopy() *TracingPolicySpec {
	if in == nil {
		return nil
	}
	out := new(TracingPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UProbeSpec) DeepCopyInto(out *UProbeSpec) {
	*out = *in
	if in.Symbols != nil {
		in, out := &in.Symbols, &out.Symbols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]KProbeSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]KProbeArg, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UProbeSpec.
func (in *UProbeSpec) DeepCopy() *UProbeSpec {
	if in == nil {
		return nil
	}
	out := new(UProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadObjectMeta) DeepCopyInto(out *WorkloadObjectMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadObjectMeta.
func (in *WorkloadObjectMeta) DeepCopy() *WorkloadObjectMeta {
	if in == nil {
		return nil
	}
	out := new(WorkloadObjectMeta)
	in.DeepCopyInto(out)
	return out
}