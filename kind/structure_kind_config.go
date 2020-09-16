package kind

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"sigs.k8s.io/kind/pkg/apis/config/v1alpha4"
)

// Flatteners

func flattenKindConfig(d *schema.ResourceData) *v1alpha4.Cluster {
	obj := &v1alpha4.Cluster{}

	obj.Kind = d.Get("kind").(string)
	obj.APIVersion = d.Get("apiVersion").(string)

	nodes := d.Get("nodes").([]interface{})
	if nodes != nil {
		for _, n := range nodes {
			obj.Nodes = append(obj.Nodes, flattenKindConfigNodes(n.(*schema.ResourceData)))
		}
	}

	return obj
}

func flattenKindConfigNodes(d *schema.ResourceData) v1alpha4.Node {
	obj := v1alpha4.Node{}

	role := d.Get("role")
	if role != nil {
		obj.Role = role.(v1alpha4.NodeRole)
	}
	image := d.Get("image").(string)
	if image != "" {
		obj.Image = image
	}

	extraMounts := d.Get("extraMounts").([]interface{})
	if extraMounts != nil {
		for _, m := range extraMounts {
			obj.ExtraMounts = append(obj.ExtraMounts, flattenKindConfigExtraMounts(m.(*schema.ResourceData)))
		}
	}

	extraPortMappings := d.Get("extraPortMappings").([]interface{})
	if extraPortMappings != nil {
		for _, m := range extraPortMappings {
			obj.ExtraPortMappings = append(obj.ExtraPortMappings, flattenKindConfigExtraPortMappings(m.(*schema.ResourceData)))
		}
	}

	kubeadmConfigPatches := d.Get("kubeadmConfigPatches").([]string)
	if kubeadmConfigPatches != nil {
		for _, k := range kubeadmConfigPatches {
			obj.KubeadmConfigPatches = append(obj.KubeadmConfigPatches, k)
		}
	}

	return obj
}

func flattenKindConfigExtraMounts(d *schema.ResourceData) v1alpha4.Mount {
	obj := v1alpha4.Mount{}

	containerPath := d.Get("containerPath").(string)
	if containerPath != "" {
		obj.ContainerPath = containerPath
	}
	hostPath := d.Get("hostPath").(string)
	if hostPath != "" {
		obj.HostPath = hostPath
	}
	propagation := d.Get("propagation")
	if propagation != nil {
		obj.Propagation = propagation.(v1alpha4.MountPropagation)
	}

	readonly := d.Get("readonly").(bool)
	if hostPath != "" {
		obj.Readonly = readonly
	}
	selinuxRelabel := d.Get("selinuxRelabel").(bool)
	if hostPath != "" {
		obj.SelinuxRelabel = selinuxRelabel
	}

	return obj
}

func flattenKindConfigExtraPortMappings(d *schema.ResourceData) v1alpha4.PortMapping {
	obj := v1alpha4.PortMapping{}

	containerPort := d.Get("containerPort")
	if containerPort != nil {
		obj.ContainerPort = containerPort.(int32)
	}
	hostPort := d.Get("hostPort")
	if hostPort != nil {
		obj.HostPort = hostPort.(int32)
	}
	listenAddress := d.Get("listenAddress").(string)
	if listenAddress != "" {
		obj.ListenAddress = listenAddress
	}
	protocol := d.Get("protocol")
	if protocol != nil {
		obj.Protocol = protocol.(v1alpha4.PortMappingProtocol)
	}

	return obj
}
