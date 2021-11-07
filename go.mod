module github.com/kyma-incubator/terraform-provider-kind

go 1.16

require (
	github.com/apparentlymart/go-cidr v1.1.0 // indirect
	github.com/aws/aws-sdk-go v1.31.9 // indirect
	github.com/hashicorp/go-getter v1.4.2-0.20200106182914-9813cbd4eb02 // indirect
	github.com/hashicorp/go-plugin v1.3.0 // indirect
	github.com/hashicorp/hcl/v2 v2.6.0 // indirect
	github.com/hashicorp/terraform-config-inspect v0.0.0-20191212124732-c6ae6269b9d7 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/pelletier/go-toml v1.8.1
	github.com/zclconf/go-cty v1.5.1 // indirect
	github.com/zclconf/go-cty-yaml v1.0.2 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/kind v0.11.1
)

replace k8s.io/client-go => k8s.io/client-go v0.20.2
