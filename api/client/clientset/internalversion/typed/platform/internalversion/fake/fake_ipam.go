/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	platform "tkestack.io/tke/api/platform"
)

// FakeIPAMs implements IPAMInterface
type FakeIPAMs struct {
	Fake *FakePlatform
}

var ipamsResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "", Resource: "ipams"}

var ipamsKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "", Kind: "IPAM"}

// Get takes name of the iPAM, and returns the corresponding iPAM object, and an error if there is any.
func (c *FakeIPAMs) Get(name string, options v1.GetOptions) (result *platform.IPAM, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ipamsResource, name), &platform.IPAM{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.IPAM), err
}

// List takes label and field selectors, and returns the list of IPAMs that match those selectors.
func (c *FakeIPAMs) List(opts v1.ListOptions) (result *platform.IPAMList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ipamsResource, ipamsKind, opts), &platform.IPAMList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platform.IPAMList{ListMeta: obj.(*platform.IPAMList).ListMeta}
	for _, item := range obj.(*platform.IPAMList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPAMs.
func (c *FakeIPAMs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ipamsResource, opts))
}

// Create takes the representation of a iPAM and creates it.  Returns the server's representation of the iPAM, and an error, if there is any.
func (c *FakeIPAMs) Create(iPAM *platform.IPAM) (result *platform.IPAM, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ipamsResource, iPAM), &platform.IPAM{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.IPAM), err
}

// Update takes the representation of a iPAM and updates it. Returns the server's representation of the iPAM, and an error, if there is any.
func (c *FakeIPAMs) Update(iPAM *platform.IPAM) (result *platform.IPAM, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ipamsResource, iPAM), &platform.IPAM{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.IPAM), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIPAMs) UpdateStatus(iPAM *platform.IPAM) (*platform.IPAM, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(ipamsResource, "status", iPAM), &platform.IPAM{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.IPAM), err
}

// Delete takes name of the iPAM and deletes it. Returns an error if one occurs.
func (c *FakeIPAMs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ipamsResource, name), &platform.IPAM{})
	return err
}

// Patch applies the patch and returns the patched iPAM.
func (c *FakeIPAMs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *platform.IPAM, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ipamsResource, name, pt, data, subresources...), &platform.IPAM{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platform.IPAM), err
}
