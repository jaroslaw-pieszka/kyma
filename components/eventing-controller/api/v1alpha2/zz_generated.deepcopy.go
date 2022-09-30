//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	if in.EmsSubscriptionStatus != nil {
		in, out := &in.EmsSubscriptionStatus, &out.EmsSubscriptionStatus
		*out = new(EmsSubscriptionStatus)
		**out = **in
	}
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]JetStreamTypes, len(*in))
		copy(*out, *in)
	}
	if in.EmsTypes != nil {
		in, out := &in.EmsTypes, &out.EmsTypes
		*out = make([]EventMeshTypes, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmsSubscriptionStatus) DeepCopyInto(out *EmsSubscriptionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmsSubscriptionStatus.
func (in *EmsSubscriptionStatus) DeepCopy() *EmsSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(EmsSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventMeshTypes) DeepCopyInto(out *EventMeshTypes) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventMeshTypes.
func (in *EventMeshTypes) DeepCopy() *EventMeshTypes {
	if in == nil {
		return nil
	}
	out := new(EventMeshTypes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventType) DeepCopyInto(out *EventType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventType.
func (in *EventType) DeepCopy() *EventType {
	if in == nil {
		return nil
	}
	out := new(EventType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JetStreamTypes) DeepCopyInto(out *JetStreamTypes) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JetStreamTypes.
func (in *JetStreamTypes) DeepCopy() *JetStreamTypes {
	if in == nil {
		return nil
	}
	out := new(JetStreamTypes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtocolSettings) DeepCopyInto(out *ProtocolSettings) {
	*out = *in
	if in.ContentMode != nil {
		in, out := &in.ContentMode, &out.ContentMode
		*out = new(string)
		**out = **in
	}
	if in.ExemptHandshake != nil {
		in, out := &in.ExemptHandshake, &out.ExemptHandshake
		*out = new(bool)
		**out = **in
	}
	if in.Qos != nil {
		in, out := &in.Qos, &out.Qos
		*out = new(string)
		**out = **in
	}
	if in.WebhookAuth != nil {
		in, out := &in.WebhookAuth, &out.WebhookAuth
		*out = new(WebhookAuth)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtocolSettings.
func (in *ProtocolSettings) DeepCopy() *ProtocolSettings {
	if in == nil {
		return nil
	}
	out := new(ProtocolSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subscription) DeepCopyInto(out *Subscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subscription.
func (in *Subscription) DeepCopy() *Subscription {
	if in == nil {
		return nil
	}
	out := new(Subscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionList) DeepCopyInto(out *SubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionList.
func (in *SubscriptionList) DeepCopy() *SubscriptionList {
	if in == nil {
		return nil
	}
	out := new(SubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionSpec) DeepCopyInto(out *SubscriptionSpec) {
	*out = *in
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionSpec.
func (in *SubscriptionSpec) DeepCopy() *SubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(SubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionStatus) DeepCopyInto(out *SubscriptionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]EventType, len(*in))
		copy(*out, *in)
	}
	in.Backend.DeepCopyInto(&out.Backend)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionStatus.
func (in *SubscriptionStatus) DeepCopy() *SubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookAuth) DeepCopyInto(out *WebhookAuth) {
	*out = *in
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookAuth.
func (in *WebhookAuth) DeepCopy() *WebhookAuth {
	if in == nil {
		return nil
	}
	out := new(WebhookAuth)
	in.DeepCopyInto(out)
	return out
}