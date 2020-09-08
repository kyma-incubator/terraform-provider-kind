package kind

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	kindDefaults "sigs.k8s.io/kind/pkg/apis/config/defaults"
	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/cmd"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceKindClusterCreate,
		Read:   resourceKindClusterRead,
		// Update: resourceKindClusterUpdate,
		Delete: resourceKindClusterDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(defaultCreateTimeout),
			Update: schema.DefaultTimeout(defaultUpdateTimeout),
			Delete: schema.DefaultTimeout(defaultDeleteTimeout),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The kind name that is given to the created cluster.",
				Required:    true,
				ForceNew:    true,
			},
			"node_image": {
				Type:        schema.TypeString,
				Description: `The node_image that kind will use (ex: kindest/node:v1.15.3).`,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
			},
			"wait_for_ready": {
				Type:        schema.TypeBool,
				Description: `Defines wether or not the provider will wait for the control plane to be ready. Defaults to false`,
				Default:     false,
				ForceNew:    true, // TODO remove this once we have the update method defined.
				Optional:    true,
			},
			"kind_config": {
				Type:        schema.TypeString,
				Description: `The kind_config that kind will use.`,
				Optional:    true,
				ForceNew:    true,
			},
			"kubeconfig": {
				Type:        schema.TypeString,
				Description: `Kubeconfig set after the the cluster is created.`,
				Computed:    true,
			},
			"kubeconfig_path": {
				Type:        schema.TypeString,
				Description: `Kubeconfig path set after the the cluster is created.`,
				Computed:    true,
			},
			"client_certificate": {
				Type:        schema.TypeString,
				Description: `Client certificate for authenticating to cluster.`,
				Computed:    true,
			},
			"client_key": {
				Type:        schema.TypeString,
				Description: `Client key for authenticating to cluster.`,
				Computed:    true,
			},
			"cluster_ca_certificate": {
				Type:        schema.TypeString,
				Description: `Client verifies the server certificate with this CA cert.`,
				Computed:    true,
			},
			"endpoint": {
				Type:        schema.TypeString,
				Description: `Kubernetes APIServer endpoint.`,
				Computed:    true,
			},
		},
	}
}

func resourceKindClusterCreate(d *schema.ResourceData, meta interface{}) error {
	log.Println("Creating local Kubernetes cluster...")
	name := d.Get("name").(string)
	nodeImage := d.Get("node_image").(string)
	config := d.Get("kind_config").(string)
	waitForReady := d.Get("wait_for_ready").(bool)

	var copts []cluster.CreateOption
	if config != "" {
		copts = append(copts, cluster.CreateWithRawConfig([]byte(config)))
	}

	if nodeImage != "" {
		copts = append(copts, cluster.CreateWithNodeImage(nodeImage))
		log.Printf("Using defined node_image: %s\n", nodeImage)
	} else {
		d.Set("node_image", kindDefaults.Image) // set image to k/kind default image.
		nodeImage = kindDefaults.Image
	}

	if waitForReady {
		copts = append(copts, cluster.CreateWithWaitForReady(defaultCreateTimeout))
		log.Printf("Will wait for cluster nodes to report ready: %t\n", waitForReady)
	}

	log.Println("=================== Creating Kind Cluster ==================")
	provider := cluster.NewProvider(cluster.ProviderWithLogger(cmd.NewLogger()))
	err := provider.Create(name, copts...)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s-%s", name, nodeImage))
	return resourceKindClusterRead(d, meta)
}

func resourceKindClusterRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	provider := cluster.NewProvider(cluster.ProviderWithLogger(cmd.NewLogger()))
	id := d.Id()
	log.Printf("ID: %s\n", id)

	kconfig, err := provider.KubeConfig(name, false)
	if err != nil {
		d.SetId("")
		return err
	}
	d.Set("kubeconfig", kconfig)

	currentPath, err := os.Getwd()
	if err != nil {
		d.SetId("")
		return err
	}
	exportPath := fmt.Sprintf("%s%s%s-config", currentPath, string(os.PathSeparator), name)
	err = provider.ExportKubeConfig(name, exportPath)
	if err != nil {
		d.SetId("")
		return err
	}
	d.Set("kubeconfig_path", exportPath)

	// use the current context in kubeconfig
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kconfig))
	if err != nil {
		return err
	}

	d.Set("client_certificate", string(config.CertData))
	d.Set("client_key", string(config.KeyData))
	d.Set("cluster_ca_certificate", string(config.CAData))
	d.Set("endpoint", string(config.Host))

	d.Set("completed", true)

	return nil
}

func resourceKindClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Println("")
	d.Partial(true)

	if d.HasChange("node_image") {
		d.SetPartial("node_image")
	}
	if d.HasChange("name") {
		d.SetPartial("name")
	}
	if d.HasChange("kind_config") {
		d.SetPartial("kind_config")
	}

	d.Partial(false)
	return resourceKindClusterRead(d, meta)
}

func resourceKindClusterDelete(d *schema.ResourceData, meta interface{}) error {
	log.Println("Deleting local Kubernetes cluster...")
	name := d.Get("name").(string)
	kubeconfigPath := d.Get("kubeconfig_path").(string)
	provider := cluster.NewProvider(cluster.ProviderWithLogger(cmd.NewLogger()))

	log.Println("=================== Deleting Kind Cluster ==================")
	err := provider.Delete(name, kubeconfigPath)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
