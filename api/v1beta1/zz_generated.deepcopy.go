//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscCluster) DeepCopyInto(out *OscCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscCluster.
func (in *OscCluster) DeepCopy() *OscCluster {
	if in == nil {
		return nil
	}
	out := new(OscCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OscCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscClusterList) DeepCopyInto(out *OscClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OscCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscClusterList.
func (in *OscClusterList) DeepCopy() *OscClusterList {
	if in == nil {
		return nil
	}
	out := new(OscClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OscClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscClusterSpec) DeepCopyInto(out *OscClusterSpec) {
	*out = *in
	out.Network = in.Network
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscClusterSpec.
func (in *OscClusterSpec) DeepCopy() *OscClusterSpec {
	if in == nil {
		return nil
	}
	out := new(OscClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscClusterStatus) DeepCopyInto(out *OscClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscClusterStatus.
func (in *OscClusterStatus) DeepCopy() *OscClusterStatus {
	if in == nil {
		return nil
	}
	out := new(OscClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscLoadBalancer) DeepCopyInto(out *OscLoadBalancer) {
	*out = *in
	out.Listener = in.Listener
	out.HealthCheck = in.HealthCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscLoadBalancer.
func (in *OscLoadBalancer) DeepCopy() *OscLoadBalancer {
	if in == nil {
		return nil
	}
	out := new(OscLoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscLoadBalancerHealthCheck) DeepCopyInto(out *OscLoadBalancerHealthCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscLoadBalancerHealthCheck.
func (in *OscLoadBalancerHealthCheck) DeepCopy() *OscLoadBalancerHealthCheck {
	if in == nil {
		return nil
	}
	out := new(OscLoadBalancerHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscLoadBalancerListener) DeepCopyInto(out *OscLoadBalancerListener) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscLoadBalancerListener.
func (in *OscLoadBalancerListener) DeepCopy() *OscLoadBalancerListener {
	if in == nil {
		return nil
	}
	out := new(OscLoadBalancerListener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscMachine) DeepCopyInto(out *OscMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscMachine.
func (in *OscMachine) DeepCopy() *OscMachine {
	if in == nil {
		return nil
	}
	out := new(OscMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OscMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscMachineList) DeepCopyInto(out *OscMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OscMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscMachineList.
func (in *OscMachineList) DeepCopy() *OscMachineList {
	if in == nil {
		return nil
	}
	out := new(OscMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OscMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscMachineSpec) DeepCopyInto(out *OscMachineSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscMachineSpec.
func (in *OscMachineSpec) DeepCopy() *OscMachineSpec {
	if in == nil {
		return nil
	}
	out := new(OscMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscMachineStatus) DeepCopyInto(out *OscMachineStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscMachineStatus.
func (in *OscMachineStatus) DeepCopy() *OscMachineStatus {
	if in == nil {
		return nil
	}
	out := new(OscMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscNetwork) DeepCopyInto(out *OscNetwork) {
	*out = *in
	out.LoadBalancer = in.LoadBalancer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscNetwork.
func (in *OscNetwork) DeepCopy() *OscNetwork {
	if in == nil {
		return nil
	}
	out := new(OscNetwork)
	in.DeepCopyInto(out)
	return out
}
