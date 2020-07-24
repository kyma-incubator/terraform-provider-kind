package kind

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	defaultCreateTimeout = time.Minute * 5
	defaultUpdateTimeout = time.Minute * 5
	defaultDeleteTimeout = time.Minute * 5
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"kind_cluster": resourceCluster(),
		},
	}
}
