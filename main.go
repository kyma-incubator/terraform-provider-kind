package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/kyma-incubator/terraform-provider-kind/kind"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kind.Provider})
}
