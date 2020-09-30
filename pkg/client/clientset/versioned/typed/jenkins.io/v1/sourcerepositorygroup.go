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

package v1

import (
	"context"
	"time"

	v1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	scheme "github.com/jenkins-x/jx-api/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SourceRepositoryGroupsGetter has a method to return a SourceRepositoryGroupInterface.
// A group's client should implement this interface.
type SourceRepositoryGroupsGetter interface {
	SourceRepositoryGroups(namespace string) SourceRepositoryGroupInterface
}

// SourceRepositoryGroupInterface has methods to work with SourceRepositoryGroup resources.
type SourceRepositoryGroupInterface interface {
	Create(ctx context.Context, sourceRepositoryGroup *v1.SourceRepositoryGroup, opts metav1.CreateOptions) (*v1.SourceRepositoryGroup, error)
	Update(ctx context.Context, sourceRepositoryGroup *v1.SourceRepositoryGroup, opts metav1.UpdateOptions) (*v1.SourceRepositoryGroup, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SourceRepositoryGroup, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SourceRepositoryGroupList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SourceRepositoryGroup, err error)
	SourceRepositoryGroupExpansion
}

// sourceRepositoryGroups implements SourceRepositoryGroupInterface
type sourceRepositoryGroups struct {
	client rest.Interface
	ns     string
}

// newSourceRepositoryGroups returns a SourceRepositoryGroups
func newSourceRepositoryGroups(c *JenkinsV1Client, namespace string) *sourceRepositoryGroups {
	return &sourceRepositoryGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sourceRepositoryGroup, and returns the corresponding sourceRepositoryGroup object, and an error if there is any.
func (c *sourceRepositoryGroups) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SourceRepositoryGroup, err error) {
	result = &v1.SourceRepositoryGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SourceRepositoryGroups that match those selectors.
func (c *sourceRepositoryGroups) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SourceRepositoryGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SourceRepositoryGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sourceRepositoryGroups.
func (c *sourceRepositoryGroups) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sourceRepositoryGroup and creates it.  Returns the server's representation of the sourceRepositoryGroup, and an error, if there is any.
func (c *sourceRepositoryGroups) Create(ctx context.Context, sourceRepositoryGroup *v1.SourceRepositoryGroup, opts metav1.CreateOptions) (result *v1.SourceRepositoryGroup, err error) {
	result = &v1.SourceRepositoryGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sourceRepositoryGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sourceRepositoryGroup and updates it. Returns the server's representation of the sourceRepositoryGroup, and an error, if there is any.
func (c *sourceRepositoryGroups) Update(ctx context.Context, sourceRepositoryGroup *v1.SourceRepositoryGroup, opts metav1.UpdateOptions) (result *v1.SourceRepositoryGroup, err error) {
	result = &v1.SourceRepositoryGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		Name(sourceRepositoryGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sourceRepositoryGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sourceRepositoryGroup and deletes it. Returns an error if one occurs.
func (c *sourceRepositoryGroups) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sourceRepositoryGroups) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sourceRepositoryGroup.
func (c *sourceRepositoryGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SourceRepositoryGroup, err error) {
	result = &v1.SourceRepositoryGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sourcerepositorygroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
