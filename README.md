Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Maintainers
-----------

Daniel Roth (@tehcyx)

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Usage
---------------------

```
provider "kind" {

}
# creating a cluster with kind of the name "test-cluster" with kubernetes version v1.15.3
resource "kind" "my-cluster" {
    name = "test-cluster"
    k8s_version = "v1.15.3"
}
```

Building and Testing The Provider
---------------------

You will need go 1.11+ installed to use go modules.


```bash
# Build
go build -o example/terraform.d/plugins/darwin_amd64/terraform-provider-kind
# Test
cd example
terraform plan
```