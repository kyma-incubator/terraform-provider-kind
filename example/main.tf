provider "kind" {
}

resource "kind_cluster" "default" {
  name           = "test-cluster"
  node_image     = "kindest/node:v1.19.1"
  wait_for_ready = true
  kind_config = {
    kind       = "Cluster"
    apiVersion = "kind.x-k8s.io/v1alpha4"
    nodes = {
      "role" = "control-plane"
    }
    nodes = {
      "role" = "worker"
    }
  }
}

