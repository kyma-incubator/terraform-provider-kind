# Kind Provider

The Kind provider is used to interact with [Kubernetes IN Docker
(kind)](https://github.com/kubernetes-sigs/kind) to provision local
[Kubernetes](https://kubernetes.io) clusters.

## Example Usage

```hcl
# Configure the Kind Provider
provider "kind" {}

# Create a cluster
resource "kind_cluster" "default" {
    name = "test-cluster"
}
```
