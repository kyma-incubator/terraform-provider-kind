
# Terraform Provider for kind


## Overview

The Terraform Provider for kind enables [Terraform](https://www.terraform.io) to provision local [Kubernetes](https://kubernetes.io) clusters on base of [Kubernetes IN Docker (kind)](https://github.com/kubernetes-sigs/kind).

## Prerequisites

- [Terraform](https://www.terraform.io/downloads.html) 0.10+
- [Go](https://golang.org/doc/install) 1.12 or higher

## Development

Perform the following steps to build the providers:

1. Build the provider:
    ```bash
    go build -o terraform-provider-kind
    ```
2. Move the provider binary into the terraform plugins folder.

    >**NOTE**: For details on Terraform plugins see [this](https://www.terraform.io/docs/plugins/basics.html#installing-plugins) document.

## Usage

Perform the following steps to use the provider:

1. Go to the provider [example](https://github.com/kyma-incubator/terraform-provider-kind/tree/master/example) folder:
    ```bash
    cd example
    ```
2. Edit the `main.tf` file and provide the following kind configuration:
```
provider "kind" {}

# creating a cluster with kind of the name "test-cluster" with kubernetes version hardcoded in kind defaults https://github.com/kubernetes-sigs/kind/blob/master/pkg/apis/config/defaults/image.go#L21
resource "kind" "my-cluster" {
    name = "test-cluster"
}
```

To override the node image used, you can specify the `node_image` like so:
```
provider "kind" {}

# creating a cluster with kind of the name "test-cluster" with kubernetes version v1.16.1
resource "kind" "my-cluster" {
    name = "test-cluster"
    node_image = "kindest/node:v1.16.1"
}
```

To override the default kind config, you can specify the `kind_config` with HEREDOC:
```
provider "kind" {}

# creating a cluster with kind of the name "test-cluster" with kubernetes version v1.15.7 and two nodes
resource "kind" "my-cluster" {
    name = "test-cluster"
    node_image = "kindest/node:v1.15.7"
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
