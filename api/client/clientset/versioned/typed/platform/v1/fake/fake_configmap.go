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
	platformv1 "tkestack.io/tke/api/platform/v1"
)

// FakeConfigMaps implements ConfigMapInterface
type FakeConfigMaps struct {
	Fake *FakePlatformV1
}

var configmapsResource = schema.GroupVersionResource{Group: "platform.tkestack.io", Version: "v1", Resource: "configmaps"}

var configmapsKind = schema.GroupVersionKind{Group: "platform.tkestack.io", Version: "v1", Kind: "ConfigMap"}

// Get takes name of the configMap, and returns the corresponding configMap object, and an error if there is any.
func (c *FakeConfigMaps) Get(name string, options v1.GetOptions) (result *platformv1.ConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(configmapsResource, name), &platformv1.ConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.ConfigMap), err
}

// List takes label and field selectors, and returns the list of ConfigMaps that match those selectors.
func (c *FakeConfigMaps) List(opts v1.ListOptions) (result *platformv1.ConfigMapList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(configmapsResource, configmapsKind, opts), &platformv1.ConfigMapList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &platformv1.ConfigMapList{ListMeta: obj.(*platformv1.ConfigMapList).ListMeta}
	for _, item := range obj.(*platformv1.ConfigMapList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configMaps.
func (c *FakeConfigMaps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(configmapsResource, opts))
}

// Create takes the representation of a configMap and creates it.  Returns the server's representation of the configMap, and an error, if there is any.
func (c *FakeConfigMaps) Create(configMap *platformv1.ConfigMap) (result *platformv1.ConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(configmapsResource, configMap), &platformv1.ConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.ConfigMap), err
}

// Update takes the representation of a configMap and updates it. Returns the server's representation of the configMap, and an error, if there is any.
func (c *FakeConfigMaps) Update(configMap *platformv1.ConfigMap) (result *platformv1.ConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(configmapsResource, configMap), &platformv1.ConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.ConfigMap), err
}

// Delete takes name of the configMap and deletes it. Returns an error if one occurs.
func (c *FakeConfigMaps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(configmapsResource, name), &platformv1.ConfigMap{})
	return err
}

// Patch applies the patch and returns the patched configMap.
func (c *FakeConfigMaps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *platformv1.ConfigMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(configmapsResource, name, pt, data, subresources...), &platformv1.ConfigMap{})
	if obj == nil {
		return nil, err
	}
	return obj.(*platformv1.ConfigMap), err
}
