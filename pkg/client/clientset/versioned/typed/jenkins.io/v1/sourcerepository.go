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

// SourceRepositoriesGetter has a method to return a SourceRepositoryInterface.
// A group's client should implement this interface.
type SourceRepositoriesGetter interface {
	SourceRepositories(namespace string) SourceRepositoryInterface
}

// SourceRepositoryInterface has methods to work with SourceRepository resources.
type SourceRepositoryInterface interface {
	Create(ctx context.Context, sourceRepository *v1.SourceRepository, opts metav1.CreateOptions) (*v1.SourceRepository, error)
	Update(ctx context.Context, sourceRepository *v1.SourceRepository, opts metav1.UpdateOptions) (*v1.SourceRepository, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SourceRepository, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SourceRepositoryList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SourceRepository, err error)
	SourceRepositoryExpansion
}

// sourceRepositories implements SourceRepositoryInterface
type sourceRepositories struct {
	client rest.Interface
	ns     string
}

// newSourceRepositories returns a SourceRepositories
func newSourceRepositories(c *JenkinsV1Client, namespace string) *sourceRepositories {
	return &sourceRepositories{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sourceRepository, and returns the corresponding sourceRepository object, and an error if there is any.
func (c *sourceRepositories) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SourceRepository, err error) {
	result = &v1.SourceRepository{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourcerepositories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SourceRepositories that match those selectors.
func (c *sourceRepositories) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SourceRepositoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SourceRepositoryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourcerepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sourceRepositories.
func (c *sourceRepositories) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sourcerepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sourceRepository and creates it.  Returns the server's representation of the sourceRepository, and an error, if there is any.
func (c *sourceRepositories) Create(ctx context.Context, sourceRepository *v1.SourceRepository, opts metav1.CreateOptions) (result *v1.SourceRepository, err error) {
	result = &v1.SourceRepository{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sourcerepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sourceRepository).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sourceRepository and updates it. Returns the server's representation of the sourceRepository, and an error, if there is any.
func (c *sourceRepositories) Update(ctx context.Context, sourceRepository *v1.SourceRepository, opts metav1.UpdateOptions) (result *v1.SourceRepository, err error) {
	result = &v1.SourceRepository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sourcerepositories").
		Name(sourceRepository.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sourceRepository).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sourceRepository and deletes it. Returns an error if one occurs.
func (c *sourceRepositories) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourcerepositories").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sourceRepositories) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourcerepositories").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sourceRepository.
func (c *sourceRepositories) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SourceRepository, err error) {
	result = &v1.SourceRepository{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sourcerepositories").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
