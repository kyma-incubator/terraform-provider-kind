package kind

import (
	"sigs.k8s.io/kind/pkg/apis/config/v1alpha4"
)

// Flatteners

func flattenKindConfig(d map[string]interface{}) *v1alpha4.Cluster {
	obj := &v1alpha4.Cluster{}

	obj.Kind = mapKeyIfExists(d, "kind").(string)
	obj.APIVersion = mapKeyIfExists(d, "api_version").(string)

	nodes := mapKeyIfExists(d, "nodes").([]interface{})
	if nodes != nil {
		for _, n := range nodes {
			data := n.(map[string]interface{})
			obj.Nodes = append(obj.Nodes, flattenKindConfigNodes(data))
		}
	}

	return obj
}

func flattenKindConfigNodes(d map[string]interface{}) v1alpha4.Node {
	obj := v1alpha4.Node{}

	role := mapKeyIfExists(d, "role").(string)
	if role != "" {
		switch role {
		case string(v1alpha4.ControlPlaneRole):
			obj.Role = v1alpha4.ControlPlaneRole
		case string(v1alpha4.WorkerRole):
			obj.Role = v1alpha4.WorkerRole
		}
	}
	image := mapKeyIfExists(d, "image").(string)
	if image != "" {
		obj.Image = image
	}

	extraMounts := mapKeyIfExists(d, "extra_mounts").([]interface{})
	if extraMounts != nil {
		for _, m := range extraMounts {
			data := m.(map[string]interface{})
			obj.ExtraMounts = append(obj.ExtraMounts, flattenKindConfigExtraMounts(data))
		}
	}

	extraPortMappings := mapKeyIfExists(d, "extra_port_mappings").([]interface{})
	if extraPortMappings != nil {
		for _, m := range extraPortMappings {
			data := m.(map[string]interface{})
			obj.ExtraPortMappings = append(obj.ExtraPortMappings, flattenKindConfigExtraPortMappings(data))
		}
	}

	kubeadmConfigPatches := mapKeyIfExists(d, "kubeadm_config_patches").([]interface{})
	if kubeadmConfigPatches != nil {
		for _, k := range kubeadmConfigPatches {
			data := k.(string)
			obj.KubeadmConfigPatches = append(obj.KubeadmConfigPatches, data)
		}
	}

	return obj
}

func flattenKindConfigExtraMounts(d map[string]interface{}) v1alpha4.Mount {
	obj := v1alpha4.Mount{}

	containerPath := mapKeyIfExists(d, "container_path").(string)
	if containerPath != "" {
		obj.ContainerPath = containerPath
	}
	hostPath := mapKeyIfExists(d, "host_path").(string)
	if hostPath != "" {
		obj.HostPath = hostPath
	}
	propagation := mapKeyIfExists(d, "propagation").(string)
	if propagation != "" {
		switch propagation {
		case string(v1alpha4.MountPropagationBidirectional):
			obj.Propagation = v1alpha4.MountPropagationBidirectional
		case string(v1alpha4.MountPropagationHostToContainer):
			obj.Propagation = v1alpha4.MountPropagationHostToContainer
		case string(v1alpha4.MountPropagationNone):
			obj.Propagation = v1alpha4.MountPropagationNone
		}
	}

	readonly := mapKeyIfExists(d, "readonly").(bool)
	if hostPath != "" {
		obj.Readonly = readonly
	}
	selinuxRelabel := mapKeyIfExists(d, "selinux_relabel").(bool)
	if hostPath != "" {
		obj.SelinuxRelabel = selinuxRelabel
	}

	return obj
}

func flattenKindConfigExtraPortMappings(d map[string]interface{}) v1alpha4.PortMapping {
	obj := v1alpha4.PortMapping{}

	containerPort := mapKeyIfExists(d, "container_port")
	if containerPort != nil {
		obj.ContainerPort = containerPort.(int32)
	}
	hostPort := mapKeyIfExists(d, "host_port")
	if hostPort != nil {
		obj.HostPort = hostPort.(int32)
	}
	listenAddress := mapKeyIfExists(d, "listen_address").(string)
	if listenAddress != "" {
		obj.ListenAddress = listenAddress
	}
	protocol := mapKeyIfExists(d, "protocol")
	if protocol != nil {
		switch protocol {
		case v1alpha4.PortMappingProtocolSCTP:
			obj.Protocol = v1alpha4.PortMappingProtocolSCTP
		case v1alpha4.PortMappingProtocolTCP:
			obj.Protocol = v1alpha4.PortMappingProtocolTCP
		case v1alpha4.PortMappingProtocolUDP:
			obj.Protocol = v1alpha4.PortMappingProtocolUDP
		}
	}

	return obj
}

func mapKeyIfExists(m map[string]interface{}, key string) interface{} {
	if val, ok := m[key]; ok {
		return val
	}
	return nil
}
