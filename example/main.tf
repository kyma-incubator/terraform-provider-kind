provider "kind" {}

resource "kind" "my-cluster" {
    name = "test-cluster"
}