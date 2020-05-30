package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a schema.Provider for VScale.
func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"vscale_ssh_key": dataSourceVScaleSSHKey(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"vscale_scalet":  resourceScalet(),
			"vscale_ssh_key": resourceSSHKey(),
			"vscale_domain":  resourceDomain(),
			"vscale_record":  resourceRecord(),
		},
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSCALE_API_TOKEN", nil),
				Description: "The token key for API operations.",
			},
		},
	}

	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := p.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}
	return p
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	c := Config{
		Token:            d.Get("token").(string),
		TerraformVersion: terraformVersion,
	}
	client := c.Client()
	return client, nil
}
