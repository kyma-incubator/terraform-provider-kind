package kind

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
	kubernetesVersion := d.Get("k8s_version").(string)

	log.Println("=================== Creating Kind Cluster ==================")
	k8Setup := create.SetupKubernetes(true)
	k8sVersion := create.WithNodeImage(fmt.Sprintf("kindest/node:%s", kubernetesVersion))

	err := ctx.Create(k8Setup, k8sVersion)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s-%s", name, kubernetesVersion))
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
	log.Println("update not yet implemented")
	return nil
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
