/*
Copyright 2020 The Jenkins X Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	jenkinsiov1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBuildPacks implements BuildPackInterface
type FakeBuildPacks struct {
	Fake *FakeJenkinsV1
	ns   string
}

var buildpacksResource = schema.GroupVersionResource{Group: "jenkins.io", Version: "v1", Resource: "buildpacks"}

var buildpacksKind = schema.GroupVersionKind{Group: "jenkins.io", Version: "v1", Kind: "BuildPack"}

// Get takes name of the buildPack, and returns the corresponding buildPack object, and an error if there is any.
func (c *FakeBuildPacks) Get(ctx context.Context, name string, options v1.GetOptions) (result *jenkinsiov1.BuildPack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(buildpacksResource, c.ns, name), &jenkinsiov1.BuildPack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.BuildPack), err
}

// List takes label and field selectors, and returns the list of BuildPacks that match those selectors.
func (c *FakeBuildPacks) List(ctx context.Context, opts v1.ListOptions) (result *jenkinsiov1.BuildPackList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(buildpacksResource, buildpacksKind, c.ns, opts), &jenkinsiov1.BuildPackList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &jenkinsiov1.BuildPackList{ListMeta: obj.(*jenkinsiov1.BuildPackList).ListMeta}
	for _, item := range obj.(*jenkinsiov1.BuildPackList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested buildPacks.
func (c *FakeBuildPacks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(buildpacksResource, c.ns, opts))

}

// Create takes the representation of a buildPack and creates it.  Returns the server's representation of the buildPack, and an error, if there is any.
func (c *FakeBuildPacks) Create(ctx context.Context, buildPack *jenkinsiov1.BuildPack, opts v1.CreateOptions) (result *jenkinsiov1.BuildPack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(buildpacksResource, c.ns, buildPack), &jenkinsiov1.BuildPack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.BuildPack), err
}

// Update takes the representation of a buildPack and updates it. Returns the server's representation of the buildPack, and an error, if there is any.
func (c *FakeBuildPacks) Update(ctx context.Context, buildPack *jenkinsiov1.BuildPack, opts v1.UpdateOptions) (result *jenkinsiov1.BuildPack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(buildpacksResource, c.ns, buildPack), &jenkinsiov1.BuildPack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.BuildPack), err
}

// Delete takes name of the buildPack and deletes it. Returns an error if one occurs.
func (c *FakeBuildPacks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(buildpacksResource, c.ns, name), &jenkinsiov1.BuildPack{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuildPacks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(buildpacksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &jenkinsiov1.BuildPackList{})
	return err
}

// Patch applies the patch and returns the patched buildPack.
func (c *FakeBuildPacks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *jenkinsiov1.BuildPack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildpacksResource, c.ns, name, pt, data, subresources...), &jenkinsiov1.BuildPack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.BuildPack), err
}
