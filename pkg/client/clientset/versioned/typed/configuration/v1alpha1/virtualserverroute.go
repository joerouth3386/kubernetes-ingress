// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/nginxinc/kubernetes-ingress/pkg/apis/configuration/v1alpha1"
	scheme "github.com/nginxinc/kubernetes-ingress/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualServerRoutesGetter has a method to return a VirtualServerRouteInterface.
// A group's client should implement this interface.
type VirtualServerRoutesGetter interface {
	VirtualServerRoutes(namespace string) VirtualServerRouteInterface
}

// VirtualServerRouteInterface has methods to work with VirtualServerRoute resources.
type VirtualServerRouteInterface interface {
	Create(*v1alpha1.VirtualServerRoute) (*v1alpha1.VirtualServerRoute, error)
	Update(*v1alpha1.VirtualServerRoute) (*v1alpha1.VirtualServerRoute, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VirtualServerRoute, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualServerRouteList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualServerRoute, err error)
	VirtualServerRouteExpansion
}

// virtualServerRoutes implements VirtualServerRouteInterface
type virtualServerRoutes struct {
	client rest.Interface
	ns     string
}

// newVirtualServerRoutes returns a VirtualServerRoutes
func newVirtualServerRoutes(c *K8sV1alpha1Client, namespace string) *virtualServerRoutes {
	return &virtualServerRoutes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualServerRoute, and returns the corresponding virtualServerRoute object, and an error if there is any.
func (c *virtualServerRoutes) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualServerRoute, err error) {
	result = &v1alpha1.VirtualServerRoute{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualserverroutes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualServerRoutes that match those selectors.
func (c *virtualServerRoutes) List(opts v1.ListOptions) (result *v1alpha1.VirtualServerRouteList, err error) {
	result = &v1alpha1.VirtualServerRouteList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualserverroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualServerRoutes.
func (c *virtualServerRoutes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualserverroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a virtualServerRoute and creates it.  Returns the server's representation of the virtualServerRoute, and an error, if there is any.
func (c *virtualServerRoutes) Create(virtualServerRoute *v1alpha1.VirtualServerRoute) (result *v1alpha1.VirtualServerRoute, err error) {
	result = &v1alpha1.VirtualServerRoute{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualserverroutes").
		Body(virtualServerRoute).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualServerRoute and updates it. Returns the server's representation of the virtualServerRoute, and an error, if there is any.
func (c *virtualServerRoutes) Update(virtualServerRoute *v1alpha1.VirtualServerRoute) (result *v1alpha1.VirtualServerRoute, err error) {
	result = &v1alpha1.VirtualServerRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualserverroutes").
		Name(virtualServerRoute.Name).
		Body(virtualServerRoute).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualServerRoute and deletes it. Returns an error if one occurs.
func (c *virtualServerRoutes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualserverroutes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualServerRoutes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualserverroutes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualServerRoute.
func (c *virtualServerRoutes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualServerRoute, err error) {
	result = &v1alpha1.VirtualServerRoute{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualserverroutes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}