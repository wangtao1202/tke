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
	auth "tkestack.io/tke/api/auth"
)

// FakeCategories implements CategoryInterface
type FakeCategories struct {
	Fake *FakeAuth
}

var categoriesResource = schema.GroupVersionResource{Group: "auth.tkestack.io", Version: "", Resource: "categories"}

var categoriesKind = schema.GroupVersionKind{Group: "auth.tkestack.io", Version: "", Kind: "Category"}

// Get takes name of the category, and returns the corresponding category object, and an error if there is any.
func (c *FakeCategories) Get(name string, options v1.GetOptions) (result *auth.Category, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(categoriesResource, name), &auth.Category{})
	if obj == nil {
		return nil, err
	}
	return obj.(*auth.Category), err
}

// List takes label and field selectors, and returns the list of Categories that match those selectors.
func (c *FakeCategories) List(opts v1.ListOptions) (result *auth.CategoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(categoriesResource, categoriesKind, opts), &auth.CategoryList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &auth.CategoryList{ListMeta: obj.(*auth.CategoryList).ListMeta}
	for _, item := range obj.(*auth.CategoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested categories.
func (c *FakeCategories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(categoriesResource, opts))
}

// Create takes the representation of a category and creates it.  Returns the server's representation of the category, and an error, if there is any.
func (c *FakeCategories) Create(category *auth.Category) (result *auth.Category, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(categoriesResource, category), &auth.Category{})
	if obj == nil {
		return nil, err
	}
	return obj.(*auth.Category), err
}

// Update takes the representation of a category and updates it. Returns the server's representation of the category, and an error, if there is any.
func (c *FakeCategories) Update(category *auth.Category) (result *auth.Category, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(categoriesResource, category), &auth.Category{})
	if obj == nil {
		return nil, err
	}
	return obj.(*auth.Category), err
}

// Delete takes name of the category and deletes it. Returns an error if one occurs.
func (c *FakeCategories) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(categoriesResource, name), &auth.Category{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCategories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(categoriesResource, listOptions)

	_, err := c.Fake.Invokes(action, &auth.CategoryList{})
	return err
}

// Patch applies the patch and returns the patched category.
func (c *FakeCategories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *auth.Category, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(categoriesResource, name, pt, data, subresources...), &auth.Category{})
	if obj == nil {
		return nil, err
	}
	return obj.(*auth.Category), err
}
