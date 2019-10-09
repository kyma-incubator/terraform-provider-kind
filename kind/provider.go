package kind

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"kind": resourceKind(),
		},
	}
}

func resourceKind() *schema.Resource {
	return &schema.Resource{
		Create: resourceKindCreate,
		Read:   resourceKindRead,
		Update: resourceKindUpdate,
		Delete: resourceKindDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The kind name that is given to the created cluster",
				Required:    true,
			},
			"k8s_version": &schema.Schema{
				Type:        schema.TypeString,
				Description: `The kubernetes version that the kind will use (ex: v1.15.3) valid values are tags from https://hub.docker.com/r/kindest/node/tags`,
				Required:    true,
			},
			"k8s_kubeconfig_path": &schema.Schema{
				Type:        schema.TypeString,
				Description: `Kubeconfig path set after the the cluster is created.`,
				Computed:    true,
			},
		},
	}
}
