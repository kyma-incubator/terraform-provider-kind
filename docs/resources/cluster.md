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
        apiVersion = "kind.x-k8s.io/v1alpha4"
        nodes {
            role = "control-plane"
        }
        nodes {
            role: worker
        }
    }
}
```

## Argument Reference

* `name` - (Required) The kind name that is given to the created cluster.
* `node_image` - (Optional) The node_image that kind will use (ex: kindest/node:v1.15.3).
* `wait_for_ready` - (Optional) Defines wether or not the provider will wait for the control plane to be ready. Defaults to false.
* `kind_config` - (Optional) The kind_config that kind will use.
