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

func TestAccCluster(t *testing.T) {
	resourceName := "kind_cluster.test"
	clusterName := fmt.Sprintf("terraform-test-%s", acctest.RandString(10))
	nodeImage := "kindest/node:v1.18.4"

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
					resource.TestCheckResourceAttr(resourceName, "node_image", kindDefaults.Image),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "kind_config"),
				),
			},
			{
				Config: testAccBasicWaitForReadyClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "node_image", kindDefaults.Image),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "true"),
					resource.TestCheckNoResourceAttr(resourceName, "kind_config"),
				),
			},
			{
				Config: testAccBasicExtraConfigClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "node_image", kindDefaults.Image),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "false"),
					resource.TestCheckResourceAttr(resourceName, "kind_config", `kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
`),
				),
			},
			{
				Config: testAccBasicWaitForReadyExtraConfigClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "node_image", kindDefaults.Image),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "true"),
					resource.TestCheckResourceAttr(resourceName, "kind_config", `kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
`),
				),
			},
			{
				Config: testAccNodeImageClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "node_image", nodeImage),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "kind_config"),
				),
			},
			{
				Config: testAccNodeImageWaitForReadyClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "node_image", nodeImage),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "true"),
					resource.TestCheckNoResourceAttr(resourceName, "kind_config"),
				),
			},
			{
				Config: testAccNodeImageExtraConfigClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "node_image", nodeImage),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "false"),
					resource.TestCheckResourceAttr(resourceName, "kind_config", `kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
`),
				),
			},
			{
				Config: testAccNodeImageWaitForReadyExtraConfigClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterCreate(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "node_image", nodeImage),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "true"),
					resource.TestCheckResourceAttr(resourceName, "kind_config", `kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
`),
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
  node_image = "kindest/node:v1.18.4"
}
`, name)
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
  node_image = "kindest/node:v1.18.4"
  wait_for_ready = true
}
`, name)
}

func testAccBasicExtraConfigClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  kind_config = <<KIONF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
KIONF
}
`, name)
}

func testAccBasicWaitForReadyExtraConfigClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  wait_for_ready = true
  kind_config = <<KIONF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
KIONF
}
`, name)
}

func testAccNodeImageExtraConfigClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  node_image = "kindest/node:v1.18.4"
  kind_config = <<KIONF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
KIONF
}
`, name)
}

func testAccNodeImageWaitForReadyExtraConfigClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  node_image = "kindest/node:v1.18.4"
  wait_for_ready = true
  kind_config = <<KIONF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
KIONF
}
`, name)
}
