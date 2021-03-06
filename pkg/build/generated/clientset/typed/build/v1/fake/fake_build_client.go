package fake

import (
	v1 "github.com/openshift/origin/pkg/build/generated/clientset/typed/build/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeBuildV1 struct {
	*testing.Fake
}

func (c *FakeBuildV1) Builds(namespace string) v1.BuildResourceInterface {
	return &FakeBuilds{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBuildV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
