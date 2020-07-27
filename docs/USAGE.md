# Using The Provider

## Usage

Perform the following steps to use the provider:

1. Go to the provider [example](https://github.com/kyma-incubator/terraform-provider-kind/tree/master/example) folder:
    ```bash
    cd example
    ```
2. Edit the `main.tf` file and provide the following kind configuration:

```hcl
provider "kind" {}

# creating a cluster with kind of the name "test-cluster" with kubernetes version hardcoded in kind defaults https://github.com/kubernetes-sigs/kind/blob/master/pkg/apis/config/defaults/image.go#L21
resource "kind" "my-cluster" {
    name = "test-cluster"
}
```

To override the node image used, you can specify the `node_image` like so:

```hcl
provider "kind" {}

# creating a cluster with kind of the name "test-cluster" with kubernetes version v1.16.1
resource "kind" "my-cluster" {
    name = "test-cluster"
    node_image = "kindest/node:v1.16.1"
}
```

To override the default kind config, you can specify the `kind_config` with HEREDOC:

```hcl
provider "kind" {}

# creating a cluster with kind of the name "test-cluster" with kubernetes version v1.18.4 and two nodes
resource "kind" "my-cluster" {
    name = "test-cluster"
    node_image = "kindest/node:v1.18.4"
    kind_config =<<KIONF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
KIONF
}
```

1. Initialize Terraform:
    ```bash
    terraform init
    ```
2. Plan the provisioning:
    ```bash
    terraform plan
    ```
3. Deploy the cluster:
    ```bash
    terraform apply
    ```