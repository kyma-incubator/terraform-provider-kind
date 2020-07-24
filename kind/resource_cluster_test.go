package kind

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
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
		CheckDestroy: testAccCheckKindClusterResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSimpleClusterConfig(clusterName),
				Check: resource.ComposeTestCheckFunc(
					testAccSimpleCluster(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", clusterName),
					resource.TestCheckResourceAttr(resourceName, "node_image", nodeImage),
					resource.TestCheckResourceAttr(resourceName, "wait_for_ready", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "kind_config"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// testAccCheckKindClusterResourceDestroy verifies the kind cluster
// has been destroyed
func testAccCheckKindClusterResourceDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	// conn := testAccProvider.Meta().(*Kind)

	// loop through the resources in state, verifying each widget
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		panic(rs)
	}

	return nil
}

func testAccSimpleCluster(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", name)
		}

		return nil
	}
}

func testAccSimpleClusterConfig(name string) string {
	return fmt.Sprintf(`
resource "kind_cluster" "test" {
  name = "%s"
  node_image = "kindest/node:v1.18.4"
}
`, name)
}
