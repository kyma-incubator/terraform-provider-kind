package kind

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Schemas

func kindConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"kind": {
			Type:     schema.TypeString,
			Optional: false,
		},
		"apiVersion": {
			Type:     schema.TypeString,
			Optional: false,
		},
		"nodes": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: kindConfigNodeFields(),
			},
		},
		"networking": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Resource{
				Schema: kindConfigNetworkingFields(),
			},
		},
	}
	return s
}

func kindConfigNodeFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"role": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"image": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"extraMounts": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: kindConfigNodeMountFields(),
			},
		},
		"extraPortMappings": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: kindConfigNodeExtraPortMappingsFields(),
			},
		},
		"kubeadmConfigPatches": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
	return s
}

func kindConfigNetworkingFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"ipFamily": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"apiServerAddress": {
			Type:        schema.TypeString,
			Description: `WARNING: It is _strongly_ recommended that you keep this the default (127.0.0.1) for security reasons. However it is possible to change this.`,
			Optional:    true,
		},
		"apiServerPort": {
			Type:        schema.TypeInt,
			Description: `By default the API server listens on a random open port. You may choose a specific port but probably don't need to in most cases. Using a random port makes it easier to spin up multiple clusters.`,
			Optional:    true,
		},
		"podSubnet": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"serviceSubnet": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"disableDefaultCNI": {
			Type:     schema.TypeBool,
			Default:  false,
			Optional: true,
		},
		"kubeProxyMode": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	return s
}

func kindConfigNodeMountFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"hostPath": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"containerPath": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	return s
}

func kindConfigNodeExtraPortMappingsFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"containerPort": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"hostPort": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"listenAddress": {
			Type:        schema.TypeString,
			Description: `optional: set the bind address on the host, 0.0.0.0 is the current default`,
			Optional:    true,
		},
		"protocol": {
			Type:        schema.TypeString,
			Description: `optional: set the protocol to one of TCP, UDP, SCTP. TCP is the default`,
			Optional:    true,
		},
	}
	return s
}
