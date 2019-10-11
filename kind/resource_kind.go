package kind

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	kindDefaults "sigs.k8s.io/kind/pkg/apis/config/defaults"
	cluster "sigs.k8s.io/kind/pkg/cluster"
	create "sigs.k8s.io/kind/pkg/cluster/create"
)

var (
	profile string = "kind"
)

func resourceKindCreate(d *schema.ResourceData, meta interface{}) error {
	log.Println("Creating local Kubernetes cluster...")
	name := d.Get("name").(string)
	ctx := cluster.NewContext(name)
	baseImage := d.Get("node_image").(string)

	log.Println("=================== Creating Kind Cluster ==================")
	var opts []create.ClusterOption
	opts = append(opts, create.SetupKubernetes(true))
	if baseImage != "" {
		opts = append(opts, create.WithNodeImage(baseImage))
		log.Printf("Using defined base image: %s\n", baseImage)
	} else {
		d.Set("node_image", kindDefaults.Image) // set image to k/kind default image.
		baseImage = kindDefaults.Image
	}

	err := ctx.Create(opts...)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s-%s", name, baseImage))
	return resourceKindRead(d, meta)
}

func resourceKindRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	ctx := cluster.NewContext(name)
	id := d.Id()
	log.Printf("ID: %s\n", id)

	err := ctx.Validate()
	if err != nil {
		d.SetId("")
		return err
	}

	k8sKubeconfigPath := ctx.KubeConfigPath()
	d.Set("k8s_kubeconfig_path", k8sKubeconfigPath)

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

	d.Partial(false)
	return resourceKindRead(d, meta)
}

func resourceKindDelete(d *schema.ResourceData, meta interface{}) error {
	log.Println("Deleting local Kubernetes cluster...")
	name := d.Get("name").(string)
	ctx := cluster.NewContext(name)

	log.Println("=================== Deleting Kind Cluster ==================")
	err := ctx.Delete()
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
