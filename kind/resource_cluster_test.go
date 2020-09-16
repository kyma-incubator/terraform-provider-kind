package kind

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	kindDefaults "sigs.k8s.io/kind/pkg/apis/config/defaults"
	"sigs.k8s.io/kind/pkg/cluster"
)

func init() {
	resource.AddTestSweepers("kind_cluster", &resource.Sweeper{
		Name: "kind_cluster",
		F:    testSweepKindCluster,
	})
}

func testSweepKindCluster(name string) error {
	//TODO: needs code to cleanup test clusters

	return nil
}

const nodeImage = "kindest/node:latest"

func TestAccCluster(t *testing.T) {
	resourceName := "kind_cluster.test"
	clusterName := acctest.RandomWithPrefix("tf-acc-test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKindClusterResourceDestroy(clusterName),
		Steps: []resource.TestStep{
			{
				Config: testAccBasicClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckNoResourceAttr(resourceName, "node_image"),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "kind_config"),
				),
			},
			{
				Config: testAccBasicWaitForReadyClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckNoResourceAttr(resourceName, "node_image"),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "true"),
					resource.TestCheckNoResourceAttr(resourceName, "kind_config"),
				),
			},
			{
				Config: testAccBasicExtraConfigClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckNoResourceAttr(resourceName, "node_image"),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "false"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.kind", "Cluster"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.api_version", "kind.x-k8s.io/v1alpha4"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.0.role", "control-plane"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.1.role", "worker"),
				),
			},
			{
				Config: testAccBasicWaitForReadyExtraConfigClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckNoResourceAttr(resourceName, "node_image"),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "true"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.kind", "Cluster"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.api_version", "kind.x-k8s.io/v1alpha4"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.0.role", "control-plane"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.1.role", "worker"),
				),
			},
			{
				Config: testAccNodeImageClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "node_image", kindDefaults.Image),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "kind_config"),
				),
			},
			{
				Config: testAccNodeImageWaitForReadyClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "node_image", kindDefaults.Image),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "true"),
					resource.TestCheckNoResourceAttr(resourceName, "kind_config"),
				),
			},
			{
				Config: testAccNodeImageExtraConfigClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "node_image", kindDefaults.Image),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "false"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.kind", "Cluster"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.api_version", "kind.x-k8s.io/v1alpha4"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.0.role", "control-plane"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.1.role", "worker"),
					// 					resource.TestCheckResourceAttr(resourceName, "kind_config", `kind: Cluster
					// apiVersion: kind.x-k8s.io/v1alpha4
					// nodes:
					// - role: control-plane
					// - role: worker
					// `),
				),
			},
			{
				Config: testAccNodeImageWaitForReadyExtraConfigClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "node_image", kindDefaults.Image),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "true"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.kind", "Cluster"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.api_version", "kind.x-k8s.io/v1alpha4"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.0.role", "control-plane"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.1.role", "worker"),
				),
			},
			{
				Config: testAccThreeNodesClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "node_image", kindDefaults.Image),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "true"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.kind", "Cluster"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.api_version", "kind.x-k8s.io/v1alpha4"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.0.role", "control-plane"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.1.role", "worker"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.2.role", "worker"),
				),
			},
			{
				Config: testAccThreeNodesImageOnNodeClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckNoResourceAttr(resourceName, "node_image"),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "true"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.kind", "Cluster"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.api_version", "kind.x-k8s.io/v1alpha4"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.0.role", "control-plane"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.0.image", nodeImage),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.1.role", "worker"),
					resource.TestCheckResourceAttr(resourceName, "kind_config.0.nodes.2.role", "worker"),
				),
			},
			// TODO: add this for when resource update is implemented
			// {
			// 	ResourceName:      resourceName,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

// testAccCheckKindClusterResourceDestroy verifies the kind cluster
// has been destroyed
func testAccCheckKindClusterResourceDestroy(clusterName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		prov := cluster.NewProvider()
		list, err := prov.List()
		if err != nil {
			return fmt.Errorf("cannot get kind provider cluster list")
		}
		for _, c := range list {
			if c == clusterName {
				return fmt.Errorf("list cannot contain cluster of name %s", clusterName)
			}
		}

		return nil
	}
}

func testAccCheckClusterCreate(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", name)
		}

		return nil
	}
}

func testAccBasicClusterConfig(name string) string {

	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
}
`, name)
}

func testAccNodeImageClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  node_image = "%s"
}
`, name, kindDefaults.Image)
}

func testAccBasicWaitForReadyClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  wait_for_ready = true
}
`, name)
}

func testAccNodeImageWaitForReadyClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  node_image = "%s"
  wait_for_ready = true
}
`, name, kindDefaults.Image)
}

func testAccBasicExtraConfigClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  kind_config {
	kind = "Cluster"
	api_version = "kind.x-k8s.io/v1alpha4"

	nodes {
		role = "control-plane"
	}

	nodes {
		role = "worker"
	}
  }
}
`, name)
}

func testAccBasicWaitForReadyExtraConfigClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  wait_for_ready = true
  kind_config {
	kind = "Cluster"
	api_version = "kind.x-k8s.io/v1alpha4"

	nodes {
		role = "control-plane"
	}

	nodes {
		role = "worker"
	}
  }
}
`, name)
}

func testAccNodeImageExtraConfigClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  node_image = "%s"
  kind_config {
	kind = "Cluster"
	api_version = "kind.x-k8s.io/v1alpha4"

	nodes {
		role = "control-plane"
	}

	nodes {
		role = "worker"
	}
  }
}
`, name, kindDefaults.Image)
}

func testAccNodeImageWaitForReadyExtraConfigClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  node_image = "%s"
  wait_for_ready = true
  kind_config {
	kind = "Cluster"
	api_version = "kind.x-k8s.io/v1alpha4"

	nodes {
		role = "control-plane"
	}

	nodes {
		role = "worker"
	}
  }
}
`, name, kindDefaults.Image)
}

func testAccThreeNodesClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  node_image = "%s"
  wait_for_ready = true
  kind_config {
	kind = "Cluster"
	api_version = "kind.x-k8s.io/v1alpha4"

	nodes {
		role = "control-plane"
	}

	nodes {
		role = "worker"
	}

	nodes {
		role = "worker"
	}
  }
}
`, name, kindDefaults.Image)
}

func testAccThreeNodesImageOnNodeClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  wait_for_ready = true
  kind_config {
	kind = "Cluster"
	api_version = "kind.x-k8s.io/v1alpha4"

	nodes {
		role = "control-plane"
		image = "%s"
	}

	nodes {
		role = "worker"
	}

	nodes {
		role = "worker"
	}
  }
}
`, name, nodeImage)
}
