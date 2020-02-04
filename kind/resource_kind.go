package kind

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	kindDefaults "sigs.k8s.io/kind/pkg/apis/config/defaults"
	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/cmd"
)

var (
	profile string = "kind"
)

func resourceKindCreate(d *schema.ResourceData, meta interface{}) error {
	log.Println("Creating local Kubernetes cluster...")
	name := d.Get("name").(string)
	nodeImage := d.Get("node_image").(string)
	config := d.Get("kind_config").(string)

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

	log.Println("=================== Creating Kind Cluster ==================")
	provider := cluster.NewProvider(cluster.ProviderWithLogger(cmd.NewLogger()))
	err := provider.Create(name, copts...)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s-%s", name, nodeImage))
	return resourceKindRead(d, meta)
}

func resourceKindRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	provider := cluster.NewProvider(cluster.ProviderWithLogger(cmd.NewLogger()))
	id := d.Id()
	log.Printf("ID: %s\n", id)

	kconfig, err := provider.KubeConfig(name, true)
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

func resourceKindUpdate(d *schema.ResourceData, meta interface{}) error {
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
	return resourceKindRead(d, meta)
}

func resourceKindDelete(d *schema.ResourceData, meta interface{}) error {
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
