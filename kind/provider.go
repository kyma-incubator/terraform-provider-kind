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
