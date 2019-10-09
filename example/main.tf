provider "kind" {
    
}

resource "kind" "my-cluster" {
    name = "test-cluster"
    k8s_version = "v1.15.3"
}