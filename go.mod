module github.com/kyma-incubator/terraform-provider-kind

go 1.16

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.10.1
	github.com/pelletier/go-toml v1.8.1
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/kind v0.11.1
)

replace k8s.io/client-go => k8s.io/client-go v0.20.2
