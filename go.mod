module github.com/kyma-incubator/terraform-provider-kind

go 1.13

require (
	github.com/hashicorp/terraform v0.12.28
	github.com/imdario/mergo v0.3.9 // indirect
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/utils v0.0.0-20200619165400-6e3d28b6ed19 // indirect
	sigs.k8s.io/kind v0.8.1
)

replace k8s.io/client-go => k8s.io/client-go v0.18.5
