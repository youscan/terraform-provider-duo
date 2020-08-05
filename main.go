package main

import (
	"github.com/youscan/terraform-provider-duo/duo"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return duo.Provider()
		},
	})
}
