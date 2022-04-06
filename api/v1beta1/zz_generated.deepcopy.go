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
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscCluster) DeepCopyInto(out *OscCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
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
	in.Network.DeepCopyInto(&out.Network)
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
	in.Network.DeepCopyInto(&out.Network)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
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
func (in *OscInternetService) DeepCopyInto(out *OscInternetService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscInternetService.
func (in *OscInternetService) DeepCopy() *OscInternetService {
	if in == nil {
		return nil
	}
	out := new(OscInternetService)
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
func (in *OscNatService) DeepCopyInto(out *OscNatService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscNatService.
func (in *OscNatService) DeepCopy() *OscNatService {
	if in == nil {
		return nil
	}
	out := new(OscNatService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscNet) DeepCopyInto(out *OscNet) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscNet.
func (in *OscNet) DeepCopy() *OscNet {
	if in == nil {
		return nil
	}
	out := new(OscNet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscNetwork) DeepCopyInto(out *OscNetwork) {
	*out = *in
	out.LoadBalancer = in.LoadBalancer
	out.Net = in.Net
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]OscSubnet, len(*in))
		copy(*out, *in)
	}
	out.InternetService = in.InternetService
	out.NatService = in.NatService
	if in.RouteTables != nil {
		in, out := &in.RouteTables, &out.RouteTables
		*out = make([]OscRouteTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PublicIps != nil {
		in, out := &in.PublicIps, &out.PublicIps
		*out = make([]OscPublicIp, len(*in))
		copy(*out, *in)
	}
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscNetworkResource) DeepCopyInto(out *OscNetworkResource) {
	*out = *in
	in.LoadbalancerRef.DeepCopyInto(&out.LoadbalancerRef)
	in.NetRef.DeepCopyInto(&out.NetRef)
	in.SubnetRef.DeepCopyInto(&out.SubnetRef)
	in.InternetServiceRef.DeepCopyInto(&out.InternetServiceRef)
	in.RouteTablesRef.DeepCopyInto(&out.RouteTablesRef)
	in.LinkRouteTableRef.DeepCopyInto(&out.LinkRouteTableRef)
	in.RouteRef.DeepCopyInto(&out.RouteRef)
	in.PublicIpRef.DeepCopyInto(&out.PublicIpRef)
	in.NatServiceRef.DeepCopyInto(&out.NatServiceRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscNetworkResource.
func (in *OscNetworkResource) DeepCopy() *OscNetworkResource {
	if in == nil {
		return nil
	}
	out := new(OscNetworkResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscPublicIp) DeepCopyInto(out *OscPublicIp) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscPublicIp.
func (in *OscPublicIp) DeepCopy() *OscPublicIp {
	if in == nil {
		return nil
	}
	out := new(OscPublicIp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscResourceMapReference) DeepCopyInto(out *OscResourceMapReference) {
	*out = *in
	if in.ResourceMap != nil {
		in, out := &in.ResourceMap, &out.ResourceMap
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscResourceMapReference.
func (in *OscResourceMapReference) DeepCopy() *OscResourceMapReference {
	if in == nil {
		return nil
	}
	out := new(OscResourceMapReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscRoute) DeepCopyInto(out *OscRoute) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscRoute.
func (in *OscRoute) DeepCopy() *OscRoute {
	if in == nil {
		return nil
	}
	out := new(OscRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscRouteTable) DeepCopyInto(out *OscRouteTable) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]OscRoute, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscRouteTable.
func (in *OscRouteTable) DeepCopy() *OscRouteTable {
	if in == nil {
		return nil
	}
	out := new(OscRouteTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OscSubnet) DeepCopyInto(out *OscSubnet) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OscSubnet.
func (in *OscSubnet) DeepCopy() *OscSubnet {
	if in == nil {
		return nil
	}
	out := new(OscSubnet)
	in.DeepCopyInto(out)
	return out
}
