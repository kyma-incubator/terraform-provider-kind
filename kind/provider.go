package kind

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

const (
	defaultCreateTimeout = time.Minute * 30
	defaultUpdateTimeout = time.Minute * 30
	defaultDeleteTimeout = time.Minute * 20
)

func Provider() terraform.ResourceProvider {
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
		// Update: resourceKindUpdate,
		Delete: resourceKindDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(defaultCreateTimeout),
			Update: schema.DefaultTimeout(defaultUpdateTimeout),
			Delete: schema.DefaultTimeout(defaultDeleteTimeout),
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The kind name that is given to the created cluster.",
				Required:    true,
				ForceNew:    true,
			},
			"node_image": &schema.Schema{
				Type:        schema.TypeString,
				Description: `The node_image that kind will use (ex: kindest/node:v1.15.3).`,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
			},
			"kind_config": &schema.Schema{
				Type:        schema.TypeString,
				Description: `The kind_config that kind will use.`,
				Optional:    true,
				ForceNew:    true,
			},
			"kubeconfig": &schema.Schema{
				Type:        schema.TypeString,
				Description: `Kubeconfig set after the the cluster is created.`,
				Computed:    true,
			},
			"kubeconfig_path": &schema.Schema{
				Type:        schema.TypeString,
				Description: `Kubeconfig path set after the the cluster is created.`,
				Computed:    true,
			},
			"client_certificate": &schema.Schema{
				Type:        schema.TypeString,
				Description: `Client certificate for authenticating to cluster.`,
				Computed:    true,
			},
			"client_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: `Client key for authenticating to cluster.`,
				Computed:    true,
			},
			"cluster_ca_certificate": &schema.Schema{
				Type:        schema.TypeString,
				Description: `Client verifies the server certificate with this CA cert.`,
				Computed:    true,
			},
			"endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Description: `Kubernetes APIServer endpoint.`,
				Computed:    true,
			},
		},
	}
}
