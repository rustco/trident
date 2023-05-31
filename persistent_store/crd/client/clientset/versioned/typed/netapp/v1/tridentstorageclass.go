// Copyright 2023 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	scheme "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TridentStorageClassesGetter has a method to return a TridentStorageClassInterface.
// A group's client should implement this interface.
type TridentStorageClassesGetter interface {
	TridentStorageClasses(namespace string) TridentStorageClassInterface
}

// TridentStorageClassInterface has methods to work with TridentStorageClass resources.
type TridentStorageClassInterface interface {
	Create(ctx context.Context, tridentStorageClass *v1.TridentStorageClass, opts metav1.CreateOptions) (*v1.TridentStorageClass, error)
	Update(ctx context.Context, tridentStorageClass *v1.TridentStorageClass, opts metav1.UpdateOptions) (*v1.TridentStorageClass, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TridentStorageClass, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TridentStorageClassList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentStorageClass, err error)
	TridentStorageClassExpansion
}

// tridentStorageClasses implements TridentStorageClassInterface
type tridentStorageClasses struct {
	client rest.Interface
	ns     string
}

// newTridentStorageClasses returns a TridentStorageClasses
func newTridentStorageClasses(c *TridentV1Client, namespace string) *tridentStorageClasses {
	return &tridentStorageClasses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tridentStorageClass, and returns the corresponding tridentStorageClass object, and an error if there is any.
func (c *tridentStorageClasses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TridentStorageClass, err error) {
	result = &v1.TridentStorageClass{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentstorageclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TridentStorageClasses that match those selectors.
func (c *tridentStorageClasses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TridentStorageClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TridentStorageClassList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentstorageclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tridentStorageClasses.
func (c *tridentStorageClasses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tridentstorageclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tridentStorageClass and creates it.  Returns the server's representation of the tridentStorageClass, and an error, if there is any.
func (c *tridentStorageClasses) Create(ctx context.Context, tridentStorageClass *v1.TridentStorageClass, opts metav1.CreateOptions) (result *v1.TridentStorageClass, err error) {
	result = &v1.TridentStorageClass{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tridentstorageclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentStorageClass).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tridentStorageClass and updates it. Returns the server's representation of the tridentStorageClass, and an error, if there is any.
func (c *tridentStorageClasses) Update(ctx context.Context, tridentStorageClass *v1.TridentStorageClass, opts metav1.UpdateOptions) (result *v1.TridentStorageClass, err error) {
	result = &v1.TridentStorageClass{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tridentstorageclasses").
		Name(tridentStorageClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentStorageClass).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tridentStorageClass and deletes it. Returns an error if one occurs.
func (c *tridentStorageClasses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentstorageclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tridentStorageClasses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentstorageclasses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tridentStorageClass.
func (c *tridentStorageClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentStorageClass, err error) {
	result = &v1.TridentStorageClass{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tridentstorageclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
