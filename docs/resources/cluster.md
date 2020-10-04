# kind_cluster

Provides a Kind cluster resource. This can be used to create and delete Kind
clusters. It does NOT support modification to an existing kind cluster.

## Example Usage

```hcl
# Create a kind cluster of the name "test-cluster" with default kubernetes
# version specified in kind
# ref: https://github.com/kubernetes-sigs/kind/blob/master/pkg/apis/config/defaults/image.go#L21
resource "kind_cluster" "default" {
    name = "test-cluster"
}
```

To override the node image used:

```hcl
provider "kind" {}

# Create a cluster with kind of the name "test-cluster" with kubernetes version v1.16.1
resource "kind_cluster" "default" {
    name = "test-cluster"
    node_image = "kindest/node:v1.16.1"
}
```

To override the default kind config:

```hcl
provider "kind" {}

# creating a cluster with kind of the name "test-cluster" with kubernetes version v1.18.4 and two nodes
resource "kind_cluster" "default" {
    name = "test-cluster"
    node_image = "kindest/node:v1.18.4"
    kind_config  {
        kind = "Cluster"
        api_version = "kind.x-k8s.io/v1alpha4"
        node {
            role = "control-plane"
        }
        node {
            role: worker
        }
    }
}
```


```hcl
provider "kind" {}

# Create a cluster with patches applied to the containerd config
resource "kind_cluster" "default" {
    name = "test-cluster"
    node_image = "kindest/node:v1.16.1"
    kind_config = {
        containerd_config_patches = [
            <<-TOML
            [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:5000"]
                endpoint = ["http://kind-registry:5000"]
            TOML
        ]
    }
}
```

## Argument Reference

* `name` - (Required) The kind name that is given to the created cluster.
* `node_image` - (Optional) The node_image that kind will use (ex: kindest/node:v1.15.3).
* `wait_for_ready` - (Optional) Defines wether or not the provider will wait for the control plane to be ready. Defaults to false.
* `kind_config` - (Optional) The kind_config that kind will use.
